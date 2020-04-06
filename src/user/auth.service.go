package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// GetProfile handler get user profile.
func GetProfile(c echo.Context) jwt.MapClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return claims
}

// IsAdmin get user is admin.
func IsAdmin(c echo.Context) bool {
	claims := GetProfile(c)
	return claims["admin"].(bool)
}
