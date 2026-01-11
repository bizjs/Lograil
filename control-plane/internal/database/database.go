package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/bizjs/Lograil/pkg/data"

	_ "github.com/mattn/go-sqlite3"
)

func NewConnection(databaseURL string) (*sql.DB, error) {
	var driverName string

	// Determine driver based on URL
	if strings.HasPrefix(databaseURL, "sqlite://") {
		driverName = "sqlite3"
		databaseURL = strings.TrimPrefix(databaseURL, "sqlite://")
	} else if strings.HasPrefix(databaseURL, "postgres://") || strings.HasPrefix(databaseURL, "postgresql://") {
		driverName = "postgres"
		// Keep PostgreSQL import for backward compatibility
		_ = "github.com/lib/pq"
	} else {
		// Default to SQLite for development
		driverName = "sqlite3"
		if databaseURL == "" {
			databaseURL = "file:lograil.db?cache=shared&_fk=1"
		}
	}

	db, err := sql.Open(driverName, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Configure connection pool
	if driverName == "sqlite3" {
		// SQLite specific settings
		db.SetMaxOpenConns(1) // SQLite only supports one writer
	} else {
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
	}

	return db, nil
}

func NewEntClient(db *sql.DB) (*data.Client, error) {
	drv := entsql.OpenDB(dialect.SQLite, db)
	return data.NewClient(data.Driver(drv)), nil
}

func RunMigrations(db *sql.DB) error {
	// Create Ent client
	client, err := NewEntClient(db)
	if err != nil {
		return fmt.Errorf("failed to create ent client: %w", err)
	}
	defer client.Close()

	// Run auto-migration
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	// Alternative: Use migration tables for version control
	// if err := migrate.NamedDiff(ctx, client.Schema, "initial"); err != nil {
	//     return fmt.Errorf("failed to run migration: %w", err)
	// }

	return nil
}
