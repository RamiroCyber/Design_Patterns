package main

import (
	"fmt"
	"sync"
)

const (
	DatabaseURL = "localhost:5432"
	Port        = 8080
)

type Config struct {
	DatabaseURL string
	Port        int
}

var singletonInstance *Config

var once sync.Once

func GetConfig() *Config {
	if singletonInstance == nil {
		once.Do(func() {
			singletonInstance = &Config{
				DatabaseURL: DatabaseURL,
				Port:        Port,
			}
		})
	}
	return singletonInstance
}

func main() {
	config := GetConfig()

	fmt.Printf("Database URL: %s\n", config.DatabaseURL)
	fmt.Printf("Port: %d\n", config.Port)

	anotherConfig := GetConfig()
	if config == anotherConfig {
		fmt.Println("As duas instâncias são as mesmas.")
	}
}
