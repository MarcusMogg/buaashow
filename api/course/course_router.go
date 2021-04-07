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

	tr := Router.Group("terms")
	{
		tr.GET("", middleware.JWTAuth(entity.Student), GetTerms)
		tr.GET("all", GetAllTerms)
		tr.POST("", middleware.JWTAuth(entity.Admin), CreateTerm)
		tr.POST(":tid", middleware.JWTAuth(entity.Admin), UpdateTerm)
		tr.DELETE(":tid", middleware.JWTAuth(entity.Admin), DeleteTerm)
	}

	cr := Router.Group("coursename")
	{
		cr.POST(":name", middleware.JWTAuth(entity.Admin), CreateCourseName)
		cr.DELETE(":name", middleware.JWTAuth(entity.Admin), DeleteCourseName)
		cr.GET("", GetCourseNames)
	}
}
