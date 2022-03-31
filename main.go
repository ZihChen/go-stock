package main

import (
	"go-stock/app/global"
	"go-stock/internal/database"
	"go-stock/internal/server"

	_ "go-stock/docs"
)

func init() {
	global.Start()
	db := database.NewDBConnect()
	db.CheckTable()
}

// @title Go-Stock swagger

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	server.Run()
}
