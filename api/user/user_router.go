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
		UserRouter.POST("email", middleware.JWTAuth(entity.Student, false), UpdateEmail)
		UserRouter.POST("name", middleware.JWTAuth(entity.Student, false), UpdateName)
		UserRouter.POST("password", middleware.JWTAuth(entity.Student, false), UpdatePassword)

		UserRouter.GET("info", middleware.JWTAuth(entity.Student, false), GetUserInfo)
		UserRouter.GET("info/:id", GetUserInfoByID)

		UserRouter.POST("teacher", middleware.JWTAuth(entity.Admin, false), CreateTeacher)

		UserRouter.GET("infolist", middleware.JWTAuth(entity.Admin, false), GetUserInfoList)
		UserRouter.DELETE("del/:id", middleware.JWTAuth(entity.Admin, false), DeleteUser)
		UserRouter.POST("reset/:id", middleware.JWTAuth(entity.Admin, false), ResetUser)
		UserRouter.GET("captcha", Captcha)
	}
	tr := Router.Group("test")
	{
		tr.POST("admin", middleware.JWTAuth(entity.Admin, false), TestUser)
		tr.POST("teacher", middleware.JWTAuth(entity.Teacher, false), TestUser)
		tr.POST("user", middleware.JWTAuth(entity.Student, false), TestUser)
	}
}
