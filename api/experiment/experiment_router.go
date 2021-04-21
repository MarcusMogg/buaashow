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
		CourseRouter.POST(":id/file", middleware.JWTAuth(entity.Teacher), AddExpFile)
		CourseRouter.DELETE(":id/file/:filename", middleware.JWTAuth(entity.Teacher), DeleteExpFile)
		CourseRouter.DELETE(":id", middleware.JWTAuth(entity.Teacher), DeleteExp)
		CourseRouter.GET(":id", GetExp)

		CourseRouter.GET(":id/dl/:account", middleware.JWTAuth(entity.Student), DownloadSubmit)
		CourseRouter.GET(":id/stat", middleware.JWTAuth(entity.Teacher), AllSubmitInfo)
		CourseRouter.POST(":id/stat", middleware.JWTAuth(entity.Teacher), TSubmitInfo)
		CourseRouter.GET(":id/dlall", middleware.JWTAuth(entity.Teacher), DownloadAll)

		CourseRouter.GET(":id/team", middleware.JWTAuth(entity.Student), MyTeamInfo)
		CourseRouter.GET(":id/team/:gid", middleware.JWTAuth(entity.Student), JoinTeam)
		CourseRouter.DELETE(":id/team/:gid", middleware.JWTAuth(entity.Student), QuitTeam)

		CourseRouter.POST(":id/submit", middleware.JWTAuth(entity.Student), SubmitExp)
		CourseRouter.GET(":id/submit", middleware.JWTAuth(entity.Student), SubmitInfo)

		CourseRouter.POST(":id/rec", middleware.JWTAuth(entity.Teacher), Reccommend)
	}
}
