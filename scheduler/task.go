package scheduler

import (
	"fmt"
	"io"
	"time"

	"github.com/satori/go.uuid"
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/eplan"
	"github.com/xitongsys/guery/optimizer"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/plan"
)

type Task struct {
	TaskId   string
	Status   pb.TaskStatus
	Progress float64
	Infos    []*pb.LogInfo

	Query   string
	Runtime *config.ConfigRuntime

	LogicalPlanTree plan.PlanNode
	EPlanNodes      []eplan.ENode
	Agents          []pb.Location

	CommitTime, BeginTime, EndTime time.Time

	AggNode eplan.ENode
	Output  io.Writer

	DoneChan chan int
}

func NewTask(runtime *config.ConfigRuntime, query string, output io.Writer) (res *Task, err error) {
	t := time.Now()
	res = &Task{
		TaskId:     fmt.Sprintf("%v_%v", t.Format("20060102150405"), uuid.Must(uuid.NewV4()).String()),
		Status:     pb.TaskStatus_TODO,
		Query:      query,
		Runtime:    runtime,
		CommitTime: t,
		Output:     output,
		DoneChan:   make(chan int),
	}

	if res.LogicalPlanTree, err = optimizer.CreateLogicalTree(runtime, query); err != nil {
		res.Status = pb.TaskStatus_ERROR
		res.Infos = []*pb.LogInfo{pb.NewErrLogInfo(fmt.Sprintf("%v", err))}
	}
	return
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
