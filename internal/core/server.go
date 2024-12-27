package core

import (
	"fmt"
	"net/http"

	"flash/internal/httpflash"
	"flash/internal/modules/request"

	"net"
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

type Server struct {
	// Timeout config
	DefaultType string `yaml:"default_type"`
	SendFile    string `yaml:"sendfile"`

	Setup Setup `yaml:"setup"`
}


func NewServer () (Server, error){
  return Server{
  }, nil
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

	fmt.Println(yamlFile)

	err = yaml.Unmarshal(yamlFile, &s)
	if err != nil {
		fmt.Printf("Could not parse the yaml - %v", err)
	}

	fmt.Println(s)
}

// Carregar rotas a partir de um arquivo de configuração
func (s *Server) Start() {
	port := fmt.Sprintf(":%d", s.Setup.Listen)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Failed to bind to port %v", port)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Listening on port :%s...", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection from", conn.RemoteAddr())

	requestBytes := make([]byte, 1024)
	conn.Read(requestBytes)

	r := request.HttpRequest{}
	req := r.ParseStringToRequest(string(requestBytes))
	response := httpflash.RouteHandler(req)

	conn.Write([]byte(response))
}
