package membership

import (
	"fmt"

	"github.com/ggalihpp/rbmc-backend/database"
	"gopkg.in/mgo.v2/bson"
)

// This layer, will store any Object’s Struct and its method.

type h map[string]interface{}

// Member - Will contains data of the member
type Member struct {
	Username      string       `json:"username" bson:"username"`
	Email         string       `json:"email" bson:"email"`
	Password      string       `json:"password" bson:"password"`
	Territory     string       `json:"territory" bson:"territory"`
	IsAdmin       bool         `json:"is_admin" bson:"is_admin"`
	IsCoordinator bool         `json:"is_coordinator" bson:"is_coordinator"`
	PersonalData  personalData `json:"personal_data" bson:"personal_data"`
	Images        []imageData  `json:"image_data" bson:"image_data"`
}

type personalData struct {
	Address     string `json:"address" bson:"address"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
}

type imageData struct {
	Title     string `json:"title" bson:"address"`
	URL       string `json:"url" bson:"url"`
	SecureURL string `json:"secure_url" bson:"secure_url"`
}

// IsExist - Will check is the username exist or not
func (m *Member) IsExist() (isExist bool) {
	var r Member

	dbSession, col, err := database.Connection("membership")
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	defer database.Disconnect(dbSession)

	err = col.Find(bson.M{"username": m.Username}).One(&r)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	isExist = true

	return

}

// VerifyPassword - Will check either the password is correct or not
func (m *Member) VerifyPassword() (correct bool) {
	var r Member

	dbSession, col, err := database.Connection("membership")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer database.Disconnect(dbSession)

	err = col.Find(bson.M{"username": m.Username}).One(&r)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return checkHash(m.Password, r.Password)

}
