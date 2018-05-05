package pb

import (
	"encoding/base64"
)

func (self *Instruction) Base64Encode() {
	self.EncodedEPlanNodeBytes = base64.StdEncoding.EncodeToString([]byte(self.EncodedEPlanNodeBytes))
}

func (self *Instruction) Base64Decode() error {
	bs, err := base64.StdEncoding.DecodeString(self.EncodedEPlanNodeBytes)
	self.EncodedEPlanNodeBytes = string(bs)
	return err
}
