package pb

func NewErrLogInfo(info string) *LogInfo {
	return &LogInfo{
		Level: LogLevel_ERR,
		Info:  []byte(info),
	}
}

func NewInfoLogInfo(info string) *LogInfo {
	return &LogInfo{
		Level: LogLevel_INFO,
		Info:  []byte(info),
	}
}

func NewWarnLogInfo(info string) *LogInfo {
	return &LogInfo{
		Level: LogLevel_WARN,
		Info:  []byte(info),
	}
}
