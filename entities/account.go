package entities

type Account struct {
	ID           string        `json:"id" gorm:"primary_key"`
	OwnerID      string        `json:"owner_id" gorm:"index;not null;unique"`
	Owner        User          `json:"owner" gorm:"foreignKey:OwnerID"`
	Transactions []Transaction `json:"transactions"`
}
