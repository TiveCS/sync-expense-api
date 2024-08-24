package usecase

import (
	"errors"
	"net/http"

	"github.com/TiveCS/sync-expense/api/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TransactionDeleteUsecase interface {
	Execute(transactionID string) error
}

type transactionDeleteUsecase struct {
	transactionRepo repositories.TransactionRepository
}

// Execute implements TransactionDeleteUsecase.
func (t *transactionDeleteUsecase) Execute(transactionID string) error {
	err := t.transactionRepo.DeleteById(transactionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Transaction not found")
		}
		return err
	}

	return nil
}

func NewTransactionDeleteUsecase(tr repositories.TransactionRepository) TransactionDeleteUsecase {
	return &transactionDeleteUsecase{
		transactionRepo: tr,
	}
}
