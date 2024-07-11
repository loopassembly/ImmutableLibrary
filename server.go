package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"IMULIB/initializers"
	"IMULIB/routes"
	"github.com/gofiber/template/html/v2"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

func main() {
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
        Views: engine,
    })
	// app := fiber.New()
	micro := fiber.New()


	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))


	app.Use(logger.New())
	app.Use(helmet.New())


	app.Mount("/api", micro)


	micro.Route("/auth", func(router fiber.Router) {
		routes.SetupAuthRoutes(router)
	})

	
	micro.Route("/users", func(router fiber.Router) {
		routes.SetupUserRoutes(router)
	})

	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and GORM",
		})
	})

	app.Get("/", func(c *fiber.Ctx) error {

        return c.Render("index", fiber.Map{
            "Title": "Hello, World!",
        })
    })

	micro.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exist on this server", path),
		})
	})

	log.Fatal(app.Listen(":3000"))
}
