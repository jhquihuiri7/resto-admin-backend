package auth

import (
	"context"
	"errors"
	"firebase.google.com/go/auth"
	"fmt"
	"resto-admin-backend/config"
)

func VerifyToken(idToken string) (*auth.Token, error) {
	if idToken == "" {
		return nil, errors.New("se requiere un token")
	}
	token, err := config.AuthClient.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, fmt.Errorf("error al verificar el token: %v", err)
	}
	return token, nil
}
