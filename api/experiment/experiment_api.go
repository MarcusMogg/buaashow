package experiment

import (
	"buaashow/entity"
	"buaashow/global"
	"buaashow/response"
	"buaashow/service"
	"fmt"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetMyExps godc
// @Tags exp
// @Summary 获取自己的实验列表，需登录
// 等价于获取自己的所有课程，然后获取课程的所有实验
// @Produce application/json
// @Success 200 {array} entity.ExperimentResponse
// @Router /exp [get]
func GetMyExps(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	courses := service.GetMyCourses(u)
	var res []*entity.ExperimentResponse
	for _, i := range courses {
		tmp, err := service.GetExpsByCID(i.ID)
		if err != nil {
			zap.S().Debug(err.Error())
		} else {
			res = append(res, tmp...)
		}
	}
	response.OkWithData(res, c)
}

// GetExp godc
// @Tags exp
// @Summary 根据id获取指定实验信息
// @Produce application/json
// @Success 200 {object} entity.ExperimentResponse
// @Router /exp/{id} [get]
func GetExp(c *gin.Context) {
	eid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	exp, err := service.GetExp(uint(eid))
	if err != nil {
		zap.S().Debug(err)
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(exp, c)
	}
}

// EditExp godc
// @Tags exp
// @Summary 修改指定实验
// @Produce application/json
// @Success 200
// @Router /exp/{id} [post]
func EditExp(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	eid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	exp, err := service.GetMExp(uint(eid))
	if err != nil {
		zap.S().Debug(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	var req entity.ExperimentReq
	if err := c.ShouldBindJSON(&req); err == nil {
		var begin, end time.Time
		if begin, err = time.ParseInLocation(global.TimeTemplateSec, req.BeginTime, time.Local); err != nil {
			response.FailValidate(c)
			zap.S().Debug(err)
			return
		}
		if end, err = time.ParseInLocation(global.TimeTemplateSec, req.EndTime, time.Local); err != nil {
			response.FailValidate(c)
			zap.S().Debug(err)
			return
		}
		exp.Name = req.Name
		exp.Info = req.Info
		exp.BeginTime = begin
		exp.EndTime = end
		exp.Team = req.Team
		if err = service.UpdateExp(exp, u.Account); err != nil {
			response.FailWithMessage(err.Error(), c)
			zap.S().Debug(err)
			return
		}
		response.Ok(c)
	} else {
		response.FailValidate(c)
		zap.S().Debug(err)
	}

}

// AddExpFile godc
// @Tags exp
// @Summary 添加实验资源文件,文件存在的话会被替换
// @Produce application/json
// @Param file formData file true "选择上传文件"
// @Accept multipart/form-data
// @Success 200
// @Router /exp/{id}/file [post]
func AddExpFile(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	eid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		response.FailValidate(c)
		return
	}

	filename := fmt.Sprintf("%d-%s", eid, file.Filename)
	c.SaveUploadedFile(file, path.Join(global.GTmpPath, filename))

	if err = service.AddExpFile(uint(eid), u.Account, filename); err == nil {
		response.Ok(c)
	} else {
		response.FailWithMessage(err.Error(), c)
		zap.S().Debug(err)
	}
}

// DeleteExpFile godc
// @Tags exp
// @Summary 删除实验资源文件
// @Produce application/json
// @Success 200
// @Router /exp/{id}/file/{filename} [delete]
func DeleteExpFile(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	eid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	filename := c.Param("filename")

	if err = service.DeleteExpFile(uint(eid), u.Account, filename); err != nil {
		response.Ok(c)
	} else {
		response.FailWithMessage(err.Error(), c)
		zap.S().Debug(err)
	}
}

// DeleteExp godc
// @Tags exp
// @Summary 删除指定实验
// @Produce application/json
// @Success 200
// @Router /exp/{id} [delete]
func DeleteExp(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	eid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	if err = service.DeleteExp(uint(eid), u.Account); err != nil {
		response.Ok(c)
	} else {
		response.FailWithMessage(err.Error(), c)
		zap.S().Debug(err)
	}
}

// SubmitExp godc
// @Tags exp
// @Summary 提交作业
// @Produce application/json
// @Param id path int true "Exp ID"
// @Param exp body SubmissionReq true "实验信息"
// @Success 200
// @Router /exp/{id}/submit [post]
func SubmitExp(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	eid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	var req SubmissionReq
	if err := c.ShouldBindJSON(&req); err == nil {
		now := time.Now()
		if req.Type < entity.HTML || req.Type > entity.URL {
			response.FailWithMessage("类型错误", c)
			return
		}
		submission := entity.MSubmission{
			EID:       uint(eid),
			Name:      req.Name,
			Info:      req.Info,
			Type:      req.Type,
			URL:       req.URL,
			Readme:    req.Readme,
			UpdatedAt: now,
		}
		if err = service.Submit(&submission, u.Account); err != nil {
			response.FailWithMessage(err.Error(), c)
			zap.S().Debug(err)
		} else {
			response.Ok(c)
		}
	} else {
		response.FailValidate(c)
		zap.S().Debug(err)
	}
}

// SubmitInfo godc
// @Tags exp
// @Summary 学生提交信息
// @Produce application/json
// @Success 200 {object} entity.SubmissionResp
// @Router /exp/{id}/submit [get]
func SubmitInfo(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	eid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	var res entity.SubmissionResp
	if err = service.GetSubmission(uint(eid), u.Account, &res); err != nil {
		zap.S().Debug(err)
		res.Status = false
	}
	response.OkWithData(res, c)
}

// AllSubmitInfo godc
// @Tags exp
// @Summary 学生提交信息列表
// @Produce application/json
// @Success 200 {array} entity.SubmissionResp
// @Router /exp/{id}/stat [get]
func AllSubmitInfo(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	eid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	res, err := service.GetAllSubmission(uint(eid), u.Account)
	if err == nil {
		response.OkWithData(res, c)
	} else {
		zap.S().Debug(err)
		response.FailWithMessage(err.Error(), c)
	}
}

// DownloadSubmit godc
// @Tags exp
// @Summary 下载提交过的作业
// @Produce application/json
// @Router /exp/{id}/dl/{account} [get]
func DownloadSubmit(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	eid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	ac := c.Param("account")
	if ac != u.Account && u.Role == entity.Student {
		response.FailWithMessage("权限不足", c)
		return
	}
	filename, err := service.DownloadSubmission(uint(eid), ac)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		zap.S().Debug(err)
	} else {
		zap.S().Debug(filename)
		c.File(filename)
	}

}

// DownloadSubmit godc
// @Tags exp
// @Summary 下载所有提交过的作业
// @Produce application/json
// @Router /exp/{id}/dlall [get]
func DownloadAll(c *gin.Context) {
	eid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	filename, err := service.DownloadAllSubmission(uint(eid))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		zap.S().Debug(err)
	} else {
		c.File(filename)
	}
}

// Reccommend godc
// @Tags exp
// @Summary 将作业设为推荐
// @Produce application/json
// @Router /exp/{id}/rec/{account} [get]
func Reccommend(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	eid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	ac := c.Param("account")
	if err := service.Reccommend(uint(eid), ac, u.Account); err == nil {
		response.Ok(c)
	} else {
		zap.S().Debug(err)
		response.Fail(c)
	}
}
