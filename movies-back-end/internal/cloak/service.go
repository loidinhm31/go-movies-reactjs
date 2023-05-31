package cloak

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
)

type GoCloakClientInterface interface {
	GetUsers(ctx context.Context, token string, realm string, params gocloak.GetUsersParams) ([]*gocloak.User, error)
}

// GoCloakClientWrapper is a wrapper around the gocloak.GoCloak client
type GoCloakClientWrapper struct {
	client *gocloak.GoCloak
}

func (w *GoCloakClientWrapper) GetUsers(ctx context.Context, token string, realm string, params gocloak.GetUsersParams) ([]*gocloak.User, error) {
	return w.client.GetUsers(ctx, token, realm, params)
}
