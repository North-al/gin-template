package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	Application Application `yaml:"application"`
	Database    Database    `yaml:"database"`
	Redis       Redis       `yaml:"redis"`
	JWT         JWT         `yaml:"jwt"`
}

type Application struct {
	Port int `yaml:"port"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type JWT struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"`
}

func initConfig() {
	env := GetEnv()

	configFilePath := filepath.Join("config", env, "config.yml")
	yamlFile, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("Failed to read config file:", err)
		panic(fmt.Sprintf("Failed to read config file: %v", err))
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("Failed to unmarshal config file:", err)
		panic(fmt.Sprintf("Failed to unmarshal config file: %v", err))
	}

	fmt.Println("Config:", config)
}

func GetEnv() string {
	go_env := os.Getenv("GO_ENV")
	if len(go_env) > 0 {
		return go_env
	}

	return "dev"
}

func GetConfig() *Config {
	once.Do(initConfig)
	return config
}
