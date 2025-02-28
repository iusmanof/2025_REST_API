package handlers

import (
	"os"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/2025_REST_API/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

Host: os.Getenv("POSTGRES_HOST"),
		Port: os.Getenv("POSTGRES_PORT"),
		User: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName: os.Getenv("POSTGRES_DB"),
		SSLMode: os.Getenv("POSTGRES_SSLMODE"),

import (
	"fmt"
	"os"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	sslmode := os.Getenv("POSTGRES_SSLMODE")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" || sslmode == "" {
		log.Fatal("One or more environment variables for DB connection are missing")
	}

	// Формируем строку подключения
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
		host, port, user, password, dbname, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to test database:", err)
	}

	db.AutoMigrate(&models.Genre{}) // Миграция таблицы для тестов
	return db
}


func TestCreateGenre(t *testing.T) {
	app := fiber.New()
	db := setupTestDB()
	repo := Repository{DB: db}
	repo.SetupRoutes(app)

	genre := models.Genre{Name: "Horror"}
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

	// Добавляем тестовый жанр в базу
	db.Create(&models.Genre{Name: "Comedy"})

	req := httptest.NewRequest("GET", "/api/genre", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
