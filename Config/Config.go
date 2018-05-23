package Config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type FileConnectorConfig struct {
	Catalog      string
	Schema       string
	Table        string
	FileType     string
	FileMD       string
	FilePathList []string
}

type Config struct {
	FileConnectorConfigs map[string]*FileConnectorConfig
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
