package suprabase

import (
	"context"
	"fmt"
	"resto-admin-backend/config"
	"github.com/gin-gonic/gin"
)

func GetRestaurantByUser(id int8) (string,error){
	conn, err := config.Conn.Acquire(context.Background())
	if err != nil {
	    return "", err
	}
	defer conn.Release()
	conn.Conn().DeallocateAll(context.Background())
	var restaurant string
	err = conn.QueryRow(
		context.Background(), 
		"select * from get_restaurant_name_by_user_id($1);", 
		id).Scan(&restaurant)
	if err != nil {
		return "", err
	}
	return restaurant, nil	
}

func GetUser(id int8) (gin.H,error){
	conn, err := config.Conn.Acquire(context.Background())
	if err != nil {
		return gin.H{}, err
	}
	defer conn.Release()
	conn.Conn().DeallocateAll(context.Background())

	var name string
	var restaurantID int
	
	err = conn.QueryRow(context.Background(), "select concat(first_name, ' ', last_name) as name, restaurant_id from users where id = $1;", id).Scan(&name, &restaurantID)
	
	if err != nil {
		return gin.H{}, err
	}
	fmt.Println(name, restaurantID)
	return gin.H{"name":name, "restaurant_id":restaurantID}, nil	
	
}