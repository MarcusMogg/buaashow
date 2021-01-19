package global

import (
	"buaashow/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	// GConfig 全局配置内容
	GConfig config.Config
	// GVP 读取配置
	GVP *viper.Viper
	// GDB 数据库连接
	GDB *gorm.DB
)

// TimeTemplateDay 时间转换模板，到天
const TimeTemplateDay = "2006-01-02"

// TimeTemplateSec 时间转换模板，到秒
const TimeTemplateSec = "2006-01-02 15:04:05"
