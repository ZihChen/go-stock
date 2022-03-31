package daily_record

import (
	"github.com/gin-gonic/gin"
	"go-stock/app/global/helper"
	"go-stock/app/global/request"
	"net/http"
)

func (h *Handler) CreateDailyRecord(c *gin.Context) {
	option := request.CreateDailyRecordOption{}
	if err := c.ShouldBindJSON(&option); err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}

	if err := helper.ValidateParams(option); err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}

	h.dailyRecordBus.CreateDailyRecord(option)
}
