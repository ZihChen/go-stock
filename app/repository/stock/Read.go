package stock

import (
	"go-stock/app/model"
	"gorm.io/gorm"
)

func (r *repo) ReadStocks() (stocks []model.Stock, err error) {
	db, err := r.DB.DBConnect()
	if err != nil {
		return
	}

	if err1 := db.Preload("DailyRecord", func(db *gorm.DB) *gorm.DB {
		return db.Order("daily_records.date DESC").Where("daily_records.id = 2")
	}).Find(&stocks).Error; err1 != nil {
		err = err1
		return stocks, err
	}

	return
}
