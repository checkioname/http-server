package main

import (
	"flash/internal/server"
)

func main() {

  s:= server.Server{}
  s.LoadConfig()
	s.Start()
}
