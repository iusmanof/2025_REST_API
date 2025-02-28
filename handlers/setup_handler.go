package handlers

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

type Repository struct{
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api")

	api.Post("/books", r.CreateBooks)
	api.Delete("/books/:id", r.DeleteBook)
	api.Get("/books/:id", r.GetBookByID)
	api.Put("books/:id", r.ReCreateBookByID)
	api.Get("/books", r.GetBooks)

	api.Post("/genre", r.CreateGenre)
	api.Get("/genre", r.GetGenres)
}