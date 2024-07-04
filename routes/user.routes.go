// user.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"hackit/controllers"
	"hackit/middleware"
)

func SetupUserRoutes(router fiber.Router) {
	router.Get("/me", middleware.DeserializeUser, controllers.GetMe)
	router.Post("/books",middleware.DeserializeUser, controllers.UploadBook)
	router.Get("/books", middleware.DeserializeUser,controllers.ShowAllBooks)
}