package service

import (
	"buaashow/entity"
	"buaashow/global"
	"errors"

	"gorm.io/gorm"
)

// table terms

// CreateTerm 创建一个新学期
func CreateTerm(term *entity.Term) error {
	t := &entity.MTerm{
		Term: (*term),
	}
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("year = ? and season = ?", t.Year, t.Season).
			First(&entity.MTerm{})
		if result.Error == nil {
			return errors.New("该学期已存在")
		}
		return tx.Create(t).Error
	})
}

// DeleteTerm 删除一个学期
func DeleteTerm(term *entity.Term) error {
	t := &entity.MTerm{
		Term: (*term),
	}
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("year = ? and season = ?", t.Year, t.Season).
			First(t)
		if result.Error != nil {
			return errors.New("该学期不存在")
		}
		return tx.Delete(t).Error
	})
}

// GetTerms 获取从某年开始的所有学期
func GetTerms(year int) (res []entity.Term) {
	global.GDB.Model(&entity.MTerm{}).
		Where("year >= ?", year).Find(&res)
	return res
}
