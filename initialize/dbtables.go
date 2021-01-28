package initialize

import (
	"buaashow/entity"
	"buaashow/global"
	"buaashow/service"
)

// DBTables 迁移 schema
func DBTables() {
	global.GDB.AutoMigrate(&entity.MUser{})
	u := entity.MUser{
		Account:  global.GConfig.Admin.Username,
		Password: global.GConfig.Admin.Password,
		Role:     entity.Admin,
		Name:     "ADMIN",
	}
	service.Register(&u)

	global.GDB.AutoMigrate(&entity.MCourse{})
	global.GDB.AutoMigrate(&entity.RCourseStudent{})
}
