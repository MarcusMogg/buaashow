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
		CourseRouter.GET("", middleware.JWTAuth(entity.Student, false), GetMyExps)
		CourseRouter.POST(":id", middleware.JWTAuth(entity.Teacher, false), EditExp)
		CourseRouter.POST(":id/file", middleware.JWTAuth(entity.Teacher, false), AddExpFile)
		CourseRouter.DELETE(":id/file/:filename", middleware.JWTAuth(entity.Teacher, false), DeleteExpFile)
		CourseRouter.DELETE(":id", middleware.JWTAuth(entity.Teacher, false), DeleteExp)
		CourseRouter.GET(":id", GetExp)

		CourseRouter.GET(":id/dl/:account", middleware.JWTAuth(entity.Student, false), DownloadSubmit)
		CourseRouter.GET(":id/stat", middleware.JWTAuth(entity.Teacher, false), AllSubmitInfo)
		CourseRouter.POST(":id/stat", middleware.JWTAuth(entity.Teacher, false), TSubmitInfo)
		CourseRouter.GET(":id/dlall", middleware.JWTAuth(entity.Teacher, false), DownloadAll)

		CourseRouter.GET(":id/team", middleware.JWTAuth(entity.Student, false), MyTeamInfo)
		CourseRouter.POST(":id/team", middleware.JWTAuth(entity.Student, false), JoinTeam)
		CourseRouter.DELETE(":id/team", middleware.JWTAuth(entity.Student, false), QuitTeam)

		CourseRouter.POST(":id/submit", middleware.JWTAuth(entity.Student, false), SubmitExp)
		CourseRouter.GET(":id/submit", middleware.JWTAuth(entity.Student, false), SubmitInfo)

		CourseRouter.POST(":id/rec", middleware.JWTAuth(entity.Teacher, false), Reccommend)
	}
}
