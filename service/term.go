package service

import (
	"buaashow/entity"
	"buaashow/global"
	"time"

	"gorm.io/gorm"
)

// table terms

// CreateTerm 创建一个新学期
func CreateTerm(term *entity.Term) error {
	begin, err := time.ParseInLocation(global.TimeTemplateDay, term.Begin, time.Local)
	if err != nil {
		return err
	}
	end, err := time.ParseInLocation(global.TimeTemplateDay, term.End, time.Local)
	if err != nil {
		return err
	}
	t := &entity.MTerm{
		TName: term.TName,
		Begin: begin,
		End:   end,
	}
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(t).Error
		if err != nil {
			return err
		}
		term.TID = t.ID
		return nil
	})
}

// DeleteTerm 删除一个学期
func DeleteTerm(id uint) error {
	t := entity.MTerm{
		ID: id,
	}
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&t).Error
	})
}

// GetTerms 获取从某年开始的所有学期
func GetTerms(year int) (res []*entity.Term) {
	begin := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)
	global.GDB.Model(&entity.MTerm{}).
		Select(`id as t_id,t_name,
			date_format(begin,'%Y-%m-%d') as begin,
			date_format(end,'%Y-%m-%d') as end`).
		Where("begin >= ?", begin.Format(global.TimeTemplateDay)).Find(&res)
	return res
}
