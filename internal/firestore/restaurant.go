package firestore

import (
	"context"
	"log"
	"resto-admin-backend/config"
)

func CreateRestaurant(restaurant any)error{
	collectionName := "restaurants"
    _, _,  err := config.FirestoreClient.Collection(collectionName).Add(context.Background(), restaurant)
	if err != nil {
		log.Fatalf("could not create restaurant: %v", err)
		return err
	}
	return nil   
}
func GetRestaurants()([]map[string]any, error){
	collectionName := "restaurants"
	var list_data []map[string]any
    restaurants ,  err := config.FirestoreClient.Collection(collectionName).DocumentRefs(context.Background()).GetAll()
	if err != nil {
		log.Fatalf("could not create restaurant: %v", err)
		return nil, err
	}
	
	for _, v := range restaurants {
        doc, err := v.Get(context.Background())
		if err != nil {
			log.Fatalf("could not get user info: %v", err)
			return nil, err
		}
		data := doc.Data()
		data["id"] = doc.Ref.ID
		list_data = append(list_data, data)  
	} 
	return list_data, nil   
}
func DeleteRestaurant(id string)error{
	collectionName := "restaurants"
	_, err := config.FirestoreClient.Collection(collectionName).Doc(id).Delete(context.Background())
	if err != nil {
		log.Fatalf("could not create restaurant: %v", err)
		return err
	}
	return nil
}