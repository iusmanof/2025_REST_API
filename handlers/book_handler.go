package handlers

import (
	"net/http"
	"fmt"
	
	"github.com/2025_REST_API/models"
    "github.com/gofiber/fiber/v2"
)

type Book struct {
	Author 		string		`json:"author"`
	Title 		string    	`json:"title"`
	Publisher 	string		`json:"publisher"`
	GenreID 	uint 		`json:"genre_id"`
}

func (r *Repository) CreateBooks(context *fiber.Ctx) error{
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{ "message": "request failed"},)
			return err
	}

	err = r.DB.Create(&book).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "book has been added"})

	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Book{}

	err := r.DB.Find(bookModels).Error
	if err != nil{
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return nil
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "books fetched", 
			"data": bookModels,
		})
	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Book{}
	id := context.Params("id")
	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(bookModel)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete book",
		})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book delete successfuly",
	})
	return nil
}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {

	id := context.Params("id")
	bookModel := &models.Book{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ? ", id).First(bookModel).Error 
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the book"},
		)
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book fetch successfuly",
		"data": bookModel,
	})
	return nil
}

func (r *Repository) ReCreateBookByID(context *fiber.Ctx) error{
	id := context.Params("id")
	fmt.Println("the ID ewfwefwef is", id)
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "test"},
	)
	return nil
}
