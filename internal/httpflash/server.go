package httpflash

import (
	"flash/internal/config"
	"flash/internal/core"
	"fmt"
	"net/http"

	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Server struct {
	// Timeout config
	DefaultType string `yaml:"default_type"`
	SendFile    string `yaml:"sendfile"`

	Setup config.Setup `yaml:"setup"`
}

func NewServer() (Server, error) {
	return Server{}, nil
}

// ServeHTTP implementa o handler HTTP
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to %s, you've hit %s\n", s.Setup.Name, r.URL.Path)
}

func (s *Server) LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Could not load env variables %v", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Could not read the .env file - %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &s)
	if err != nil {
		fmt.Printf("Could not parse the yaml - %v", err)
	}

	fmt.Println(s)
}

// Carregar rotas a partir de um arquivo de configuração
// Carregar os plugins
func (s *Server) Start() error {
	plugins := core.GetAllPlugins()

	router := NewRouter(plugins)
	addr := fmt.Sprintf(":%d", s.Setup.Listen)
	fmt.Printf("Listening on %s...\n", addr)
	return http.ListenAndServe(addr, router)
}

func (s *Server) StartWithRouter(handler http.Handler) error {
	port := fmt.Sprintf(":%d", s.Setup.Listen)
	return http.ListenAndServe(port, handler)
}
