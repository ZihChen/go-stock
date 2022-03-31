package user

import (
	"go-stock/app/global/helper"
	"go-stock/app/global/request"
)

func (b *business) CreateUser(option request.CreateUserOption) (err error) {
	m, _ := helper.StructToMap(option)
	err = b.user.InsertUser(m)

	return
}