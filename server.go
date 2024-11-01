package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalinchkma/mrk69/mailers"
)

func main() {
	app := fiber.New()

	mailer := mailers.AuthMailer{}
	app.Get("/hello-world", func(c *fiber.Ctx) error {
		return c.SendString("hello world test")
	})

	app.Get("/name/:name", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("name"))
	})

	app.Get("/age/:age?", func(c *fiber.Ctx) error {
		return c.SendString("Age: " + c.Params("age"))
	})

	app.Get("/random/*", func(c *fiber.Ctx) error {
		return c.SendString("Test" + c.Params("*"))
	})

	app.Get("/sample", func(c *fiber.Ctx) error {
		return c.SendString(mailer.Welcome())
	})

	app.Static("/", "./public") // static files

	app.Listen(":3000")
}
