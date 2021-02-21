package experiment

import (
	"buaashow/entity"
	"buaashow/middleware"

	"github.com/gin-gonic/gin"
)

//InitRouter 初始化experiment路由组
func InitRouter(Router *gin.RouterGroup) {
	CourseRouter := Router.Group("exp")
	{
		CourseRouter.POST("", middleware.JWTAuth(entity.Teacher), CreateExp)
		CourseRouter.GET("", middleware.JWTAuth(entity.Student), GetMyExps)
		CourseRouter.GET(":id", GetExp)
		CourseRouter.GET(":id/submit", middleware.JWTAuth(entity.Student), SubmitExp)
	}
}
