package config

type ConfigStruct struct {
	Input struct {
		ConnectionString string            `yaml:"connectionString"`
		RequestConfig    map[string]string `yaml:"requestConfig"`
	} `yaml:"input"`
	Repository map[string]string `yaml:"repository"`
	Output     map[string]string `yaml:"output"`
}

func NewConfigStruct() *ConfigStruct {
	return &ConfigStruct{}
}
