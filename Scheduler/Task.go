package Scheduler

import (
	"time"

	"github.com/xitongsys/guery/Plan"
)

type TaskStatusType int32

const (
	_ TaskStatusType = iota
	TODO
	DOING
	DONE
)

type Task struct {
	TaskId     int64
	TaskStatus TaskStatusType
	Executors  []string
	Query      string
	Priority   int32

	LogicalPlanTree PlanNode
	ExecutorNumber  int32

	BeginTime, EndTime time.Time
}

type TaskList []*Task

func (self TaskList) Len() int           { return len(self) }
func (self TaskList) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }
func (self TaskList) Less(i, j int) bool { return self[i].Priority < self[i].Priority }
