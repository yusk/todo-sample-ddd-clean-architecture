package database

import (
	"github.com/jinzhu/gorm"
	"github.com/yusk/todo-sample-ddd-clean-architecture/adapter/database/dao"
	"github.com/yusk/todo-sample-ddd-clean-architecture/domain/entity"
	"github.com/yusk/todo-sample-ddd-clean-architecture/domain/repository"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return UserRepositoryImpl{
		db: db,
	}
}

func (r UserRepositoryImpl) Get(id uint) (*entity.User, error) {
	var _user dao.User

	res := r.db.Where(entity.User{ID: id}).First(&_user)
	if len(res.GetErrors()) > 0 {
		return nil, res.GetErrors()[0]
	}

	user := _user.ToUser()
	return &user, nil
}

func (r UserRepositoryImpl) GetByEmail(email string) (*entity.User, error) {
	var _user dao.User

	res := r.db.Where(entity.User{Email: email}).First(&_user)
	if len(res.GetErrors()) > 0 {
		return nil, res.GetErrors()[0]
	}

	user := _user.ToUser()
	return &user, nil
}

func (r UserRepositoryImpl) AttachTodos(user *entity.User) (*entity.User, error) {
	var _todos []*dao.Todo

	res := r.db.Where(dao.Todo{UserID: user.ID}).Find(&_todos)
	if len(res.GetErrors()) > 0 {
		return nil, res.GetErrors()[0]
	}

	var todos []*entity.Todo

	for _, _todo := range _todos {
		todo := _todo.ToTodo()
		todos = append(todos, &todo)
	}

	user.Todos = todos

	return user, nil
}

func (r UserRepositoryImpl) GetWithTodos(id uint) (*entity.User, error) {
	user, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	user, err = r.AttachTodos(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r UserRepositoryImpl) Create(email string, password string) (*entity.User, error) {
	_user := dao.User{Email: email, Password: password}

	tx := r.db.Begin()

	res := tx.Create(&_user)
	if len(res.GetErrors()) > 0 {
		tx.Rollback()
		return nil, res.GetErrors()[0]
	}

	tx.Commit()

	user := _user.ToUser()
	return &user, nil
}

func (r UserRepositoryImpl) AddToDo(user *entity.User, title string, content string) (*entity.User, *entity.Todo, error) {
	_todo := dao.Todo{UserID: user.ID, Title: title, Content: content}

	tx := r.db.Begin()

	res := tx.Create(&_todo)
	if len(res.GetErrors()) > 0 {
		tx.Rollback()
		return nil, nil, res.GetErrors()[0]
	}

	tx.Commit()

	todo := _todo.ToTodo()
	user.Todos = append(user.Todos, &todo)

	return user, &todo, nil
}
