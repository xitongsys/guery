package Master

import (
	"bytes"
	"context"
	"encoding/gob"

	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

func (self *Master) Query(ctx context.Context, sql *pb.QuerySQL) (*pb.QueryResponse, error) {
	sqlStr := sql.SQL
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
			freeExecutors = append(freeExecutors, einfo.Heartbeat.Location)
		}
	}

	if err = EPlan.CreateEPlan(logicalPlanTree, &ePlanNodes, freeExecutors, 2); err == nil {
		for _, enode := range ePlanNodes {
			buf := bytes.Buffer
			enc := gob.NewEncoder(buf)
			enc.Encode(enode)
			instruction := pb.Instruction{
				TaskId:                1,
				TaskType:              int32(enode.GetEPlanNodeType()),
				EncodedEPlanNodeBytes: buf.String(),
			}

			grpcConn, err := grpc.Dial(enode.GetLocation().GetAddress(), grpc.WithInsecure())

		}
	}

	self.Topology.Unlock()

	return nil, err
}
