package course

import (
	"buaashow/response"
	"buaashow/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateCourseName gdoc
// @Tags coursename
// @Summary 添加一个coursename only ADMIN
// @Router /coursename/{name} [post]
func CreateCourseName(c *gin.Context) {
	name := c.Param("name")
	if len(name) == 0 {
		response.FailValidate(c)
		return
	}
	err := service.CreateCourseNmae(name)
	if err == nil {
		response.Ok(c)
	} else {
		zap.S().Debug(err.Error())
		response.Fail(c)
	}
}

// DeleteCourseName gdoc
// @Tags coursename
// @Summary 添加一个coursename only ADMIN
// @Router /coursename/{name} [delete]
func DeleteCourseName(c *gin.Context) {
	name := c.Param("name")
	err := service.DeleteCourseNmae(name)
	if err == nil {
		response.Ok(c)
	} else {
		zap.S().Debug(err.Error())
		response.Fail(c)
	}
}

// GetCourseNames gdoc
// @Tags coursename
// @Summary 获取coursenames
// @Router /coursename [get]
func GetCourseNames(c *gin.Context) {
	res := service.GetAllCourseNmae()
	response.OkWithData(res, c)
}
