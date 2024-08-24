package usecase

import (
	"github.com/TiveCS/sync-expense/api/entities"
	"github.com/TiveCS/sync-expense/api/repositories"
)

type TransactionGetManyUsecase interface {
	Execute(dto *entities.GetTransactionsDTO) ([]entities.Transaction, error)
}

type transactionGetManyUsecase struct {
	transactionRepo repositories.TransactionRepository
}

// Execute implements TransactionGetManyUsecase.
func (u *transactionGetManyUsecase) Execute(dto *entities.GetTransactionsDTO) ([]entities.Transaction, error) {
	transactions, err := u.transactionRepo.FindByAccountID(dto)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func NewTransactionGetManyUsecase(tr repositories.TransactionRepository) TransactionGetManyUsecase {
	return &transactionGetManyUsecase{
		transactionRepo: tr,
	}
}
