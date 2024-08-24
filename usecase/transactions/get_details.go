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

type TransactionGetDetailsUsecase interface {
	Execute(transactionID string) (*entities.Transaction, error)
}

type transactionGetDetailsUsecase struct {
	transactionRepo repositories.TransactionRepository
}

func (u *transactionGetDetailsUsecase) Execute(transactionID string) (*entities.Transaction, error) {
	transaction, err := u.transactionRepo.FindByID(transactionID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound, exceptions.TransactionNotFound)
		}
		return nil, err
	}

	return transaction, nil
}

func NewTransactionGetDetailsUsecase(tr repositories.TransactionRepository) TransactionGetDetailsUsecase {
	return &transactionGetDetailsUsecase{
		transactionRepo: tr,
	}
}
