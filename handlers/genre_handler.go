package handlers

import (
	"net/http"
	
	"github.com/2025_REST_API/models"
    "github.com/gofiber/fiber/v2"
)

func (r *Repository) CreateGenre(context *fiber.Ctx) error{
	var genre models.Genre

	err := context.BodyParser(&genre)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{ "message": "request failed"},)
			return err
	}

	err = r.DB.Create(&genre).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create genre"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "genre has been added"})

	return nil
}

func (r *Repository) GetGenres(context *fiber.Ctx) error {
	genreModels := &[]models.Genre{}

	err := r.DB.Find(genreModels).Error
	if err != nil{
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get genres"})
		return nil
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "genres fetched", 
			"data": genreModels,
		})
	return nil
}