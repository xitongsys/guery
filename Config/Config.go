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

type HiveConnectorConfig struct {
	Host           string
	User, Password string
}

type Config struct {
	FileConnectorConfigs map[string]*FileConnectorConfig
	HiveConnectorConfigs map[string]*HiveConnectorConfig
}

var Conf Config

func LoadConfig(fileName string) error {
	var data []byte
	var err error
	if data, err = ioutil.ReadFile(fileName); err != nil {
		log.Fatalf("Fail to load the configure file, due to %v ", err.Error())
		return err
	}

	if err = json.Unmarshal(data, &Conf); err != nil {
		log.Fatalf("Fail to load the configure file, due to %v", err.Error())
		return err
	}
	return nil
}
