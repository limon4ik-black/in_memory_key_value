package config

import (
	"os"

	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Engine struct {
		Type string `yaml:"type"`
	} `yaml:"engine"`

	Network struct {
		Address          string `yaml:"address"`
		Max_connections  int    `yaml:"max_connections"`
		Max_message_size string `yaml:"max_message_size"`
		Idle_timeout     string `yaml:"idle_timeout"`
	} `yaml:"network"`
	Logging struct {
		Level  string `yaml:"level"`
		Output string `yaml:"output"`
	} `yaml:"logging"`
}

var AppConfig *Config

func LoadConfig(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		logger.Log.Errorw("failed to read config file")
		os.Exit(1)
	}
	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		logger.Log.Errorw("parsing error of config file")
		os.Exit(1)
	}

	AppConfig = &cfg
}
