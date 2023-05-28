package helper

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"github.com/stretchr/testify/mock"
)

type MockGoCloak struct {
	mock.Mock
}

func (m *MockGoCloak) GetUsers(ctx context.Context, token string, realm string, params gocloak.GetUsersParams) ([]*gocloak.User, error) {
	args := m.Called(ctx, token, realm, params)
	return args.Get(0).([]*gocloak.User), args.Error(1)
}
