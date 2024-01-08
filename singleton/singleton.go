package singleton

import (
	"fmt"
	"sync"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	Password string
	Username string
	Email    string
}

// GetConfig function will return single instanse of Config
func GetConfig() *Config {
	once.Do(func() {
		config = &Config{
			Password: "sample",
			Username: "sample",
			Email:    "sample",
		}
	})
	return config
}

func mainSingleton() {
	// config1 and config2 will have the same address
	config1 := GetConfig()
	config2 := GetConfig()
	fmt.Println(config1 == config2)
}
