package tests

import (
	"buaashow/global"
	"buaashow/initialize"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupGin() {
	gin.SetMode(gin.TestMode)
	tr = initialize.Router()
}

func setupDB() {
	// 临时数据库
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("db init err %s", err.Error())
	}
	global.GDB = db
	initialize.DBTables()
}

func TestMain(m *testing.M) {
	setupGin()
	setupDB()
	os.Exit(m.Run())
}
