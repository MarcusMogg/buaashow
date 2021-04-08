package course

import (
	"buaashow/response"
	"buaashow/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateCourseName gdoc
// @Tags coursename
// @Summary 添加一个coursename only ADMIN
// @Router /coursename [post]
func CreateCourseName(c *gin.Context) {
	var r courseName
	if err := c.ShouldBindJSON(&r); err == nil && len(r.Name) > 0 {
		if err := service.CreateCourseNmae(r.Name); err == nil {
			response.Ok(c)
		} else {
			zap.S().Debug(err.Error())
			response.Fail(c)
		}
	} else {
		response.FailValidate(c)
		zap.S().Debug(err)
	}
}

// UpdateThumb gdoc
// @Tags coursename
// @Summary 修改大图 only ADMIN
// @Router /coursename [post]
func UpdateThumb(c *gin.Context) {
	var r courseName
	if err := c.ShouldBindJSON(&r); err == nil &&
		len(r.Thumbnail) > 0 && r.NID != 0 {
		if err := service.UpdateCourseThumb(r.NID, r.Thumbnail); err == nil {
			response.Ok(c)
		} else {
			zap.S().Debug(err.Error())
			response.Fail(c)
		}
	} else {
		response.FailValidate(c)
		zap.S().Debug(err)
	}
}

// UpdateThumb gdoc
// @Tags coursename
// @Summary 修改简介 only ADMIN
// @Router /coursename [post]
func UpdateInfo(c *gin.Context) {
	var r courseName
	if err := c.ShouldBindJSON(&r); err == nil &&
		r.NID != 0 {
		if err := service.UpdateCourseInfo(r.NID, r.Info); err == nil {
			response.Ok(c)
		} else {
			zap.S().Debug(err.Error())
			response.Fail(c)
		}
	} else {
		response.FailValidate(c)
		zap.S().Debug(err)
	}
}

// UpdateThumb gdoc
// @Tags coursename
// @Summary 修改大图 only ADMIN
// @Router /coursename [post]
func UpdateName(c *gin.Context) {
	var r courseName
	if err := c.ShouldBindJSON(&r); err == nil &&
		len(r.Name) > 0 && r.NID != 0 {
		if err := service.UpdateCourseName(r.NID, r.Name); err == nil {
			response.Ok(c)
		} else {
			zap.S().Debug(err.Error())
			response.Fail(c)
		}
	} else {
		response.FailValidate(c)
		zap.S().Debug(err)
	}
}

// DeleteCourseName gdoc
// @Tags coursename
// @Summary 添加一个coursename only ADMIN
// @Router /coursename/ [delete]
func DeleteCourseName(c *gin.Context) {
	var r courseName
	if err := c.ShouldBindJSON(&r); err == nil {
		if err := service.DeleteCourseNmae(r.NID); err == nil {
			response.Ok(c)
		} else {
			zap.S().Debug(err.Error())
			response.Fail(c)
		}
	} else {
		response.FailValidate(c)
		zap.S().Debug(err)
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

// GetCourseNameDetails gdoc
// @Tags coursename
// @Summary 获取coursenames
// @Router /coursename/detail/{id} [get]
func GetCourseNameDetails(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	res, err := service.GetCourseName(uint(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(res, c)
	}
}
