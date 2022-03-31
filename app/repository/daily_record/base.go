package daily_record

import (
	"go-stock/internal/database"
	"sync"
)

var singleton *repo
var once sync.Once

type Interface interface {
	CreateByMap(req map[string]interface{}) (err error)
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
