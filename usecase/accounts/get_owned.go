package usecase

import (
	"errors"
	"net/http"

	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/TiveCS/sync-expense/api/exceptions"
	"github.com/TiveCS/sync-expense/api/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AccountGetOwnedUsecase interface {
	Execute(ownerID string) (*entities.Account, error)
}

type accountGetOwnedUsecase struct {
	accountRepo repositories.AccountRepository
}

// Execute implements AccountGetOwnedUsecase.
func (u *accountGetOwnedUsecase) Execute(ownerID string) (*entities.Account, error) {
	account, err := u.accountRepo.FindByOwnerID(ownerID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound, exceptions.AccountNotFound)
		}
		return nil, err
	}

	return account, nil
}

func NewAccountGetOwnedUsecase(ar repositories.AccountRepository) AccountGetOwnedUsecase {
	return &accountGetOwnedUsecase{
		accountRepo: ar,
	}
}
