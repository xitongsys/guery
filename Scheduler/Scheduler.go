package Scheduler

import (
	"sync"
)

type Scheduler struct {
	sync.Mutex
	Topology *Topology

	Todos, Doings, Dones TaskList
	Allocated            map[string]int32 //executorName:taskId

	TotalTaskNumber int64
}

func (self *Scheduler) AddTask(query string, priority int32) {
	self.Lock()
	TotalTaskNumber++
	taskId = TotalTaskNumber
	self.Unlock()

}
