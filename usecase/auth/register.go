package usecase

import (
	"errors"
	"net/http"

	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/TiveCS/sync-expense/api/exceptions"
	"github.com/TiveCS/sync-expense/api/repositories"
	usecase "github.com/TiveCS/sync-expense/api/usecase/accounts"
	"github.com/labstack/echo/v4"
	"github.com/matthewhartstonge/argon2"
	"github.com/nrednav/cuid2"
	"gorm.io/gorm"
)

type authRegisterUsecase struct {
	ur                   repositories.UserRepository
	accountCreateUsecase usecase.AccountCreateUsecase
	argon                argon2.Config
}

type AuthRegisterUsecase interface {
	Execute(newUser *entities.NewUserDTO) (*map[string]interface{}, error)
}

func (u *authRegisterUsecase) Execute(newUser *entities.NewUserDTO) (*map[string]interface{}, error) {
	hashedPassword, err := u.argon.HashEncoded([]byte(newUser.Password))

	if err != nil {
		return nil, err
	}

	user := &entities.User{
		ID:       cuid2.Generate(),
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: string(hashedPassword),
	}

	err = u.ur.Create(user)

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, exceptions.AuthDuplicateCredentials)
		}

		return nil, err
	}

	account, err := u.accountCreateUsecase.Execute(&entities.NewAccountDTO{
		OwnerID: user.ID,
		Name:    "Personal",
	})

	if err != nil {
		return nil, err
	}

	return &map[string]interface{}{
		"id":         user.ID,
		"account_id": account.ID,
	}, nil
}

func NewAuthRegisterUsecase(ur repositories.UserRepository, acu usecase.AccountCreateUsecase) AuthRegisterUsecase {
	return &authRegisterUsecase{
		ur:                   ur,
		argon:                argon2.DefaultConfig(),
		accountCreateUsecase: acu,
	}
}
