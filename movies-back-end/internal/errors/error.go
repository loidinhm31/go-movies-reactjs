package errors

import (
	"errors"
	"fmt"
)

var (
	ErrResourceNotFound    = errors.New("resource not found")
	ErrUserExisted         = errors.New("user existed")
	ErrInvalidAccessToken  = errors.New("invalid access token")
	ErrInvalidClient       = errors.New("invalid client")
	ErrUnAuthorized        = errors.New("unauthorized")
	ErrInvalidInput        = errors.New("invalid input")
	ErrCannotExecuteAction = errors.New("cannot execute this action")
	ErrUserNotFound        = errors.New("user not found")
)

func ErrInvalidInputDetail(name string) error {
	return errors.New(fmt.Sprintf("invalid input %s", name))
}
