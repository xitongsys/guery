package Master

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Util"
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

	var aggNode EPlan.ENode
	if aggNode, err = EPlan.CreateEPlan(logicalPlanTree, &ePlanNodes, &freeExecutors, 1); err == nil {

		for _, enode := range ePlanNodes {
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
					Logger.Errorf("failed setup readers %v: %v", loc, err)
					grpcConn.Close()
					break
				}

				if _, err = client.Run(context.Background(), &empty); err != nil {
					Logger.Errorf("failed run %v: %v", loc, err)
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
		return
	}

	self.CollectResults(response, aggNode)

}

func (self *Master) CollectResults(response http.ResponseWriter, enode EPlan.ENode) {
	output := enode.GetLocation()
	conn, err := grpc.Dial(output.GetURL(), grpc.WithInsecure())
	if err != nil {
		return
	}
	client := pb.NewGueryExecutorClient(conn)
	inputChannelLocation, err := client.GetOutputChannelLocation(context.Background(), &output)
	if err != nil {
		return
	}
	conn.Close()

	cconn, err := net.Dial("tcp", inputChannelLocation.GetURL())
	if err != nil {
		Logger.Errorf("failed to connect to input channel %v: %v", inputChannelLocation, err)
		return
	}

	//Util.CopyBuffer(cconn, response)
	var msg []byte
	md := &Util.Metadata{}
	if err = Util.ReadObject(cconn, md); err != nil {
		return
	}

	if msg, err = json.Marshal(md); err != nil {
		Logger.Errorf("json marshal: %v", err)
		return
	}

	if n, err := response.Write(msg); n != len(msg) || err != nil {
		return
	}

	for {
		row, err := Util.ReadRow(cconn)
		Logger.Infof("====== %v", row)
		if err != nil {
			break
		}

		if msg, err = json.Marshal(row); err != nil {
			break
		}

		if n, err := response.Write(msg); n != len(msg) || err != nil {
			break
		}

	}

}
