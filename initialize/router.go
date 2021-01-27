package initialize

import (
	"buaashow/api/course"
	"buaashow/api/swagger"
	"buaashow/api/user"
	"buaashow/middleware"

	"github.com/gin-gonic/gin"
)

// Router 初始化路由列表
func Router() *gin.Engine {
	var Router = gin.Default()

	APIGroup := Router.Group("")

	Router.Use(middleware.Cors()) // 跨域

	user.InitRouter(APIGroup)
	course.InitRouter(APIGroup)
	swagger.InitRouter(APIGroup)

	return Router
}
