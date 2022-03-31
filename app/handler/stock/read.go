package stock

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ReadStock 讀取Stock
// @Summary 讀取Stock
// @Tags Stock
// @Produce json
// @Security BearerAuth
// @Success 200 string string 成功後返回的值
// @Router /api/stock/list [GET]
func (h *Handler) ReadStock(c *gin.Context) {
	list, err := h.stockBus.ReadStockList()
	if err != nil {
		c.JSON(http.StatusOK, "")
		return
	}

	c.JSON(http.StatusOK, list)
	return
}
