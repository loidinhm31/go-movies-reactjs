package middlewares

import "movies-service/internal/auth"

type MiddlewareManager struct {
	authService auth.Service
}

func NewMiddlewareManager(authService auth.Service) *MiddlewareManager {
	return &MiddlewareManager{authService}
}
