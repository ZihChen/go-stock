package user

import "go-stock/app/model"

func (r *repo) GetUserByUsername(username string) (user model.User, err error) {
	db, err := r.DB.DBConnect()
	if err != nil {
		return
	}

	if err = db.Where("username = ?", username).Find(&user).Error; err != nil {
		return
	}

	return
}
