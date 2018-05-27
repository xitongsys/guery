package Config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/xitongsys/guery/Connector/FileConnector"
	"github.com/xitongsys/guery/Connector/HiveConnector"
)

type Config struct {
	FileConnectorConfigs FileConnector.FileConnectorConfigs
	HiveConnectorConfigs HiveConnector.HiveConnectorConfigs
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

	FileConnector.Configs = Conf.FileConnectorConfigs
	HiveConnector.Configs = Conf.HiveConnectorConfigs
	return nil
}
