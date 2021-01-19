package api

import (
	"buaashow/middleware"
	"buaashow/model/entity"
	"buaashow/model/request"
	"buaashow/model/response"
	"buaashow/service"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// Register gdoc
// @Tags User
// @Summary 用户注册
// @accept application/json
// @Produce application/json
// @Param register body request.RegisterData true "Add account"
// @Success 200 {string} string "{"code":1,"data":{},"msg":"注册成功"}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	var r request.RegisterData
	if err := c.BindJSON(&r); err == nil {
		user := &entity.MUser{
			UserName: r.UserName,
			Email:    r.Email,
			Password: r.Password,
			Role:     entity.Student,
			NickName: r.UserName,
		}

		if err = service.Register(user); err == nil {
			response.OkWithMessage("注册成功", c)
		} else {
			response.FailWithMessage(fmt.Sprintf("%v", err), c)
		}

	} else {
		response.FailValidate(c)
	}
}

// Login 用户登录
func Login(c *gin.Context) {
	var r request.LoginData
	if err := c.BindJSON(&r); err == nil {
		user := &entity.MUser{UserName: r.UserName, Password: r.Password}

		if service.Login(user) {
			tokenNext(c, user)
		} else {
			response.FailWithMessage("账号或者密码错误", c)
		}

	} else {
		response.FailValidate(c)
	}
}

func tokenNext(c *gin.Context, u *entity.MUser) {
	j := middleware.NewJWT()
	claim := middleware.JWTClaim{
		UserID:   u.ID,
		UserName: u.UserName,
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
	response.OkWithData(token, c)
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
