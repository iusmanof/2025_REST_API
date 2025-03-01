package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/2025_REST_API/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/books", r.CreateBooks)
	api.Delete("/books/:id", r.DeleteBook)
	api.Get("/books/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)

	api.Post("/genre", r.CreateGenre)
	api.Get("/genre", r.GetGenres)
}

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

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to test database:", err)
	}

	db.AutoMigrate(&models.Genre{})
	return db
}
