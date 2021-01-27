package course

import (
	"buaashow/entity"
	"buaashow/response"
	"buaashow/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateTeacher gdoc
// @Tags course
// @Summary 创建教师账号 需管理员登录
// @accept application/json
// @Produce application/json
// @Param logindata body RegisterData true "账号密码必选，邮箱可选"
// @Router /course/teacher [post]
func CreateTeacher(c *gin.Context) {
	var r RegisterData
	if err := c.BindJSON(&r); err == nil {

		user := &entity.MUser{
			Account:  r.Account,
			Password: r.Password,
			Role:     entity.Teacher,
		}
		if err = service.Register(user); err == nil {
			response.OkWithMessage("注册成功", c)
			zap.S().Infof("Register Teacher %s", user.Account)
		} else {
			response.FailWithMessage(fmt.Sprintf("%v", err), c)
			zap.S().Debug(err)
		}

	} else {
		response.FailValidate(c)
		zap.S().Debug(err)
	}

}

// CreateCourse gdoc
// @Tags course
// @Summary 创建课程 需教师登录
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
