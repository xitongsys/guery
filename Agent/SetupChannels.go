package Agent

import (
	"context"
	"fmt"
	"time"

	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

func (self *Agent) SendInstruction(inst *pb.Instruction) error {
	var err error
	name := inst.Location.Name
	empty := &pb.Empty{}
	executor := self.Topology.GetExecutor(name)
	if executor == nil {
		return fmt.Errorf("executor not found")
	}
	loc := executor.Heartbeat.Location
	grpcConn, err := grpc.Dial(loc.GetURL(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer grpcConn.Close()

	client := pb.NewGueryExecutorClient(grpcConn)
	if _, err = client.SendInstruction(context.Background(), inst); err != nil {
		return err
	}

	if _, err = client.SetupWriters(context.Background(), empty); err != nil {
		return err
	}
	return nil
}

func (self *Agent) SendTask(ctx context.Context, task *pb.Task) (*pb.Empty, error) {
	res := &pb.Empty{}
	var err error
	if task == nil {
		return res, nil
	}
	if task.GetInstruction() == nil {
		return res, fmt.Errorf("[Agent] instruction is nil")
	}

	if err = self.Tasks.AddTask(task); err != nil {
		return res, err
	}
	for _, inst := range task.GetInstruction() {
		if err = self.LanchExecutor(inst.Location.Name); err != nil {
			return res, err
		}
	}

	flag := false
	timeout := time.After(TIMEOUT * time.Millisecond)
	tick := time.Tick(50 * time.Millisecond)
	for !flag {
		flag = true
		select {
		case <-timeout:
			err = fmt.Errorf("timeout")
			break
		case <-tick:
			for _, inst := range task.GetInstruction() {
				name := inst.Location.Name
				if !self.Topology.HasExecutor(name) {
					flag = false
					break
				}
			}
		}
	}
	if err != nil {
		self.KillTask(context.Background(), task)
	}

	for _, inst := range task.GetInstruction() {
		if err = self.SendInstruction(inst); err != nil {
			break
		}
	}

	if err != nil {
		self.KillTask(context.Background(), task)
	}

	return res, err
}

func (self *Agent) Run(ctx context.Context, task *pb.Task) (*pb.Empty, error) {
	var err error
	empty := &pb.Empty{}

	if task = self.Tasks.GetTask(task.TaskId); task == nil {
		return empty, fmt.Errorf("task not found")
	}

	for _, inst := range task.Instruction {
		loc := inst.Location
		grpcConn, err := grpc.Dial(loc.GetURL(), grpc.WithInsecure())
		if err != nil {
			return empty, err
		}
		client := pb.NewGueryExecutorClient(grpcConn)
		if _, err = client.SetupReaders(context.Background(), empty); err != nil {
			grpcConn.Close()
			break
		}

		if _, err = client.Run(context.Background(), empty); err != nil {
			grpcConn.Close()
			break
		}
		grpcConn.Close()
	}
	return empty, err
}

func (self *Agent) GetOutputChannelLocation(ctx context.Context, location *pb.Location) (*pb.Location, error) {
	var err error
	var res *pb.Location
	name := location.Name
	executor := self.Topology.GetExecutor(name)
	if executor == nil {
		return nil, fmt.Errorf("executor not found")
	}
	loc := executor.Heartbeat.Location
	grpcConn, err := grpc.Dial(loc.GetURL(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	client := pb.NewGueryExecutorClient(grpcConn)
	return client.GetOutputChannelLocation(ctx, location)
}
