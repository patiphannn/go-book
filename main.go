package main

import (
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/polnoy/go-book/src/book"
	db "github.com/polnoy/go-book/src/common"
	"github.com/polnoy/go-book/src/user"
)

func init() {
	_ = db.ConnectDb()
}

// CustomValidator validator struct
type CustomValidator struct {
	validator *validator.Validate
}

// Validate handle validator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// Echo instance
	e := echo.New()

	// Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Auth
	authGroup := e.Group("/auth")
	{
		authGroup.POST("/login", user.Login)
	}

	// User
	userGroup := e.Group("/user")
	userGroup.Use(middleware.JWT([]byte("secret")))
	{
		userGroup.GET("/profile", user.Profile)
		userGroup.GET("", user.Gets)
		userGroup.GET("/:id", user.Get)
		userGroup.POST("", user.Create)
		userGroup.PUT("/:id", user.Update)
		userGroup.DELETE("/:id", user.Delete)
	}

	// Book
	bookGroup := e.Group("/book")
	bookGroup.Use(middleware.JWT([]byte("secret")))
	{
		bookGroup.GET("", book.Gets)
		bookGroup.GET("/:id", book.Get)
		bookGroup.POST("", book.Create)
		bookGroup.PUT("/:id", book.Update)
		bookGroup.DELETE("/:id", book.Delete)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
