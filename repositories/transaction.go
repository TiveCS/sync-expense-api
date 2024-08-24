package repositories

import (
	"github.com/TiveCS/sync-expense/api/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepository interface {
	Create(transaction *entities.Transaction) (string, error)
	FindByID(id string) (*entities.Transaction, error)
	FindByAccountID(dto *entities.GetTransactionsDTO) ([]entities.Transaction, error)
	DeleteById(id string) error
	UpdateById(id string, transaction *entities.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

// UpdateById implements TransactionRepository.
func (t *transactionRepository) UpdateById(id string, transaction *entities.Transaction) error {
	result := t.db.Model(&entities.Transaction{}).Where("id = ?", id).Updates(transaction)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Create implements TransactionRepository.
func (t *transactionRepository) Create(transaction *entities.Transaction) (string, error) {
	result := t.db.Create(transaction)

	if result.Error != nil {
		return "", result.Error
	}

	return transaction.ID, nil
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
func (t *transactionRepository) FindByAccountID(dto *entities.GetTransactionsDTO) ([]entities.Transaction, error) {
	var transactions []entities.Transaction

	var limit int = -1
	if dto.Limit != 0 {
		limit = int(dto.Limit)
	}

	var sortDir string = "desc"
	if dto.SortDir != "" {
		sortDir = dto.SortDir
	}

	var sortBy string = "occurred_at"
	if dto.SortBy != "" {
		sortBy = dto.SortBy
	}

	result := t.db.Where("account_id = ?", dto.AccountID).Order(clause.OrderByColumn{
		Column: clause.Column{Name: sortBy}, Desc: sortDir == "desc",
	}).Limit(limit).Find(&transactions)

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
