package middlewares

import (
	"os"

	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Authenticated() echo.MiddlewareFunc {
	accessSecret, exists := os.LookupEnv("JWT_ACCESS_SECRET")

	if !exists {
		panic("JWT_ACCESS_SECRET not found")
	}

	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(accessSecret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entities.JwtUserClaims)
		},
	})
}
