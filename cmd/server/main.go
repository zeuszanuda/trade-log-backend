package main

import (
	"fmt"
	"log"
	"net/http"

	"trade_log_backend/internal/infrastructure"
)

func main() {
	cfg := infrastructure.LoadConfig()

	addr := fmt.Sprintf(":%s", cfg.App.Port)
	log.Printf("Starting %s on %s in %s mode", cfg.App.Name, addr, cfg.App.Env)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
