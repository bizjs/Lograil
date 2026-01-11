package api

import (
	"context"
	"net/http"
	"time"

	"github.com/bizjs/Lograil/ingestion/internal/config"
	"github.com/bizjs/Lograil/ingestion/internal/storage"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router       *gin.Engine
	server       *http.Server
	config       *config.Config
	victoriaLogs *storage.VictoriaLogsClient
}

func NewServer(cfg *config.Config, vl *storage.VictoriaLogsClient) *Server {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	server := &Server{
		router:       router,
		config:       cfg,
		victoriaLogs: vl,
		server: &http.Server{
			Addr:    ":" + cfg.ServerPort,
			Handler: router,
		},
	}

	server.setupRoutes()

	return server
}

func (s *Server) setupRoutes() {
	// Health check
	s.router.GET("/health", s.healthCheck)

	// Ingestion endpoints
	s.router.POST("/ingest/logs", s.ingestLogs)
	s.router.POST("/ingest/batch", s.ingestBatchLogs)
}

func (s *Server) Start(addr string) error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
