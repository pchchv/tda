package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func pingHandler(c fiber.Ctx) error {
	return c.SendString("to do app service. Version 0.0.0")
}

func indexHandler(c fiber.Ctx, db *sql.DB) (err error) {
	var rows *sql.Rows
	if rows, err = db.Query("SELECT * FROM todos"); err != nil {
		log.Fatal(err)
		c.JSON(fmt.Sprintf("database error: %e", err))
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
	var newTodo todo
	if err := c.Bind().Body(&newTodo); err != nil {
		log.Printf("error occured: %e", err)
		return c.SendString(fmt.Sprintf("error occured: %e", err))
	}

	fmt.Print(newTodo)

	if newTodo.Item != "" {
		if _, err := db.Exec("INSERT into todos VALUES ($1)", newTodo.Item); err != nil {
			log.Fatalf("error occured while executing query: %e", err)
		}
	} else {
		return c.SendString("error: empty todo")
	}

	return c.Redirect().To("/")
}

func putHandler(c fiber.Ctx, db *sql.DB) error {
	old, new := c.Query("olditem"), c.Query("newitem")
	db.Exec("UPDATE todos SET item=$1 WHERE item=$2", new, old)
	return c.Redirect().To("/")
}

func deleteHandler(c fiber.Ctx, db *sql.DB) error {
	todoToDelete := c.Query("item")
	db.Exec("DELETE from todos WHERE item=$1", todoToDelete)
	return c.SendString("deleted")
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
