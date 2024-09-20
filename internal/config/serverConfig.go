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

type Server struct {
        DefaultType string `yaml:"default_type"` 
        SendFile string `yaml:"sendfile"`
        KeepAlive int `yaml:"keepalive_timeout"`
        Listen     int    `yaml:"listen"`
        ServerName string `yaml:"server_name"`
        Location   Location `yaml:"location"`
}

type Config struct {
    Http   string   `yaml:"http"`
    Events []string `yaml:"events"`
    Server Server  `yaml:"server"`

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
    fmt.Printf("Could not read the .env file - %v",err)
  }


  err = yaml.Unmarshal(yamlFile, &r)
  if err != nil { 
    fmt.Printf("Could not parse the yaml - %v",err)
  }
  
  fmt.Println(r)
}
