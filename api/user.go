package api

import (
	"buaashow/global"
	"buaashow/middleware"
	"buaashow/model/entity"
	"buaashow/model/request"
	"buaashow/model/response"
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
// @Param logindata body request.LoginData true "账号密码"
// @Success 200 {object} response.LoginRes
// @Router /user/login [post]
func LoginByPwd(c *gin.Context) {
	var r request.LoginData
	if err := c.BindJSON(&r); err == nil {
		user := &entity.MUser{Account: r.Account, Password: r.Password}

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
// @Param ticket body request.LoginTicketData true "云平台返回的ticket"
// @Success 200 {object} response.LoginRes
// @Router /user/verify [post]
func LoginByTicket(c *gin.Context) {
	var r request.LoginTicketData
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

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	response.OkWithData(user, c)
}

// GetUserInfoByID 获取指定用户信息
func GetUserInfoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	u, err := service.GetUserInfoByID(uint(id))
	if err == nil {
		response.OkWithData(u, c)
	} else {
		response.Fail(c)
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
			Issuer:    "715worker",
		},
	}
	token, err := j.CreateToken(claim)
	if err != nil {
		response.FailWithMessage("token创建失败", c)
		return
	}
	response.OkWithData(response.LoginRes{
		ID:    u.ID,
		Role:  int(u.Role),
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
	var res response.TicketRes
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
