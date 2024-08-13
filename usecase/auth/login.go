package usecase

import (
	"errors"
	"net/http"

	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/TiveCS/sync-expense/api/exceptions"
	"github.com/TiveCS/sync-expense/api/repositories"
	"github.com/labstack/echo/v4"
	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

type AuthLoginUsecase interface {
	Execute(loginUser *entities.LoginUserDTO) (string, string, error)
}

type authLoginUsecase struct {
	userRepo             repositories.UserRepository
	generateTokenUsecase AuthGenerateTokenUsecase
}

// Execute implements AuthLoginUsecase.
func (u *authLoginUsecase) Execute(loginUser *entities.LoginUserDTO) (string, string, error) {
	user, err := u.userRepo.FindByEmail(loginUser.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", echo.NewHTTPError(http.StatusUnauthorized, exceptions.AuthInvalidCredentials)
		}

		return "", "", err
	}

	matches, err := argon2.VerifyEncoded([]byte(loginUser.Password), []byte(user.Password))

	if err != nil {
		return "", "", err
	}

	if !matches {
		return "", "", echo.NewHTTPError(http.StatusUnauthorized, exceptions.AuthInvalidCredentials)
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
