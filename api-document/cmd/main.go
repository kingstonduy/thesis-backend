package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/kingstonduy/api-document/resources/docs"
)

const (
	PORT = "8000"
	PATH = "/swagger/*"
)

func main() {
	app := fiber.New()

	swaggerConfig := swagger.Config{
		URL: "doc.json",
	}

	app.Get(PATH, swagger.New(swaggerConfig))

	app.Listen(":" + PORT)
}
