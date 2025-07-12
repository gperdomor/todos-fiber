package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Use an external setup function in order to configure the app in tests as well
	app := Setup()

	// Start the application on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}

// Setup Setup a fiber app with all of its routes
func Setup() *fiber.App {
	// Initialize a new app
	app := fiber.New()

	// Middlewares
	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(cors.New())
	app.Use(healthcheck.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Title": "Hello, World!",
		})
	})

	// Return the configured app
	return app
}
