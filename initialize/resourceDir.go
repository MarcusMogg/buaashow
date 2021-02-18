package initialize

import (
	"buaashow/global"
	"os"

	"go.uber.org/zap"
)

func mkdir(dirPath string) {
	if info, err := os.Stat(dirPath); err != nil {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			zap.S().Fatalf("创建目录失败: %s", err)
		}
	} else {
		if !info.IsDir() {
			zap.S().Fatal(dirPath + "已存在，但不是目录")
		}
	}
}

func resourceDir() {
	mkdir(global.GImgPath)
	mkdir(global.GStaticPath)
	mkdir(global.GCoursePath)
}
