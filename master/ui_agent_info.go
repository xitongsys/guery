package master

import ()

type UIAgentInfo struct {
	Name              string
	Address           string
	CpuNumber         int
	CpuUsage          []float64
	TotalMemory       int
	FreeMemory        int
	MaxExecutorNumber int
	ExecutorNumber    int
}

func (self *Master) GetUIAgentInfos() []*UIAgentInfo {
	self.Topology.Lock()
	defer self.Topology.Unlock()

	res := []*UIAgentInfo{}
	for name, info := range self.Topology.Agents {
		agent := &UIAgentInfo{
			Name:              name,
			Address:           info.Heartbeat.Location.GetURL(),
			CpuNumber:         int(info.Heartbeat.CpuNumber),
			CpuUsage:          info.Heartbeat.CpuUsage,
			TotalMemory:       int(info.Heartbeat.TotalMemory),
			FreeMemory:        int(info.Heartbeat.FreeMemory),
			MaxExecutorNumber: int(info.Heartbeat.MaxExecutorNumber),
			ExecutorNumber:    int(info.Heartbeat.ExecutorNumber),
		}
		res = append(res, agent)
	}
	return res
}
