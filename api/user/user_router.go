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
		UserRouter.POST("name", middleware.JWTAuth(entity.Student), UpdateName)
		UserRouter.POST("password", middleware.JWTAuth(entity.Student), UpdatePassword)

		UserRouter.GET("info", middleware.JWTAuth(entity.Student), GetUserInfo)
		UserRouter.GET("info/:id", GetUserInfoByID)

		UserRouter.POST("teacher", middleware.JWTAuth(entity.Admin), CreateTeacher)

		UserRouter.GET("infolist", middleware.JWTAuth(entity.Admin), GetUserInfoList)
		UserRouter.DELETE("del/:id", middleware.JWTAuth(entity.Admin), DeleteUser)
		UserRouter.POST("reset/:id", middleware.JWTAuth(entity.Admin), ResetUser)
		UserRouter.GET("captcha", Captcha)
	}
	tr := Router.Group("test")
	{
		tr.POST("admin", middleware.JWTAuth(entity.Admin), TestUser)
		tr.POST("teacher", middleware.JWTAuth(entity.Teacher), TestUser)
		tr.POST("user", middleware.JWTAuth(entity.Student), TestUser)
	}
}
