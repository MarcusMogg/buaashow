package initialize

import (
	"buaashow/middleware"
	"buaashow/router"

	"github.com/gin-gonic/gin"
)

// Router 初始化路由列表
func Router() *gin.Engine {
	var Router = gin.Default()

	APIGroup := Router.Group("")

	Router.Use(middleware.Cors()) // 跨域

	router.InitSwaggerRouter(APIGroup)
	router.InitUserRouter(APIGroup)

	return Router
}
