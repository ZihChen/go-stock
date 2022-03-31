package user

import (
	"github.com/gin-gonic/gin"
	"go-stock/app/global/helper"
	"go-stock/app/global/request"
	"go-stock/app/http/response"
	"net/http"
)

func (h *Handler) CreateUser(c *gin.Context) {
	option := request.CreateUserOption{}
	if err := c.ShouldBindJSON(&option); err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}

	if err := helper.ValidateParams(option); err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}

	if err := h.userBus.CreateUser(option); err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

func (h *Handler) Login(c *gin.Context) {
	option := request.UserLogin{}
	if err := c.ShouldBindJSON(&option); err != nil {
		response.WrapContext(c).Error(1001, err.Error())
		return
	}

	token, err := h.userBus.Login(option)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		response.WrapContext(c).Error(1002, err.Error())
		return
	}

	response.WrapContext(c).Success(token, "success")
}
