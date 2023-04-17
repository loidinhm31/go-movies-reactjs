package server

import (
	"github.com/gin-gonic/gin"
	"log"
	authHttp "movies-service/internal/auth/delivery/http"
	userRepository "movies-service/internal/auth/repository"
	authService "movies-service/internal/auth/service"

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
	uRepo := userRepository.NewUserRepository(s.db)
	mRepo := movieRepository.NewMovieRepository(s.db)
	gRepo := genreRepository.NewGenreRepository(s.db)

	// Init service
	aService := authService.NewAuthService(uRepo,
		[]byte(s.cfg.Server.SigningKey),
		s.cfg.Server.TokenTTL,
		s.cfg.Server.ClientId, s.cfg.Server.ClientSecret)
	mService := movieService.NewMovieService(mRepo)
	gService := genreService.NewGenreService(gRepo)

	// Init handler
	aHandler := authHttp.NewAuthHandler(aService)
	mHandler := movieHttp.NewMovieHandler(mService)
	gHandler := genreHttp.NewGenreHandler(gService)

	// Init middlewares
	//mw := middlewares.NewMiddlewareManager(aService)

	apiV1 := g.Group("/api/v1")

	health := apiV1.Group("/health")
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
