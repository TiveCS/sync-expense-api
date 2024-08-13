package usecase

import (
	"errors"
	"net/http"

	"github.com/TiveCS/sync-expense/api/exceptions"
	"github.com/TiveCS/sync-expense/api/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AccountDeleteUsecase interface {
	Execute(accountID string) error
}

type accountDeleteUsecase struct {
	accountRepo repositories.AccountRepository
}

// Execute implements AccountDeleteUsecase.
func (u *accountDeleteUsecase) Execute(accountID string) error {
	err := u.accountRepo.DeleteByID(accountID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			echo.NewHTTPError(http.StatusNotFound, exceptions.AccountNotFound)
		}

		return err
	}

	return nil
}

func NewAccountDeleteUsecase(ar repositories.AccountRepository) AccountDeleteUsecase {
	return &accountDeleteUsecase{
		accountRepo: ar,
	}
}
