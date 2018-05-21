package Config

type Config struct {
	Connectors map[string]ConfigConnector
}

type ConfigConnector struct {
	Name string
	KV   map[string]interface{}
}
