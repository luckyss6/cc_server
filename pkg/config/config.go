package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var Cfg = new(Config)

type Config struct {
	MysqlConfig *Mysql `yaml:"mysql"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Debug    bool   `yaml:"debug"`
}

func InitConfig() {
	var configPath string
	configPath = os.Getenv("CC_SERVER_CONFIG")
	if len(configPath) == 0 {
		configPath = "/config/config.yaml"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Sprintf("read config file error: %s", err.Error()))
	}

	err = yaml.Unmarshal(data, Cfg)
	if err != nil {
		panic(fmt.Sprintf("unmarshal config error: %s", err.Error()))
	}
}
