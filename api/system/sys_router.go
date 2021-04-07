package system

import (
	"buaashow/entity"
	"buaashow/middleware"
	"buaashow/response"
	"buaashow/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(Router *gin.RouterGroup) {
	sr := Router.Group("sys")
	{
		sr.GET("", middleware.JWTAuth(entity.Admin), func(c *gin.Context) {
			var s utils.Server
			s.InitOS()
			s.InitCPU()
			s.InitDisk()
			s.InitRAM()
			response.OkWithData(s, c)
		})
	}
}
