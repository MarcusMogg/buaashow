package service

import (
	"buaashow/entity"
	"buaashow/global"
	"fmt"
)

func GetSummary(params *entity.SearchParam) []*entity.SummaryResp {
	var res []*entity.SummaryResp
	db := global.GDB
	db = db.Model(&entity.MExperiment{}).
		Joins("INNER JOIN m_submissions ON m_experiments.id = m_submissions.eid").
		Joins("INNER JOIN m_courses ON m_courses.id = m_experiments.c_id").
		Joins("INNER JOIN m_users ON m_users.account = m_submissions.g_id").
		Select(`m_courses.name,
				m_submissions.g_id,
				m_submissions.e_id,
				m_users.name as user_name,
				m_submissions.info,
				m_submissions.type`)
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
	db.Find(&res)
	for i := range res {
		res[i].URL = res[i].ShowID.Encode()
	}
	return res
}
