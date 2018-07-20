package Agent

import (
	"fmt"
	"sync"

	"github.com/xitongsys/guery/pb"
)

type TaskMap struct {
	sync.Mutex
	Tasks map[int64]*pb.Task
}

func NewTaskMap() *TaskMap {
	return &TaskMap{
		Tasks: make(map[int64]*pb.Task),
	}
}

func (self *TaskMap) HasTask(id int64) bool {
	self.Lock()
	defer self.Unlock()
	_, ok := self.Tasks[id]
	return ok
}

func (self *TaskMap) PopTask(id int64) *pb.Task {
	self.Lock()
	defer self.Unlock()
	if _, ok := self.Tasks[id]; ok {
		res := self.Tasks[id]
		delete(self.Tasks, task.TaskId)
		return res

	} else {
		return nil
	}
}

func (self *TaskMap) AddTask(task *pb.Task) error {
	self.Lock()
	defer self.Unlock()
	if _, ok := self.Tasks[task.TaskId]; ok {
		return fmt.Errorf("task already exists")
	}
	self.Tasks[task.TaskId] = task
	return nil
}

func (self *TaskMap) DeleteTask(task *pb.Task) error {
	self.Lock()
	defer self.Unlock()
	if _, ok := self.Tasks[task.TaskId]; !ok {
		return nil
	}
	delete(self.Tasks, task.TaskId)
	return nil
}
