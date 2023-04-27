package middlewares

import (
	"github.com/Nerzal/gocloak/v13"
	"movies-service/config"
	"movies-service/internal/auth"
)

type MiddlewareManager struct {
	keycloak    config.KeycloakConfig
	gocloak     *gocloak.GoCloak
	authService auth.Service
}

func NewMiddlewareManager(keycloak config.KeycloakConfig, gocloak *gocloak.GoCloak, authService auth.Service) *MiddlewareManager {
	return &MiddlewareManager{keycloak, gocloak, authService}
}
