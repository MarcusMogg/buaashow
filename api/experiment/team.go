package experiment

import (
	"buaashow/entity"
	"buaashow/response"
	"buaashow/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// MyTeamInfo godc
// @Tags exp
// @Summary 查看当前用户的组队信息
// @Produce application/json
// @Router /exp/{id}/team [get]
func MyTeamInfo(c *gin.Context) {
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
	if gid, inTeam, err := service.TeamInfo(uint(eid), u.Account); err == nil {
		response.OkDetailed(gin.H{
			"inTeam": inTeam,
			"gid":    gid,
		}, "", c)
	} else {
		zap.S().Debug(err)
		response.Fail(c)
	}
}

// JoinTeam godc
// @Tags exp
// @Summary 加入队伍
// @Produce application/json
// @Router /exp/{id}/team/{gid} [get]
func JoinTeam(c *gin.Context) {
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
	var gid struct {
		GID string `json:"gid"`
	}
	err = c.ShouldBindJSON(&gid)
	if err != nil {
		response.FailValidate(c)
		return
	}
	if err := service.JoinTeam(uint(eid), u.Account, gid.GID); err == nil {
		response.Ok(c)
	} else {
		zap.S().Debug(err)
		response.FailWithMessage(err.Error(), c)
	}
}

// QuitTeam godc
// @Tags exp
// @Summary 退出队伍
// @Produce application/json
// @Router /exp/{id}/team/{gid} [delete]]
func QuitTeam(c *gin.Context) {
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
	var gid struct {
		GID string `json:"gid"`
		Tar string `json:"account"`
	}
	err = c.ShouldBindJSON(&gid)
	if err != nil {
		response.FailValidate(c)
		return
	}
	if gid.GID == u.Account { // for leader
		if len(gid.Tar) == 0 || gid.Tar == gid.GID {
			response.FailWithMessage("队长不能删除", c)
			return
		}
		err = service.QiutTeam(uint(eid), gid.Tar, gid.GID)
	} else { // for quit team
		if len(gid.Tar) != 0 || gid.Tar == u.Account {
			response.FailWithMessage("队员不能删除其他人", c)
			return
		}
		err = service.QiutTeam(uint(eid), u.Account, gid.GID)
	}
	if err == nil {
		response.Ok(c)
	} else {
		zap.S().Debug(err)
		response.FailWithMessage(err.Error(), c)
	}
}
