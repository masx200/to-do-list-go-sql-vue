package models

import (
	// "encoding/json"

	"gorm.io/gorm"
)

type ToDoItem struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`

	Finished bool `gorm:"not null" json:"finished"`

	ID uint `gorm:"primarykey" json:"id"`

	Author string `gorm:"not null,index" json:"author"`
}

// func (t *ToDoItem) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(map[string]any{
// 		"id":       t.ID,
// 		"finished": t.Finished,
// 		"content":  t.Content})
// }
