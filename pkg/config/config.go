package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var Cfg *Config

type Config struct {
	MysqlConfig *Mysql `yaml:"mysql"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func InitConfig() {
	var configPath string
	configPath = os.Getenv("CC_SERVER_CONFIG")
	if len(configPath) == 0 {
		configPath = "/config/config.yaml"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	Cfg = new(Config)
	err = yaml.Unmarshal(data, Cfg)
	if err != nil {
		panic(err)
	}
}
