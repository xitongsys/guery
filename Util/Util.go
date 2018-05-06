package Util

import (
	"fmt"
	"strings"
)

func GetPortFromAddress(address string) int32 {
	ss := strings.Split(address, ":")
	var port int32
	if len(ss) == 2 {
		fmt.Sscanf(ss[1], "%d", &port)
	}
	return port
}

func GetHostFromAddress(address string) string {
	ss := strings.Split(address, ":")
	if len(ss) <= 0 {
		return ""
	}
	return ss[0]

}
