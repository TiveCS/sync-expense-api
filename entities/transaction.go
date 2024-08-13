package entities

import (
	"time"
)

type TransactionCategory uint16

const (
	Food TransactionCategory = iota + 1
	Transport
	Education
	Entertainment
	Health
	Donation
	Housing
	Investment
	Utility
	Insurance
	Work
	Other
)

func (c TransactionCategory) String() string {
	return [...]string{"Food", "Transport", "Education", "Entertainment", "Health", "Donation", "Housing", "Investment", "Utility", "Insurance", "Work", "Other"}[c-1]
}

func (c TransactionCategory) Index() uint16 {
	return uint16(c)
}

type Transaction struct {
	ID         string              `json:"id" gorm:"primary_key"`
	Amount     float64             `json:"amount" gorm:"not null;check:amount >= 0.0"`
	Note       string              `json:"note"`
	AccountID  string              `json:"account_id" gorm:"index;not null"`
	Account    Account             `json:"account" gorm:"foreignKey:AccountID"`
	Category   TransactionCategory `json:"category" gorm:"not null"`
	OccurredAt time.Time           `json:"occurred_at" gorm:"not null"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}

type NewTransactionDTO struct {
	AccountID  string              `json:"account_id" validate:"required"`
	Amount     float64             `json:"amount" validate:"required"`
	Note       string              `json:"note"`
	Category   TransactionCategory `json:"category" validate:"required"`
	OccurredAt time.Time           `json:"occurred_at" validate:"required"`
}
