package Scheduler

import (
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Plan"
)

type TaskStatusType int32

const (
	_ TaskStatusType = iota
	TODO
	DOING
	DONE
	FAILED
)

func (self TaskStatusType) String() string {
	switch self {
	case TODO:
		return "TODO"
	case DOING:
		return "DOING"
	case DONE:
		return "DONE"
	case FAILED:
		return "FAILED"
	}
	return "UNKNOWN"
}

type Task struct {
	TaskId int64
	Status TaskStatusType

	Executors []string
	Query     string
	Catalog   string
	Schema    string
	Priority  int32

	LogicalPlanTree Plan.PlanNode
	ExecutorNumber  int32

	CommitTime, BeginTime, EndTime time.Time

	AggNode EPlan.ENode
	Output  io.Writer

	DoneChan chan int
	Errs     []error
}

func (self *Task) SetStatus(status TaskStatusType) {
	switch status {
	case TODO:
		self.Status = TODO
		self.CommitTime = time.Now()

	case DOING:
		self.Status = DOING
		self.BeginTime = time.Now()

	case DONE:
		self.Status = DONE
		self.EndTime = time.Now()

	case FAILED:
		self.Status = FAILED
		self.EndTime = time.Now()
	}

}

type TaskList []*Task

func (self TaskList) Len() int      { return len(self) }
func (self TaskList) Swap(i, j int) { self[i], self[j] = self[j], self[i] }
func (self TaskList) Less(i, j int) bool {
	if self[i].Priority == self[j].Priority {
		return self[i].CommitTime.After(self[j].CommitTime)
	}
	return self[i].Priority > self[j].Priority
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
	if task.Status == TODO {
		sort.Sort(*self)
	}
}
