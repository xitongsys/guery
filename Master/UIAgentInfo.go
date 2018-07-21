package Master

import (
	"fmt"
)

type UIAgentInfo struct {
	Name              string
	Address           string
	CpuNumber         int32
	CpuUsage          string
	TotalMemory       int64
	FreeMemory        int64
	MaxExecutorNumber int32
	ExecutorNumber    int32
}

func (self *Master) GetUIAgentInfos() []*UIAgentInfo {
	self.Topology.Lock()
	defer self.Topology.Unlock()

	res := []*UIAgentInfo{}
	for name, info := range self.Topology.Agents {
		agent := &UIAgentInfo{
			Name:              name,
			Address:           info.Heartbeat.Location.GetURL(),
			CpuNumber:         info.Heartbeat.CpuNumber,
			CpuUsage:          fmt.Sprintf("%v", info.Heartbeat.CpuUsage),
			TotalMemory:       info.Heartbeat.TotalMemory,
			FreeMemory:        info.Heartbeat.FreeMemory,
			MaxExecutorNumber: info.Heartbeat.MaxExecutorNumber,
			ExecutorNumber:    info.Heartbeat.ExecutorNumber,
		}
		res = append(res, agent)
	}
	return res
}
