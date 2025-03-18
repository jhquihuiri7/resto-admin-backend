package firestore

import (
	"context"
	"fmt"
	"log"
	"resto-admin-backend/config"
)

func GetUsers() []string {
	collectionName := "users" // Replace with your collection name
	documents, err := config.FirestoreClient.Collection(collectionName).Documents(context.Background()).GetAll()

	if err != nil {
		log.Fatalf("error getting documents: %v", err)
	}

	for _, v := range documents {
		fmt.Println(v.Ref.ID)
	}

	return []string{}
}
