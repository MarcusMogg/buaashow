package router

import (
	_ "buaashow/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//InitSwaggerRouter 初始化gin swagger
func InitSwaggerRouter(Router *gin.RouterGroup) {
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
