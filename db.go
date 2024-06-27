package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func base() {
	// Connect to database
	if _, err := sql.Open("postgres", getEnvValue("DATABASE")); err != nil {
		log.Fatal(err)
	}
}
