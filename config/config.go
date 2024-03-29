package config

type ServerConfig struct {
	Name             string           `mapstructure:"name" json:"name"`
	Port             int              `mapstructure:"port" json:"port"`
	Tags             []string         `mapstructure:"tags" json:"tags"`
	DBConfig         DBConfig         `mapstructure:"db" json:"db"`
	UserServerConfig UserServerConfig `mapstructure:"user-server" json:"user-server"`
	PayServerConfig  PayServerConfig  `mapstructure:"pay-server" json:"pay-server"`
}

type DBConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Database string `mapstructure:"database" json:"database"`
}

type UserServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type PayServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
