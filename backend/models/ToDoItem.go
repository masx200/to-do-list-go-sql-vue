package models

import (
	"time"

	"gorm.io/gorm"
)

type DeletedAt = gorm.DeletedAt
type ToDoItem struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt DeletedAt `gorm:"index" json:"-"  `
	Content   string    `gorm:"not null" json:"content"  form:"content"`

	Completed bool `gorm:"not null;index" json:"completed" form:"completed" `

	ID uint `gorm:"primarykey" json:"id" form:"id"`

	Author string `gorm:"not null;index" json:"author" form:"author"`
}
