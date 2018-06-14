package Config

import (
	"fmt"
)

type HiveConnectorConfig struct {
	Host, DB       string
	User, Password string
}

func (self *HiveConnectorConfig) GetURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", self.User, self.Password, self.Host, self.DB)
}

type HiveConnectorConfigs map[string]*HiveConnectorConfig

func (self HiveConnectorConfigs) GetConfig(name string) *HiveConnectorConfig {
	for pattern, config := range self {
		if WildcardMatch(name, pattern) {
			return config
		}
	}
	return nil
}
