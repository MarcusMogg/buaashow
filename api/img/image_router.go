package img

import "github.com/gin-gonic/gin"

// InitRouter 初始化img 路由
func InitRouter(router *gin.RouterGroup) {
	rg := router.Group("img")
	{
		rg.POST("", Upload)
		rg.GET("/:name", Download)
	}
}
