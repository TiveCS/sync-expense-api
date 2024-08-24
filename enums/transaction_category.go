package enums

import (
	"github.com/go-playground/validator/v10"
)

type TransactionCategory uint16

const (
	Unknown TransactionCategory = iota
	Food
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

func IsValidTransactionCategory(input uint16) bool {
	return input >= uint16(Food) && input <= uint16(Other)
}

func ValidateTransactionCategory(fl validator.FieldLevel) bool {
	categoryIndex := fl.Field().Uint()

	return IsValidTransactionCategory(uint16(categoryIndex))
}
