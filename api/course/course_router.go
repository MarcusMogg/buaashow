package course

import (
	"buaashow/entity"
	"buaashow/middleware"

	"github.com/gin-gonic/gin"
)

//InitRouter 初始化course路由组
func InitRouter(Router *gin.RouterGroup) {
	CourseRouter := Router.Group("course")
	{
		CourseRouter.POST("teacher", middleware.JWTAuth(entity.Admin), CreateTeacher)

	}

	Router.GET("terms", middleware.JWTAuth(entity.Student), GetTerms)
}
