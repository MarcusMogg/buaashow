package initialize

import (
	"buaashow/global"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Mysql 函数初始化mysql连接
func Mysql() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	connect := global.GConfig.Mysql
	dsn := connect.Username + ":" + connect.Password + "@(" + connect.Path + ")/" + connect.Dbname + "?" + connect.Parm
	zap.S().Debug(dsn)
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger}); err != nil {
		zap.S().Fatalf("MySQL启动异常: %s", err)
	} else {
		global.GDB = db
	}
}
