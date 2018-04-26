package Master

import (
	"context"
	"fmt"
	"log"
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
		hb, err := stream.Recv()
		if err == nil {
			if location == nil {
				location = hb.Location
				log.Println("Add executor: %v", location)
			}

		} else {
			if location != nil {
				self.Topology.DropExecutorInfo(location)
				log.Println("Lost agent: %v", location)
			}
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
		}
		self.Topology.UpdateExecutorInfo(hb)
	}
}
