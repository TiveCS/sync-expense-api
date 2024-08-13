package repositories

import (
	"github.com/TiveCS/sync-expense/api/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	FindByID(id string) (*entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// FindByID implements UserRepository.
func (r *userRepository) FindByID(id string) (*entities.User, error) {
	var user entities.User

	err := r.db.Where("id = ?", id).First(&user)

	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}

// Create implements UserRepository.
func (r *userRepository) Create(user *entities.User) error {
	err := r.db.Create(user)

	if err.Error != nil {
		return err.Error
	}

	return nil
}

// FindByEmail implements UserRepository.
func (r *userRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	err := r.db.Where("email = ?", email).First(&user)

	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
