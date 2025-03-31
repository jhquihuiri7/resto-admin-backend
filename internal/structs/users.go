package structs

import (
	"resto-admin-backend/internal/firestore"
	"time"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) DeleteUser()error{
   return firestore.DeleteUser(u.Email)
}

type UserInfo struct {
	Id                        string    `json:"id" firestore:"id"`
	Company                   string    `json:"company" firestore:"company"`
	CreatedDatetime           time.Time `json:"created_datetime" firestore:"created_datetime"`
	FirstName                 string    `json:"first_name" firestore:"first_name"`
	LastLoginDatetime         time.Time `json:"last_login_datetime" firestore:"last_login_datetime"`
	LastName                  string    `json:"last_name" firestore:"last_name"`
	Role                      int64     `json:"role" firestore:"role"`
	Suscription               int64     `json:"suscription" firestore:"suscription"`
	SuscriptionExpireDatetime time.Time `json:"suscription_expire_datetime" firestore:"suscription_expire_datetime"`
}

func (u UserInfo) CreateUserInfo() error {
	return firestore.CreateUserInfo(u, u.Id)
}
func (u UserInfo) DeleteId() error {
	return firestore.DeleteId(u.Id)
}
func (u *UserInfo) FromMap(data map[string]interface{}) error {
	u.Id = verifyOk(data, "id").(string)
	u.Company = verifyOk(data, "company").(string)
	//u.CreatedDatetime = verifyOk(data, "id").(time.Time)
	u.FirstName = verifyOk(data, "first_name").(string)
	//u.LastLoginDatetime = verifyOk(data, "id").(time.Time)
	u.LastName = verifyOk(data, "last_name").(string)
	u.Role = verifyOk(data, "role").(int64)
	u.Suscription = verifyOk(data, "suscription").(int64)
	//u.SubscriptionExpireDatetime = verifyOk(data, "id").(time.Time)
	return nil
}

func verifyOk(data map[string]interface{}, key string) interface{} {
	value, ok := data[key]
	if !ok {
		value = "0"
	}
	return value
}
