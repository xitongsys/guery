package pb

import (
	"fmt"
)

func (self *Location) GetAddress() string {
	return fmt.Sprintf("%v:%v", self.Address, self.Port)
}
