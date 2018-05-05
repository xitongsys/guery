package Master

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/parser"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

func (self *Master) QueryHandler(response http.ResponseWriter, request *http.Request) {
	var err error

	if err = request.ParseForm(); err != nil {
		response.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	sqlStr := request.FormValue("sql")

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
		if einfo.Heartbeat.Resource == 0 {
			freeExecutors = append(freeExecutors, *einfo.Heartbeat.Location)
		}
	}

	if _, err = EPlan.CreateEPlan(logicalPlanTree, &ePlanNodes, freeExecutors, 2); err == nil {
		for _, enode := range ePlanNodes {
			buf := new(bytes.Buffer)
			enc := gob.NewEncoder(buf)
			enc.Encode(enode)
			instruction := pb.Instruction{
				TaskId:                1,
				TaskType:              int32(enode.GetNodeType()),
				EncodedEPlanNodeBytes: buf.String(),
			}

			loc := enode.GetLocation()
			grpcConn, err := grpc.Dial(loc.GetURL(), grpc.WithInsecure())
			if err != nil {
				break
			}
			client := pb.NewGueryExecutorClient(grpcConn)
			if _, err = client.SendInstruction(context.Background(), &instruction); err != nil {
				grpcConn.Close()
				break
			}
			grpcConn.Close()
		}
	}

	self.Topology.Unlock()

	response.Write([]byte("QueryHandler"))
}
