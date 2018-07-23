package Agent

import (
	"context"
	"io"
	"time"

	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

//executor heartbeat
func (self *Agent) SendHeartbeat(stream pb.GueryAgent_SendHeartbeatServer) error {
	var location *pb.Location
	for {
		hb, err := stream.Recv()
		if err == nil {
			if location == nil {
				location = hb.Location
				Logger.Infof("Add executor %v", location)
			}

		} else {
			if location != nil {
				self.Topology.DropExecutorInfo(location)
				Logger.Infof("Lost executor %v: %v", location, err)
			}
			if err == io.EOF {
				Logger.Infof("Lost executor %v: %v", location, err)
				return nil
			}
			if err != nil {
				Logger.Infof("Lost executor %v: %v", location, err)
				return err
			}
		}
		self.Topology.UpdateExecutorInfo(hb)
		self.Tasks.UpdateTaskInfo(hb)
	}
}

//master heartbeat
func (self *Agent) Heartbeat() {
	for {
		if err := self.DoHeartbeat(); err != nil {
			time.Sleep(3 * time.Second)
		}
	}
}

func (self *Agent) DoHeartbeat() error {
	grpcConn, err := grpc.Dial(self.MasterAddress, grpc.WithInsecure())
	if err != nil {
		Logger.Errorf("DoHeartBeat failed: %v", err)
		return err
	}
	defer grpcConn.Close()

	client := pb.NewGueryMasterClient(grpcConn)
	stream, err := client.SendHeartbeat(context.Background())
	if err != nil {
		return err
	}

	ticker := time.NewTicker(1 * time.Second)
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

func (self *Agent) SendOneHeartbeat(stream pb.GueryMaster_SendHeartbeatClient) error {
	hb := self.GetInfo()
	if err := stream.Send(hb); err != nil {
		Logger.Errorf("failed to SendOneHeartbeat: %v, %v", err, hb)
		return err
	}
	return nil
}
