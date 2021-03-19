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
func CreateExp(e *entity.MExperiment, uid string) error {
	if !checkMCourseAuth(e.CID, uid, entity.Owner) {
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
func UpdateExp(e *entity.MExperiment, uid string) error {
	if !checkMCourseAuth(e.CID, uid, entity.Owner) {
		return errors.New("权限不足")
	}
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(e).Error; err != nil {
			return err
		}
		updateEndtime(e)
		return nil
	})
}

func AddExpFile(eid uint, uid, filename string) error {
	if !checkMCourseAuth(eid, uid, entity.Owner) {
		return errors.New("权限不足")
	}
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		return tx.FirstOrCreate(&entity.MExperimentResource{},
			entity.MExperimentResource{
				EID:  eid,
				File: filename,
			}).Error
	})
}

func DeleteExpFile(eid uint, uid, filename string) error {
	if !checkMCourseAuth(eid, uid, entity.Owner) {
		return errors.New("权限不足")
	}
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&entity.MExperimentResource{
			EID:  eid,
			File: filename,
		}).Error
	})
}

func expToResp(i *entity.MExperiment) (*entity.ExperimentResponse, error) {
	var course entity.MCourse
	var ca entity.RCourseStudent
	var resources []string
	err := global.GDB.Where("id = ?", i.CID).First(&course).Error
	if err != nil {
		return nil, err
	}
	err = global.GDB.Where("course_id = ? AND auth = ?",
		i.CID, entity.Owner).First(&ca).Error
	if err != nil {
		return nil, err
	}
	global.GDB.Model(&entity.MExperimentResource{}).
		Select("file").Where("e_id = ?", i.ID).Find(&resources)

	return &entity.ExperimentResponse{
		ID:          i.ID,
		Name:        i.Name,
		Info:        i.Info,
		CourseID:    i.CID,
		CourseName:  course.Name,
		TeacherName: ca.UserID,
		BeginTime:   i.BeginTime.Format(global.TimeTemplateSec),
		EndTime:     i.EndTime.Format(global.TimeTemplateSec),
		Resources:   resources,
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
func DeleteExp(eid uint, uid string) error {
	exp, err := GetMExp(eid)
	if err != nil {
		return err
	}
	if !checkMCourseAuth(exp.CID, uid, entity.Owner) {
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
func Submit(s *entity.MSubmission, uid string) error {
	var mid entity.MExperimentSubmit
	var exp entity.MExperiment
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("e_id = ? AND uid = ?", s.EID, uid).
			First(&mid).Error; err != nil {
			return err
		}
		if mid.GID != uid {
			return errors.New("权限不足")
		}
		s.GID = mid.GID

		if err := tx.Where("id = ?", s.EID).
			Select("begin_time,end_time").First(&exp).Error; err != nil {
			return err
		}
		if s.UpdatedAt.Before(exp.BeginTime) || s.UpdatedAt.After(exp.EndTime) {
			return errors.New("不在时间区间")
		}

		var oldURL struct {
			OldURL string
		}
		if mid.Status {
			global.GDB.Model(&entity.MSubmission{}).
				Where("e_id = ? AND g_id = ?", s.EID, s.GID).
				Select("old_url").First(&oldURL)
			zap.S().Debug(oldURL)
		}
		// 省略一次操作
		if oldURL.OldURL != s.URL {
			if s.Type == entity.HTML {
				if err := ToUnzip(s.EID, s.GID, s.URL); err != nil {
					return err
				}
			} else if s.Type == entity.EXE {
				if err := moveExe(s.EID, s.GID, s.URL); err != nil {
					return err
				}
			}
		}
		s.OldURL = s.URL
		if s.Type == entity.HTML {
			s.URL = fmt.Sprintf("show/%d/%s/index.html", s.EID, s.GID)
		} else if s.Type == entity.EXE {
			s.URL = fmt.Sprintf("show/%d/%s/release.zip", s.EID, s.GID)
		}

		if mid.Status {
			return tx.Save(s).Error
		}
		if err := tx.Create(s).Error; err != nil {
			return err
		}
		mid.Status = true
		mid.UpdatedAt = s.UpdatedAt
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
	if err := global.GDB.Where("e_id = ? AND g_id = ?", eid, mid.GID).
		First(&sub).Error; err != nil {
		return err
	}
	var groups []entity.MExperimentSubmit
	if err := global.GDB.Model(&entity.MExperimentSubmit{}).
		Where("e_id = ? AND g_id = ?", eid, mid.GID).
		Find(&groups).Error; err != nil {
		return err
	}
	res.Status = true
	res.UpdatedAt = sub.UpdatedAt.Format(global.TimeTemplateSec)
	res.Name = sub.Name
	res.Info = sub.Info
	res.Type = int(sub.Type)
	res.URL = sub.URL
	res.Readme = sub.Readme
	for _, i := range groups {
		res.Groups = append(res.Groups, i.UID)
	}
	return nil
}

func GetAllSubmission(eid uint, uid string) ([]*entity.SubmissionResp, error) {
	exp, err := GetMExp(eid)
	if err != nil {
		return nil, err
	}
	if !checkMCourseAuth(exp.CID, uid, entity.Owner) {
		return nil, errors.New("权限不足")
	}

	var rs []entity.RCourseStudent
	global.GDB.Model(&entity.RCourseStudent{}).
		Where("course_id = ? AND auth = ? ", exp.CID, entity.Member).Find(&rs)

	getSub := func(eid uint, uid string) (*entity.SubmissionResp, error) {
		res := &entity.SubmissionResp{}
		var mid entity.MExperimentSubmit
		var sub entity.MSubmission
		if err := global.GDB.Where("e_id = ? AND uid = ?", eid, uid).
			First(&mid).Error; err != nil {
			return nil, err
		}
		if err := global.GDB.Where("e_id = ? AND g_id = ?", eid, mid.GID).
			First(&sub).Error; err != nil {
			return nil, err
		}
		res.Status = true
		res.UpdatedAt = sub.UpdatedAt.Format(global.TimeTemplateSec)
		return nil, nil
	}

	res := make([]*entity.SubmissionResp, len(rs))
	for i, v := range rs {
		tmp, err := getSub(eid, v.UserID)
		if err == nil {
			res[i] = tmp
		} else {
			zap.S().Debugf("get submission error for eid:%d,uid:%s,%s",
				eid, v.UserID, err.Error())
		}
	}
	return res, nil
}
