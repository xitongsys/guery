package config

import ()

type ConfigRuntime struct {
	Catalog                 string
	Schema                  string
	Table                   string
	Priority                int32
	S3Region                string
	ParallelNumber          int32
	MaxConcurrentTaskNumber int32
	QueueSize               int32
}

func NewConfigRuntime() *ConfigRuntime {
	return &ConfigRuntime{
		Catalog:                 "default",
		Schema:                  "default",
		Table:                   "default",
		Priority:                0,
		S3Region:                "",
		ParallelNumber:          4,
		MaxConcurrentTaskNumber: 2,
		QueueSize:               100,
	}
}
