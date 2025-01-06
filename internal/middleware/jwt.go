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

// 在文件开头添加自定义类型
type contextKey string

const (
	UserTokenKey contextKey = "user-token"
)

type JWTMiddlewareOption struct {
	Secret          string
	EmergencyHeader string
	DefaultSecret   string
}

// JWTMiddleware returns a middleware that validates JWT tokens
func JWTMiddleware(opt JWTMiddlewareOption) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {

			// Skip JWT check for login endpoint
			if tr, ok := transport.FromServerContext(ctx); ok {
				if tr.Operation() == "/api.appix.v1.Admin/Login" {
					return handler(ctx, req)
				}
			}

			header, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, status.Errorf(codes.Unauthenticated, "context error")
			}

			if header.RequestHeader().Get(opt.EmergencyHeader) != "" {
				return handler(ctx, req)
			}

			jwtHeader := header.RequestHeader().Get("Authorization")

			if jwtHeader != "" {
				auths := strings.SplitN(jwtHeader, " ", 2)
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
					if opt.Secret == "" {
						return []byte(opt.DefaultSecret), nil
					}
					return []byte(opt.Secret), nil
				})

				if err != nil {
					return nil, status.Errorf(codes.Unauthenticated, "failed to validate JWT")
				}

				if !token.Valid {
					return nil, status.Errorf(codes.Unauthenticated, "invalid JWT")
				}

				// If JWT is valid, proceed with request
				ctx = context.WithValue(ctx, UserTokenKey, jwtToken)

				return handler(ctx, req)
			}

			return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
		}
	}
}
