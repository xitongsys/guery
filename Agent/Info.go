package Agent

import (
	"log"
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
		TaskInfos:         self.GetTaskInfos(),
	}

	return res
}

func (self *Agent) GetTaskInfos() []*pb.TaskInfo {
	res := make([]*pb.TaskInfo, 0)
	for _, task := range self.Tasks.Tasks { //should copy?
		totNum := float64(len(task.Instructions))
		doneNum := float64(0)
		for _, inst := range task.Instructions {
			name := inst.Location.Name
			if self.Topology.HasExecutor(name) {
				s := self.Topology.GetExecutorStatus(name)
				if s != pb.TaskStatus_RUNNING && s != pb.TaskStatus_TODO {
					doneNum += float64(1.0)
				}
			}
		}
		task.Info.Progress = doneNum / totNum
		log.Println("======", doneNum, totNum)
		res = append(res, task.Info)
	}
	return res
}
