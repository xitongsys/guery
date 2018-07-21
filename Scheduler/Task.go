package Scheduler

import (
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/pb"
)

type Task struct {
	TaskId int64
	Status pb.TaskStatus

	Query   string
	Runtime *Config.ConfigRuntime

	LogicalPlanTree Plan.PlanNode
	EPlanNodes      []EPlan.ENode
	Agents          []pb.Location
	ExecutorNumber  int32

	CommitTime, BeginTime, EndTime time.Time

	AggNode EPlan.ENode
	Output  io.Writer

	DoneChan chan int
	Errs     []error
}

func (self *Task) SetStatus(status pb.TaskStatus) {
	switch status {
	case pb.TaskStatus_TODO:
		self.CommitTime = time.Now()

	case pb.TaskStatus_RUNNING:
		self.BeginTime = time.Now()

	case pb.TaskStatus_SUCCEED:
		self.EndTime = time.Now()

	case pb.TaskStatus_ERROR:
		self.EndTime = time.Now()
	}
	self.Status = status

}

type TaskList []*Task

func (self TaskList) Len() int      { return len(self) }
func (self TaskList) Swap(i, j int) { self[i], self[j] = self[j], self[i] }
func (self TaskList) Less(i, j int) bool {
	if self[i].Runtime.Priority == self[j].Runtime.Priority {
		return self[i].CommitTime.After(self[j].CommitTime)
	}
	return self[i].Runtime.Priority > self[j].Runtime.Priority
}

func (self *TaskList) Top() *Task {
	ln := len(*self)
	if ln <= 0 {
		return nil
	}
	return (*self)[ln-1]
}

func (self *TaskList) Pop() {
	ln := len(*self)
	if ln > 0 {
		*self = (*self)[:ln-1]
	}
}

func (self *TaskList) Delete(task *Task) error {
	i, ln := 0, len(*self)
	for i = 0; i < ln && (*self)[i].TaskId != task.TaskId; i++ {
	}
	if i >= ln {
		return fmt.Errorf("task not in this list")
	}

	for j := i; j < ln-1; j++ {
		(*self)[i] = (*self)[i+1]
	}
	self.Pop()
	return nil
}

func (self *TaskList) Add(task *Task) {
	(*self) = append((*self), task)
	if task.Status == pb.TaskStatus_TODO {
		sort.Sort(*self)
	}
}

func (self *TaskList) HasTask(taskId int64) bool {
	for _, task := range *self {
		if task.TaskId == taskId {
			return true
		}
	}
	return false
}
