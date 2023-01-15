package common

import "errors"

type HTTPError struct {
	Message string `json:"message"`
}

var (
	ErrInvalidToken            = errors.New("invalid token")
	ErrUnexpectedSigningMethod = errors.New("unexpected jwt signing method")
)
