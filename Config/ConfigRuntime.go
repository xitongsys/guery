package Config

import ()

type ConfigRuntime struct {
	MaxConcurrentNumber int
	Catalog             string
	Schema              string
	Priority            int32
}

func NewConfigRuntime() *ConfigRuntime {
	return &ConfigRuntime{
		MaxConcurrentNumber: 10,
		Catalog:             "default",
		Schema:              "default",
		Priority:            0,
	}
}
