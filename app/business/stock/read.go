package stock

import "go-stock/app/global/structs"

func (b *business) ReadStockList() (stockList []structs.StockObj, err error) {
	stocks, err := b.stock.ReadStocks()
	if err != nil {
		return
	}

	for i := range stocks {
		stockObj := structs.StockObj{}
		stockObj.ID = stocks[i].ID
		stockObj.Name = stocks[i].Name
		stockObj.Symbol = stocks[i].Symbol

		stockList = append(stockList, stockObj)
	}

	return
}
