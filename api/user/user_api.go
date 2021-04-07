package user

import (
	"buaashow/entity"
	"buaashow/response"
	"buaashow/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoginByPwd gdoc
// @Tags User
// @Summary 使用账号密码登录
// @accept application/json
// @Produce application/json
// @Param logindata body loginData true "账号密码"
// @Success 200 {object} loginRes
// @Router /user/login [post]
func LoginByPwd(c *gin.Context) {
	var r loginData
	if err := c.ShouldBindJSON(&r); err == nil {
		user := &entity.MUser{Account: r.Account, Password: r.Password}
		// TODO: Use verification code when login
		if service.Login(user) {
			tokenNext(c, user)
		} else {
			response.FailWithMessage("账号或者密码错误", c)
		}

	} else {
		response.FailValidate(c)
	}
}

// LoginByTicket gdoc
// @Tags User
// @Summary 使用云平台登录
// @accept application/json
// @Produce application/json
// @Param ticket body loginTicketData true "云平台返回的ticket"
// @Success 200 {object} loginRes
// @Router /user/verify [post]
func LoginByTicket(c *gin.Context) {
	var r loginTicketData
	if err := c.ShouldBindJSON(&r); err == nil {
		user, err := ticketVerify(r.Authorization, r.ServiceURL)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		}

		tokenNext(c, user)
	} else {
		response.FailValidate(c)
	}
}

// GetUserInfo gdoc
// @Tags User
// @Summary 获取当前用户信息，需用户登录
// @Produce application/json
// @Success 200 {object} entity.UserInfoRes
// @Router /user/info [get]
func GetUserInfo(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	response.OkWithData(entity.UserInfoRes{
		Account: u.Account,
		Role:    int(u.Role),
		Email:   u.Email,
		Name:    u.Name,
	}, c)
}

// GetUserInfoByID gdoc
// @Tags User
// @Summary 获取指定id(account)的用户信息
// @Accept  json
// @Produce  json
// @Param id path string true "User account"
// @Success 200 {object} entity.UserInfoRes
// @Router /user/info/{id} [get]
func GetUserInfoByID(c *gin.Context) {
	id := c.Param("id")
	u, err := service.GetUserInfo(id)
	if err == nil {
		response.OkWithData(entity.UserInfoRes{
			Account: u.Account,
			Role:    int(u.Role),
			Name:    u.Name,
			Email:   u.Email,
		}, c)
	} else {
		response.Fail(c)
	}
}

// UpdateEmail gdoc
// @Tags User
// @Summary 修改邮箱, 需用户登录
// @accept application/json
// @Produce application/json
// @Param ticket body emailData true "新邮箱"
// @Success 200 {object} loginRes
// @Router /user/email [post]
func UpdateEmail(c *gin.Context) {
	claim, ok := c.Get("user")
	//FIXME: 中间件信息是否需要验证
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	var email emailData
	if err := c.ShouldBindJSON(&email); err == nil {
		err = service.UpdateEmail(u, email.Email)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.Ok(c)
		}
	} else {
		response.FailValidate(c)
	}
}

// UpdateEmail gdoc
// @Tags User
// @Summary 修改name, 需用户登录
// @accept application/json
// @Produce application/json
// @Param ticket body nameData true "新name"
// @Success 200 {object} loginRes
// @Router /user/name [post]
func UpdateName(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	var name nameData
	if err := c.ShouldBindJSON(&name); err == nil {
		err = service.UpdateName(u, name.Name)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.Ok(c)
		}
	} else {
		response.FailValidate(c)
	}
}

// UpdatePassword gdoc
// @Tags User
// @Summary 修改密码, 需用户登录
// @accept application/json
// @Produce application/json
// @Param ticket body passwordData true "新旧密码"
// @Success 200 {object} loginRes
// @Router /user/password [post]
func UpdatePassword(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	var pass passwordData
	if err := c.ShouldBindJSON(&pass); err == nil {
		err = service.UpdatePassword(u, pass.OldPassword, pass.NewPassword)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.Ok(c)
		}
	} else {
		response.FailValidate(c)
	}
}

// CreateTeacher gdoc
// @Tags User
// @Summary 创建教师账号 only ADMIN
// @accept application/json
// @Produce application/json
// @Param logindata body registerData true "账号密码必选，邮箱可选"
// @Router /user/teacher [post]
func CreateTeacher(c *gin.Context) {
	var r registerData
	if err := c.ShouldBindJSON(&r); err == nil {

		user := &entity.MUser{
			Account:  r.Account,
			Password: r.Password,
			Role:     entity.Teacher,
		}
		if err = service.Register(user); err == nil {
			response.OkWithMessage("注册成功", c)
			zap.S().Infof("Register Teacher %s", user.Account)
		} else {
			response.FailWithMessage(err.Error(), c)
			zap.S().Debug(err.Error())
		}

	} else {
		response.FailValidate(c)
		zap.S().Debug(err.Error())
	}

}

// GetUserInfoList gdoc
// @Tags User
// @Summary 获取用户列表，only ADMIN
// @Produce application/json
// @Param page query int false "page number"
// @Param account query string false "account target"
// @Success 200 {array} entity.UserInfoRes
// @Router /user/infolist [get]
func GetUserInfoList(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.ParseUint(pageStr, 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	tp, err := strconv.ParseUint(c.DefaultQuery("type", "1"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	ac := c.Query("account")
	size := 16
	tot, res, err := service.GetUserInfoList(int(page), size, int(tp), ac)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(gin.H{
		"tot":   tot,
		"users": res,
	}, c)
}

// DeleteUser gdoc
// @Tags User
// @Summary 删除一个用户 only ADMIN
// @Router /user/del/{id} [delete]
func DeleteUser(c *gin.Context) {
	name := c.Param("id")
	err := service.DeleteUser(name)
	if err == nil {
		response.Ok(c)
	} else {
		zap.S().Debug(err.Error())
		response.Fail(c)
	}
}

// ResetUser gdoc
// @Tags User
// @Summary 重置用户密码 only ADMIN
// @Router /user/del/{id} [delete]
func ResetUser(c *gin.Context) {
	name := c.Param("id")
	var u entity.MUser
	u.Account = name
	err := service.ResetPassword(&u)
	if err == nil {
		response.Ok(c)
	} else {
		zap.S().Debug(err.Error())
		response.Fail(c)
	}
}

// TestAdmin gdoc
// @Tags User
// @Summary 测试管理员token
// @Router /test/admin [post]
func TestAdmin(c *gin.Context) {
	response.Ok(c)
}

// TestUser gdoc
// @Tags User
// @Summary 测试用户token
// @Router /test/user [post]
func TestUser(c *gin.Context) {
	response.Ok(c)
}
