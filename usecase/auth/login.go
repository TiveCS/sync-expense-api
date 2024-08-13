package usecase

import (
	"net/http"

	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/TiveCS/sync-expense/api/repositories"
	"github.com/labstack/echo/v4"
	"github.com/matthewhartstonge/argon2"
)

type AuthLoginUsecase interface {
	Execute(loginUser *entities.LoginUser) (string, string, error)
}

type authLoginUsecase struct {
	userRepo             repositories.UserRepository
	generateTokenUsecase AuthGenerateTokenUsecase
}

// Execute implements AuthLoginUsecase.
func (u *authLoginUsecase) Execute(loginUser *entities.LoginUser) (string, string, error) {
	user, err := u.userRepo.FindByEmail(loginUser.Email)
	if err != nil {
		return "", "", err
	}

	if user == nil {
		echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		return "", "", nil
	}

	matches, err := argon2.VerifyEncoded([]byte(loginUser.Password), []byte(user.Password))

	if err != nil {
		return "", "", err
	}

	if !matches {
		echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		return "", "", nil
	}

	accessToken, refreshToken, err := u.generateTokenUsecase.Execute(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func NewAuthLoginUsecase(ur repositories.UserRepository, gtu AuthGenerateTokenUsecase) AuthLoginUsecase {
	return &authLoginUsecase{
		userRepo:             ur,
		generateTokenUsecase: gtu,
	}
}
