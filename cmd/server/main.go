package main

import (
	"flash/internal/core"
	"flash/internal/httpflash"
	"fmt"
	"log"

	_ "flash/internal/modules/echo"
)

func main() {

	// 1. Cria e carrega a configuração do servidor
	s := httpflash.Server{}
	s.LoadConfig()

	for _, p := range core.GetAllPlugins() {
		err := p.Init(nil)
		if err != nil {
			log.Fatalf("Erro ao inicializar plugin %s: %v", p.Name(), err)
		}
	}

	plugins := core.GetAllPlugins()
	fmt.Println(plugins)
	router := httpflash.NewRouter(plugins)

	addr := fmt.Sprintf(":%d", s.Setup.Listen)
	log.Printf("Servidor iniciando na porta %s", addr)

	err := s.StartWithRouter(router)
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
