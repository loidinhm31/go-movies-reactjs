package server

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"log"
	authHttp "movies-service/internal/auth/delivery/http"
	authService "movies-service/internal/auth/service"
	managementService "movies-service/internal/control/service"

	"movies-service/internal/middlewares"

	movieHttp "movies-service/internal/movie/delivery/http"
	movieRepository "movies-service/internal/movie/repository"
	movieService "movies-service/internal/movie/service"

	seasonHttp "movies-service/internal/season/delivery/http"
	seasonRepository "movies-service/internal/season/repository"
	seasonService "movies-service/internal/season/service"

	episodeHttp "movies-service/internal/episode/delivery/http"
	episodeRepository "movies-service/internal/episode/repository"
	episodeService "movies-service/internal/episode/service"

	genreHttp "movies-service/internal/genre/delivery/http"
	genreRepository "movies-service/internal/genre/repository"
	genreService "movies-service/internal/genre/service"

	searchHttp "movies-service/internal/search/delivery/http"
	searchRepository "movies-service/internal/search/repository"
	searchService "movies-service/internal/search/service"

	analysisHttp "movies-service/internal/analysis/delivery/http"
	analysisRepository "movies-service/internal/analysis/repository"
	analysisService "movies-service/internal/analysis/service"

	viewHttp "movies-service/internal/view/delivery/http"
	viewRepository "movies-service/internal/view/repository"
	viewService "movies-service/internal/view/service"

	roleHttp "movies-service/internal/role/delivery/http"
	roleRepository "movies-service/internal/role/repository"
	roleService "movies-service/internal/role/service"

	ratingHttp "movies-service/internal/rating/delivery/http"
	ratingRepository "movies-service/internal/rating/repository"
	ratingService "movies-service/internal/rating/service"

	userHttp "movies-service/internal/user/delivery/http"
	userRepository "movies-service/internal/user/repository"
	userService "movies-service/internal/user/service"

	referenceHttp "movies-service/internal/reference/delivery/http"
	integrationService "movies-service/internal/reference/service"

	blobHttp "movies-service/internal/blob/delivery/http"
	blobService "movies-service/internal/blob/service"

	"net/http"
	"time"
)

func (s *Server) MapHandlers(g *gin.Engine) error {
	s.cloak = gocloak.NewClient(s.cfg.Keycloak.EndPoint)

	// Init repositories
	iRoleRepo := roleRepository.NewRoleRepository(s.cfg, s.db)
	iUserRepo := userRepository.NewUserRepository(s.cfg, s.db)
	iMovieRepo := movieRepository.NewMovieRepository(s.cfg, s.db)
	iSeasonRepo := seasonRepository.NewSeasonRepository(s.cfg, s.db)
	iEpisodeRepo := episodeRepository.NewEpisodeRepository(s.cfg, s.db)
	iGenreRepo := genreRepository.NewGenreRepository(s.cfg, s.db)
	iSearchRepo := searchRepository.NewSearchRepository(s.cfg, s.db)
	iAnalysisRepo := analysisRepository.NewAnalysisRepository(s.cfg, s.db)
	iViewRepo := viewRepository.NewViewRepository(s.cfg, s.db)
	iRatingRepo := ratingRepository.NewRatingRepository(s.cfg, s.db)

	// Init service
	managementCtrl := managementService.NewManagementCtrl(iUserRepo)
	iAuthService := authService.NewAuthService(s.cfg.Keycloak, s.cloak, managementCtrl, iRoleRepo, iUserRepo)
	iRoleService := roleService.NewRoleService(iRoleRepo)
	iBlobService := blobService.NewBlobService(s.cfg, managementCtrl)
	iMovieService := movieService.NewMovieService(managementCtrl, iMovieRepo, iBlobService)
	iSeasonService := seasonService.NewSeasonService(managementCtrl, iMovieRepo, iSeasonRepo, iEpisodeRepo)
	iEpisodeService := episodeService.NewEpisodeService(managementCtrl, iSeasonRepo, iEpisodeRepo)
	iGenreService := genreService.NewGenreService(managementCtrl, iGenreRepo)
	iSearchService := searchService.NewSearchService(iSearchRepo)
	iAnalysisService := analysisService.NewAnalysisService(managementCtrl, iAnalysisRepo)
	iViewService := viewService.NewViewService(managementCtrl, iViewRepo)
	iReferenceService := integrationService.NewReferenceService(s.cfg, managementCtrl)
	iUserService := userService.NewUserService(managementCtrl, iRoleRepo, iUserRepo)
	iRatingService := ratingService.NewRatingService(iRatingRepo)

	// Init handler
	iAuthHandler := authHttp.NewAuthHandler(iAuthService, s.cfg.Keycloak, s.cloak)
	iRoleHandler := roleHttp.NewRoleHandler(iRoleService)
	iMovieHandler := movieHttp.NewMovieHandler(iMovieService)
	iSeasonHandler := seasonHttp.NewSeasonHandler(iSeasonService)
	iEpisodeHandler := episodeHttp.NewEpisodeHandler(iEpisodeService)
	iGenreHandler := genreHttp.NewGenreHandler(iGenreService)
	iSearchHandler := searchHttp.NewGraphHandler(iSearchService)
	iAnalysisHandler := analysisHttp.NewAnalysisHandler(iAnalysisService)
	iViewHandler := viewHttp.NewViewHandler(iViewService)
	iReferenceHandler := referenceHttp.NewReferenceHandler(iReferenceService)
	iBlobHandler := blobHttp.NewBlobHandler(iBlobService)
	iUserHandler := userHttp.NewUserHandler(iUserService)
	iRatingHandler := ratingHttp.NewRatingHandler(iRatingService)

	// Init middlewares
	mw := middlewares.NewMiddlewareManager(s.cfg.Keycloak, s.cloak, iAuthService)

	// Init Group
	apiV1 := g.Group("/api/v1")
	health := apiV1.Group("/health")

	authGroup := apiV1.Group("/auth")
	authGroup.Use(mw.JWTValidation())
	authHttp.MapAuthRoutes(authGroup, iAuthHandler)

	movieAuthGroup := authGroup.Group("/movies")
	movieAuthGroup.Use(mw.JWTValidation())
	movieHttp.MapAuthMovieRoutes(movieAuthGroup, iMovieHandler)

	seasonAuthGroup := authGroup.Group("/seasons")
	seasonAuthGroup.Use(mw.JWTValidation())
	seasonHttp.MapAuthSeasonRoutes(seasonAuthGroup, iSeasonHandler)

	episodeAuthGroup := authGroup.Group("/episodes")
	episodeAuthGroup.Use(mw.JWTValidation())
	episodeHttp.MapAuthEpisodeRoutes(episodeAuthGroup, iEpisodeHandler)

	genreAuthGroup := authGroup.Group("/genres")
	genreAuthGroup.Use(mw.JWTValidation())
	genreHttp.MapAuthGenreRoutes(genreAuthGroup, iGenreHandler)

	authRoleGroup := authGroup.Group("/roles")
	authRoleGroup.Use(mw.JWTValidation())
	roleHttp.MapRoleRoutes(authRoleGroup, iRoleHandler)

	authUserGroup := authGroup.Group("/users")
	authUserGroup.Use(mw.JWTValidation())
	userHttp.MapUserRoutes(authUserGroup, iUserHandler)

	authAnalysisGroup := authGroup.Group("/analysis")
	authAnalysisGroup.Use(mw.JWTValidation())
	analysisHttp.MapAnalysisRoutes(authAnalysisGroup, iAnalysisHandler)

	authRefGroup := authGroup.Group("/references")
	authRefGroup.Use(mw.JWTValidation())
	referenceHttp.MapReferenceRoutes(authRefGroup, iReferenceHandler)

	authBlobGroup := authGroup.Group("/blobs")
	authBlobGroup.Use(mw.JWTValidation())
	blobHttp.MapIntegrationRoutes(authBlobGroup, iBlobHandler)

	movieGroup := apiV1.Group("/movies")

	seasonGroup := apiV1.Group("/seasons")

	episodeGroup := apiV1.Group("/episodes")

	genreGroup := apiV1.Group("/genres")

	searchGroup := apiV1.Group("/search")

	viewGroup := apiV1.Group("/views")

	ratingGroup := apiV1.Group("/ratings")

	// Map public routes
	movieHttp.MapMovieRoutes(movieGroup, iMovieHandler)
	seasonHttp.MapSeasonRoutes(seasonGroup, iSeasonHandler)
	episodeHttp.MapEpisodeRoutes(episodeGroup, iEpisodeHandler)
	genreHttp.MapGenreRoutes(genreGroup, iGenreHandler)
	searchHttp.MapGraphRoutes(searchGroup, iSearchHandler)
	viewHttp.MapViewRoutes(viewGroup, iViewHandler)
	ratingHttp.MapRatingRoutes(ratingGroup, iRatingHandler)

	health.GET("", func(c *gin.Context) {
		log.Printf("Health check: %d", time.Now().Unix())
		c.String(http.StatusOK, "up")
	})

	return nil
}
