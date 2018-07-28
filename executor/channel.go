package executor

import (
	"github.com/xitongsys/guery/util"
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
