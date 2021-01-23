package router

import (
	"buaashow/api"

	"github.com/gin-gonic/gin"
)

//InitCourseRouter 初始化course路由组
func InitCourseRouter(Router *gin.RouterGroup) {
	CourseRouter := Router.Group("course")
	{
		CourseRouter.POST("login", api.LoginByPwd)

	}
}
