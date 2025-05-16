package globals

// Config 全局配置结构体
type Config struct {
	App AppConfig `yaml:"app"`
	DB  struct {
		Type   string      `yaml:"type"`
		Config MysqlConfig `yaml:"config"`
	}
	Logger LoggerConfig `yaml:"logger"`
}

// AppConfig 应用配置
type AppConfig struct {
	Name    string `yaml:"name"`
	Mode    string `yaml:"mode"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Version string `yaml:"version"`
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	MaxCon   int    `yaml:"maxCon"`
	MaxIdle  int    `yaml:"maxIdle"`
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	Compress   bool   `yaml:"compress"`
	Console    bool   `yaml:"console"`
	MaxSize    int    `yaml:"maxSize"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxAge     int    `yaml:"maxAge"`
}
