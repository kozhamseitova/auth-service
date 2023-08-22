package utils

import "errors"

var (
	ErrInternalError = errors.New("Internal Server Error")
	ErrInvalidCredentials = errors.New("Invalid Credentials")
)