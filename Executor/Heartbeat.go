package Executor

import (
	"context"
	"time"

	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

func (self *Executor) Heartbeat() {
	for {
		if err := self.DoHeartbeat(); err != nil {
			time.Sleep(30 * time.Second)
		}
	}
}

func (self *Executor) DoHeartbeat() error {
	grpcConn, err := grpc.Dial(self.MasterAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer grpcConn.Close()

	client := pb.NewGueryMasterClient(grpcConn)
	stream, err := client.SendHeartbeat(context.Background())
	if err != nil {
		return err
	}

	ticker := time.NewTicker(10 * time.Second)
	quickTicker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-quickTicker.C:
			if self.IsStatusChanged {
				self.IsStatusChanged = false
				if err := self.SendOneHeartbeat(stream); err != nil {
					return err
				}
			}
		case <-ticker.C:
			if err := self.SendOneHeartbeat(stream); err != nil {
				return err
			}
		}
	}
}

func (self *Executor) SendOneHeartbeat(stream pb.GueryMaster_SendHeartbeatClient) error {
	hb := &pb.Heartbeat{}

	if err := stream.Send(hb); err != nil {
		return err
	}
	return nil
}
