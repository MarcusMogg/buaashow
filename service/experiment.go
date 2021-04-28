package service

import (
	"buaashow/entity"
	"buaashow/global"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// CreateExp 创建实验
func CreateExp(e *entity.MExperiment, user *entity.MUser) error {
	if !checkMCourseAuth(e.CID, user, entity.Owner) {
		return errors.New("权限不足")
	}
	var rs []entity.RCourseStudent
	global.GDB.Model(&entity.RCourseStudent{}).
		Where("course_id = ? AND auth = ? ", e.CID, entity.Member).Find(&rs)

	return global.GDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(e).Error; err != nil {
			return err
		}
		for _, i := range rs {
			if i.Auth == entity.Owner {
				continue
			}
			if err := tx.Create(&entity.MExperimentSubmit{
				EID:    e.ID,
				UID:    i.UserID,
				GID:    i.UserID,
				Status: false,
			}).Error; err != nil {
				return err
			}
		}
		return initWorker(e)
	})
}

// UpdateExp 修改实验
func UpdateExp(e *entity.MExperiment, user *entity.MUser) error {
	if !checkMCourseAuth(e.CID, user, entity.Owner) {
		return errors.New("权限不足")
	}
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(e).Error; err != nil {
			return err
		}
		//updateEndtime(e)
		return nil
	})
}

func AddExpFile(eid uint, user *entity.MUser, filename string) error {
	exp, err := GetMExp(eid)
	if err != nil {
		return err
	}
	if !checkMCourseAuth(exp.CID, user, entity.Owner) {
		return errors.New("权限不足")
	}
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		return tx.FirstOrCreate(&entity.RExperimentResource{},
			entity.RExperimentResource{
				EID:  eid,
				File: filename,
			}).Error
	})
}

func DeleteExpFile(eid uint, user *entity.MUser, filename string) error {
	exp, err := GetMExp(eid)
	if err != nil {
		return err
	}
	if !checkMCourseAuth(exp.CID, user, entity.Owner) {
		return errors.New("权限不足")
	}
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&entity.RExperimentResource{
			EID:  eid,
			File: filename,
		}).Error
	})
}

func expToResp(i *entity.MExperiment) (*entity.ExperimentResponse, error) {
	var course entity.MCourse
	var term entity.MTerm
	var ca entity.RCourseStudent
	var resources []string
	var cname string
	err := global.GDB.Where("id = ?", i.CID).First(&course).Error
	if err != nil {
		return nil, err
	}
	err = global.GDB.Where("id = ?", course.TID).First(&term).Error
	if err != nil {
		return nil, err
	}
	err = global.GDB.Where("course_id = ? AND auth = ?",
		i.CID, entity.Owner).First(&ca).Error
	if err != nil {
		return nil, err
	}
	var teacherName string
	global.GDB.Model(&entity.MUser{}).
		Where("account = ?", ca.UserID).
		Select("name").
		First(&teacherName)
	if len(teacherName) == 0 {
		teacherName = ca.UserID
	}
	global.GDB.Model(&entity.RExperimentResource{}).
		Select("file").Where("e_id = ?", i.ID).Find(&resources)
	global.GDB.Model(&entity.MCourseName{}).
		Where("id = ?", course.CID).
		Select("name").First(&cname)
	return &entity.ExperimentResponse{
		ID:          i.ID,
		Name:        i.Name,
		Info:        i.Info,
		Team:        i.Team,
		CourseID:    i.CID,
		CourseName:  cname,
		Teacher:     ca.UserID,
		TeacherName: teacherName,
		TermID:      course.TID,
		TermName:    term.TName,
		//BeginTime:   i.BeginTime.Format(global.TimeTemplateSec),
		//EndTime:     i.EndTime.Format(global.TimeTemplateSec),
		Resources: resources,
	}, nil
}

// GetExpsByCID 获取和课程CID相关联的所有实验
func GetExpsByCID(cid uint) ([]*entity.ExperimentResponse, error) {
	var res []entity.MExperiment
	var resp []*entity.ExperimentResponse
	global.GDB.Model(&entity.MExperiment{}).Where("c_id = ?", cid).Find(&res)
	for _, i := range res {
		tmp, err := expToResp(&i)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}

	return resp, nil
}

// GetMExp 获取指定实验
func GetMExp(eid uint) (*entity.MExperiment, error) {
	var res entity.MExperiment
	err := global.GDB.Where("id = ?", eid).First(&res).Error
	return &res, err
}

// GetExp 获取指定实验
func GetExp(eid uint) (*entity.ExperimentResponse, error) {
	res, err := GetMExp(eid)
	if err != nil {
		return nil, err
	}
	return expToResp(res)
}

// DeleteExp 删除指定实验
func DeleteExp(eid uint, user *entity.MUser) error {
	exp, err := GetMExp(eid)
	if err != nil {
		return err
	}
	if !checkMCourseAuth(exp.CID, user, entity.Owner) {
		return errors.New("权限不足")
	}
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", eid).Delete(&entity.MExperiment{}).Error; err != nil {
			return err
		}
		if err := tx.Where("e_id = ?", eid).Delete(&entity.MExperimentSubmit{}).Error; err != nil {
			return err
		}
		if err := tx.Where("e_id = ?", eid).Delete(&entity.MSubmission{}).Error; err != nil {
			return err
		}
		return nil
	})
}

// Submit 提交作业
// only uid == gid can submit
func Submit(s *entity.MSubmission, uid string, user *entity.MUser) error {
	if user.Role == entity.Student {
		uid = user.Account
	} else {
		if len(uid) == 0 {
			return errors.New("参数错误")
		}
	}
	var mid entity.MExperimentSubmit
	var exp entity.MExperiment
	var cid uint

	return global.GDB.Transaction(func(tx *gorm.DB) error {
		tx.Model(&entity.MExperiment{}).Select("c_id").Where("id = ?", s.EID).Scan(&cid)
		if user.Role != entity.Student {
			if !checkMCourseAuth(cid, user, entity.Owner) {
				return errors.New("权限不足")
			}
		}

		if err := tx.Where("e_id = ? AND uid = ?", s.EID, uid).
			First(&mid).Error; err != nil {
			// FIXME : really need it?
			if err == gorm.ErrRecordNotFound {
				if tx.Where("c_id = ? AND uid = ?", cid, uid).First(&entity.RCourseStudent{}).Error == nil {
					mid = entity.MExperimentSubmit{
						EID:    s.EID,
						UID:    uid,
						GID:    uid,
						Status: false,
					}
					err := tx.Create(&mid).Error
					if err != nil {
						return err
					}
				} else {
					return err
				}
			} else {
				return err
			}
		}
		if mid.GID != uid {
			return errors.New("权限不足")
		}
		s.GID = mid.GID

		if err := tx.Where("id = ?", s.EID).
			Select("id").First(&exp).Error; err != nil {
			return err
		}
		// if s.UpAt.Before(exp.BeginTime) || s.UpAt.After(exp.EndTime) {
		// 	return errors.New("不在时间区间")
		// }

		var url struct {
			SrcURL  string
			DistURL string
		}
		if mid.Status {
			global.GDB.Model(&entity.MSubmission{}).
				Where("e_id = ? AND g_id = ?", s.EID, s.GID).
				Select("src_url,dist_url").First(&url)
			zap.S().Debug(url)
		}
		// 省略一次操作
		if len(s.DistURL) != 0 && url.DistURL != s.DistURL {
			if s.Type == entity.HTML {
				if err := toWorker(s.EID, s.GID, s.DistURL, toZip); err != nil {
					return err
				}
			} else if s.Type == entity.EXE {
				if err := toWorker(s.EID, s.GID, s.DistURL, toExe); err != nil {
					return err
				}
			}
		}
		if len(s.SrcURL) != 0 && url.SrcURL != s.SrcURL {
			if err := toWorker(s.EID, s.GID, s.SrcURL, toSrc); err != nil {
				return err
			}
		}
		sid := entity.ShowID{
			EID: s.EID, GID: s.GID,
		}
		if s.Type == entity.HTML {
			s.URL = fmt.Sprintf("show/preview/x/%s/index.html", sid.Encode())
		} else if s.Type == entity.EXE {
			s.URL = fmt.Sprintf("show/preview/x/%s/release.zip", sid.Encode())
		}

		if mid.Status {
			return tx.Updates(s).Error
		}
		if err := tx.Create(s).Error; err != nil {
			return err
		}
		mid.Status = true
		mid.UpAt = s.UpAt
		return tx.Save(&mid).Error
	})
}

// GetSubmission 获取提交信息
func GetSubmission(eid uint, uid string, res *entity.SubmissionResp) error {
	var mid entity.MExperimentSubmit
	var sub entity.MSubmission
	if err := global.GDB.Where("e_id = ? AND uid = ?", eid, uid).
		First(&mid).Error; err != nil {
		return err
	}
	var groups []*entity.UserInfoSimple

	if err := global.GDB.Model(&entity.MExperimentSubmit{}).
		Select("m_users.account,m_users.name").
		Joins("INNER JOIN m_users ON m_experiment_submits.uid = m_users.account").
		Where("m_experiment_submits.e_id = ? AND m_experiment_submits.g_id = ?", eid, mid.GID).
		Find(&groups).Error; err != nil {
		return err
	}
	res.Groups = groups
	res.GID = mid.GID

	if !mid.Status {
		res.Status = false
		return nil
	}
	if err := global.GDB.Where("e_id = ? AND g_id = ?", eid, mid.GID).
		First(&sub).Error; err != nil {
		return err
	}

	res.Status = true
	res.UpAt = sub.UpAt.Format(global.TimeTemplateSec)
	res.Name = sub.Name
	res.Info = sub.Info
	res.Type = int(sub.Type)
	res.URL = sub.URL
	res.Readme = sub.Readme
	res.Thumbnail = sub.Thumbnail
	sid := entity.ShowID{
		EID: eid,
		GID: mid.GID,
	}
	res.ShowID = sid.Encode()
	return nil
}

func GetAllSubmission(eid uint, user *entity.MUser) ([]*entity.SubmissionResp, error) {
	exp, err := GetMExp(eid)
	if err != nil {
		return nil, err
	}
	if !checkMCourseAuth(exp.CID, user, entity.Owner) {
		return nil, errors.New("权限不足")
	}

	var rs []entity.RCourseStudent
	global.GDB.Model(&entity.RCourseStudent{}).
		Where("course_id = ? AND auth = ? ", exp.CID, entity.Member).Find(&rs)

	getSub := func(eid uint, uid string) (*entity.SubmissionResp, error) {
		res := &entity.SubmissionResp{}
		res.StudentID = uid
		var mid entity.MExperimentSubmit

		if err := global.GDB.Where("e_id = ? AND uid = ?", eid, uid).
			First(&mid).Error; err != nil {
			return res, err
		}

		var groups []*entity.UserInfoSimple

		if err := global.GDB.Model(&entity.MExperimentSubmit{}).
			Select("m_users.account,m_users.name").
			Joins("INNER JOIN m_users ON m_experiment_submits.uid = m_users.account").
			Where("m_experiment_submits.e_id = ? AND m_experiment_submits.g_id = ?", eid, mid.GID).
			Find(&groups).Error; err != nil {
			return nil, err
		}
		res.Status = mid.Status
		res.Groups = groups
		res.GID = mid.GID

		if mid.Status {
			var sub entity.MSubmission
			if err := global.GDB.Where("e_id = ? AND g_id = ?", eid, mid.GID).
				Select("up_at").
				First(&sub).Error; err != nil {
				return res, err
			}
			res.UpAt = sub.UpAt.Format(global.TimeTemplateSec)

			var rec entity.MRecSubmission
			if err := global.GDB.Where("e_id = ? AND g_id = ? AND rec = ?", eid, mid.GID, true).
				Select("up_at").
				First(&rec).Error; err == nil {
				res.RecAt = rec.UpAt.Format(global.TimeTemplateSec)
			}
		}
		sid := entity.ShowID{
			EID: eid,
			GID: mid.GID,
		}
		res.ShowID = sid.Encode()
		return res, nil
	}

	res := make([]*entity.SubmissionResp, len(rs))
	for i, v := range rs {
		tmp, err := getSub(eid, v.UserID)
		res[i] = tmp
		if err != nil {
			zap.S().Debugf("get submission error for eid:%d,uid:%s,%s",
				eid, v.UserID, err.Error())
		}
	}
	return res, nil
}

func TeamInfo(eid uint, uid string) (string, bool, error) {
	var mid entity.MExperimentSubmit
	var exp entity.MExperiment
	if err := global.GDB.Where("id = ?", eid).
		Select("team").
		First(&exp).Error; err != nil {
		return "", false, err
	}
	if !exp.Team {
		return "", false, errors.New("此课程不允许组队")
	}
	if err := global.GDB.Where("e_id = ? AND uid = ?", eid, uid).
		First(&mid).Error; err != nil {
		return "", false, err
	}
	res := mid.GID
	var teamMember int64
	inTeam := true
	if res == uid {
		global.GDB.Model(&entity.MExperimentSubmit{}).
			Where("e_id = ? AND g_id = ?", eid, res).
			Count(&teamMember)
		inTeam = (teamMember > 1)
	}
	return res, inTeam, nil
}

func JoinTeam(eid uint, uid, gid string) error {
	_, in, err := TeamInfo(eid, uid)
	if err != nil {
		return err
	}
	if in {
		return errors.New("已经组队")
	}
	return global.GDB.Model(&entity.MExperimentSubmit{}).
		Where("e_id = ? AND uid = ?", eid, uid).
		Update("g_id", gid).Error
}

func QiutTeam(eid uint, uid, gid string) error {
	return global.GDB.Model(&entity.MExperimentSubmit{}).
		Where("e_id = ? AND uid = ?", eid, uid).
		Update("g_id", uid).Error
}

func AttrSubmitStatus(res []*entity.ExperimentResponse, uid string) {
	for i := range res {
		var mid entity.MExperimentSubmit
		if err := global.GDB.Where("e_id = ? AND uid = ?", res[i].ID, uid).
			First(&mid).Error; err == nil {
			res[i].Submit = mid.Status
		}
	}
}
