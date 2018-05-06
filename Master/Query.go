package Master

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/parser"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

func (self *Master) QueryHandler(response http.ResponseWriter, request *http.Request) {
	Logger.Infof("QueryHandler")
	var err error

	if err = request.ParseForm(); err != nil {
		response.Write([]byte(fmt.Sprintf("Request Error: %v", err)))
		return
	}
	sqlStr := request.FormValue("sql")
	catalog := request.FormValue("catalog")
	schema := request.FormValue("schema")

	is := antlr.NewInputStream(sqlStr)
	lexer := parser.NewSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSqlParser(stream)
	tree := p.SingleStatement()

	self.Topology.Lock()

	logicalPlanTree := Plan.NewPlanNodeFromSingleStatement(tree)
	ePlanNodes := []EPlan.ENode{}
	freeExecutors := []pb.Location{}

	for _, einfo := range self.Topology.Executors {
		if einfo.Heartbeat.Status == 0 {
			freeExecutors = append(freeExecutors, *einfo.Heartbeat.Location)
		}
	}

	if _, err = EPlan.CreateEPlan(logicalPlanTree, &ePlanNodes, &freeExecutors, 1); err == nil {
		for _, enode := range ePlanNodes {
			Logger.Infof("======%v, %v", enode, len(ePlanNodes))
			var buf bytes.Buffer
			gob.NewEncoder(&buf).Encode(enode)

			instruction := pb.Instruction{
				TaskId:                1,
				TaskType:              int32(enode.GetNodeType()),
				Catalog:               catalog,
				Schema:                schema,
				EncodedEPlanNodeBytes: buf.String(),
			}
			instruction.Base64Encode()

			loc := enode.GetLocation()
			var grpcConn *grpc.ClientConn
			grpcConn, err = grpc.Dial(loc.GetURL(), grpc.WithInsecure())
			if err != nil {
				break
			}
			client := pb.NewGueryExecutorClient(grpcConn)
			if _, err = client.SendInstruction(context.Background(), &instruction); err != nil {
				grpcConn.Close()
				break
			}

			empty := pb.Empty{}
			if _, err = client.SetupWriters(context.Background(), &empty); err != nil {
				grpcConn.Close()
				break
			}
			grpcConn.Close()
		}

		if err == nil {
			for _, enode := range ePlanNodes {
				loc := enode.GetLocation()
				var grpcConn *grpc.ClientConn
				grpcConn, err = grpc.Dial(loc.GetURL(), grpc.WithInsecure())
				if err != nil {
					break
				}
				client := pb.NewGueryExecutorClient(grpcConn)
				empty := pb.Empty{}
				if _, err = client.SetupReaders(context.Background(), &empty); err != nil {
					grpcConn.Close()
					break
				}
				grpcConn.Close()
			}
		}

	}

	self.Topology.Unlock()

	if err != nil {
		Logger.Errorf("%v", err)
		response.Write([]byte(err.Error()))
	}
}
