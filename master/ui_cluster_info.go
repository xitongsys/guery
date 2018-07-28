package master

import ()

type UIClusterInfo struct {
	Running, Queued, Finished int
	Agent, Busy, Total        int
	StartTime                 string
}

func (self *Master) GetUIClusterInfo() *UIClusterInfo {
	info := &UIClusterInfo{
		Running:  len(self.Scheduler.Doings),
		Queued:   len(self.Scheduler.Todos),
		Finished: len(self.Scheduler.Dones) + len(self.Scheduler.Fails),

		StartTime: self.StartTime.Format("2006-01-02 15:04:05"),
	}
	return info
}
