package Master

import (
	"fmt"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Scheduler"
)

type UITaskInfo struct {
	TaskId     int64
	Status     string
	Query      string
	PlanTree   *EPlan.OutputNode
	Priority   int32
	CommitTime string
	ErrInfo    string
}

func NewUITaskInfoFromTask(task *Scheduler.Task) *UITaskInfo {
	res := &UITaskInfo{
		TaskId:     task.TaskId,
		Status:     task.Status.String(),
		Query:      task.Query,
		PlanTree:   EPlan.EPlanOutput(task.EPlanNodes),
		Priority:   task.Priority,
		CommitTime: task.CommitTime.Format("2006-01-02 15:04:05"),
		ErrInfo:    fmt.Sprintf("%v", task.Errs),
	}
	return res
}

func (self *Master) GetUITaskInfos() map[string][]*UITaskInfo {
	res := make(map[string][]*UITaskInfo)

	res["TODO"] = []*UITaskInfo{}
	for _, t := range self.Scheduler.Todos {
		res["TODO"] = append(res["TODO"], NewUITaskInfoFromTask(t))
	}

	res["DOING"] = []*UITaskInfo{}
	for _, t := range self.Scheduler.Doings {
		res["DOING"] = append(res["DOING"], NewUITaskInfoFromTask(t))
	}

	res["DONE"] = []*UITaskInfo{}
	for _, t := range self.Scheduler.Dones {
		res["DONE"] = append(res["DONE"], NewUITaskInfoFromTask(t))
	}

	res["FAILED"] = []*UITaskInfo{}
	for _, t := range self.Scheduler.Fails {
		res["FAILED"] = append(res["FAILED"], NewUITaskInfoFromTask(t))
	}

	return res
}
