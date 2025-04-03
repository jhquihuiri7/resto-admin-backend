package structs

import (
	"fmt"
	"resto-admin-backend/internal/firestore"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) DeleteUser()error{
   return firestore.DeleteUser(u.Email)
}

type UserInfo struct {
	Id           string `json:"id" firestore:"id"`
	RestaurantId string `json:"restaurant_id" firestore:"restaurant_id"`
	FirstName    string `json:"first_name" firestore:"first_name"`
	LastName     string `json:"last_name" firestore:"last_name"`
	Role         string `json:"role" firestore:"role"`
	AvatarRef    string `json:"avatar_ref" firestore:"avatar_ref"`
}

func (u UserInfo) CreateUserInfo() error {
	return firestore.CreateUserInfo(u, u.Id)
}
func (u UserInfo) DeleteId() error {
	return firestore.DeleteId(u.Id)
}
func (u *UserInfo) FromMap(data map[string]interface{}) error {
	u.Id = verifyOk(data, "id").(string)
	u.RestaurantId =  verifyOk(data,"restaurant_id").(string)
	u.FirstName = verifyOk(data, "first_name").(string)
	u.LastName = verifyOk(data, "last_name").(string)
	u.Role = verifyOk(data, "role").(string)
	u.AvatarRef = verifyOk(data, "avatar_ref").(string)
	return nil
}

func verifyOk(data map[string]interface{}, key string) interface{} {
	value, ok := data[key]
	if !ok {
		value = "0"
	}
	return value
}

type Restaurant struct {
	Id       string   `json:"id" firestore:"id"`
	Name     string   `json:"name" firestore:"name"`
	Branches []string `json:"branches" firestore:"branches"`
	Icon_ref string   `json:"icon_ref" firestore:"icon_ref"`
}
func (r Restaurant) CreateRestaurant() error {
	r.Branches = []string{fmt.Sprintf("%v_branch001",r.Name)}
	return firestore.CreateRestaurant(r)
}
func (r *Restaurant) FromMap(data map[string]any) error {
	branches := verifyOk(data, "branches").([]interface{})

	r.Id = verifyOk(data, "id").(string)
	r.Name =  verifyOk(data,"name").(string)
	r.Branches = make([]string, len(branches))
    
	for i, v := range branches {
		if str, ok := v.(string); ok {
			r.Branches[i] = str
		} else {
			return fmt.Errorf("invalid type for branches element")
		}
	}
	return nil
}