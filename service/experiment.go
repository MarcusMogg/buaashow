package service

import (
	"buaashow/entity"
	"buaashow/global"

	"gorm.io/gorm"
)

// CreateExp 创建实验
func CreateExp(e *entity.MExperiment) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(e).Error; err != nil {
			return err
		}
		// TODO: 实验相关的初始化内容，GID……
		return nil
	})
}

// GetExpsByCID 获取和课程CID相关联的所有实验
func GetExpsByCID(cid uint) {

}
