package controllers

import (
	"fmt"
	"IMULIB/initializers"
	"IMULIB/models"
	"IMULIB/utils"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": user,
		},
	})
}

// UploadBook handles uploading a book
// UploadBook handles uploading a book
// UploadBook handles uploading a book\




// finally  finall dsiifmfgsknjsnjgnsf
func UploadBook(c *fiber.Ctx) error {
	
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Failed to parse multipart form"})
	}

	var payload models.BookInput
	if titles, ok := form.Value["title"]; ok && len(titles) > 0 {
		payload.Title = titles[0]
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Title is required"})
	}

	if authors, ok := form.Value["author"]; ok && len(authors) > 0 {
		payload.Author = authors[0]
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Author is required"})
	}

	if descriptions, ok := form.Value["discription"]; ok && len(descriptions) > 0 {
		payload.Description = descriptions[0]
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Description is required"})
	}

	if genres, ok := form.Value["genre"]; ok && len(genres) > 0 {
		payload.Genre = genres[0]
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Genre is required"})
	}

	
	if files, ok := form.File["bookcontent"]; ok && len(files) > 0 {
		payload.Bookcontent = files[0]
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Book content file is required"})
	}

	
	fileContent, err := payload.Bookcontent.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to open book content file"})
	}
	defer fileContent.Close()

	
	tempDir := "uploads/temp"
	tempFile := filepath.Join(tempDir, payload.Bookcontent.Filename)

	
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to create temporary directory"})
	}

	
	if err := c.SaveFile(payload.Bookcontent, tempFile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to save book content file"})
	}

	
	cid, err := utils.UploadFileAndGetCID(tempFile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to upload book content"})
	}

	
	defer func() {
		if err := os.Remove(tempFile); err != nil {
			log.Printf("Failed to remove temporary file: %v", err)
		}
	}()

	
	if files, ok := form.File["bookthumbnail"]; ok && len(files) > 0 {
		payload.BookThumbnail = files[0]
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Book thumbnail file is required"})
	}

	
	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config", err)
	}

	
	filePath := fmt.Sprintf("uploads/%s", payload.BookThumbnail.Filename)
	if err := c.SaveFile(payload.BookThumbnail, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to save thumbnail file"})
	}

	
	user := c.Locals("user").(models.UserResponse)
	newBook := models.Book{
		Title:         payload.Title,
		Author:        payload.Author,
		Description:   payload.Description,
		BookThumbnail: config.ClientOrigin + filePath, 
		Bookcontent:   "https://"+cid+".ipfs.w3s.link",
		Genre:         payload.Genre,
		UserID:        user.ID.String(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	
	result := initializers.DB.Create(&newBook)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Failed to save book record"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"book": newBook}})
}


func ShowAllBooks(c *fiber.Ctx) error {
	var books []models.Book
	result := initializers.DB.Find(&books)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"books": books}})
}
