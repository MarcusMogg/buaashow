package experiment

import "github.com/gin-gonic/gin"

// CreateExp godc
// @Tags exp
// @Summary 创建实验 需教师登录
// @Router /exp [post]
func CreateExp(c *gin.Context) {

}

// GetMyExps godc
// @Tags exp
// @Summary 获取自己的实验列表，需登录
// 等价于获取自己的所有课程，然后获取课程的所有实验
// @Router /exp [get]
func GetMyExps(c *gin.Context) {

}

// GetExp godc
// @Tags exp
// @Summary 根据id获取指定实验信息
// @Router /exp/{id} [get]
func GetExp(c *gin.Context) {

}

// SubmitExp godc
// @Tags exp
// @Summary 提交作业
// @Router /exp/{id}/submit [post]
func SubmitExp(c *gin.Context) {

}
