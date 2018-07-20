package Agent

import (
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/pb"
)

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

	}
}
