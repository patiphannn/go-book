package bookhandler

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string]*Book{
		// "1": &Book{"1", "FIRE and BLOOD", "George R. R. Martin", 1200},
	}
	jsonData   = `{"_id":"1","name":"FIRE and BLOOD","author":"George R. R. Martin","pages":1200}`
	jsonUpdate = `{"_id":"1","name":"FIRE and BLOOD","author":"George R. R. Martin","pages":1300}`
)

func TestCreateBook(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/book", strings.NewReader(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.createBook(c)) {
		re := regexp.MustCompile(`\n`)
		res := rec.Body.String()
		res = re.ReplaceAllString(res, ``)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, jsonData, res)
	}
}

func TestGetBooks(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.getBooks(c)) {
		re := regexp.MustCompile(`\n`)
		res := rec.Body.String()
		res = re.ReplaceAllString(res, ``)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"1":`+jsonData+`}`, res)
	}
}

func TestGetBook(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/book/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.getBook(c)) {
		re := regexp.MustCompile(`\n`)
		res := rec.Body.String()
		res = re.ReplaceAllString(res, ``)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, jsonData, res)
	}
}

func TestUpdateBook(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(jsonUpdate))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/book/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.updateBook(c)) {
		re := regexp.MustCompile(`\n`)
		res := rec.Body.String()
		res = re.ReplaceAllString(res, ``)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, jsonUpdate, res)
	}
}

func TestDeleteBook(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/book/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.deleteBook(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
