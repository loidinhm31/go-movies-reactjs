package errors

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrWrongPassword      = errors.New("wrong password")
	ErrUserExisted        = errors.New("user existed")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrInvalidClient      = errors.New("invalid client")
)
