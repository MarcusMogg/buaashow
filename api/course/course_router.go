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
		CourseRouter.POST("", middleware.JWTAuth(entity.Teacher, false), CreateCourse)
		CourseRouter.GET("", middleware.JWTAuth(entity.Student, false), GetMyCourses)
		CourseRouter.GET(":id", middleware.JWTAuth(entity.Student, false), GetCourseInfo)
		CourseRouter.POST(":id/exp", middleware.JWTAuth(entity.Teacher, false), CreateExp)
		CourseRouter.GET(":id/exps", GetExps)
		CourseRouter.GET(":id/students", GetStudents)
		CourseRouter.POST(":id/students", middleware.JWTAuth(entity.Teacher, false), CreateStudents)
		CourseRouter.DELETE(":cid/student/:uid", middleware.JWTAuth(entity.Teacher, false), DeleteStudent)
		CourseRouter.DELETE(":cid/students", middleware.JWTAuth(entity.Teacher, false), DeleteAllStudent)
		CourseRouter.DELETE(":cid", middleware.JWTAuth(entity.Teacher, false), DeleteCourse)
	}

	tr := Router.Group("terms")
	{
		tr.GET("", middleware.JWTAuth(entity.Student, false), GetTerms)
		tr.GET("all", GetAllTerms)
		tr.POST("", middleware.JWTAuth(entity.Admin, false), CreateTerm)
		tr.POST(":tid", middleware.JWTAuth(entity.Admin, false), UpdateTerm)
		tr.DELETE(":tid", middleware.JWTAuth(entity.Admin, false), DeleteTerm)
	}

	cr := Router.Group("coursename")
	{
		cr.POST("", middleware.JWTAuth(entity.Admin, false), CreateCourseName)
		cr.POST("thumb", middleware.JWTAuth(entity.Admin, false), UpdateThumb)
		cr.POST("info", middleware.JWTAuth(entity.Admin, false), UpdateInfo)
		cr.POST("name", middleware.JWTAuth(entity.Admin, false), UpdateName)
		cr.DELETE("", middleware.JWTAuth(entity.Admin, false), DeleteCourseName)
		cr.GET("detail/:id", GetCourseNameDetails)
		cr.GET("", GetCourseNames)
	}
}
