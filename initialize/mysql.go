package initialize

import (
	"buaashow/global"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mysql 函数初始化mysql连接
func Mysql() {
	connect := global.GConfig.Mysql
	dsn := connect.Username + ":" + connect.Password + "@(" + connect.Path + ")/" + connect.Dbname + "?" + connect.Parm
	fmt.Println(dsn)
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(fmt.Errorf("MySQL启动异常: %s", err))
	} else {
		global.GDB = db
	}
}
