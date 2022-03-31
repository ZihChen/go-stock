package user

import "go-stock/app/business/user"

type Handler struct {
	userBus user.Interface
}

func NewsHandler() *Handler {
	return &Handler{
		userBus: user.NewBusiness(),
	}
}
