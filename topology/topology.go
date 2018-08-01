package topology

import (
	"container/heap"
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

func (self *Topology) DropAgentInfo(hb *pb.AgentHeartbeat) {
	self.Lock()
	defer self.Unlock()

	location := hb.Location
	if _, ok := self.Agents[location.Name]; !ok {
		return
	} else {
		delete(self.Agents, location.Name)
	}
	self.AgentNumber = int32(len(self.Agents))
}

func (self *Topology) GetExecutors(number int) ([]pb.Location, []pb.Location) {
	self.Lock()
	defer self.Unlock()

	agents, executors := []pb.Location{}, []pb.Location{}
	agentMap := map[string]pb.Location{}
	pq := &Heap{}
	heap.Init(pq)

	for _, info := range self.Agents {
		item := NewItem(*info.Heartbeat.Location, int(info.Heartbeat.ExecutorNumber))
		heap.Push(pq, item)
	}

	for i := 0; i < number; i++ {
		item := heap.Pop(pq).(*Item)
		exe := pb.Location{
			Name:    "executor_" + uuid.Must(uuid.NewV4()).String(),
			Address: item.Location.Address,
			Port:    item.Location.Port,
		}
		executors = append(executors, exe)
		agentMap[item.Location.Name] = item.Location
		item.ExecutorNumber++
		heap.Push(pq, item)
	}
	for _, v := range agentMap {
		agents = append(agents, v)
	}
	return agents, executors
}

///////////////////////////
