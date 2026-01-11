package storage

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type VictoriaLogsClient struct {
	baseURL    string
	httpClient *http.Client
}

type LogEntry struct {
	Timestamp time.Time              `json:"timestamp"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Source    string                 `json:"source"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

func NewVictoriaLogsClient(baseURL string) (*VictoriaLogsClient, error) {
	return &VictoriaLogsClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}

func (v *VictoriaLogsClient) Close() error {
	// HTTP client doesn't need explicit closing
	return nil
}

func (v *VictoriaLogsClient) WriteLogs(logs []LogEntry) error {
	if len(logs) == 0 {
		return nil
	}

	// Convert logs to VictoriaLogs format
	// VictoriaLogs expects logs in a specific JSON format
	var buffer bytes.Buffer

	for _, log := range logs {
		// Format log entry for VictoriaLogs
		// This is a simplified implementation - actual VictoriaLogs format may vary
		logLine := fmt.Sprintf(
			"%s [%s] %s: %s",
			log.Timestamp.Format(time.RFC3339),
			log.Level,
			log.Source,
			log.Message,
		)

		// Add structured fields if present
		if len(log.Fields) > 0 {
			logLine += " | "
			for k, v := range log.Fields {
				logLine += fmt.Sprintf("%s=%v ", k, v)
			}
		}

		buffer.WriteString(logLine + "\n")
	}

	// Send to VictoriaLogs ingest endpoint
	url := fmt.Sprintf("%s/insert/jsonl", v.baseURL)
	req, err := http.NewRequest("POST", url, &buffer)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/stream+json")

	resp, err := v.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send logs to VictoriaLogs: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("VictoriaLogs returned status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (v *VictoriaLogsClient) HealthCheck() error {
	url := fmt.Sprintf("%s/health", v.baseURL)
	resp, err := v.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("health check failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("health check returned status %d", resp.StatusCode)
	}

	return nil
}
