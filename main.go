package main

import (
	"log"

	"github.com/21satvik/go_fiber_tut/database"
	"github.com/21satvik/go_fiber_tut/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API!")
}

func setupRoutes(app *fiber.App) {
	//welcome endpoint
	app.Get("/api", welcome)
	//user endpoints
	app.Post("/api/user", routes.CreateUser)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
