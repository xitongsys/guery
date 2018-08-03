package topology

import (
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

type Heap []*Item

func NewHeap() *Heap {
	return &Heap{}
}

func (self Heap) Len() int               { return len(self) }
func (self Heap) Less(i, j int) bool     { return self[i].ExecutorNumber < self[j].ExecutorNumber }
func (self Heap) Swap(i, j int)          { self[i], self[j] = self[j], self[i] }
func (self *Heap) Push(item interface{}) { *self = append(*self, item.(*Item)) }
func (self *Heap) Pop() interface{} {
	n := len(*self)
	x := (*self)[n-1]
	*self = (*self)[:n-1]
	return x
}

func (self *Heap) GetExecutorLoc() pb.Location {
	item := self.Pop().(*Item)
	item.ExecutorNumber++
	self.Push(item)
	return item.Location
}
