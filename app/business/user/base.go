package user

import (
	"go-stock/app/global/request"
	"go-stock/app/repository/user"
	"go-stock/app/service/jwtservice"
	"sync"
)

var singleton *business
var once sync.Once

type Interface interface {
	CreateUser(option request.CreateUserOption) (err error)
	Login(option request.UserLogin) (token string, err error)
}

type business struct {
	user user.Interface
	jwt  jwtservice.Interface
}

func NewBusiness() Interface {
	once.Do(func() {
		singleton = &business{
			user: user.NewRepo(),
			jwt:  jwtservice.NewService(),
		}
	})
	return singleton
}
