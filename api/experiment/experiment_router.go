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
		CourseRouter.GET("", middleware.JWTAuth(entity.Student), GetMyExps)
		CourseRouter.POST(":id", middleware.JWTAuth(entity.Teacher), EditExp)
		CourseRouter.DELETE(":id", middleware.JWTAuth(entity.Teacher), DeleteExp)
		CourseRouter.GET(":id", GetExp)
		CourseRouter.POST(":id/submit", middleware.JWTAuth(entity.Student), SubmitExp)
		CourseRouter.GET(":id/submit", middleware.JWTAuth(entity.Student), SubmitInfo)
		//TODO: statistics for teacher
	}
}
