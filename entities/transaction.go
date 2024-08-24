package entities

import (
	"time"

	"github.com/TiveCS/sync-expense/api/enums"
)

type Transaction struct {
	ID         string                    `json:"id" gorm:"primary_key"`
	Amount     float64                   `json:"amount" gorm:"not null;check:amount >= 0.0"`
	Note       string                    `json:"note"`
	AccountID  string                    `json:"account_id" gorm:"index;not null"`
	Account    *Account                  `json:"account" gorm:"foreignKey:AccountID"`
	Category   enums.TransactionCategory `json:"category" gorm:"not null"`
	OccurredAt time.Time                 `json:"occurred_at" gorm:"not null"`
	CreatedAt  time.Time                 `json:"created_at"`
	UpdatedAt  time.Time                 `json:"updated_at"`
}

type NewTransactionDTO struct {
	AccountID  string                    `json:"account_id" validate:"required"`
	Amount     float64                   `json:"amount" validate:"required"`
	Note       string                    `json:"note"`
	Category   enums.TransactionCategory `json:"category" validate:"required,transaction_category"`
	OccurredAt time.Time                 `json:"occurred_at" validate:"required"`
}

type EditTransactionDTO struct {
	TransactionID string                    `json:"transaction_id" validate:"required"`
	Amount        float64                   `json:"amount" validate:"required"`
	Note          string                    `json:"note"`
	Category      enums.TransactionCategory `json:"category" validate:"required,transaction_category"`
	OccurredAt    time.Time                 `json:"occurred_at" validate:"required"`
}

type GetTransactionsDTO struct {
	AccountID string `query:"account_id" validate:"required"`
	Limit     uint16 `query:"limit" validate:"omitempty,gte=1"`
	Cursor    string `query:"cursor"`
	SortDir   string `query:"sort_dir" validate:"omitempty,oneof=asc desc"`
	SortBy    string `query:"sort_by" validate:"omitempty,oneof=occurred_at amount"`
}
