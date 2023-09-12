package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Server struct {
		IP   string
		Port int
	}
	Log struct {
		Directory string
	}
	Blacklist struct {
		Keywords []string
	}
}

var logDir = "logs/"

func LoadConfig(path string) (*Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
