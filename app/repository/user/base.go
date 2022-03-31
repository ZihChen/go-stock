package user

import (
	"go-stock/app/model"
	"go-stock/internal/database"
	"sync"
)

var singleton *repo
var once sync.Once

type Interface interface {
	InsertUser(fields map[string]interface{}) (err error)
	GetUserByUsername(username string) (user model.User, err error)
}

type repo struct {
	DB database.DBInterface
}

func NewRepo() Interface {
	once.Do(func() {
		singleton = &repo{
			DB: database.NewDBConnect(),
		}
	})
	return singleton
}
