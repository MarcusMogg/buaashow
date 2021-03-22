package initialize

import (
	"buaashow/api/course"
	"buaashow/api/experiment"
	"buaashow/api/files"
	"buaashow/api/img"
	"buaashow/api/show"
	"buaashow/api/swagger"
	"buaashow/api/user"
	"buaashow/middleware"

	"github.com/gin-gonic/gin"
)

// Router 初始化路由列表
func Router() *gin.Engine {
	var Router = gin.Default()
	// 设置multipart forms最大内存限制 4MB
	Router.MaxMultipartMemory = 4 << 20

	APIGroup := Router.Group("")

	Router.Use(middleware.Cors()) // 跨域

	user.InitRouter(APIGroup)
	course.InitRouter(APIGroup)
	swagger.InitRouter(APIGroup)
	img.InitRouter(APIGroup)
	experiment.InitRouter(APIGroup)
	show.InitRouter(APIGroup)
	files.InitRouter(APIGroup)
	return Router
}
