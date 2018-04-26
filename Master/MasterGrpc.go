package Master

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/xitongsys/guery/pb"
)

type MasterGrpc struct {
}

func (self *MasterGrpc) SendHeartbeat(stream pb.GueryMaster_SendHeartbeatServer) error {
	var location *pb.Location
	for {
		heartbeat, err := stream.Recv()
		if err == nil {
		} else {
		}
	}
}
