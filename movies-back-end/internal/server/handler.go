package server

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"log"
	authHttp "movies-service/internal/auth/delivery/http"
	userRepository "movies-service/internal/auth/repository"
	authService "movies-service/internal/auth/service"
	managementService "movies-service/internal/control/service"
	"movies-service/internal/middlewares"
	movieHttp "movies-service/internal/movies/delivery/http"
	movieRepository "movies-service/internal/movies/repository"
	movieService "movies-service/internal/movies/service"

	genreHttp "movies-service/internal/genres/delivery/http"
	genreRepository "movies-service/internal/genres/repository"
	genreService "movies-service/internal/genres/service"

	searchHttp "movies-service/internal/search/delivery/http"
	searchRepository "movies-service/internal/search/repository"
	searchService "movies-service/internal/search/service"

	analysisHttp "movies-service/internal/analysis/delivery/http"
	analysisRepository "movies-service/internal/analysis/repository"
	analysisService "movies-service/internal/analysis/service"

	"net/http"
	"time"
)

func (s *Server) MapHandlers(g *gin.Engine) error {
	// Init repositories
	uRepo := userRepository.NewUserRepository(s.cfg, s.db)
	mRepo := movieRepository.NewMovieRepository(s.cfg, s.db)
	gRepo := genreRepository.NewGenreRepository(s.cfg, s.db)
	sRepo := searchRepository.NewSearchRepository(s.cfg, s.db)
	anRepo := analysisRepository.NewAnalysisRepository(s.cfg, s.db)

	// Init service
	managementCtrl := managementService.NewManagementCtrl(uRepo)
	aService := authService.NewAuthService(uRepo)
	mService := movieService.NewMovieService(managementCtrl, mRepo)
	gService := genreService.NewGenreService(gRepo)
	sService := searchService.NewSearchService(sRepo)
	anService := analysisService.NewAnalysisService(managementCtrl, anRepo)

	// Init handler
	s.cloak = gocloak.NewClient(s.cfg.Keycloak.EndPoint)
	aHandler := authHttp.NewAuthHandler(aService, s.cfg.Keycloak, s.cloak)
	mHandler := movieHttp.NewMovieHandler(mService)
	gHandler := genreHttp.NewGenreHandler(gService)
	sHandler := searchHttp.NewGraphHandler(sService)
	anHandler := analysisHttp.NewAnalysisHandler(anService)

	// Init middlewares
	mw := middlewares.NewMiddlewareManager(s.cfg.Keycloak, s.cloak, aService)

	// Init Group
	apiV1 := g.Group("/api/v1")
	health := apiV1.Group("/health")

	privateMovieGroup := apiV1.Group("/private/movies")
	privateMovieGroup.Use(mw.JWTValidation())
	movieHttp.MapRoleMovieRoutes(privateMovieGroup, mHandler)

	authGroupPublic := apiV1.Group("/auth")

	movieGroup := apiV1.Group("/movies")
	genreGroup := apiV1.Group("/genres")
	searchGroup := apiV1.Group("/search")
	analysisGroup := apiV1.Group("/analysis")

	// Map routes
	authHttp.MapAuthRoutes(authGroupPublic, aHandler)
	movieHttp.MapMovieRoutes(movieGroup, mHandler)
	genreHttp.MapGenreRoutes(genreGroup, gHandler)
	searchHttp.MapGraphRoutes(searchGroup, sHandler)
	analysisHttp.MapAnalysisRoutes(analysisGroup, anHandler)

	health.GET("", func(c *gin.Context) {
		log.Printf("Health check: %d", time.Now().Unix())
		c.String(http.StatusOK, "up")
	})

	return nil
}
