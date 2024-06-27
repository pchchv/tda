package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v3"
)

func pingHandler(c fiber.Ctx) error {
	return c.SendString("To do app service. Version 0.0.0")
}

func indexHandler(c fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func postHandler(c fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func putHandler(c fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func deleteHandler(c fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func server() {
	app := fiber.New()

	app.Get("/ping", pingHandler)

	app.Get("/", func(c fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Post("/", func(c fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Put("/update", func(c fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/delete", func(c fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	log.Fatal(app.Listen(":" + getEnvValue("PORT")))
}
