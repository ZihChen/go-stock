package user

import "go-stock/app/model"

func (r *repo) InsertUser(fields map[string]interface{}) (err error) {
	db, err := r.DB.DBConnect()
	if err != nil {
		return
	}

	tx := db.Begin()

	if err = tx.Model(&model.User{}).Create(fields).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return nil
}
