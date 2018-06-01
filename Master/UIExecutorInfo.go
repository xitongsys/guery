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
		Location: e.Heartbeat.Location.GetAddress(),
		Status: func(s int32) string {
			switch s {
			case 0:
				return "IDLE"
			case 1:
				return "BUSY"
			}
			return "UNKNOWN"
		}(e.Heartbeat.Status),
		TaskId: e.Heartbeat.Instruction.TaskId,
	}
	return res
}

func (self *Master) GetUIExecutorInfos() []*UIExecutorInfo {
	res := []*UIExecutorInfo{}
	for _, e := range self.Topology.Executors {
		res = append(res, NewUIExecutorInfoFromExecutorInfo(e))
	}
	return res
}
