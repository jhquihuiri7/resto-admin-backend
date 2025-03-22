package auth

import (
	"context"
	"errors"
	"fmt"
	"resto-admin-backend/config"

	"firebase.google.com/go/auth"
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

func CreateUserAuth(email string, password string) (*auth.UserRecord, error) {
	credentials := (&auth.UserToCreate{}).Email(email).Password(password)
	user, err := config.AuthClient.CreateUser(context.Background(), credentials)
	if err != nil {
		return nil, fmt.Errorf("error al crear usuario: %v", err)
	}
	return user, nil
}
