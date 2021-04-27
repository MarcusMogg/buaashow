package service

import (
	"buaashow/entity"
	"buaashow/global"
	"fmt"

	"go.uber.org/zap"
)

func GetSummary(params *entity.SearchParam) (int64, []*entity.SummaryResp) {
	var res []*entity.SummaryResp
	db := global.GDB
	db = db.Table("m_rec_submissions").
		Joins("INNER JOIN m_experiments ON m_experiments.id = m_rec_submissions.e_id").
		Joins("INNER JOIN m_courses ON m_courses.id = m_experiments.c_id").
		Joins("INNER JOIN m_users ON m_users.account = m_rec_submissions.g_id").
		Joins("INNER JOIN m_course_names ON m_courses.c_id = m_course_names.id")
	zap.S().Debug(db.Statement.SQL.String())
	//Where("m_submissions.show = true")
	if params.NameID != 0 {
		db = db.Where("m_courses.c_id = ?", params.NameID)
	}
	//if len(params.Recommend) != 0 {
	db = db.Where("m_rec_submissions.rec = ?", true)
	//}
	if params.TermID != 0 {
		db = db.Where("m_courses.t_id = ?", params.TermID)
	}
	if len(params.Title) != 0 {
		db = db.Where("m_rec_submissions.name LIKE ? ", fmt.Sprintf("%%%s%%", params.Title))
	}
	if params.PageNum <= 0 {
		params.PageNum = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 16
	}
	var tot int64
	offset := (params.PageNum - 1) * params.PageSize
	db.Count(&tot)
	db.Select(`m_course_names.name as course_name,
		m_rec_submissions.e_id,
		m_rec_submissions.g_id,
		m_users.name as user_name,
		m_rec_submissions.name,
		m_rec_submissions.info,
		m_rec_submissions.type,
		m_rec_submissions.thumbnail,
		m_courses.teacher`).
		Offset(offset).Limit(params.PageSize).Scan(&res)

	//db.Scan(&res)
	for i := range res {
		//zap.S().Debug(res[i])
		res[i].URL = res[i].ShowID.Encode()
		tid := res[i].Teacher
		global.GDB.Model(&entity.MUser{}).
			Where("account = ?", tid).Select("name").Scan(&res[i].Teacher)
	}
	return tot, res
}
