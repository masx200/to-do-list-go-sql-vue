package todoitem

import (
	"gorm.io/gorm"
)

type ToDoItem struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`

	Finished bool `gorm:"not null" json:"finished"`

	ID uint `gorm:"primarykey" json:"id"`
}
