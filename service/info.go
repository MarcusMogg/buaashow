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
	db = db.Table("m_submissions").
		Select(`m_courses.name as course_name,
        m_submissions.e_id,
        m_submissions.g_id,
        m_users.name as user_name,
        m_submissions.name,
        m_submissions.info,
        m_submissions.type`).
		Joins("INNER JOIN m_experiments ON m_experiments.id = m_submissions.e_id").
		Joins("INNER JOIN m_courses ON m_courses.id = m_experiments.c_id").
		Joins("INNER JOIN m_users ON m_users.account = m_submissions.g_id")
	zap.S().Debug(db.Statement.SQL.String())
	//Where("m_submissions.show = true")
	if len(params.CourseName) != 0 {
		db = db.Where("m_courses.name = ?", params.CourseName)
	}
	if len(params.Recommend) != 0 {
		db = db.Where("m_submissions.recommend = true")
	}
	if params.TermID != 0 {
		db = db.Where("m_courses.t_id = ?", params.TermID)
	}
	if len(params.Title) != 0 {
		db = db.Where("m_submissions.name LIKE ? ", fmt.Sprintf("%%%s%%", params.Title))
	}
	if params.PageNum <= 0 {
		params.PageNum = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 16
	}
	var tot int64
	offset := (params.PageNum - 1) * params.PageSize

	db.Offset(offset).Limit(params.PageSize).Scan(&res).Count(&tot)
	//db.Scan(&res)
	for i := range res {
		//zap.S().Debug(res[i])
		res[i].URL = res[i].ShowID.Encode()
	}
	return tot, res
}
