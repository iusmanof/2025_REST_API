package main 

import (
    "fmt"
    "log"
	"net/http"
	"os"

	"github.com/iusmanof/2025_go_rest_api-PET-/models"
	"github.com/iusmanof/2025_go_rest_api-PET-/storage"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)
type Book struct {
	Author string		`json:"author"`
	Title string    	`json:"title"`
	Publisher string	`json:"publisher"`
}

type Repository struct{
	DB *gorm.DB
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
	bookModels := &[]models.Books{}

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
	bookModel := models.Books{}
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
	bookModel := &models.Books{}
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

func (r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	 
	api.Post("/create_books", r.CreateBooks)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBookByID)
	api.Get("/get", r.GetBooks)
}

func main(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
	}
	config := &storage.Config{
		Host: os.Getenv("POSTGRES_HOST"),
		Port: os.Getenv("POSTGRES_PORT"),
		User: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName: os.Getenv("POSTGRES_DB"),
		SSLMode: os.Getenv("POSTGRES_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

    if err != nil {
		 log.Fatal("could not load the database")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
	app.Listen(":8080")
}