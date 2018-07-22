package Agent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

type ExecutorInfo struct {
	Name              string
	Heartbeat         pb.ExecutorHeartbeat
	LastHeartBeatTime time.Time
}

func NewExecutorInfo(hb *pb.ExecutorHeartbeat) *ExecutorInfo {
	return &ExecutorInfo{
		Name:              hb.Location.Name,
		Heartbeat:         *hb,
		LastHeartBeatTime: time.Now(),
	}
}

//Topology/////////////////
type Topology struct {
	sync.RWMutex
	Executors      map[string]*ExecutorInfo
	ExecutorNumber int32
}

func NewTopology() *Topology {
	return &Topology{
		Executors: make(map[string]*ExecutorInfo),
	}
}

func (self *Topology) RestartExecutor(name string) error {
	self.Lock()
	self.Unlock()
	if _, ok := self.Executors[name]; !ok {
		return fmt.Errorf("executor not found")
	}
	loc := self.Executors[name].Heartbeat.GetLocation()
	grpcConn, err := grpc.Dial(loc.GetURL(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewGueryExecutorClient(grpcConn)
	_, err = client.Restart(context.Background(), &pb.Empty{})
	grpcConn.Close()
	return err
}

func (self *Topology) DuplicateExecutor(name string) error {
	self.Lock()
	self.Unlock()
	if _, ok := self.Executors[name]; !ok {
		return fmt.Errorf("executor not found")
	}
	loc := self.Executors[name].Heartbeat.GetLocation()
	grpcConn, err := grpc.Dial(loc.GetURL(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewGueryExecutorClient(grpcConn)
	_, err = client.Duplicate(context.Background(), &pb.Empty{})
	grpcConn.Close()
	return err
}

func (self *Topology) KillExecutor(name string) error {
	self.Lock()
	defer self.Unlock()
	if _, ok := self.Executors[name]; !ok {
		return fmt.Errorf("executor not found")
	}
	loc := self.Executors[name].Heartbeat.GetLocation()
	grpcConn, err := grpc.Dial(loc.GetURL(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewGueryExecutorClient(grpcConn)
	_, err = client.Quit(context.Background(), &pb.Empty{})
	grpcConn.Close()
	return err
}

func (self *Topology) KillAllExecutors() (err error) {
	self.Lock()
	defer self.Unlock()

	for _, info := range self.Executors {
		loc := info.Heartbeat.GetLocation()
		grpcConn, err := grpc.Dial(loc.GetURL(), grpc.WithInsecure())
		if err != nil {
			continue
		}
		client := pb.NewGueryExecutorClient(grpcConn)
		_, err = client.Quit(context.Background(), &pb.Empty{})
		grpcConn.Close()
	}
	return err
}

func (self *Topology) GetExecutors() []pb.Location {
	self.Lock()
	defer self.Unlock()

	res := []pb.Location{}
	for _, einfo := range self.Executors {
		res = append(res, *einfo.Heartbeat.Location)
	}
	return res
}

func (self *Topology) HasExecutor(name string) bool {
	self.Lock()
	defer self.Unlock()
	_, ok := self.Executors[name]
	return ok
}

func (self *Topology) GetExecutor(name string) *ExecutorInfo {
	self.Lock()
	defer self.Unlock()
	if _, ok := self.Executors[name]; ok {
		return self.Executors[name]
	}
	return nil
}

func (self *Topology) UpdateExecutorInfo(hb *pb.ExecutorHeartbeat) {
	ts := time.Now()

	self.Lock()
	defer self.Unlock()
	exeInfo := NewExecutorInfo(hb)
	if ts.After(exeInfo.LastHeartBeatTime) {
		return
	}
	self.Executors[hb.Location.Name] = exeInfo
	self.ExecutorNumber = int32(len(self.Executors))
}

func (self *Topology) DropExecutorInfo(location *pb.Location) {
	self.Lock()
	defer self.Unlock()

	if _, ok := self.Executors[location.Name]; !ok {
		return
	} else {
		delete(self.Executors, location.Name)
	}
	self.ExecutorNumber = int32(len(self.Executors))
}

///////////////////////////
