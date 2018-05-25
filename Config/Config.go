package Config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/xitongsys/guery/Util"
)

func WildcardMatch(s, p string) bool {
	ls, lp := len(s), len(p)
	dp := make([][]bool, ls+1)
	for i := 0; i < ls+1; i++ {
		dp[i] = make([]bool, lp+1)
	}
	dp[0][0] = true
	for i := 1; i <= lp; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i <= ls; i++ {
		for j := 1; j <= lp; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else {
				if p[j-1] == s[i-1] || p[j-1] == '?' {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}
	return dp[ls][lp]
}

type Config struct {
	FileConnectorConfigs FileConnectorConfigs
	HiveConnectorConfigs HiveConnectorConfigs
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
