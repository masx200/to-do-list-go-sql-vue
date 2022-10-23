package models

import (
	// "encoding/json"

	"time"

	"gorm.io/gorm"
)

type DeletedAt = gorm.DeletedAt
type ToDoItem struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt DeletedAt `gorm:"index" json:"deleted_at"`
	Content   string    `gorm:"not null" json:"content"`

	Finished bool `gorm:"not null" json:"finished" `

	ID uint `gorm:"primarykey" json:"id"`

	Author string `gorm:"not null;index" json:"author" form:"author"`
}

