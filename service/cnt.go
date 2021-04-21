package service

import (
	"buaashow/entity"
	"buaashow/global"
)

func Total() *entity.Total {
	var res entity.Total
	global.GDB.Model(&entity.MCourse{}).Count(&res.Course)
	global.GDB.Model(&entity.MUser{}).Count(&res.User)
	global.GDB.Model(&entity.MExperiment{}).Count(&res.Exp)
	global.GDB.Model(&entity.MSubmission{}).Count(&res.Submit)
	global.GDB.Model(&entity.MRecSubmission{}).Count(&res.Rec)
	return &res
}
