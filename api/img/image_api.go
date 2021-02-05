package img

import (
	"buaashow/response"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Upload gdoc
// @Tags Img
// @Summary 上传图片
// @Param file formData file true "选择上传文件"
// @Accept multipart/form-data
// @Produce application/json
// @Router /img [post]
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailValidate(c)
		return
	}
	start := time.Now()
	iimg, imgType, err := decode(file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	zap.S().Debug(time.Now().Sub(start))
	baseName, err := rename(file, imgType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	zap.S().Debug(time.Now().Sub(start))
	res, err := resizeFile(file, iimg, imgType, baseName)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	zap.S().Debug(time.Now().Sub(start))
	response.OkWithData(res, c)
}

// Download gdoc
// @Tags Img
// @Summary 获取图片
// @Produce application/json
// @Router /img/{name} [get]
func Download(c *gin.Context) {
	name := c.Param("name")
	c.File(imgDir + name)
}
