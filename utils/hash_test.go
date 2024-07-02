package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Error hashing password: %v", err)
	}

	if len(hash) == 0 {
		t.Fatal("Expected hashed password to be non-empty")
	}

}

func BenchmarkHashPassword(b *testing.B) {
	password := "mysecretpassword"
	for i := 0; i < b.N; i++ {
		HashPassword(password)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "mysecretpassword"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Error hashing password: %v", err)
	}

	if !CheckPasswordHash(password, hash) {
		t.Fatal("Expected password hash to match")
	}

	if CheckPasswordHash("wrongpassword", hash) {
		t.Fatal("Expected password hash not to match")
	}
}
