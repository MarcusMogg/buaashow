package initialize

import (
	"buaashow/global"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type myGormLogger struct{}

func (*myGormLogger) Printf(fm string, param ...interface{}) {
	zap.S().Debugf(fm, param...)
}

// Mysql 函数初始化mysql连接
func Mysql() {
	var cfg logger.Config
	if gin.Mode() == gin.ReleaseMode {
		cfg = logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Warn, // Log level
			Colorful:      false,       // 禁用彩色打印
		}
	} else {
		cfg = logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		}
	}

	newLogger := logger.New(
		&myGormLogger{}, // io writer
		cfg)
	connect := global.GConfig.Mysql
	dsn := connect.Username + ":" + connect.Password + "@(" + connect.Path + ")/" + connect.Dbname + "?" + connect.Parm
	zap.S().Debug(dsn)
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger}); err != nil {
		zap.S().Fatalf("MySQL启动异常: %s", err)
	} else {
		global.GDB = db
	}
}
