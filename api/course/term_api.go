package course

import (
	"buaashow/entity"
	"buaashow/response"
	"buaashow/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetTerms gdoc
// @Tags term
// @Summary 获取学期信息，从用户创建时间到当前时间段的所有学期数,需用户登录
// @Produce application/json
// @Success 200 {array} entity.Term
// @Router /terms [get]
func GetTerms(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	year := u.CreatedAt.Year()
	response.OkWithData(service.GetTerms(year), c)
}

// CreateTerm gdoc
// @Tags term
// @Summary 新增一个学期,需管理员登录
// @Produce application/json
// @Param newTermData body entity.Term true "学期信息"
// @Router /terms [post]
func CreateTerm(c *gin.Context) {
	var t entity.Term
	if err := c.ShouldBindJSON(&t); err == nil {
		if err = service.CreateTerm(&t); err == nil {
			response.OkWithData(t, c)
		} else {
			response.FailWithMessage(err.Error(), c)
		}
	} else {
		response.FailValidate(c)
		zap.S().Debug(err)
	}

}

// DeleteTerm gdoc
// @Tags term
// @Summary 删除一个学期,需管理员登录，注意，会同步删除该学期相关的所有课程、实验
// @Produce application/json
// @Param newTermData body entity.Term true "学期信息"
// @Router /terms/{tid} [delete]
func DeleteTerm(c *gin.Context) {
	tid, err := strconv.ParseUint(c.Param("tid"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	if err = service.DeleteTerm(uint(tid)); err == nil {
		response.Ok(c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}

// GetAllTerms gdoc
// @Tags term
// @Summary 获取所有学期信息
// @Produce application/json
// @Success 200 {array} entity.Term
// @Router /terms/all [get]
func GetAllTerms(c *gin.Context) {
	response.OkWithData(service.GetTerms(0), c)
}
