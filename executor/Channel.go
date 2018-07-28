package Executor

import (
	"github.com/xitongsys/guery/Util"
)

type Channel struct {
	Name     string
	Port     int32
	IsClosed bool
	Piper    *Util.Piper
}

func NewChannel() *Channel {
	return nil
}
