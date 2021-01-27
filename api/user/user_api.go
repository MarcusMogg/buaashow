package user

import (
	"buaashow/entity"
	"buaashow/global"
	"buaashow/middleware"
	"buaashow/response"
	"buaashow/service"
	"buaashow/utils"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// LoginByPwd gdoc
// @Tags User
// @Summary 使用账号密码登录
// @accept application/json
// @Produce application/json
// @Param logindata body LoginData true "账号密码"
// @Success 200 {object} response.LoginRes
// @Router /user/login [post]
func LoginByPwd(c *gin.Context) {
	var r LoginData
	if err := c.BindJSON(&r); err == nil {
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
// @Param ticket body LoginTicketData true "云平台返回的ticket"
// @Success 200 {object} response.LoginRes
// @Router /user/verify [post]
func LoginByTicket(c *gin.Context) {
	var r LoginTicketData
	if err := c.BindJSON(&r); err == nil {
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
// @Success 200 {object} response.UserInfoRes
// @Router /user/info [get]
func GetUserInfo(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	response.OkWithData(InfoRes{
		ID:    u.ID,
		Role:  int(u.Role),
		Email: u.Email,
	}, c)
}

// GetUserInfoByID gdoc
// @Tags User
// @Summary 获取指定id的用户信息
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} response.UserInfoRes
// @Router /user/info/{id} [get]
func GetUserInfoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	u, err := service.GetUserInfoByID(uint(id))
	if err == nil {
		response.OkWithData(InfoRes{
			ID:    u.ID,
			Role:  int(u.Role),
			Email: u.Email,
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
// @Param ticket body EmailData true "新邮箱"
// @Success 200 {object} response.LoginRes
// @Router /user/email [post]
func UpdateEmail(c *gin.Context) {
	claim, ok := c.Get("user")
	//FIXME: 中间件信息是否需要验证
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	var email EmailData
	if err := c.BindJSON(&email); err == nil {
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

// UpdatePassword gdoc
// @Tags User
// @Summary 修改密码, 需用户登录
// @accept application/json
// @Produce application/json
// @Param ticket body PasswordData true "新旧密码"
// @Success 200 {object} response.LoginRes
// @Router /user/password [post]
func UpdatePassword(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	var pass PasswordData
	if err := c.BindJSON(&pass); err == nil {
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

func tokenNext(c *gin.Context, u *entity.MUser) {
	j := middleware.NewJWT()
	claim := middleware.JWTClaim{
		UserID:   u.ID,
		UserName: u.Account,
		Role:     u.Role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 60*60*24*7,
			Issuer:    "Mogg",
		},
	}
	token, err := j.CreateToken(claim)
	if err != nil {
		response.FailWithMessage("token创建失败", c)
		return
	}
	response.OkWithData(LoginRes{
		InfoRes: InfoRes{
			ID:    u.ID,
			Role:  int(u.Role),
			Email: u.Email,
		},
		Token: token,
	}, c)
}

func ticketVerify(ticket string, serviceURL string) (user *entity.MUser, err error) {
	str := []byte(ticket)
	i := len(str) - 1
	for ; i >= 0; i-- {
		if str[i] != '#' {
			break
		}
	}
	str = str[0 : i+1]
	data := fmt.Sprintf("token=%s&service=%s", str, serviceURL)
	resp, err := utils.Post(global.GConfig.SSOServer, "application/x-www-form-urlencoded", data)
	if err != nil {
		return
	}
	var res TicketRes
	json.Unmarshal(resp, &res)
	if res.Code != 1003 {
		err = errors.New(res.Msg)
		return
	}
	user, err = service.GetUserInfoByAccount(res.Data.ID)
	if err != nil {
		return
	}
	if user.Role != entity.Role(res.Data.Role) {
		err = errors.New("角色不匹配")
	}
	return
}
