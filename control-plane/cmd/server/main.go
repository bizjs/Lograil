package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bizjs/Lograil/control-plane/internal/api"
	"github.com/bizjs/Lograil/control-plane/internal/config"
	"github.com/bizjs/Lograil/control-plane/internal/database"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	db, err := database.NewConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run database migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize API server
	server := api.NewServer(cfg, db)

	// Start server in a goroutine
	go func() {
		log.Printf("Starting Control Plane server on port %s", cfg.ServerPort)
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
