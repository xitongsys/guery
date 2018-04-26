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
		Name:              hb.Location.ExecutorName,
		Heartbeat:         *hb,
		LastHeartBeatTime: time.Now(),
	}

}

//////////////////////////////

//rack////////////////////////
type Rack struct {
	sync.RWMutex
	Name             string
	Executors        map[string]*ExecutorInfo
	IdleExecutorNum  int32
	TotalExecutorNum int32
}

func NewRack(name string) *Rack {
	return &Rack{
		Name:      name,
		Executors: make(map[string]*Rack),
	}
}

func (self *Rack) GetExecutor(name string) *ExecutorInfo {
	self.RLock()
	defer self.RUnlock()
	return self.Executors[name]
}

func (self *Rack) AddExecutor(e *ExecutorInfo) {
	self.Lock()
	defer self.Unlock()
	self.Executors[e.Name] = e
}

func (self *Rack) DropExecutor(name string) {
	self.Lock()
	defer self.Unlock()
	delete(self.Executor, name)
}

func (self *Rack) AddResource(dIdleNum int32, dTotalNum int32) {
	self.Lock()
	defer self.Unlock()
	self.IdleExecutorNum += dIdleNum
	self.TotalExecutorNum += dTotalNum
}

//////////////////////////////

//DataCenter/////////////////
type DataCenter struct {
	sync.RWMutex
	Name             string
	Racks            map[string]*Rack
	IdleExecutorNum  int32
	TotalExecutorNum int32
}

func NewDataCenter(name string) *DataCenter {
	return &DataCenter{
		Name:  name,
		Racks: make(map[string]*Rack),
	}
}

func (self *DataCenter) GetRack(name string) *Rack {
	self.RLock()
	defer self.RUnlock()
	return self.Racks[name]
}

func (self *DataCenter) AddRack(rack *Rack) {
	self.Lock()
	defer self.Unlock()
	self.Racks[rack.Name] = rack
}

func (self *DataCenter) AddResource(dIdleNum int32, dTotalNum int32) {
	self.Lock()
	defer self.Unlock()
	self.IdleExecutorNum += dIdleNum
	self.dTotalNum += dTotalNum
}

///////////////////////////

//Topology/////////////////
type Topology struct {
	sync.RWMutex
	DataCenters      map[string]*DataCenter
	IdleExecutorNum  int32
	TotalExecutorNum int32
}

func NewTopology() *Topology {
	return &Topology{
		DataCenters: make(map[string]*DataCenter),
	}
}

func (self *Topology) GetDataCenter(name string) *DataCenter {
	self.RLock()
	defer self.RUnlock()
	return self.DataCenters[name]
}

func (self *Topology) AddDataCenter(dc *DataCenter) {
	self.Lock()
	defer self.Unlock()
	self.DataCenters[dc.Name] = dc
}

func (self *Topology) AddResource(dIdleNum int32, dTotalNum int32) {
	self.Lock()
	defer self.Unlock()
	self.IdleExecutorNum += dIdleNum
	self.TotalExecutorNum += dTotalNum
}

func (self *Topology) UpdateExecutorInfo(hb *pb.Heartbeat) {
	dc := self.GetDataCenter(hb.Location.DataCenter)
	if dc == nil {
		dc = NewDataCenter(hb.Location.DataCenter)
		hb.AddDataCenter(dc)
	}

	rack := dc.GetRack(hb.Location.Rack)
	if rack == nil {
		rack = NewRack(hb.Location.Rack)
	}

	dIdleNum, dTotalNum := 0, 0
	if exeInfo := rack.GetExecutor(hb.Location.ExecutorName); exeInfo == nil {
		exeInfo = NewExecutorInfo(hb)
		dIdleNum, dTotalNum = 1, 1
		rack.AddExecutor(exeInfo)

	} else {
		dIdleNum, dTotalNum := exeInfo.Heartbeat.Resource-hb.Resource, 0
	}

	self.Lock()
	defer self.Unlock()

	rack.AddExecutor(dIdleNum, dTotalNum)
	dc.AddResource(dIdleNum, dTotalNum)
	self.AddResource(dIdleNum, dTotalNum)
}

func (self *Topology) DropExecutorInfo(location *pb.Location) {
	dc := self.GetDataCenter(location.DataCenter)
	if dc == nil {
		return
	}

	rack := dc.GetRack(hb.Location.Rack)
	if rack == nil {
		return
	}

	dIdleNum, dTotalNum := 0, 0
	if exeInfo := rack.GetExecutor(hb.Location.ExecutorName); exeInfo == nil {
		return

	} else {
		dIdleNum, dTotalNum := exeInfo.Heartbeat.Resource-1, -1
	}

	self.Lock()
	defer self.Unlock()

	rack.AddExecutor(dIdleNum, dTotalNum)
	dc.AddResource(dIdleNum, dTotalNum)
	self.AddResource(dIdleNum, dTotalNum)
}

///////////////////////////
