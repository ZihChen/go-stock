package stock

import "go-stock/app/global/request"

func (b *business) CreateStock(option request.CreateStockOption) (err error) {

	b.stock.CreateByMap(option)

	return
}
