package module

type YamlConfig struct {
	SqlConfig        SqlConfig        `yaml:"SqlConfig"`
	FileServerConfig FileServerConfig `yaml:"FileServerConfig"`
}

type SqlConfig struct {
	DbAddress  string `yaml:"DbAddress"`
	DbName     string `yaml:"DbName"`
	DbUser     string `yaml:"DbUser"`
	DbPassword string `yaml:"DbPassword"`
}
type FileServerConfig struct {
	FileServerAddress string `yaml:"FileServerAddress"`
}
