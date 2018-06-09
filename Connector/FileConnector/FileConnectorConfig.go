package FileConnector

import (
	"github.com/xitongsys/guery/Util"
)

var (
	Configs FileConnectorConfigs
)

type FileConnectorConfig struct {
	Catalog  string
	Schema   string
	Table    string
	FileType string
	FileMD   Util.JsonMetadata
	PathList []string
}
type FileConnectorConfigs map[string]*FileConnectorConfig

func (self FileConnectorConfigs) GetConfig(name string) *FileConnectorConfig {
	for pattern, config := range self {
		if Util.WildcardMatch(name, pattern) {
			return config
		}
	}
	return nil
}
