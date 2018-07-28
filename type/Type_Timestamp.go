package Type

import (
	"time"
)

type Timestamp struct {
	Sec int64
}

func (self Timestamp) String() string {
	return time.Unix(self.Sec, 0).Format("2006-01-02 15:04:05")
}
