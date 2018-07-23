package Agent

import (
	"fmt"
	"sync"

	"github.com/xitongsys/guery/pb"
)

////////////////
type Task struct {
	TaskId       int64
	Instructions []*pb.Instruction
	Status       pb.TaskStatus
	LogInfos     []*pb.LogInfo

	ExecutorInfos map[string]*pb.ExecutorHeartbeat
}

func NewTask(pbtask *pb.Task) *Task {
	task := &Task{
		TaskId:       pbtask.TaskId,
		Status:       pb.TaskStatus_RUNNING,
		Instructions: pbtask.Instructions,

		ExecutorInfos: map[string]*pb.ExecutorHeartbeat{},
	}
	return task
}

////////////////

type TaskMap struct {
	sync.Mutex
	Tasks map[int64]*Task
}

func NewTaskMap() *TaskMap {
	return &TaskMap{
		Tasks: make(map[int64]*Task),
	}
}

func (self *TaskMap) GetTaskInfos() []*pb.TaskInfo {
	self.Lock()
	defer self.Unlock()
	res := []*pb.TaskInfo{}
	for taskId, task := range self.Tasks {
		info := &pb.TaskInfo{
			TaskId: taskId,
			Infos:  task.LogInfos,
		}
		totalNum := len(task.Instructions)
		doneNum := 0
		for _, e := range task.ExecutorInfos {
			if e.Status != pb.TaskStatus_RUNNING {
				doneNum++
			}
		}
		info.Progress = float64(doneNum) / float64(totalNum)
		res = append(res, info)
	}
	return res
}

func (self *TaskMap) UpdateTaskInfo(hb *pb.ExecutorHeartbeat) {
	if hb == nil {
		return
	}
	self.Lock()
	defer self.Unlock()
	taskId, executorName := hb.TaskId, hb.Location.Name
	if task, ok := self.Tasks[id]; ok {
		task.ExecutorInfos[executorName] = hb
		if hb.Status == pb.TaskStatus_ERROR {
			task.Status = hb.Status
		}
		task.LogInfos = append(task.LogInfos, hb.Infos...)
	}
}

func (self *TaskMap) HasTask(id int64) bool {
	self.Lock()
	defer self.Unlock()
	_, ok := self.Tasks[id]
	return ok
}

func (self *TaskMap) GetTask(id int64) *Task {
	self.Lock()
	defer self.Unlock()
	if _, ok := self.Tasks[id]; ok {
		res := self.Tasks[id]
		return res
	} else {
		return nil
	}
}

func (self *TaskMap) GetTaskNumber() int32 {
	self.Lock()
	defer self.Unlock()
	return int32(len(self.Tasks))
}

func (self *TaskMap) PopTask(id int64) *Task {
	self.Lock()
	defer self.Unlock()
	if task, ok := self.Tasks[id]; ok {
		delete(self.Tasks, id)
		return task

	} else {
		return nil
	}
}

func (self *TaskMap) AddTask(task *Task) error {
	self.Lock()
	defer self.Unlock()
	if _, ok := self.Tasks[task.TaskId]; ok {
		return fmt.Errorf("task already exists")
	}
	self.Tasks[task.TaskId] = task
	return nil
}

func (self *TaskMap) DeleteTask(task *Task) error {
	self.Lock()
	defer self.Unlock()
	if _, ok := self.Tasks[task.TaskId]; !ok {
		return nil
	}
	delete(self.Tasks, task.TaskId)
	return nil
}
