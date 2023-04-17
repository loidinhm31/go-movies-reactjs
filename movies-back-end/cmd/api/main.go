package main

import (
	"log"
	"movies-service/config"
	"movies-service/internal/server"
	database "movies-service/pkg/database/postgres"
	"movies-service/pkg/utils"
	"os"
)

func main() {
	log.Println("Starting API server")

	envProfile := os.Getenv("profile")
	configPath := utils.GetConfigPath(envProfile)
	cfgFile, err := config.LoadConfig(configPath, envProfile)

	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	psqlDB, err := database.NewPsqlDB(cfg, false)
	if err != nil {
		log.Fatalf("Postgresql init: %s", err)
	} else {
		log.Println("Postgres connected")
	}

	s := server.NewServer(cfg, psqlDB)
	s.Run()

	log.Println("Starting Application on port", cfg.Server.Port)
}
