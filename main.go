package main

import (
	_ "github.com/mattn/go-sqlite3"

	"github.com/yusk/todo-sample-ddd-clean-architecture/adapter/router"
	"github.com/yusk/todo-sample-ddd-clean-architecture/infrastructure"
)

func main() {
	db, err := infrastructure.Sqlite()
	if err != nil {
		panic(err)
	}
	err = infrastructure.Migrate(db)
	if err != nil {
		panic(err)
	}
	store, err := infrastructure.RedisStore()
	if err != nil {
		panic(err)
	}

	router.Router(db, store)
}
