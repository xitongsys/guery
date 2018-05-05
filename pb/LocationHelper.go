package pb

import (
	"fmt"
)

func (self *Location) GetURL() string {
	return fmt.Sprintf("%v:%v", self.Address, self.Port)
}
