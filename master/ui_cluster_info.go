package master

import ()

type UIClusterInfo struct {
	Running, Queued, Finished int
	Agent, Busy, Total        int
	StartTime                 string
}

func (self *Master) GetUIClusterInfo() *UIClusterInfo {
	info := &UIClusterInfo{
		Running:  self.Scheduler.RunningQueue.Tasks.Len(),
		Queued:   self.Scheduler.TodoQueue.Tasks.Len(),
		Finished: self.Scheduler.SucceedQueue.Tasks.Len() + self.Scheduler.ErrorQueue.Tasks.Len(),

		StartTime: self.StartTime.Format("2006-01-02 15:04:05"),
	}
	return info
}
