package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	store := session.New()
	app := fiber.New()

	// Serve index.html as the landing page
	app.Static("/", "Views")

	// controllers

	app.Listen(":3000")
}
