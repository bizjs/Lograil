package api

import (
	"net/http"
	"time"

	"github.com/bizjs/Lograil/ingestion/internal/storage"
	"github.com/gin-gonic/gin"
)

// Health check handler
func (s *Server) healthCheck(c *gin.Context) {
	// Check VictoriaLogs health
	if err := s.victoriaLogs.HealthCheck(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "unhealthy",
			"service": "ingestion",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "ingestion",
	})
}

// Single log ingestion handler
func (s *Server) ingestLogs(c *gin.Context) {
	var req struct {
		Timestamp *time.Time             `json:"timestamp,omitempty"`
		Level     string                 `json:"level" binding:"required"`
		Message   string                 `json:"message" binding:"required"`
		Source    string                 `json:"source" binding:"required"`
		Fields    map[string]interface{} `json:"fields,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default timestamp if not provided
	timestamp := time.Now()
	if req.Timestamp != nil {
		timestamp = *req.Timestamp
	}

	logEntry := storage.LogEntry{
		Timestamp: timestamp,
		Level:     req.Level,
		Message:   req.Message,
		Source:    req.Source,
		Fields:    req.Fields,
	}

	// Write log to VictoriaLogs
	if err := s.victoriaLogs.WriteLogs([]storage.LogEntry{logEntry}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to write log",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Log ingested successfully",
	})
}

// Batch log ingestion handler
func (s *Server) ingestBatchLogs(c *gin.Context) {
	var req struct {
		Logs []struct {
			Timestamp *time.Time             `json:"timestamp,omitempty"`
			Level     string                 `json:"level" binding:"required"`
			Message   string                 `json:"message" binding:"required"`
			Source    string                 `json:"source" binding:"required"`
			Fields    map[string]interface{} `json:"fields,omitempty"`
		} `json:"logs" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(req.Logs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No logs provided"})
		return
	}

	// Convert to LogEntry slice
	logEntries := make([]storage.LogEntry, len(req.Logs))
	for i, log := range req.Logs {
		timestamp := time.Now()
		if log.Timestamp != nil {
			timestamp = *log.Timestamp
		}

		logEntries[i] = storage.LogEntry{
			Timestamp: timestamp,
			Level:     log.Level,
			Message:   log.Message,
			Source:    log.Source,
			Fields:    log.Fields,
		}
	}

	// Write logs to VictoriaLogs in batches
	batchSize := s.config.BatchSize
	for i := 0; i < len(logEntries); i += batchSize {
		end := i + batchSize
		if end > len(logEntries) {
			end = len(logEntries)
		}

		batch := logEntries[i:end]
		if err := s.victoriaLogs.WriteLogs(batch); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":     "Failed to write log batch",
				"details":   err.Error(),
				"processed": i,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logs ingested successfully",
		"count":   len(logEntries),
	})
}
