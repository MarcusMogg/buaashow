package system

import (
	"buaashow/entity"
	"buaashow/middleware"
	"buaashow/response"
	"buaashow/service"
	"buaashow/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(Router *gin.RouterGroup) {
	sr := Router.Group("sys")
	{
		sr.GET("s", middleware.JWTAuth(entity.Admin), func(c *gin.Context) {
			var s utils.Server
			s.InitOS()
			s.InitCPU()
			s.InitDisk()
			s.InitRAM()
			response.OkWithData(s, c)
		})
		sr.GET("i", middleware.JWTAuth(entity.Admin), func(c *gin.Context) {
			s := service.Total()
			response.OkWithData(s, c)
		})
	}
}
