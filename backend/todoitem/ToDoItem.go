package todoitem

import (
	"gorm.io/gorm"
)

type ToDoItem struct {
	gorm.Model
	Content string ` gorm:"not null"`

	Finished bool ` gorm:"not null"`
}
