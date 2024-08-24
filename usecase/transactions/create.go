package usecase

import (
	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/TiveCS/sync-expense/api/repositories"
	"github.com/nrednav/cuid2"
)

type TransactionCreateUsecase interface {
	Execute(dto *entities.NewTransactionDTO) (string, error)
}

type transactionCreateUsecase struct {
	transactionRepo repositories.TransactionRepository
}

// Execute implements TransactionCreateUsecase.
func (u *transactionCreateUsecase) Execute(dto *entities.NewTransactionDTO) (string, error) {
	transaction := &entities.Transaction{
		ID:         cuid2.Generate(),
		AccountID:  dto.AccountID,
		Amount:     dto.Amount,
		Note:       dto.Note,
		OccurredAt: dto.OccurredAt,
		Category:   dto.Category,
	}

	transactionID, err := u.transactionRepo.Create(transaction)
	if err != nil {
		return "", err
	}

	return transactionID, nil
}

func NewTransactionCreateUsecase(tr repositories.TransactionRepository) TransactionCreateUsecase {
	return &transactionCreateUsecase{
		transactionRepo: tr,
	}
}
