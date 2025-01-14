package repo

import "errors"

type TokenClaims map[string]string

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrMissingToken = errors.New("missing token")
	ErrExpiredToken = errors.New("expired token")
)

type ApiKey struct {
	ID          uint32 `gorm:"primaryKey;autoIncrement"`
	Key         string `gorm:"type:varchar(255);unique"`
	Description string `gorm:"type:varchar(255);"`
}

type ApiKeysFilter struct {
	Ids      []uint32
	Keys     []string
	Page     uint32
	PageSize uint32
}
