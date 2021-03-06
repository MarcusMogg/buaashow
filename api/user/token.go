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
	"go.uber.org/zap"
)

// 生成token
func tokenNext(c *gin.Context, u *entity.MUser) {
	j := middleware.NewJWT()
	claim := middleware.JWTClaim{
		Account: u.Account,
		Role:    u.Role,
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
	response.OkWithData(loginRes{
		UserInfoRes: entity.UserInfoRes{
			Account: u.Account,
			Role:    int(u.Role),
			Name:    u.Name,
			Email:   u.Email,
		},
		Token: token,
	}, c)
}

// 使用cloud.beihangsoft.cn 验证登录
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
	var res ticketRes
	json.Unmarshal(resp, &res)
	if res.Code != 1003 {
		err = errors.New(res.Msg)
		return
	}
	user, err = service.GetUserInfo(res.Data.ID)
	if err != nil {
		return
	}
	r, err := strconv.Atoi(res.Data.Role)
	if err != nil {
		return
	}
	if r > 3 {
		r = 3
	}
	if user.Role != entity.Role(r) {
		err = errors.New("角色不匹配")
		zap.S().Debug(string(resp))
	}
	return
}
