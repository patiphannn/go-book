package book

import (
	"github.com/Kamva/mgm/v2"
	"gopkg.in/mgo.v2/bson"
)

// Book is Book database model
type Book struct {
	mgm.DefaultModel `bson:",inline"`

	Name   string `json:"name" bson:"name" validate:"required"`
	Author string `json:"author" bson:"author" validate:"required"`
	Pages  int    `json:"pages" bson:"pages" validate:"required,gte=10"`
}

// NewBook is Set Book data
func NewBook(name string, author string, pages int) *Book {
	return &Book{
		Name:   name,
		Author: author,
		Pages:  pages,
	}
}

// Define our errors:
var internalError = bson.M{"message": "internal error"}
var bookNotFound = bson.M{"message": "book not found"}

// Request is Request struct
type Request struct {
	Name   string `json:"name" form:"name" query:"name" validate:"required"`
	Author string `json:"author" form:"author" query:"author" validate:"required"`
	Pages  int    `json:"pages" form:"pages" query:"pages" validate:"required,gte=10"`
}
