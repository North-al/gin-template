package config

import (
	"flag"
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
	Docs        Docs        `yaml:"docs"`
}

type Application struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
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
	Database int    `yaml:"database"`
}

type JWT struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"`
}

type Docs struct {
	Host string `yaml:"host"`
	Path string `yaml:"path"`
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
}

func GetEnv() string {
	// 解析命令行参数
	env := flag.String("env", "dev", "set environment (dev, prod, test)")
	flag.Parse()

	return *env
}

func GetConfig() *Config {
	once.Do(initConfig)
	return config
}
