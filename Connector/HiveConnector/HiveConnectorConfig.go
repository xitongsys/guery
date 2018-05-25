package HiveConnector

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/xitongsys/guery/Util"
)

type HiveConnectorConfig struct {
	Host, DB       string
	User, Password string
}

type HiveConnectorConfigs map[string]*HiveConnectorConfig

func (self HiveConnectorConfigs) GetConfig(name string) *HiveConnectorConfig {
	for pattern, config := range self {
		if Util.WildcardMatch(name, pattern) {
			return config
		}
	}
	return nil
}
