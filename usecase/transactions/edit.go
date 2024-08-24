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

type TransactionEditUsecase interface {
	Execute(dto *entities.EditTransactionDTO) error
}

type transactionEditUsecase struct {
	transactionRepo repositories.TransactionRepository
}

func (u *transactionEditUsecase) Execute(dto *entities.EditTransactionDTO) error {
	newTransaction := &entities.Transaction{
		Amount:     dto.Amount,
		Note:       dto.Note,
		Category:   dto.Category,
		OccurredAt: dto.OccurredAt,
	}

	err := u.transactionRepo.UpdateById(dto.TransactionID, newTransaction)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, exceptions.TransactionNotFound)
		}
		return err
	}

	return nil
}

func NewTransactionEditUsecase(tr repositories.TransactionRepository) TransactionEditUsecase {
	return &transactionEditUsecase{
		transactionRepo: tr,
	}
}
