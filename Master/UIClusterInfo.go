package Master

import ()

type UIClusterInfo struct {
	Running, Queued, Finished int
	Active, Busy, Free        int
	StartTime                 string
}

func (self *Master) GetUIClusterInfo() *UIClusterInfo {
	info := &UIClusterInfo{
		Running:  len(self.Scheduler.Doings),
		Queued:   len(self.Scheduler.Todos),
		Finished: len(self.Scheduler.Dones) + len(self.Scheduler.Fails),

		Active: int(self.Scheduler.Topology.TotalExecutorNum),
		Busy:   int(self.Scheduler.Topology.TotalExecutorNum - self.Scheduler.Topology.IdleExecutorNum),
		Free:   int(self.Scheduler.Topology.IdleExecutorNum),

		StartTime: self.StartTime.Format("2006-01-02 15:04:05"),
	}
	return info
}
