package main

import (
	"buaashow/global"
	"buaashow/initialize"
	"fmt"
)

// @title buaashow
// @version 1.0
// @description buaashow is a sample RESTful api server.
// @contact.name Mogg
// @contact.url https://github.com/MarcusMogg
// @BasePath /
func main() {
	initialize.Mysql()
	initialize.DBTables()
	runServer()
}

func runServer() {
	Router := initialize.Router()
	Router.Static("source", "./source")

	address := fmt.Sprintf(":%d", global.GConfig.Port)
	Router.Run(address)
}
