package util

import (
	"github.com/satori/go.uuid"
	"github.com/xitongsys/guery/pb"
)

type Item struct {
	Location       pb.Location
	ExecutorNumber int
}

func NewItem(loc pb.Location, num int) *Item {
	return &Item{
		Location:       loc,
		ExecutorNumber: num,
	}
}

type Heap struct {
	Items    []*Item
	AgentMap map[string]pb.Location
}

func NewHeap() *Heap {
	return &Heap{
		Items:    []*Item{},
		AgentMap: map[string]pb.Location{},
	}
}

func (self Heap) Len() int { return len(self.Items) }
func (self Heap) Less(i, j int) bool {
	return self.Items[i].ExecutorNumber < self.Items[j].ExecutorNumber
}
func (self Heap) Swap(i, j int)          { self.Items[i], self.Items[j] = self.Items[j], self.Items[i] }
func (self *Heap) Push(item interface{}) { self.Items = append(self.Items, item.(*Item)) }
func (self *Heap) Pop() interface{} {
	n := len(self.Items)
	x := self.Items[n-1]
	self.Items = self.Items[:n-1]
	return x
}

func (self *Heap) GetExecutorLoc() pb.Location {
	item := self.Pop().(*Item)
	item.ExecutorNumber++
	self.Push(item)
	exe := pb.Location{
		Name:    "executor_" + uuid.Must(uuid.NewV4()).String(),
		Address: item.Location.Address,
		Port:    item.Location.Port,
	}
	self.AgentMap[item.Location.Name] = item.Location
	return exe
}

func (self *Heap) GetAgents() []pb.Location {
	res := []pb.Location{}
	for _, loc := range self.AgentMap {
		res = append(res, loc)
	}
	return res
}
