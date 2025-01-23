package data

import (
	"opspillar/internal/conf"
	"opspillar/internal/data/repo"
	"context"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtMemRepo struct {
	conf *conf.Admin
}

func NewJwtMemRepo(conf *conf.Admin) repo.TokenRepo {
	return &JwtMemRepo{
		conf: conf,
	}
}

func (r *JwtMemRepo) CreateToken(ctx context.Context, claims repo.TokenClaims) (string, error) {
	jwtClaims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(r.conf.JwtExpireHours)).Unix(),
	}
	for k, v := range claims {
		jwtClaims[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	tokenString, err := token.SignedString([]byte(r.conf.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (r *JwtMemRepo) DeleteToken(ctx context.Context, token string) error {
	return nil
}

func (r *JwtMemRepo) ValidateToken(ctx context.Context, token string) (repo.TokenClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, repo.ErrInvalidToken
		}
		return []byte(r.conf.JwtSecret), nil
	})

	if err != nil {
		return nil, repo.ErrInvalidToken
	}

	if !parsedToken.Valid {
		return nil, repo.ErrInvalidToken
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, repo.ErrInvalidToken
	}

	tokenClaims := make(repo.TokenClaims)
	for key, value := range claims {
		if v, ok := value.(string); ok {
			tokenClaims[key] = v
		}
	}

	return tokenClaims, nil

}
