package repositories

import (
	"github.com/TiveCS/sync-expense/api/entities"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(account *entities.Account) error
	FindByID(id string) (*entities.Account, error)
	FindByOwnerID(ownerID string) (*entities.Account, error)
	UpdateByID(id string, account *entities.Account) error
	DeleteByID(id string) error
}

type accountRepository struct {
	db *gorm.DB
}

// DeleteByID implements AccountRepository.
func (r *accountRepository) DeleteByID(id string) error {
	result := r.db.Where("id = ?", id).Delete(&entities.Account{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdateByID implements AccountRepository.
func (r *accountRepository) UpdateByID(id string, account *entities.Account) error {
	result := r.db.Model(&entities.Account{}).Where("id = ?", id).Updates(account)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Create implements AccountRepository.
func (r *accountRepository) Create(account *entities.Account) error {
	result := r.db.Create(account)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindByID implements AccountRepository.
func (r *accountRepository) FindByID(id string) (*entities.Account, error) {
	var account entities.Account

	result := r.db.Where("id = ?", id).First(&account)

	if result.Error != nil {
		return nil, result.Error
	}

	return &account, nil
}

// FindByOwnerID implements AccountRepository.
func (r *accountRepository) FindByOwnerID(ownerID string) (*entities.Account, error) {
	var account entities.Account

	result := r.db.Where("owner_id = ?", ownerID).First(&account)

	if result.Error != nil {
		return nil, result.Error
	}

	return &account, nil
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
}
