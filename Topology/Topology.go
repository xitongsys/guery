package Topology

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/satori/go.uuid"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

type AgentInfo struct {
	Name              string
	Heartbeat         pb.AgentHeartbeat
	LastHeartBeatTime time.Time
}

func NewAgentInfo(hb *pb.AgentHeartbeat) *AgentInfo {
	return &AgentInfo{
		Name:              hb.Location.Name,
		Heartbeat:         *hb,
		LastHeartBeatTime: time.Now(),
	}
}

//Topology/////////////////
type Topology struct {
	sync.RWMutex
	Agents      map[string]*AgentInfo
	AgentNumber int32
}

func NewTopology() *Topology {
	return &Topology{
		Agents: make(map[string]*AgentInfo),
	}
}

func (self *Topology) RestartAgent(name string) error {
	self.Lock()
	self.Unlock()
	if _, ok := self.Agents[name]; !ok {
		return fmt.Errorf("agent not found")
	}
	loc := self.Agents[name].Heartbeat.GetLocation()
	grpcConn, err := grpc.Dial(loc.GetURL(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewGueryAgentClient(grpcConn)
	_, err = client.Restart(context.Background(), &pb.Empty{})
	grpcConn.Close()
	return err
}

func (self *Topology) DuplicateAgent(name string) error {
	self.Lock()
	self.Unlock()
	if _, ok := self.Agents[name]; !ok {
		return fmt.Errorf("agent not found")
	}
	loc := self.Agents[name].Heartbeat.GetLocation()
	grpcConn, err := grpc.Dial(loc.GetURL(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewGueryAgentClient(grpcConn)
	_, err = client.Duplicate(context.Background(), &pb.Empty{})
	grpcConn.Close()
	return err
}

func (self *Topology) KillAgent(name string) error {
	self.Lock()
	self.Unlock()
	if _, ok := self.Agents[name]; !ok {
		return fmt.Errorf("agent not found")
	}
	loc := self.Agents[name].Heartbeat.GetLocation()
	grpcConn, err := grpc.Dial(loc.GetURL(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewGueryAgentClient(grpcConn)
	_, err = client.Quit(context.Background(), &pb.Empty{})
	grpcConn.Close()
	return err
}

func (self *Topology) GetAgents() []pb.Location {
	self.Lock()
	defer self.Unlock()

	res := []pb.Location{}
	for _, einfo := range self.Agents {
		res = append(res, *einfo.Heartbeat.Location)
	}
	return res
}

func (self *Topology) UpdateAgentInfo(hb *pb.AgentHeartbeat) {
	ts := time.Now()

	self.Lock()
	defer self.Unlock()
	agentInfo := NewAgentInfo(hb)
	if ts.After(agentInfo.LastHeartBeatTime) {
		return
	}
	self.Agents[hb.Location.Name] = agentInfo
	self.AgentNumber = int32(len(self.Agents))
}

func (self *Topology) DropAgentInfo(location *pb.Location) {
	self.Lock()
	defer self.Unlock()

	if _, ok := self.Agents[location.Name]; !ok {
		return
	} else {
		delete(self.Agents, location.Name)
	}
	self.AgentNumber = int32(len(self.Agents))
}

func (self *Topology) GetFreeExecutorNumber() int32 {
	self.Lock()
	defer self.Unlock()
	res := int32(0)
	for _, info := range self.Agents {
		res += int32(info.Heartbeat.MaxExecutorNumber - info.Heartbeat.ExecutorNumber)
	}
	return res
}

func (self *Topology) GetFreeExecutors(number int32) []pb.Location {
	self.Lock()
	defer self.Unlock()
	res := []pb.Location{}
	for _, info := range self.Agents {
		num := (info.Heartbeat.MaxExecutorNumber - info.Heartbeat.ExecutorNumber)
		for i := 0; i < int(num) && len(res) < int(number); i++ {
			exe := pb.Location{
				Name:    "executor_" + uuid.Must(uuid.NewV4()).String(),
				Address: info.Heartbeat.Location.Address,
				Port:    info.Heartbeat.Location.Port,
			}
			res = append(res, exe)
		}
		if len(res) >= int(number) {
			break
		}
	}
	return res
}

///////////////////////////
