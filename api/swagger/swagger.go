package swagger

import (
	// fow gin-swagger
	_ "buaashow/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//InitRouter 初始化gin swagger
func InitRouter(Router *gin.RouterGroup) {
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
