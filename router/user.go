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
		UserRouter.POST("email", middleware.JWTAuth(1), api.UpdateEmail)
		UserRouter.POST("password", middleware.JWTAuth(1), api.UpdatePassword)

		UserRouter.GET("info", middleware.JWTAuth(1), api.GetUserInfo)
		UserRouter.GET("info/:id", api.GetUserInfoByID)
	}
}
