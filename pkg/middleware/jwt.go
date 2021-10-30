package middleware

import (
	"errors"
	"fmt"
	"ifoodish-store/pkg/jwt"
	"ifoodish-store/pkg/resperr"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type JWTHeaderMiddleware struct {
	secret string
}

func NewJWTHeaderMiddleware(secret string) JWTHeaderMiddleware {
	return JWTHeaderMiddleware{
		secret: secret,
	}
}

func (m JWTHeaderMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoCtx echo.Context) (err error) {
		authorizationHeader := echoCtx.Request().Header.Get("Authorization")
		if authorizationHeader == "" {
			return resperr.WithStatusCode(
				errors.New("no authorization header"),
				http.StatusUnauthorized,
			)
		}
		components := strings.Split(authorizationHeader, " ")
		if len(components) != 2 {
			return resperr.WithStatusCode(
				errors.New("authorization header bad format"),
				http.StatusUnauthorized,
			)
		}

		if components[0] != "Bearer" {
			return resperr.WithStatusCode(
				errors.New("no Bearer token in authorization header"),
				http.StatusUnauthorized,
			)
		}

		parentCtx := echoCtx.Request().Context()

		jwtCtx, err := jwt.ParseJWT(parentCtx, m.secret, components[1])
		if err != nil {
			return fmt.Errorf("error parsering jwt from authorization header: %w", err)
		}

		echoCtx.SetRequest(echoCtx.Request().WithContext(jwtCtx))

		return next(echoCtx)
	}
}
