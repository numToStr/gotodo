package todo

import (
	"gorm.io/gorm"
)

// Todo struct defines the Todo Model
type Todo struct {
	gorm.Model
	Task      string `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
	User      *uint  `gorm:"not null"`
}

// Response struct contains the todo field which should be returned in a response
type Response struct {
	ID        uint   `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

// CreateDTO struct defines the /todo/create payload
type CreateDTO struct {
	Task string `json:"task" validate:"required,min=3,max=150"`
}

// CreateRes struct defines the /todo/create response
type CreateRes struct {
	Todo *Response `json:"todo"`
}

// ListRes defines the todos list
type ListRes struct {
	Todos *[]Response `json:"todos"`
}
