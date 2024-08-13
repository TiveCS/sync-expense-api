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

type AccountEditUsecase interface {
	Execute(dto *entities.EditAccountDTO) error
}

type accountEditUsecase struct {
	accountRepo repositories.AccountRepository
}

// Execute implements AccountEditUsecase.
func (u *accountEditUsecase) Execute(dto *entities.EditAccountDTO) error {
	err := u.accountRepo.UpdateByID(dto.AccountID, &entities.Account{
		Name: dto.Name,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			echo.NewHTTPError(http.StatusNotFound, exceptions.AccountNotFound)
		}

		return err
	}

	return nil
}

func NewAccountEditUsecase(ar repositories.AccountRepository) AccountEditUsecase {
	return &accountEditUsecase{
		accountRepo: ar,
	}
}
