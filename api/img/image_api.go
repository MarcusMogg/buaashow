package img

import (
	"buaashow/global"
	"buaashow/response"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
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
	height, err := strconv.ParseUint(c.DefaultQuery("height", "0"), 10, 0)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	width, err := strconv.ParseUint(c.DefaultQuery("width", "0"), 10, 0)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	iimg, imgType, err := decode(file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	baseName, err := rename(imgType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := resizeFile(file, iimg, imgType, baseName,
		size{width: uint(width), height: uint(height)})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(res, c)
}

// Download gdoc
// @Tags Img
// @Summary 获取图片
// @Produce application/json
// @Router /img/{name} [get]
func Download(c *gin.Context) {
	name := c.Param("name")
	c.File(path.Join(global.GImgPath, name))
}
