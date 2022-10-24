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
	DeletedAt DeletedAt `gorm:"index" json:"-"  form:"-"`
	Content   string    `gorm:"not null" json:"content"  form:"content"`

	Finished bool `gorm:"not null;index" json:"finished" form:"finished" `

	ID uint `gorm:"primarykey" json:"id" form:"id"`

	Author string `gorm:"not null;index" json:"author" form:"author"`
}
