package repositories

import (
	"github.com/TiveCS/sync-expense/api/entities"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *entities.Transaction) error
	FindByID(id string) (*entities.Transaction, error)
	FindByAccountID(accountID string) ([]entities.Transaction, error)
	DeleteById(id string) error
}

type transactionRepository struct {
	db *gorm.DB
}

// Create implements TransactionRepository.
func (t *transactionRepository) Create(transaction *entities.Transaction) error {
	result := t.db.Create(transaction)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteById implements TransactionRepository.
func (t *transactionRepository) DeleteById(id string) error {
	result := t.db.Where("id = ?", id).Delete(&entities.Transaction{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindByAccountID implements TransactionRepository.
func (t *transactionRepository) FindByAccountID(accountID string) ([]entities.Transaction, error) {
	var transactions []entities.Transaction

	result := t.db.Where("account_id = ?", accountID).Find(&transactions)

	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}

// FindByID implements TransactionRepository.
func (t *transactionRepository) FindByID(id string) (*entities.Transaction, error) {
	var transaction entities.Transaction

	result := t.db.Where("id = ?", id).First(&transaction)

	if result.Error != nil {
		return nil, result.Error
	}

	return &transaction, nil
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
