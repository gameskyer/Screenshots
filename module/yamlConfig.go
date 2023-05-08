package module

type YamlConfig struct {
	SqlConfig        SqlConfig
	FileServerConfig FileServerConfig
}

type SqlConfig struct {
	DbAddress  string
	DbName     string
	DbUser     string
	DbPassword string
}
type FileServerConfig struct {
	FileServerAddress string
}
