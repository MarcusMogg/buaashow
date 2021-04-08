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
func DeleteCourseNmae(id uint) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&entity.MCourseName{
			ID: id,
		}).Error
	})
}

// GetAllCourseNmae 获取所有coursename
func GetAllCourseNmae() []*entity.MCourseName {
	var res []*entity.MCourseName
	global.GDB.Find(&res)
	return res
}

func GetCourseName(id uint) (*entity.MCourseName, error) {
	res := &entity.MCourseName{
		ID: id,
	}
	err := global.GDB.First(res).Error
	return res, err
}

func UpdateCourseInfo(id uint, info string) error {
	return global.GDB.Model(&entity.MCourseName{}).
		Where("id = ?", id).
		Update("info", info).Error
}

func UpdateCourseName(id uint, name string) error {
	return global.GDB.Model(&entity.MCourseName{}).
		Where("id = ?", id).
		Update("name", name).Error
}

func UpdateCourseThumb(id uint, path string) error {
	return global.GDB.Model(&entity.MCourseName{}).
		Where("id = ?", id).
		Update("thumbnail", path).Error
}
