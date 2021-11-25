package routes

import (
	"os"

	"github.com/deepakgonda/go-serverless-test/src/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.

	var isRunningWithoutServerless = os.Getenv("IS_RUNNING_WO_SERVERLESS")

	var route fiber.Router

	if isRunningWithoutServerless == "true" {
		route = a.Group("/dev/api/v1")
	} else {
		route = a.Group("/api/v1")
	}

	route.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! from test")
	})

	route.Get("/json-test", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"success": true,
			"message": "Hello World!",
		})
	})

	// Routes for POST method:
	route.Post("/user/sign/up", controllers.UserSignUp) // register a new user
	route.Post("/user/sign/in", controllers.UserSignIn) // auth, return Access & Refresh tokens
}
