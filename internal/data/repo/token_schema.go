package repo

import "errors"

type TokenClaims map[string]any

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrMissingToken = errors.New("missing token")
	ErrExpiredToken = errors.New("expired token")
)
