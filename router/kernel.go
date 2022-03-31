package router

import (
	"github.com/gin-gonic/gin"
	"go-stock/app/handler/daily_record"
	"go-stock/app/handler/stock"
	"go-stock/app/handler/user"
	"go-stock/app/middleware/jwt"
	"go-stock/app/middleware/logwriter"
)

func RouteProvider(r *gin.Engine) {
	stockHandler := stock.NewsHandler()
	dailyRecordHandler := daily_record.NewsHandler()
	userHandler := user.NewsHandler()

	api := r.Group("/api", logwriter.RequestLog)
	{
		user := api.Group("user")
		{
			user.POST("/create", userHandler.CreateUser)
			user.POST("/login", userHandler.Login)
		}

		stock := api.Group("stock", jwt.VerifyToken)
		{
			stock.GET("/list", stockHandler.ReadStock)
			stock.POST("/create", stockHandler.CreateStock)
		}

		dailyRecord := api.Group("daily_record")
		{
			dailyRecord.POST("/create", dailyRecordHandler.CreateDailyRecord)
		}
	}
}
