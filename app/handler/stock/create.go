package stock

import (
	"github.com/gin-gonic/gin"
	"go-stock/app/global/helper"
	"go-stock/app/global/request"
	"net/http"
)

func (h *Handler) CreateStock(c *gin.Context) {
	reqOption := request.CreateStockOption{}
	if err := c.ShouldBindJSON(&reqOption); err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}

	if err := helper.ValidateParams(reqOption); err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}

	h.stockBus.CreateStock(reqOption)

}
