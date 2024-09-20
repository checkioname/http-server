package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Location struct {
	Root  string   `yaml:"root"`
	Index []string `yaml:"index"`
}

type Server struct {
  Listen int64 `yaml:"listen"`
  ServerName string `yaml:"server_name"`
  Location Location `yaml:"location"`
}


type Config struct {
  Events []string `yaml:"events"`
  Http string `yaml:"http"`
  Include string `yaml:"include"`
  DefaultType string  `yaml:"default_type"`
  SendFile string `yaml:"sendfile"`
  KeepAlive int64 `yaml:"keepalive_timeout"`
}


// Carregar rotas a partir de um arquivo de configuração
func (r *Config) LoadConfig() {

  err := godotenv.Load()
  if err != nil {
    fmt.Printf("Could not load env variables %v", err)
  }

  
  configPath := os.Getenv("CONFIG_PATH")
  yamlFile, err := os.ReadFile(configPath)
  if err != nil {
    fmt.Println(err)
  }
  
  fmt.Println(yamlFile)
}
