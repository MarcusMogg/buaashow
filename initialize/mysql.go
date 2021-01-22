package initialize

import (
	"buaashow/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mysql 函数初始化mysql连接
func Mysql() {
	connect := global.GConfig.Mysql
	dsn := connect.Username + ":" + connect.Password + "@(" + connect.Path + ")/" + connect.Dbname + "?" + connect.Parm
	zap.S().Debug(dsn)
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		zap.S().Fatalf("MySQL启动异常: %s", err)
	} else {
		global.GDB = db
	}
}
