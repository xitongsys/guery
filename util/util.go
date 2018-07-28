package util

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

func DuplicateName(name string) string {
	nbs := []byte(name)
	ln := len(nbs)
	i := 0
	for i = ln - 1; i >= 0; i-- {
		if nbs[i] < '0' || nbs[i] > '9' {
			break
		}
	}
	if i >= 0 && nbs[i] == '_' {
		num := 0
		fmt.Sscanf(name[i+1:], "%d", &num)
		if num < 0 {
			return name + "_1"
		}
		return fmt.Sprintf("%v_%v", name[:i], num+1)
	} else {
		return name + "_1"
	}
}
