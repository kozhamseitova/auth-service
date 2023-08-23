package utils

import "errors"

var (
	ErrInternalError = errors.New("Internal Server Error")
	ErrInvalidCredentials = errors.New("Invalid Credentials")
	ErrUserAlreadyExists = errors.New("User is alredy exists")
)