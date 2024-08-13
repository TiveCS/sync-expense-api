package usecase

import (
	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/TiveCS/sync-expense/api/repositories"
)

type AuthMeUsecase interface {
	Execute(claim *entities.JwtUserClaims) (*entities.User, error)
}

type authMeUsecase struct {
	userRepo repositories.UserRepository
}

// Execute implements AuthMeUsecase.
func (u *authMeUsecase) Execute(claim *entities.JwtUserClaims) (*entities.User, error) {
	user, err := u.userRepo.FindByID(claim.Subject)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewAuthMeUsecase(userRepo repositories.UserRepository) AuthMeUsecase {
	return &authMeUsecase{
		userRepo: userRepo,
	}
}
