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

type Rack struct {
	sync.RWMutex
	Name             string
	Executors        map[string]*ExecutorInfo
	IdleExecutorNum  int32
	TotalExecutorNum int32
}

type DataCenter struct {
	sync.RWMutex
	Name             string
	Racks            map[string]*Rack
	IdleExecutorNum  int32
	TotalExecutorNum int32
}

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

func NewDataCenter(name string) *DataCenter {
	return &DataCenter{
		Name:  name,
		Racks: make(map[string]*Rack),
	}
}

func NewRack(name string) *Rack {
	return &Rack{
		Name:      name,
		Executors: make(map[string]*Rack),
	}
}

func (self *Topology) GetDataCenter(name string) *DataCenter {
	self.RLock()
	defer self.RUnlock()
	return self.DataCenters[name]
}

func (self *DataCenter) GetRack(name string) *Rack {
	self.RLock()
	defer self.RUnlock()
	return self.Racks[name]
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
