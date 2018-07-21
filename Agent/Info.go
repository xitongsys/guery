package Agent

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Agent) GetInfo() *pb.AgentHeartbeat {
	m, _ := mem.VirtualMemory()
	cpuUsage, _ := cpu.Percent(time.Duration(0), true)

	res := &pb.AgentHeartbeat{
		Location: &pb.Location{
			Name:    self.Name,
			Address: Util.GetHostFromAddress(self.Address),
			Port:    Util.GetPortFromAddress(self.Address),
		},
		TotalMemory:       int64(m.Total),
		FreeMemory:        int64(m.Free),
		CpuNumber:         int32(len(cpuUsage)),
		CpuUsage:          cpuUsage,
		ExecutorNumber:    self.Topology.ExecutorNumber,
		MaxExecutorNumber: self.MaxExecutorNumber,
		RunningTaskNumber: self.Tasks.GetTaskNumber(),
		TaskInfos:         self.Tasks.GetTaskInfos(),
	}
	return res
}
