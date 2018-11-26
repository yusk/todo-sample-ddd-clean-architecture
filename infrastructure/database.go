package infrastructure

import (
	"fmt"

	session "github.com/daisuke310vvv/echo-session"
	"github.com/jinzhu/gorm"
	"github.com/yusk/todo-sample-ddd-clean-architecture/adapter/database/dao"
)

func Sqlite() (*gorm.DB, error) {
	return gorm.Open("sqlite3", "db.sqlite3")
}

func Migrate(db *gorm.DB) error {
	if res := db.AutoMigrate(
		&dao.User{},
		&dao.Todo{},
	); len(res.GetErrors()) > 0 {
		return res.GetErrors()[0]
	}
	if res := db.Model(&dao.User{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE"); len(res.GetErrors()) > 0 {
		fmt.Println(res.GetErrors()[0].Error())
	}
	return nil
}

func RedisStore() (session.RedisStore, error) {
	return session.NewRedisStore(32, "tcp", "localhost:6379", "", []byte("secret"))
}
