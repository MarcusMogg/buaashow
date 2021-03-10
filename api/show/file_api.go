package show

import (
	"buaashow/global"
	"buaashow/response"
	"buaashow/utils"
	"fmt"
	"net/http"
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
	c.File(path.Join(global.GImgPath, name))
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

// Show gdoc
// @Tags show
// @Summary 图片展示
// @Produce application/json
// @Router /show/{eid}/{gid}/{filepath} [get]
func Show() func(c *gin.Context) {
	fs := gin.Dir(global.GCoursePath, false)
	// fileServer := http.StripPrefix("/show", http.FileServer(fs))
	canShow := true
	return func(c *gin.Context) {
		eid := c.Param("eid")
		gid := c.Param("gid")
		dir := fmt.Sprintf("%s/%s/show", eid, gid)
		file := c.Param("filepath")
		// default index
		/*if file == "/" {
			file = "index.html"
		}*/
		if !canShow {
			c.Status(http.StatusNotFound)
			return
		}
		f, err := fs.Open(path.Join(dir, file))
		if err != nil {
			// default 404
			file = "404.html"
			f, err = fs.Open(path.Join(dir, file))
			if err != nil {

				c.Status(http.StatusNotFound)
				return
			}
		}
		f.Close()
		/* FIXME: 是否屏蔽列出文件夹内容？
		info, err := f.Stat()
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		if info.IsDir() {
			c.Status(http.StatusNotFound)
			return
		}*/
		c.File(path.Join(global.GCoursePath, dir, file))
	}
}
