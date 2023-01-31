package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi User")
	})

	app.Get("/add", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to App")
	})

	app.Listen(":3000")
}
