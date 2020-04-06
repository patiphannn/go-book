package user

import (
	"context"
	"net/http"
	"time"

	"github.com/Kamva/mgm/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// Login handler login.
func Login(c echo.Context) (err error) {
	form := &RequestLogin{}

	// skip checking bind errors.
	if err = c.Bind(form); err != nil {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	user := &User{}
	errFind := mgm.Coll(user).First(bson.M{"username": form.Username, "password": form.Password}, user)
	if errFind != nil {
		return c.JSON(http.StatusBadRequest, userNotFound)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["_id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["admin"] = user.Admin
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	user.Token = t

	return c.JSON(http.StatusOK, user)
}

// Profile handler get user profile.
func Profile(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, GetProfile(c))
}

// Gets handler get users.
func Gets(c echo.Context) (err error) {
	// is admin
	if admin := IsAdmin(c); admin == false {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	result := []User{}
	errFind := mgm.Coll(&User{}).SimpleFind(&result, bson.M{})

	if errFind != nil {
		return c.JSON(http.StatusBadRequest, internalError)
	}

	return c.JSON(http.StatusOK, result)
}

// Get handler get user.
func Get(c echo.Context) (err error) {
	// is admin
	if admin := IsAdmin(c); admin == false {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	id := c.Param("id")
	result := &User{}
	errFind := mgm.Coll(result).FindByID(id, result)

	if errFind != nil {
		return c.JSON(http.StatusBadRequest, userNotFound)
	}

	return c.JSON(http.StatusOK, result)
}

// Create handler create new user.
func Create(c echo.Context) (err error) {
	// is admin
	if admin := IsAdmin(c); admin == false {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	user := &Request{}
	// skip checking bind errors.
	if _ = c.Bind(user); err != nil {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	// Validate our data:
	if err = c.Validate(user); err != nil {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	coll := mgm.Coll(&User{})

	// Create index
	if _, err := coll.Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{
			{
				Keys:    bson.M{"username": 1},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys:    bson.M{"email": 1},
				Options: options.Index().SetUnique(true),
			},
		},
	); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": "Create user index error."})
	}

	data := NewUser(
		user.Username,
		user.Password,
		user.Name,
		user.Email,
		user.Admin,
	)

	errCreate := mgm.Coll(&User{}).Create(data)
	if errCreate != nil {
		return c.JSON(http.StatusBadRequest, internalError)
	}

	return c.JSON(http.StatusCreated, data)
}

// Update handler update user.
func Update(c echo.Context) (err error) {
	// is admin
	if admin := IsAdmin(c); admin == false {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	id := c.Param("id")
	user := &Request{}
	// skip checking bind errors.
	if _ = c.Bind(user); err != nil {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	// Validate our data:
	if err = c.Validate(user); err != nil {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	data := &User{}
	coll := mgm.Coll(data)

	errFind := coll.FindByID(id, data)
	if errFind != nil {
		return c.JSON(http.StatusNotFound, userNotFound)
	}
	data.Username = user.Username
	data.Password = user.Password
	data.Name = user.Name
	data.Email = user.Email
	data.Admin = user.Admin

	if err = coll.Update(data); err != nil {
		return c.JSON(http.StatusBadRequest, internalError)
	}

	return c.JSON(http.StatusOK, data)
}

// Delete handler delete user.
func Delete(c echo.Context) error {
	// is admin
	if admin := IsAdmin(c); admin == false {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	user := &User{}
	coll := mgm.Coll(user)
	err := coll.FindByID(c.Param("id"), user)

	if err != nil {
		return c.JSON(http.StatusNotFound, userNotFound)
	}

	if err := coll.Delete(user); err != nil {
		return c.JSON(http.StatusBadRequest, internalError)
	}

	return c.JSON(http.StatusNoContent, nil)
}
