package dao

import (
	"time"

	"github.com/yusk/todo-sample-ddd-clean-architecture/domain/entity"
)

type Todo struct {
	ID        uint      `json:"id" gorm:"primary_key;unique;index;not null;"`
	UserID    uint      `json:"user_id" gorm:"index;not null;"`
	Title     string    `json:"title" gorm:"not null;"`
	Content   string    `json:"content"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t Todo) ToTodo() entity.Todo {
	return entity.Todo{
		ID:        t.ID,
		Title:     t.Title,
		Content:   t.Content,
		IsDone:    t.IsDone,
		CreatedAt: t.CreatedAt,
	}
}
