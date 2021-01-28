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
		CourseRouter.GET(":id/students", GetStudents)
		CourseRouter.POST(":id/students", middleware.JWTAuth(entity.Teacher), CreateStudents)
		CourseRouter.DELETE(":cid/student/:uid", middleware.JWTAuth(entity.Teacher), DeleteStudent)
	}

	Router.GET("terms", middleware.JWTAuth(entity.Student), GetTerms)
}
