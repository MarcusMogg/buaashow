package service

import (
	"buaashow/entity"
	"buaashow/global"

	"gorm.io/gorm"
)

// CreateCourseNmae 创建一个coursename
func CreateCourseNmae(name string) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&entity.MCourseName{
			Name: name,
		}).Error
	})
}

// DeleteCourseNmae 删除一个coursename
func DeleteCourseNmae(name string) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&entity.MCourseName{
			Name: name,
		}).Error
	})
}

// GetAllCourseNmae 获取所有coursename
func GetAllCourseNmae() []string {
	var res []string
	global.GDB.Model(&entity.MCourseName{}).Select("name").Find(&res)
	return res
}
