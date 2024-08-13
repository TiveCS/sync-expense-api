package usecase

import (
	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/TiveCS/sync-expense/api/repositories"
	"github.com/matthewhartstonge/argon2"
	"github.com/nrednav/cuid2"
)

type authRegisterUsecase struct {
	ur    repositories.UserRepository
	argon argon2.Config
}

type AuthRegisterUsecase interface {
	Execute(newUser *entities.NewUser) (*map[string]interface{}, error)
}

func (u *authRegisterUsecase) Execute(newUser *entities.NewUser) (*map[string]interface{}, error) {
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
		return nil, err
	}

	return &map[string]interface{}{
		"id": user.ID,
	}, nil
}

func NewAuthRegisterUsecase(ur repositories.UserRepository) AuthRegisterUsecase {
	return &authRegisterUsecase{
		ur:    ur,
		argon: argon2.DefaultConfig(),
	}
}
