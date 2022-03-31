package user

import (
	"go-stock/app/global/request"
)

func (b *business) Login(option request.UserLogin) (token string, err error) {
	user, err := b.user.GetUserByUsername(option.Username)
	if err != nil {
		return
	}

	if user.Password != option.Password {
		panic("password error")
	}

	token, err = b.jwt.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return
}
