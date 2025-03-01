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


func TestCreateGenre(t *testing.T) {
	app := fiber.New()
	db := setupTestDB()
	repo := Repository{DB: db}
	repo.SetupRoutes(app)

	name := "Horror"
	genre := models.Genre{Name: &name }
	body, _ := json.Marshal(genre)

	req := httptest.NewRequest("POST", "/api/genre", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetGenres(t *testing.T) {
	app := fiber.New()
	db := setupTestDB()
	repo := Repository{DB: db}
	repo.SetupRoutes(app)

	name := "Comedy"
	db.Create(&models.Genre{Name: &name})

	req := httptest.NewRequest("GET", "/api/genre", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
