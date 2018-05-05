package Master

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

func (self *Topology) UpdateExecutorInfo(hb *pb.Heartbeat) {
	ts := time.Now()

	self.Lock()
	defer self.Unlock()

	dIdleNum, dTotalNum := int32(0), int32(0)
	if exeInfo, ok := self.Executors[hb.Location.Name]; !ok {
		exeInfo = NewExecutorInfo(hb)
		dIdleNum, dTotalNum = 1, 1
		self.Executors[hb.Location.Name] = exeInfo
	} else {
		if ts.Before(exeInfo.LastHeartBeatTime) {
			return
		}
		dIdleNum, dTotalNum = exeInfo.Heartbeat.Resource-hb.Resource, 0
	}

	self.IdleExecutorNum += dIdleNum
	self.TotalExecutorNum += dTotalNum

}

func (self *Topology) DropExecutorInfo(location *pb.Location) {
	self.Lock()
	defer self.Unlock()

	dIdleNum, dTotalNum := int32(0), int32(0)
	if exeInfo, ok := self.Executors[location.Name]; !ok {
		return
	} else {
		dIdleNum, dTotalNum = exeInfo.Heartbeat.Resource-1, -1
	}

	self.IdleExecutorNum += dIdleNum
	self.TotalExecutorNum += dTotalNum
}

///////////////////////////
