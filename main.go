package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/go-playground/validator/v10"
	
	book "github.com/polnoy/go-book/src/book"
	db "github.com/polnoy/go-book/src/common"
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

	bookGroup := e.Group("/book")
	{
		bookGroup.POST("", book.Create)
		bookGroup.GET("", book.Gets)
		bookGroup.GET("/:id", book.Get)
		bookGroup.PUT("/:id", book.Update)
		bookGroup.DELETE("/:id", book.Delete)
	}

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
