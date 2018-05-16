package Util

func GetFuncType(name string) FuncType {
	if name == "MAX" || name == "MIN" ||
		name == "SUM" || name == "AVG" {
		return AGGREGATE
	} else {
		return NORMAL
	}
}
