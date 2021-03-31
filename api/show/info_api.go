package show

import (
	"buaashow/entity"
	"buaashow/global"
	"buaashow/response"
	"buaashow/service"
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Show gdoc
// @Tags show
// @Summary 图片展示
// @Produce application/json
// @Router /show/x/{showid}/{filepath} [get]
func Show() func(c *gin.Context) {
	fs := gin.Dir(global.GCoursePath, false)
	// fileServer := http.StripPrefix("/show", http.FileServer(fs))
	canShow := true
	return func(c *gin.Context) {
		sid := c.Param("showid")
		s, err := entity.DecodeShowID(sid)
		if err != nil {
			zap.S().Debug(err)
			c.Status(http.StatusNotFound)
			return
		}
		dir := fmt.Sprintf("%d/%s/show", s.EID, s.GID)
		file := c.Param("filepath")
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

// Search gdoc
// @Tags show
// @Summary 简略信息
// @Produce application/json
// @Router /show/search [get]
func Search(c *gin.Context) {
	var sp entity.SearchParam
	if err := c.ShouldBindQuery(&sp); err == nil {
		tot, res := service.GetSummary(&sp)
		response.OkWithData(gin.H{
			"tot": tot,
			"res": res,
		}, c)
	} else {
		response.FailValidate(c)
		zap.S().Debug(err)
	}
}

// Readme gdoc
// @Tags show
// @Summary 简介
// @Produce application/json
// @Router /show/readme/{showid} [get]
func Readme(c *gin.Context) {
	sid := c.Param("showid")
	s, err := entity.DecodeShowID(sid)
	if err != nil {
		zap.S().Debug(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	var res entity.SubmissionResp
	if err := service.GetSubmission(s.EID, s.GID, &res); err == nil {
		for i := range res.Groups {
			res.Groups[i].Account = ""
		}
		response.OkWithData(res, c)
	} else {
		zap.S().Debug(err)
		response.FailWithMessage(err.Error(), c)
	}
}
