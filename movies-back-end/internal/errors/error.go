package errors

import "errors"

var (
	ErrResourceNotFound   = errors.New("resource not found")
	ErrUserExisted        = errors.New("user existed")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrInvalidClient      = errors.New("invalid client")
	ErrUnAuthorized       = errors.New("unauthorized")
	ErrInvalidInput       = errors.New("invalid input")
)
