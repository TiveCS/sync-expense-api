package usecase

import (
	"os"
	"time"

	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/golang-jwt/jwt/v5"
)

type AuthGenerateTokenUsecase interface {
	sign(claims *entities.JwtUserClaims, isAccess bool) (string, error)
	Execute(user *entities.User) (string, string, error)
}

type authGenerateTokenUsecase struct {
	accessSecret  string
	refreshSecret string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

// sign implements AuthGenerateTokenUsecase.
func (u *authGenerateTokenUsecase) sign(claims *entities.JwtUserClaims, isAccess bool) (string, error) {
	var secret []byte
	if isAccess {
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(u.accessExpiry))
		secret = []byte(u.accessSecret)
	} else {
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(u.refreshExpiry))
		secret = []byte(u.refreshSecret)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return signed, nil
}

// Execute implements AuthGenerateTokenUsecase.
func (u *authGenerateTokenUsecase) Execute(user *entities.User) (string, string, error) {
	baseClaims := &entities.JwtUserClaims{
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: user.ID,
		},
	}

	accessToken, err := u.sign(baseClaims, true)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := u.sign(baseClaims, false)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func NewAuthGenerateTokenUsecase() AuthGenerateTokenUsecase {
	accessSecret, exists := os.LookupEnv("JWT_ACCESS_SECRET")
	if !exists {
		panic("JWT_ACCESS_SECRET not found")
	}

	refreshSecret, exists := os.LookupEnv("JWT_REFRESH_SECRET")
	if !exists {
		panic("JWT_REFRESH_SECRET not found")
	}

	accessExpiryStr, exists := os.LookupEnv("JWT_ACCESS_EXPIRY")
	if !exists {
		panic("JWT_ACCESS_EXPIRY not found")
	}

	refreshExpiryStr, exists := os.LookupEnv("JWT_REFRESH_EXPIRY")
	if !exists {
		panic("JWT_REFRESH_EXPIRY not found")
	}

	accessExpiry, err := time.ParseDuration(accessExpiryStr)
	if err != nil {
		panic("JWT_ACCESS_EXPIRY is not a valid duration")
	}

	refreshExpiry, err := time.ParseDuration(refreshExpiryStr)
	if err != nil {
		panic("JWT_REFRESH_EXPIRY is not a valid duration")
	}

	return &authGenerateTokenUsecase{
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
	}
}
