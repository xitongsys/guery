package Config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/xitongsys/guery/Util"
)

type FileConnectorConfig struct {
	Catalog      string
	Schema       string
	Table        string
	FileType     string
	FileMD       Util.JsonMetadata
	FilePathList []string
}
type FileConnectorConfigs map[string]*FileConnectorConfig

func (self FileConnectorConfig) GetConfig(name string) *FileConnectorConfig {
	for pattern, config := range self {
		if WildcardMatch(name, pattern) {
			return config
		}
	}
}
