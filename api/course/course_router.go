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
		CourseRouter.POST("", middleware.JWTAuth(entity.Teacher), CreateCourse)
		CourseRouter.GET("", middleware.JWTAuth(entity.Student), GetMyCourses)
		CourseRouter.GET(":id", middleware.JWTAuth(entity.Student), GetCourseInfo)
		CourseRouter.POST(":id/exp", middleware.JWTAuth(entity.Teacher), CreateExp)
		CourseRouter.GET(":id/exps", GetExps)
		CourseRouter.GET(":id/students", GetStudents)
		CourseRouter.POST(":id/students", middleware.JWTAuth(entity.Teacher), CreateStudents)
		CourseRouter.DELETE(":cid/student/:uid", middleware.JWTAuth(entity.Teacher), DeleteStudent)
		CourseRouter.DELETE(":cid", middleware.JWTAuth(entity.Teacher), DeleteCourse)
	}

	Router.GET("terms", middleware.JWTAuth(entity.Student), GetTerms)
	Router.GET("terms/all", GetAllTerms)
	Router.POST("terms", middleware.JWTAuth(entity.Admin), CreateTerm)
	Router.DELETE("terms", middleware.JWTAuth(entity.Admin), DeleteTerm)
}
