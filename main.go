package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func welcome(c *fiber.Ctx) {
	c.Send("Welcome!")
}

func main() {
	app := fiber.New()

	app.Get("/", welcome)

	log.Fatal(app.Listen(3000))
}
