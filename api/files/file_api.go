package files

import (
	"buaashow/global"
	"buaashow/response"
	"buaashow/utils"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

// Download gdoc
// @Tags file
// @Summary 获取文件
// @Produce application/json
// @Router /file/{name} [get]
func Download(c *gin.Context) {
	name := c.Param("name")
	c.File(path.Join(global.GTmpPath, name))
}

// DownloadStatic gdoc
// @Tags file
// @Summary 获取文件
// @Produce application/json
// @Router /static/{name} [get]
func DownloadStatic(c *gin.Context) {
	name := c.Param("name")
	c.File(path.Join(global.GConfig.Static, name))
}

// Upload gdoc
// @Tags file
// @Summary 上传文件
// @Param file formData file true "选择上传文件"
// @Accept multipart/form-data
// @Produce application/json
// @Router /file [post]
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailValidate(c)
		return
	}
	fileType := path.Ext(file.Filename)
	baseName := utils.NextID()
	realPath := path.Join(global.GTmpPath, baseName+fileType)
	_, err = os.Stat(realPath)
	if err == nil {
		response.FailWithMessage("文件已存在", c)
		return
	}
	err = c.SaveUploadedFile(file, realPath)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(baseName+fileType, c)
}
