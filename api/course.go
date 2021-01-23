package api

import (
	"github.com/gin-gonic/gin"
)

// CreateTeacher gdoc
// @Tags course
// @Summary 创建教师账号 仅管理员
// @accept application/json
// @Produce application/json
// @Param logindata body request.LoginData true "账号密码"
// @Success 200 {object} response.LoginRes
// @Router /course/teacher [post]
func CreateTeacher(c *gin.Context) {

}

// CreateCourse gdoc
// @Tags course
// @Summary 创建课程 仅教师
// @accept application/json
// @Produce application/json
// @Param logindata body request.LoginData true "账号密码"
// @Success 200 {object} response.LoginRes
// @Router /course [post]
func CreateCourse(c *gin.Context) {

}

// GetMyCourses gdoc
// @Tags course
// @Summary 获取与当前用户相关的课程(教师创建、学生加入)
// @accept application/json
// @Produce application/json
// @Param logindata body request.LoginData true "账号密码"
// @Success 200 {object} response.LoginRes
// @Router /course [get]
func GetMyCourses(c *gin.Context) {

}

// GetCourseInfo gdoc
// @Tags course
// @Summary 获取课程信息, 当前用户需要与课程相关
// @Produce application/json
// @Success 200 {object} response.LoginRes
// @Router /course/{id} [get]
func GetCourseInfo(c *gin.Context) {
}
