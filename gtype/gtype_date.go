package gtype

import (
	"time"
)

type Date struct {
	Sec int64
}

func (self Date) String() string {
	return time.Unix(self.Sec, 0).Format("2006-01-02")
}
