package show

import (
	"buaashow/entity"
	"buaashow/middleware"

	"github.com/gin-gonic/gin"
)

//InitRouter 初始化路由组
func InitRouter(Router *gin.RouterGroup) {
	sr := Router.Group("show")
	{
		// 简略信息
		sr.GET("search", Search)
		sr.GET("readme/:showid", Readme)
		sr.GET("x/:showid/*filepath", Show())

		sr.GET("preview/readme/:showid", middleware.JWTAuth(entity.Student), PreReadme)
		sr.GET("preview/x/:showid/*filepath", middleware.JWTAuth(entity.Student), Preview())
	}
}
