package router

import (
	"buaashow/api"
	"buaashow/middleware"

	"github.com/gin-gonic/gin"
)

//InitUserRouter 初始化user路由组
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", api.LoginByPwd)
		UserRouter.POST("verify", api.LoginByTicket)
		UserRouter.GET("userinfo", middleware.JWTAuth(), api.GetUserInfo)
		UserRouter.GET("info", api.GetUserInfoByID)
	}
}
