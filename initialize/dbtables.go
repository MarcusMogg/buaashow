package initialize

import (
	"buaashow/global"
	"buaashow/model/entity"
	"buaashow/service"
)

// DBTables 迁移 schema
func DBTables() {
	global.GDB.AutoMigrate(&entity.MUser{})
	u := entity.MUser{
		Account:  global.GConfig.Admin.Username,
		Password: global.GConfig.Admin.Password,
		Role:     entity.Admin,
	}
	service.Register(&u)

}
