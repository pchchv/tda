package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func server() {
	app := fiber.New()
	log.Fatal(app.Listen(":" + getEnvValue("PORT")))
}
