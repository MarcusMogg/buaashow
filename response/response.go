package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 通用返回值
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	// SUCCESS 操作成功默认返回码
	SUCCESS = 200
	// ERROR 操作失败默认返回码
	ERROR = 100
	//ValidateError 参数验证失败返回码
	ValidateError = 102
	//TokenError token错误返回码
	TokenError = 103
	//AuthError 权限不足
	AuthError = 104
)

// Result 将结果以json的形式输出至响应
func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// Ok 默认操作成功值
func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

// OkWithMessage 操作成功,自定义返回信息
func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

// OkWithData 操作成功,自定义返回数据
func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

// OkDetailed  操作成功,自定义返回数据和消息
func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

// Fail 默认操作失败值
func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

// FailWithMessage 操作失败,自定义返回信息
func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

// FailDetailed 操作失败,自定义返回码,数据,消息
func FailDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}

// FailValidate 接口传入的参数错误
func FailValidate(c *gin.Context) {
	FailDetailed(ValidateError, map[string]interface{}{}, "传入参数错误", c)
}

// FailToken token错误
func FailToken(c *gin.Context) {
	FailDetailed(TokenError, map[string]interface{}{}, "token错误", c)
}

// FailAuth 权限不足
func FailAuth(c *gin.Context) {
	FailDetailed(AuthError, map[string]interface{}{}, "权限不足", c)
}

// Fail404 文件未找到
func Fail404(c *gin.Context) {
	c.String(http.StatusNotFound, "file not found")
}
