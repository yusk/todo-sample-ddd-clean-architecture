package database

import (
	"fmt"

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
		fmt.Println(res.GetErrors())
		return nil, res.GetErrors()[0]
	}

	user := _user.ToUser()
	return &user, nil
}

func (r UserRepositoryImpl) GetByEmail(email string) (*entity.User, error) {
	var _user dao.User

	res := r.db.Where(entity.User{Email: email}).First(&_user)
	if len(res.GetErrors()) > 0 {
		fmt.Println(res.GetErrors())
		return nil, res.GetErrors()[0]
	}

	user := _user.ToUser()
	return &user, nil
}

func (r UserRepositoryImpl) Create(email string, password string) (*entity.User, error) {
	_user := dao.User{Email: email, Password: password}

	res := r.db.Create(&_user)
	if len(res.GetErrors()) > 0 {
		r.db.Rollback()
		return nil, res.GetErrors()[0]
	}

	r.db.Commit()

	user := _user.ToUser()
	return &user, nil
}
