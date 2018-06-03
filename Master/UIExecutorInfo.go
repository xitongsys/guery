package Master

import (
	"github.com/xitongsys/guery/Topology"
)

type UIExecutorInfo struct {
	Name     string
	Location string
	Status   string
	TaskId   int64
}

func NewUIExecutorInfoFromExecutorInfo(e *Topology.ExecutorInfo) *UIExecutorInfo {
	res := &UIExecutorInfo{
		Name:     e.Name,
		Location: e.Heartbeat.Location.GetURL(),
		Status: func(s int32) string {
			switch s {
			case 0:
				return "Idle"
			case 1:
				return "Busy"
			}
			return "UNKNOWN"
		}(e.Heartbeat.Status),
	}
	if e.Heartbeat.Instruction != nil {
		res.TaskId = e.Heartbeat.Instruction.TaskId
	}
	return res
}

func (self *Master) GetUIExecutorInfos() []*UIExecutorInfo {
	self.Topology.Lock()
	defer self.Topology.Unlock()
	res := []*UIExecutorInfo{}
	for _, e := range self.Topology.Executors {
		res = append(res, NewUIExecutorInfoFromExecutorInfo(e))
	}
	return res
}
