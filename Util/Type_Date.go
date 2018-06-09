package Util

import (
	"time"
)

type Date time.Time

func (self Date) String() string {
	return time.Time(self).Format("2006-01-02")
}
