package auth

import (
	"testing"

	"github.com/google/uuid"
)

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")
	userID := uuid.New()

	token, err := CreateJWT(secret, userID)
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Error("expected token to be not empty")
	}
}
