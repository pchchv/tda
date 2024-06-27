package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func pingHandler(c fiber.Ctx) error {
	return c.SendString("To do app service. Version 0.0.0")
}

func server() {
	app := fiber.New()

	app.Get("/ping", pingHandler)

	log.Fatal(app.Listen(":" + getEnvValue("PORT")))
}
