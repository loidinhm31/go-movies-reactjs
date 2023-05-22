package server

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"log"
	authHttp "movies-service/internal/auth/delivery/http"
	authService "movies-service/internal/auth/service"
	managementService "movies-service/internal/control/service"

	"movies-service/internal/middlewares"

	movieHttp "movies-service/internal/movies/delivery/http"
	movieRepository "movies-service/internal/movies/repository"
	movieService "movies-service/internal/movies/service"

	seasonHttp "movies-service/internal/seasons/delivery/http"
	seasonRepository "movies-service/internal/seasons/repository"
	seasonService "movies-service/internal/seasons/service"

	episodeHttp "movies-service/internal/episodes/delivery/http"
	episodeRepository "movies-service/internal/episodes/repository"
	episodeService "movies-service/internal/episodes/service"

	genreHttp "movies-service/internal/genres/delivery/http"
	genreRepository "movies-service/internal/genres/repository"
	genreService "movies-service/internal/genres/service"

	searchHttp "movies-service/internal/search/delivery/http"
	searchRepository "movies-service/internal/search/repository"
	searchService "movies-service/internal/search/service"

	analysisHttp "movies-service/internal/analysis/delivery/http"
	analysisRepository "movies-service/internal/analysis/repository"
	analysisService "movies-service/internal/analysis/service"

	viewHttp "movies-service/internal/views/delivery/http"
	viewRepository "movies-service/internal/views/repository"
	viewService "movies-service/internal/views/service"

	roleHttp "movies-service/internal/roles/delivery/http"
	roleRepository "movies-service/internal/roles/repository"
	roleService "movies-service/internal/roles/service"

	userHttp "movies-service/internal/users/delivery/http"
	userRepository "movies-service/internal/users/repository"
	userService "movies-service/internal/users/service"

	integrationHttp "movies-service/internal/integration/delivery/http"
	integrationService "movies-service/internal/integration/service"

	"net/http"
	"time"
)

func (s *Server) MapHandlers(g *gin.Engine) error {
	s.cloak = gocloak.NewClient(s.cfg.Keycloak.EndPoint)

	// Init repositories
	rRepo := roleRepository.NewRoleRepository(s.cfg, s.db)
	uRepo := userRepository.NewUserRepository(s.cfg, s.db)
	mRepo := movieRepository.NewMovieRepository(s.cfg, s.db)
	iSeasonRepo := seasonRepository.NewSeasonRepository(s.cfg, s.db)
	iEpisodeRepo := episodeRepository.NewEpisodeRepository(s.cfg, s.db)
	gRepo := genreRepository.NewGenreRepository(s.cfg, s.db)
	sRepo := searchRepository.NewSearchRepository(s.cfg, s.db)
	anRepo := analysisRepository.NewAnalysisRepository(s.cfg, s.db)
	vRepo := viewRepository.NewViewRepository(s.cfg, s.db)

	// Init service
	managementCtrl := managementService.NewManagementCtrl(uRepo)
	aService := authService.NewAuthService(s.cfg.Keycloak, s.cloak, managementCtrl, rRepo, uRepo)
	rService := roleService.NewRoleService(rRepo)
	mService := movieService.NewMovieService(managementCtrl, mRepo)
	iSeasonService := seasonService.NewSeasonService(managementCtrl, mRepo, iSeasonRepo, iEpisodeRepo)
	iEpisodeService := episodeService.NewSeasonService(managementCtrl, iSeasonRepo, iEpisodeRepo)
	gService := genreService.NewGenreService(managementCtrl, gRepo)
	sService := searchService.NewSearchService(sRepo)
	anService := analysisService.NewAnalysisService(managementCtrl, anRepo)
	vService := viewService.NewViewService(managementCtrl, vRepo)
	iService := integrationService.NewIntegrationService(s.cfg, managementCtrl, mRepo)
	uService := userService.NewUserService(managementCtrl, rRepo, uRepo)

	// Init handler
	aHandler := authHttp.NewAuthHandler(aService, s.cfg.Keycloak, s.cloak)
	rHandler := roleHttp.NewRoleHandler(rService)
	mHandler := movieHttp.NewMovieHandler(mService)
	iSeasonHandler := seasonHttp.NewSeasonHandler(iSeasonService)
	iEpisodeHandler := episodeHttp.NewEpisodeHandler(iEpisodeService)
	gHandler := genreHttp.NewGenreHandler(gService)
	sHandler := searchHttp.NewGraphHandler(sService)
	anHandler := analysisHttp.NewAnalysisHandler(anService)
	vHandler := viewHttp.NewViewHandler(vService)
	iHandler := integrationHttp.NewIntegrationHandler(iService)
	uHandler := userHttp.NewUserHandler(uService)

	// Init middlewares
	mw := middlewares.NewMiddlewareManager(s.cfg.Keycloak, s.cloak, aService)

	// Init Group
	apiV1 := g.Group("/api/v1")
	health := apiV1.Group("/health")

	authGroup := apiV1.Group("/auth")
	authGroup.Use(mw.JWTValidation())
	authHttp.MapAuthRoutes(authGroup, aHandler)

	movieAuthGroup := authGroup.Group("/movies")
	movieAuthGroup.Use(mw.JWTValidation())
	movieHttp.MapAuthMovieRoutes(movieAuthGroup, mHandler)

	seasonAuthGroup := authGroup.Group("/seasons")
	seasonAuthGroup.Use(mw.JWTValidation())
	seasonHttp.MapAuthSeasonRoutes(seasonAuthGroup, iSeasonHandler)

	episodeAuthGroup := authGroup.Group("/episodes")
	episodeAuthGroup.Use(mw.JWTValidation())
	episodeHttp.MapAuthEpisodeRoutes(episodeAuthGroup, iEpisodeHandler)

	genreAuthGroup := authGroup.Group("/genres")
	genreAuthGroup.Use(mw.JWTValidation())
	genreHttp.MapAuthGenreRoutes(genreAuthGroup, gHandler)

	authRoleGroup := authGroup.Group("/roles")
	authRoleGroup.Use(mw.JWTValidation())
	roleHttp.MapRoleRoutes(authRoleGroup, rHandler)

	authUserGroup := authGroup.Group("/users")
	authUserGroup.Use(mw.JWTValidation())
	userHttp.MapUserRoutes(authUserGroup, uHandler)

	authAnalysisGroup := authGroup.Group("/analysis")
	authAnalysisGroup.Use(mw.JWTValidation())
	analysisHttp.MapAnalysisRoutes(authAnalysisGroup, anHandler)

	authIntegrationGroup := authGroup.Group("/integration")
	authIntegrationGroup.Use(mw.JWTValidation())
	integrationHttp.MapIntegrationRoutes(authIntegrationGroup, iHandler)

	movieGroup := apiV1.Group("/movies")

	seasonGroup := apiV1.Group("/seasons")

	episodeGroup := apiV1.Group("/episodes")

	genreGroup := apiV1.Group("/genres")

	searchGroup := apiV1.Group("/search")

	viewGroup := apiV1.Group("/views")

	// Map public routes
	movieHttp.MapMovieRoutes(movieGroup, mHandler)
	seasonHttp.MapSeasonRoutes(seasonGroup, iSeasonHandler)
	episodeHttp.MapEpisodeRoutes(episodeGroup, iEpisodeHandler)
	genreHttp.MapGenreRoutes(genreGroup, gHandler)
	searchHttp.MapGraphRoutes(searchGroup, sHandler)
	viewHttp.MapViewRoutes(viewGroup, vHandler)

	health.GET("", func(c *gin.Context) {
		log.Printf("Health check: %d", time.Now().Unix())
		c.String(http.StatusOK, "up")
	})

	return nil
}
