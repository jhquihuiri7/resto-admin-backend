package firestore

import (
	"context"
	"log"
	"resto-admin-backend/config"

	"cloud.google.com/go/firestore"
)

func GetUsers() ([]map[string]interface{}, error) {
	var list_data []map[string]interface{}
	collectionName := "users" // Replace with your collection name
	documents, err := config.FirestoreClient.Collection(collectionName).Documents(context.Background()).GetAll()

	if err != nil {
		log.Fatalf("error getting documents: %v", err)
	}

	for _, v := range documents {
		doc, err := v.Ref.Get(context.Background())
		if err != nil {
			log.Fatalf("could not get user info: %v", err)
			return nil, err
		}
		data :=  doc.Data()
		data["id"] = doc.Ref.ID
		list_data = append(list_data, data)
	}

	return list_data, err
}
func GetUser(id string) (map[string]interface{}, error) {
	doc, err := config.FirestoreClient.Collection("users").Doc(id).Get(context.Background())
	if err != nil {
		log.Fatalf("could not get user info: %v", err)
		return nil, err
	}
	data := doc.Data()
	data["id"] = doc.Ref.ID
	return doc.Data(), nil
}

func CreateUserInfo(userInfo interface{}, id string) error {
	collectionName := "users"
	_, err := config.FirestoreClient.Collection(collectionName).Doc(id).Set(context.Background(), userInfo)
	if err != nil {
		log.Fatalf("could not create user info: %v", err)
		return err
	}
	return nil
}
func DeleteId(id string) error {
	collectionName := "users"
	update := []firestore.Update{
		{
			Path:  "id",             // Field name you want to delete
			Value: firestore.Delete, // Deleting the field
		},
	}
	_, err := config.FirestoreClient.Collection(collectionName).Doc(id).Update(context.Background(), update)
	if err != nil {
		log.Fatalf("could not create user info: %v", err)
		return err
	}
	return nil
}
func DeleteUser(id string) error {
	user, err := config.AuthClient.GetUserByEmail(context.Background(),id)
	if err != nil {
		log.Fatalf("could not delete auth user: %v", err)
		return err
	}
	err = config.AuthClient.DeleteUser(context.Background(),user.UID)
	if err != nil {
		log.Fatalf("could not delete auth user: %v", err)
		return err
	}
	collectionName := "users"
	_, err = config.FirestoreClient.Collection(collectionName).Doc(id).Delete(context.Background())
	if err != nil {
		log.Fatalf("could not delete user info: %v", err)
		return err
	}
	return nil
}
func ValidateToken (token string)error{
    _, err := config.AuthClient.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Fatalf("could not validate token: %v", err)
		return err
	}
	return nil
}