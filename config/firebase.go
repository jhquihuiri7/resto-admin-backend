package config

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"sync"

	"google.golang.org/api/option"
)

var (
	FirebaseApp     *firebase.App
	AuthClient      *auth.Client
	FirestoreClient *firestore.Client
	once            sync.Once
)

func InitFirebase() {
	once.Do(func() {
		opt := option.WithCredentialsFile("./resto-admin-backend-d3252-firebase-adminsdk-fbsvc-8f853fb34f.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Printf("error initializing app: %v", err)
		}
		FirebaseApp = app
		auth, err := FirebaseApp.Auth(context.Background())
		if err != nil {
			fmt.Printf("Error al obtener el cliente de autenticación: %v", err)
		}
		AuthClient = auth
		client, err := FirebaseApp.Firestore(context.Background())
		if err != nil {
			fmt.Printf("Error al obtener el cliente de firestore: %v", err)
		}
		FirestoreClient = client
	})
}