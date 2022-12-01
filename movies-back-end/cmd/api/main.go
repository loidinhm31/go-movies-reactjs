package main

import (
	"flag"
	"fmt"
	"log"
	"movies/backend/internal/repository"
	"movies/backend/internal/repository/dbrepo"
	"net/http"
	"time"
)

const port = 9090

type Application struct {
	DSN          string
	Domain       string
	DB           repository.DatabaseRepo
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
}

func main() {
	// Set Application config
	var app Application

	// Read from command line
	dbStr := "host=localhost port=5432 user=postgres password =postgrespw dbname=movies sslmode=disable timezone=UTC connect_timeout=5"
	flag.StringVar(&app.DSN, "dsn", dbStr, "Postgres connection")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "verysecret", "signing secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "signing issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "signing audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&app.Domain, "domain", "example.com", "domain")
	flag.Parse()

	// Connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.auth = Auth{
		Secret:        app.JWTSecret,
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieName:    "__Host-refresh_token",
		CookieDomain:  app.CookieDomain,
	}

	log.Println("Starting Application on port", port)

	// Start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
