package Topology

import (
	"sync"
	"time"

	"github.com/xitongsys/guery/pb"
)

type ExecutorInfo struct {
	Name              string
	Heartbeat         pb.Heartbeat
	LastHeartBeatTime time.Time
}

func NewExecutorInfo(hb *pb.Heartbeat) *ExecutorInfo {
	return &ExecutorInfo{
		Name:              hb.Location.Name,
		Heartbeat:         *hb,
		LastHeartBeatTime: time.Now(),
	}
}

//Topology/////////////////
type Topology struct {
	sync.RWMutex
	Executors        map[string]*ExecutorInfo
	IdleExecutorNum  int32
	TotalExecutorNum int32
}

func NewTopology() *Topology {
	return &Topology{
		Executors: make(map[string]*ExecutorInfo),
	}
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

func (self *Topology) UpdateExecutorInfo(hb *pb.Heartbeat) {
	ts := time.Now()

	self.Lock()
	defer self.Unlock()

	if exeInfo, ok := self.Executors[hb.Location.Name]; !ok {
		exeInfo = NewExecutorInfo(hb)
		self.Executors[hb.Location.Name] = exeInfo
	} else {
		if ts.Before(exeInfo.LastHeartBeatTime) {
			return
		}
	}

	self.IdleExecutorNum, self.TotalExecutorNum = 0, 0
	for _, exeInfo := range self.Executors {
		self.TotalExecutorNum++
		if exeInfo.Heartbeat.Status == 0 {
			self.IdleExecutorNum++
		}
	}

}

func (self *Topology) DropExecutorInfo(location *pb.Location) {
	self.Lock()
	defer self.Unlock()

	dIdleNum, dTotalNum := int32(0), int32(0)
	if exeInfo, ok := self.Executors[location.Name]; !ok {
		return
	} else {
		delete(self.Executors, location.Name)
		dIdleNum, dTotalNum = exeInfo.Heartbeat.Status-1, -1
	}

	self.IdleExecutorNum += dIdleNum
	self.TotalExecutorNum += dTotalNum
}

///////////////////////////
