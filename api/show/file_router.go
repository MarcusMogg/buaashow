package show

import "github.com/gin-gonic/gin"

// InitRouter 初始化img 路由
func InitRouter(router *gin.RouterGroup) {
	rg := router.Group("file")
	{
		rg.POST("", Upload)
		rg.GET("/:name", Download)
	}
	sg := router.Group("show")
	{
		// TODO: 中间件或者啥判断可见性
		sg.GET("/:eid/:gid/*filepath", Show())
	}
}
