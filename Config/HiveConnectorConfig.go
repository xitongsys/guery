package Config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/xitongsys/guery/Util"
)

type HiveConnectorConfig struct {
	Host           string
	User, Password string
}

type HiveConnectorConfigs map[string]*HiveConnectorConfig

func (self HiveConnectorConfig) GetConfig(name string) *HiveConnectorConfig {
	for pattern, config := range self {
		if WildcardMatch(name, pattern) {
			return config
		}
	}
}
