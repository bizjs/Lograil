package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bizjs/Lograil/ingestion/internal/api"
	"github.com/bizjs/Lograil/ingestion/internal/config"
	"github.com/bizjs/Lograil/ingestion/internal/storage"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize VictoriaLogs storage
	victoriaLogs, err := storage.NewVictoriaLogsClient(cfg.VictoriaLogsURL)
	if err != nil {
		log.Fatalf("Failed to connect to VictoriaLogs: %v", err)
	}
	defer victoriaLogs.Close()

	// Initialize API server
	server := api.NewServer(cfg, victoriaLogs)

	// Start server in a goroutine
	go func() {
		log.Printf("Starting Ingestion server on port %s", cfg.ServerPort)
		if err := server.Start(":" + cfg.ServerPort); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	if err := server.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
