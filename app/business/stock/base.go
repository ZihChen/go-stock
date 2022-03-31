package stock

import (
	"go-stock/app/global/request"
	"go-stock/app/global/structs"
	"go-stock/app/repository/stock"
	"sync"
)

var singleton *business
var once sync.Once

type Interface interface {
	CreateStock(option request.CreateStockOption) (err error)
	ReadStockList() (stockList []structs.StockObj, err error)
}

type business struct {
	stock stock.Interface
}

func NewBusiness() Interface {
	once.Do(func() {
		singleton = &business{
			stock: stock.NewRepo(),
		}
	})
	return singleton
}
