package bookhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	// Book is book struct
	Book struct {
		ID     string `json:"_id" form:"_id"`
		Name   string `json:"name" form:"name" validate:"required"`
		Author string `json:"author" form:"author" validate:"required"`
		Pages  int    `json:"pages" form:"pages" validate:"required,gte=10"`
	}
	handler struct {
		db map[string]*Book
	}
)

func (h *handler) createBook(c echo.Context) error {
	u := &Book{}
	if err := c.Bind(u); err != nil {
		return err
	}
	h.db[u.ID] = u
	return c.JSON(http.StatusCreated, u)
}

func (h *handler) getBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, h.db)
}

func (h *handler) getBook(c echo.Context) error {
	id := c.Param("id")
	data := h.db[id]
	if data == nil {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	return c.JSON(http.StatusOK, data)
}

func (h *handler) updateBook(c echo.Context) error {
	id := c.Param("id")

	u := &Book{}
	if err := c.Bind(u); err != nil {
		return err
	}
	h.db[id] = u

	return c.JSON(http.StatusOK, u)
}

func (h *handler) deleteBook(c echo.Context) error {
	id := c.Param("id")
	delete(h.db, id)

	return c.JSON(http.StatusNoContent, nil)
}
