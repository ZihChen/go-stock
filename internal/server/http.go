package server

import (
	"github.com/gin-gonic/gin"
	"go-stock/router"
)

func Run() {
	var r *gin.Engine
	r = gin.New()

	router.SetupRouter(r)

	router.RouteProvider(r)

	r.Run(":8080")
}
