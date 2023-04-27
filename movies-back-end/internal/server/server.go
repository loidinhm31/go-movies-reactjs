package server

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"movies-service/config"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	gin   *gin.Engine
	cfg   *config.Config
	db    *gorm.DB
	cloak *gocloak.GoCloak
	ready chan bool
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	return &Server{cfg: cfg, db: db}
}

func (s *Server) Run() {
	r := gin.New()

	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	if err := s.MapHandlers(r); err != nil {
		log.Panic(err)
	}

	go func() {
		log.Printf("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := r.Run(s.cfg.Server.Port); err != nil {
			log.Fatalf("Error starting Server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	_, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	log.Println("Server Exited Properly")
}
