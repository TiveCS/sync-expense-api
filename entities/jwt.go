package entities

import "github.com/golang-jwt/jwt/v5"

type JwtUserClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}
