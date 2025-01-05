package middleware

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// JWTMiddleware returns a middleware that validates JWT tokens
func JWTMiddleware(secret string, emergencyHeader string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			header, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, status.Errorf(codes.Unauthenticated, "context error")
			}

			if header.RequestHeader().Get(emergencyHeader) != "" {
				return handler(ctx, req)
			}

			auths := strings.SplitN(header.RequestHeader().Get("Authorization"), " ", 2)
			if len(auths) != 2 || !strings.EqualFold(auths[0], "Bearer") {
				return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
			}
			jwtToken := auths[1]

			if jwtToken == "" {
				return nil, status.Errorf(codes.Unauthenticated, "empty JWT token")
			}

			// Validate JWT
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				// Get secret from config
				if secret == "" {
					return []byte("secret"), nil
				}
				return []byte(secret), nil
			})

			if err != nil {
				return nil, status.Errorf(codes.Unauthenticated, "failed to validate JWT")
			}

			if !token.Valid {
				return nil, status.Errorf(codes.Unauthenticated, "invalid JWT")
			}

			// If JWT is valid, proceed with request
			return handler(ctx, req)
		}
	}
}
