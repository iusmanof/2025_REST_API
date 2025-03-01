package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/2025_REST_API/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateBooks(t *testing.T) {
	app := fiber.New()
	db := setupTestDB()
	repo := Repository{DB: db}
	repo.SetupRoutes(app)

	name := "Fantasy"
	genre := models.Genre{Name: &name}
	body, _ := json.Marshal(genre)

	req := httptest.NewRequest("POST", "/api/genre", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	author := "Rowling"
	title := "Harry Potter"
	publisher := "ACT"

	bookRowling := models.Book{
		Author:    &author,
		Title:     &title,
		Publisher: &publisher,
		GenreID:   1}

	body_book, _ := json.Marshal(bookRowling)

	req = httptest.NewRequest("POST", "/api/books", bytes.NewReader(body_book))
	req.Header.Set("Content-Type", "aplication/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetBooks(t *testing.T) {
	app := fiber.New()
	db := setupTestDB()
	repo := Repository{DB: db}
	repo.SetupRoutes(app)

	author := "Rowling"
	title := "Harry Potter"
	publisher := "ACT"

	db.Create(&models.Book{Author: &author, Title: &title, Publisher: &publisher, GenreID: 1})

	req := httptest.NewRequest("GET", "/api/books", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteBook(t *testing.T) {
	app := fiber.New()
	db := setupTestDB()
	repo := Repository{DB: db}
	repo.SetupRoutes(app)

	author := "Rowling"
	title := "Harry"
	publisher := "ACT"

	db.Create(&models.Book{Author: &author, Title: &title, Publisher: &publisher, GenreID: 1})

	req := httptest.NewRequest("DELETE", "/api/books/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetBookByID(t *testing.T) {
	app := fiber.New()
	db := setupTestDB()
	repo := Repository{DB: db}
	repo.SetupRoutes(app)

	author := "Rowling"
	title := "Harry Potter"
	publisher := "ACT"

	db.Create(&models.Book{Author: &author, Title: &title, Publisher: &publisher, GenreID: 1})

	req := httptest.NewRequest("GET", "/api/books/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateBook_InvalidJSON(t *testing.T) {
	app := fiber.New()
	db := setupTestDB()
	repo := Repository{DB: db}
	repo.SetupRoutes(app)

	req := httptest.NewRequest("POST", "/api/books", bytes.NewReader([]byte("{invalid json")))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)

	var response map[string]string
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "request failed", response["message"])
}

func TestCreateBook_DBError(t *testing.T) {
	app := fiber.New()
	db := setupTestDB()
	repo := Repository{DB: db}
	repo.SetupRoutes(app)

	body, _ := json.Marshal(models.Book{})
	req := httptest.NewRequest("POST", "/api/books", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var response map[string]string
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "could not create book", response["message"])
}
