package auth

import (
	"numtostr/gotodo/app/todo"

	"gorm.io/gorm"
)

// User struct defines the user
type User struct {
	gorm.Model
	Name     string
	Email    string      `gorm:"uniqueIndex;not null"`
	Password string      `gorm:"not null"`
	Todos    []todo.Todo `gorm:"foreignKey:User"`
}

// LoginDTO defined the /login payload
type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"password"`
}

// SignupDTO defined the /login payload
type SignupDTO struct {
	LoginDTO
	Name string `json:"name" validate:"required,min=3"`
}

// UserRes todo
type UserRes struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// AccessRes todo
type AccessRes struct {
	Token string `json:"token"`
}

// Response todo
type Response struct {
	User *UserRes   `json:"user"`
	Auth *AccessRes `json:"auth"`
}
