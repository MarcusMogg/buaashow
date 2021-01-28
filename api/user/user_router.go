package user

import (
	"buaashow/entity"
	"buaashow/middleware"

	"github.com/gin-gonic/gin"
)

//InitRouter 初始化user路由组
func InitRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", LoginByPwd)
		UserRouter.POST("verify", LoginByTicket)
		UserRouter.POST("email", middleware.JWTAuth(entity.Student), UpdateEmail)
		UserRouter.POST("password", middleware.JWTAuth(entity.Student), UpdatePassword)

		UserRouter.GET("info", middleware.JWTAuth(entity.Student), GetUserInfo)
		UserRouter.GET("info/:id", GetUserInfoByID)

		UserRouter.POST("teacher", middleware.JWTAuth(entity.Admin), CreateTeacher)
	}
}
