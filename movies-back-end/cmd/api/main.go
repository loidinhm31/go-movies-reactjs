package main

import (
	"flag"
	"fmt"
	"log"
	"movies/backend/internal/repository"
	"movies/backend/internal/repository/dbrepo"
	"net/http"
)

const port = 9090

type Application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {
	// set Application config
	var app Application

	// read from command line
	dbStr := "host=localhost port=5432 user=postgres password =postgrespw dbname=movies sslmode=disable timezone=UTC connect_timeout=5"
	flag.StringVar(&app.DSN, "dsn", dbStr, "Postgres connection")
	flag.Parse()

	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.Domain = "example.com"

	log.Println("Starting Application on port", port)

	// start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
