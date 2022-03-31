package stock

import (
	"go-stock/app/global/helper"
	"go-stock/app/global/request"
	"go-stock/app/model"
)

func (r *repo) CreateByMap(req request.CreateStockOption) (err error) {
	m, _ := helper.StructToMap(req)
	db, err := r.DB.DBConnect()
	if err != nil {
		return
	}
	stock := model.Stock{}

	tx := db.Begin()

	if err = tx.Where("symbol = ?", m["symbol"]).Assign(m).FirstOrCreate(&stock).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return nil
}
