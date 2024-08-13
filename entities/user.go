package entities

type User struct {
	ID       string `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-" gorm:"not null"`
}

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
