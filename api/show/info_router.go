package show

import "github.com/gin-gonic/gin"

//InitRouter 初始化路由组
func InitRouter(Router *gin.RouterGroup) {
	sr := Router.Group("show")
	{
		// 简略信息
		sr.GET("search", Search)
		// 详细介绍
		sr.GET("readme/:showid", Readme)
		// TODO: 中间件或者啥判断可见性
		sr.GET("x/:showid/*filepath", Show())
	}
}
