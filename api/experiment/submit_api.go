package experiment

import (
	"buaashow/entity"
	"buaashow/response"
	"buaashow/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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
			SrcURL:    req.SrcURL,
			DistURL:   req.DistURL,
			Readme:    req.Readme,
			UpAt:      now,
			Thumbnail: req.Thumbnail,
		}
		if err = service.Submit(&submission, req.Account, u); err != nil {
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

// TSubmitInfo godc
// @Tags exp
// @Summary 教师获取学生提交的详细信息
// @Produce application/json
// @Success 200 {object} entity.SubmissionResp
// @Router /exp/{id}/stat [post]
func TSubmitInfo(c *gin.Context) {
	eid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	var req struct {
		Account string `json:"account"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		var res entity.SubmissionResp
		if err = service.GetSubmission(uint(eid), req.Account, &res); err != nil {
			zap.S().Debug(err)
			res.Status = false
		}
		response.OkWithData(res, c)
	} else {
		response.FailValidate(c)
	}
}
