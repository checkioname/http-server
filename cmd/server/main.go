package main

import (
	"flash/internal/server"
  "flash/internal/config"
)

func main() {

  c:= config.Config{}
  c.LoadConfig()
	server.Start()
}
