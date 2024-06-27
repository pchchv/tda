package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func pingHandler(c fiber.Ctx) error {
	return c.SendString("To do app service. Version 0.0.0")
}

func indexHandler(c fiber.Ctx) error {
	return c.SendString("Hello")
}

func postHandler(c fiber.Ctx) error {
	return c.SendString("Hello")
}

func putHandler(c fiber.Ctx) error {
	return c.SendString("Hello")
}

func deleteHandler(c fiber.Ctx) error {
	return c.SendString("Hello")
}

func server() {
	app := fiber.New()

	app.Get("/ping", pingHandler)
	app.Get("/", indexHandler)
	app.Post("/", postHandler)
	app.Put("/update", putHandler)
	app.Delete("/delete", deleteHandler)

	log.Fatal(app.Listen(":" + getEnvValue("PORT")))
}
