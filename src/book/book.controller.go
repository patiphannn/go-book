package book

import (
	"net/http"

	"github.com/Kamva/mgm/operator"
	"github.com/Kamva/mgm/v2"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)

// Gets handler get books.
func Gets(c echo.Context) (err error) {
	result := []Book{}
	errFind := mgm.Coll(&Book{}).SimpleFind(&result, bson.M{"pages": bson.M{operator.Gt: 0}})

	if errFind != nil {
		return c.JSON(http.StatusBadRequest, internalError)
	}

	return c.JSON(http.StatusOK, result)
}

// Get handler get book.
func Get(c echo.Context) (err error) {
	id := c.Param("id")
	result := &Book{}
	errFind := mgm.Coll(result).FindByID(id, result)

	if errFind != nil {
		return c.JSON(http.StatusBadRequest, bookNotFound)
	}

	return c.JSON(http.StatusOK, result)
}

// Create handler create new book.
func Create(c echo.Context) (err error) {
	book := &Request{}

	// skip checking bind errors.
	if _ = c.Bind(book); err != nil {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	// Validate our data:
	if err = c.Validate(book); err != nil {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	data := NewBook(
		book.Name,
		book.Author,
		book.Pages,
	)

	errCreate := mgm.Coll(&Book{}).Create(data)

	if errCreate != nil {
		return c.JSON(http.StatusBadRequest, internalError)
	}

	return c.JSON(http.StatusCreated, data)
}

// Update handler update book.
func Update(c echo.Context) (err error) {
	id := c.Param("id")
	book := &Request{}

	// skip checking bind errors.
	if _ = c.Bind(book); err != nil {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	// Validate our data:
	if err = c.Validate(book); err != nil {
		return c.JSON(http.StatusInternalServerError, internalError)
	}

	data := &Book{}
	coll := mgm.Coll(data)

	errFind := coll.FindByID(id, data)
	if errFind != nil {
		return c.JSON(http.StatusNotFound, bookNotFound)
	}
	data.Name = book.Name
	data.Author = book.Author
	data.Pages = book.Pages

	if err = coll.Update(data); err != nil {
		return c.JSON(http.StatusBadRequest, internalError)
	}

	return c.JSON(http.StatusOK, data)
}

// Delete handler delete book.
func Delete(c echo.Context) error {
	book := &Book{}
	coll := mgm.Coll(book)
	err := coll.FindByID(c.Param("id"), book)

	if err != nil {
		return c.JSON(http.StatusNotFound, bookNotFound)
	}

	if err := coll.Delete(book); err != nil {
		return c.JSON(http.StatusBadRequest, internalError)
	}

	return c.JSON(http.StatusNoContent, nil)
}
