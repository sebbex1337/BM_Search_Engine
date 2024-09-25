package security_test

import (
	"testing"

	"github.com/UpsDev42069/BM_Search_Engine/backend/security"
)

func TestHashPassword(t *testing.T) {
	password := "123"
	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if hashedPassword == "" {
		t.Fatalf("Expected hashed password to be non-empty")
	}

	t.Logf("Hashed password: %v", hashedPassword)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "123"
	hashedPassword := "$2a$10$vrBLFYpHt4UOaG8sECgqdOl7JHSSftInsfmiu1C/ggoQWn78BsG86"
	if !security.CheckPasswordHash(hashedPassword, password) {
		t.Fatalf("Expected password to match hashed password")
	}

	if !security.CheckPasswordHash(hashedPassword, password) {
		t.Fatalf("Expected password to match hashed password")
	}

	wrongPassword := "wrongpassword"
	if security.CheckPasswordHash(hashedPassword, wrongPassword) {
		t.Fatalf("Expected password not to match hashed password")
	}
}
