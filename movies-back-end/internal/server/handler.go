package server

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"log"
	authHttp "movies-service/internal/auth/delivery/http"
	userRepository "movies-service/internal/auth/repository"
	authService "movies-service/internal/auth/service"
	"movies-service/internal/middlewares"
	movieHttp "movies-service/internal/movies/delivery/http"
	movieRepository "movies-service/internal/movies/repository"
	movieService "movies-service/internal/movies/service"

	genreHttp "movies-service/internal/genres/delivery/http"
	genreRepository "movies-service/internal/genres/repository"
	genreService "movies-service/internal/genres/service"

	"net/http"
	"time"
)

func (s *Server) MapHandlers(g *gin.Engine) error {
	// Init repositories
	uRepo := userRepository.NewUserRepository(s.cfg, s.db)
	mRepo := movieRepository.NewMovieRepository(s.cfg, s.db)
	gRepo := genreRepository.NewGenreRepository(s.cfg, s.db)

	// Init service
	aService := authService.NewAuthService(uRepo)
	mService := movieService.NewMovieService(mRepo)
	gService := genreService.NewGenreService(gRepo)

	// Init handler
	s.cloak = gocloak.NewClient(s.cfg.Keycloak.EndPoint)
	aHandler := authHttp.NewAuthHandler(aService, s.cfg.Keycloak, s.cloak)
	mHandler := movieHttp.NewMovieHandler(mService)
	gHandler := genreHttp.NewGenreHandler(gService)

	// Init middlewares
	mw := middlewares.NewMiddlewareManager(s.cfg.Keycloak, s.cloak, aService)

	apiV1 := g.Group("/api/v1")

	health := apiV1.Group("/health")
	health.Use(mw.JWTValidation())
	authGroupPublic := apiV1.Group("/auth")

	movieGroup := apiV1.Group("/movies")
	genreGroup := apiV1.Group("/genres")

	// Map routes
	authHttp.MapAuthRoutes(authGroupPublic, aHandler)
	movieHttp.MapMovieRoutes(movieGroup, mHandler)
	genreHttp.MapGenreRoutes(genreGroup, gHandler)

	health.GET("", func(c *gin.Context) {
		log.Printf("Health check: %d", time.Now().Unix())
		c.String(http.StatusOK, "up")
	})

	return nil
}
