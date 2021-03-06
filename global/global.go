package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	// GConfig 全局配置内容
	GConfig Config
	// GVP 读取配置
	GVP *viper.Viper
	// GDB 数据库连接
	GDB *gorm.DB
)

// TimeTemplateDay 时间转换模板，到天
const TimeTemplateDay = "2006-01-02"

// TimeTemplateSec 时间转换模板，到秒
const TimeTemplateSec = "2006-01-02 15:04:05"

// GResourcesPath 资源文件路径
var GResourcesPath = ""

// GImgPath 图片资源文件路径
var GImgPath = GResourcesPath + "img/"

// GStaticPath 静态资源文件路径
var GStaticPath = GResourcesPath + "static/"

// GCoursePath 课程文件路径
var GCoursePath = GResourcesPath + "course/"

// GTmpPath 临时文件路径
var GTmpPath = GResourcesPath + "tmp/"

// GShowPath 展示文件路径
var GShowPath = GResourcesPath + "show/"
