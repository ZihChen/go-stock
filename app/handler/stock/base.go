package stock

import (
	"go-stock/app/business/stock"
)

type Handler struct {
	stockBus stock.Interface
}

func NewsHandler() *Handler {
	return &Handler{
		stockBus: stock.NewBusiness(),
	}
}
