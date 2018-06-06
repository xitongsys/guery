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
