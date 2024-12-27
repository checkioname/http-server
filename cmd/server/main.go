package main

import (
	"flash/internal/core"
)

func main() {

  s, _ := core.NewServer()
  s.LoadConfig()
	s.Start()
}
