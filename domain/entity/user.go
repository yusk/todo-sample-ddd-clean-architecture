package entity

import "github.com/pkg/errors"

type User struct {
	ID       uint    `json:"id"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Todos    []*Todo `json:"todos"`
}

func (u User) IsAuthenticated(password string) bool {
	return u.Password == password
}

func (u User) GetTodo(todoID uint) (*Todo, error) {
	for _, todo := range u.Todos {
		if todo.ID == todoID {
			return todo, nil
		}
	}
	return nil, errors.Errorf("Todo Not Found")
}
