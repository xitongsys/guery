package Scheduler

import (
	"sync"

	"github.com/xitongsys/guery/Plan"
)

type Scheduler struct {
	sync.Mutex
	Topology *Topology

	Todos, Doings, Dones, Fails TaskList
	Allocated                   map[string]int32 //executorName:taskId

	TotalTaskNumber int64
}

func (self *Scheduler) AddTask(query string, priority int32) error {
	var err error
	self.Lock()
	TotalTaskNumber++
	taskId = TotalTaskNumber

	task := &Task{
		TaskId:     taskId,
		TaskStatus: TODO,
		Executors:  []string{},
		Query:      query,
		Priority:   priority,

		LogicalPlanTree: Plan.CreateLogicalTree(query),
	}
	self.Unlock()

}
