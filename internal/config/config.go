package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Location struct {
	Root  string   `yaml:"root"`
	Index []string `yaml:"index"`
}

type Setup struct {
	KeepAlive  int      `yaml:"keepalive_timeout"`
	Listen     int      `yaml:"listen"`
	ServerName string   `yaml:"server_name"`
	Location   Location `yaml:"location"`
	Name       string   `yaml:"name"`
	Events     []string `yaml:"events"`
}

type Config interface {
	LoadConfig() (Setup, error)
}

func NewConfig() Config {
	return &config{}
}

type config struct {
  Setup Setup `yaml:"setup"`
}

func (c *config) LoadConfig() (Setup, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Could not load env variables %v", err)
		return Setup{}, err
	}

	configPath := os.Getenv("CONFIG_PATH")
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Could not read the .env file - %v", err)
		return Setup{}, err
	}

	err = yaml.Unmarshal(yamlFile, &c.Setup)
	if err != nil {
		fmt.Printf("Could not parse the yaml - %v", err)
		return Setup{}, err
	}

  if err = validateConfig(); err != nil {
    return Setup{}, nil
  }

	fmt.Println(c.Setup)
	return c.Setup, nil
}


func validateConfig() error {
  return nil
}
