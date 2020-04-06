package user

import (
	"github.com/Kamva/mgm/v2"
	"gopkg.in/mgo.v2/bson"
)

// User is User database model
type User struct {
	mgm.DefaultModel `bson:",inline"`

	Username string `json:"username" bson:"username" validate:"required,gte=4,lte=20"`
	Password string `json:"-" bson:"password" validate:"required,gte=6,lte=20"`
	Name     string `json:"name" bson:"name" validate:"required"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Admin    bool   `json:"admin" bson:"admin" validate:"required"`
	Token    string `json:"token" bson:"token"`
}

// NewUser is Set User data
func NewUser(username string, password string, name string, email string, admin bool) *User {
	return &User{
		Username: username,
		Password: password,
		Name:     name,
		Email:    email,
		Admin:    admin,
	}
}

// Define our errors:
var internalError = bson.M{"message": "internal error"}
var userNotFound = bson.M{"message": "user not found"}

// Request is Request struct
type Request struct {
	Username string `json:"username" form:"username" validate:"required,gte=4,lte=20"`
	Password string `json:"password" form:"password" validate:"required,gte=6,lte=20"`
	Confirm  string `json:"confirm" form:"confirm" validate:"eqfield=Password"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Admin    bool   `json:"admin" form:"admin"`
}

// RequestLogin is RequestLogin struct
type RequestLogin struct {
	Username string `json:"username" form:"username" validate:"required,gte=4,lte=20"`
	Password string `json:"password" form:"password" validate:"required,gte=6,lte=20"`
}
