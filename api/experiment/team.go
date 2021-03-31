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
	gid := c.Param("gid")
	if err != nil {
		response.FailValidate(c)
		return
	}
	if err := service.JoinTeam(uint(eid), u.Account, gid); err == nil {
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
	gid := c.Param("gid")
	if err != nil {
		response.FailValidate(c)
		return
	}
	if gid == u.Account {
		response.FailWithMessage("you are leader", c)
		return
	}
	if err := service.QiutTeam(uint(eid), u.Account, gid); err == nil {
		response.Ok(c)
	} else {
		zap.S().Debug(err)
		response.FailWithMessage(err.Error(), c)
	}
}
