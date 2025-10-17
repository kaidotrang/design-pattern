package config

import (
	"sync"
)

type Config struct {
	AppName string
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{AppName: "MyApp"}
	})
	return instance
}
