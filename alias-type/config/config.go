package config

import (
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type ApplicationConfig = struct {
	Application application `toml:"application"`
}

type application = struct {
	Name string
}

func GetConfig() ApplicationConfig {
	var ac ApplicationConfig
	current, _ := filepath.Abs(".")
	configPath := current + "/config/config.toml"

	if _, err := toml.DecodeFile(configPath, &ac); err != nil {
		panic(err)
	}

	return ac
}
