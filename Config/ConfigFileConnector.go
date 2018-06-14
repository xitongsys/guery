package Config

type FileConnectorConfig struct {
	Catalog     string
	Schema      string
	Table       string
	FileType    string
	ColumnNames []string
	ColumnTypes []string
	PathList    []string
}
type FileConnectorConfigs map[string]*FileConnectorConfig

func (self FileConnectorConfigs) GetConfig(name string) *FileConnectorConfig {
	for pattern, config := range self {
		if WildcardMatch(name, pattern) {
			return config
		}
	}
	return nil
}
