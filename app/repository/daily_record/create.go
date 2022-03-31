package daily_record

import (
	"go-stock/app/model"
)

func (r *repo) CreateByMap(req map[string]interface{}) (err error) {
	db, err := r.DB.DBConnect()
	if err != nil {
		return
	}

	dailyRecord := model.DailyRecord{}

	tx := db.Begin()

	if err = tx.Model(&dailyRecord).Create(req).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}
