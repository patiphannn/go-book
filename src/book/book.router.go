package book

import (
	"github.com/labstack/echo/v4"
)

// Router handle book router
func Router(e *echo.Echo) *echo.Echo {
	bookGroup := e.Group("/book")
	{
		bookGroup.POST("", Create)
		bookGroup.GET("", Gets)
		bookGroup.GET("/:id", Get)
		bookGroup.PUT("/:id", Update)
		bookGroup.DELETE("/:id", Delete)
	}

	return e
}
