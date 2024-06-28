package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func pingHandler(c fiber.Ctx) error {
	return c.SendString("To do app service. Version 0.0.0")
}

func indexHandler(c fiber.Ctx, db *sql.DB) (err error) {
	var rows *sql.Rows
	if rows, err = db.Query("SELECT * FROM todos"); err != nil {
		log.Fatal(err)
		c.JSON(fmt.Sprintf("Database error: %v", err))
	}

	defer rows.Close()

	var res string
	var todos []string
	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}

	return c.Render("index", fiber.Map{
		"Todos": todos,
	})
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
