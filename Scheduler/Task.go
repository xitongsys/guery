package Scheduler

import (
	"io"
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
}

type TaskList []*Task

func (self TaskList) Len() int           { return len(self) }
func (self TaskList) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }
func (self TaskList) Less(i, j int) bool { return self[i].Priority > self[i].Priority }

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
