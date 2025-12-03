package main

import (
	"go-be/internal/config"
	"go-be/internal/database"
	"go-be/internal/route"
	"go-be/internal/worker"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// init worker
	w := worker.NewWorker(database.DB)
	go w.Run()

	r := route.Mount(database.DB)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
