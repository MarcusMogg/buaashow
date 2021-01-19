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
		UserName: "admin",
		Password: "123456",
		Role:     entity.Admin,
	}
	service.Register(&u)

}
