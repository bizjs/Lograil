package database

import (
	"context"
	"testing"
)

func TestSQLiteConnection(t *testing.T) {
	// Test SQLite connection
	db, err := NewConnection("file::memory:?cache=shared&_fk=1")
	if err != nil {
		t.Fatalf("Failed to create SQLite connection: %v", err)
	}
	defer db.Close()

	// Test Ent client creation
	client, err := NewEntClient(db)
	if err != nil {
		t.Fatalf("Failed to create Ent client: %v", err)
	}
	defer client.Close()

	// Test schema creation
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		t.Fatalf("Failed to create schema: %v", err)
	}

	// Test basic user operations
	user := client.User.Create().
		SetUsername("testuser").
		SetEmail("test@example.com").
		SetPasswordHash("hashedpassword").
		SetRole("user")

	createdUser, err := user.Save(ctx)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	if createdUser.Username != "testuser" {
		t.Errorf("Expected username 'testuser', got '%s'", createdUser.Username)
	}

	// Test user retrieval
	foundUser, err := client.User.Get(ctx, createdUser.ID)
	if err != nil {
		t.Fatalf("Failed to retrieve user: %v", err)
	}

	if foundUser.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", foundUser.Email)
	}

	t.Logf("SQLite database test passed! Created and retrieved user: %s", foundUser.Username)
}
