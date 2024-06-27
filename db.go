package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func base() {
	// Connect to database
	var err error
	if db, err = sql.Open("postgres", getEnvValue("DATABASE")); err != nil {
		log.Fatal(err)
	}
}
