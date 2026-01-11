package api

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/bizjs/Lograil/control-plane/internal/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	server *http.Server
	db     *sql.DB
	config *config.Config
}

func NewServer(cfg *config.Config, db *sql.DB) *Server {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	server := &Server{
		router: router,
		db:     db,
		config: cfg,
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

	// API v1 routes
	v1 := s.router.Group("/api/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/login", s.login)
			auth.POST("/register", s.register)
		}

		// Protected routes
		// TODO: Add JWT middleware here
		protected := v1.Group("")
		{
			// User routes
			users := protected.Group("/users")
			{
				users.GET("", s.getUsers)
				users.POST("", s.createUser)
			}

			// Project routes
			projects := protected.Group("/projects")
			{
				projects.GET("", s.getProjects)
				projects.POST("", s.createProject)
				projects.GET("/:id", s.getProject)
				projects.PUT("/:id", s.updateProject)
				projects.DELETE("/:id", s.deleteProject)

				// Project logs
				projects.GET("/:id/logs", s.getProjectLogs)
			}

			// Configuration routes
			config := protected.Group("/config")
			{
				config.GET("/retention", s.getRetentionPolicies)
				config.PUT("/retention", s.updateRetentionPolicy)
			}
		}
	}
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
