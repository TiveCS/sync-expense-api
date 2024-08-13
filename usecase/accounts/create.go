package usecase

import (
	"errors"
	"net/http"

	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/TiveCS/sync-expense/api/exceptions"
	"github.com/TiveCS/sync-expense/api/repositories"
	"github.com/labstack/echo/v4"
	"github.com/nrednav/cuid2"
	"gorm.io/gorm"
)

type AccountCreateUsecase interface {
	Execute(dto *entities.NewAccountDTO) (*entities.Account, error)
}

type accountCreateUsecase struct {
	accountRepo repositories.AccountRepository
}

// Execute implements AccountCreateUsecase.
func (u *accountCreateUsecase) Execute(dto *entities.NewAccountDTO) (*entities.Account, error) {
	account := &entities.Account{
		ID:      cuid2.Generate(),
		OwnerID: dto.OwnerID,
		Name:    dto.Name,
	}

	err := u.accountRepo.Create(account)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, echo.NewHTTPError(http.StatusConflict, exceptions.AccountOnePerUser)
		}
		return nil, err
	}

	return account, nil
}

func NewAccountCreateUsecase(ar repositories.AccountRepository) AccountCreateUsecase {
	return &accountCreateUsecase{
		accountRepo: ar,
	}
}
