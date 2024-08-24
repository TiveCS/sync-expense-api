package entities

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID           string         `json:"id" gorm:"primary_key"`
	Name         string         `json:"name" gorm:"not null"`
	OwnerID      string         `json:"owner_id" gorm:"index;not null;unique"`
	Owner        *User          `json:"owner" gorm:"foreignKey:OwnerID"`
	Transactions []Transaction  `json:"transactions"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type NewAccountDTO struct {
	OwnerID string
	Name    string `json:"name" validate:"required"`
}

type EditAccountDTO struct {
	AccountID string `param:"account_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type GetManyAccountsDTO struct {
	OwnerID string `query:"owner_id" validate:"required"`
}
