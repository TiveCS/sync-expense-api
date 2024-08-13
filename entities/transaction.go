package entities

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
	ID        string              `json:"id" gorm:"primary_key"`
	Amount    float64             `json:"amount"`
	AccountID string              `json:"account_id"`
	Account   Account             `json:"account" gorm:"foreignKey:AccountID"`
	Category  TransactionCategory `json:"category"`
}
