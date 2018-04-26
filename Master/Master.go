package Master

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/xitongsys/guery/pb"
)

type Master struct {
	Topology  *Topology
	StartTime time.Time
}

func NewMaster() *Master {
	m := &Master{
		Topology:  NewTopology(),
		StartTime: time.Now(),
	}
	return m
}

func (self *Master) SendHeartbeat(stream pb.GueryMaster_SendHeartbeatServer) error {
	var location *pb.Location
	for {
		heartbeat, err := stream.Recv()
		if err == nil {
		} else {
		}
	}
}
