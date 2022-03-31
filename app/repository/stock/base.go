package stock

import (
	"go-stock/app/global/request"
	"go-stock/app/model"
	"go-stock/internal/database"
	"sync"
)

var singleton *repo
var once sync.Once

type Interface interface {
	CreateByMap(req request.CreateStockOption) (err error)
	ReadStocks() (stocks []model.Stock, err error)
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
