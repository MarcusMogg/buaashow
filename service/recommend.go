package service

import (
	"buaashow/entity"
	"buaashow/global"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func sub2rec(s *entity.MSubmission, r *entity.MRecSubmission) {
	r.EID = s.EID
	r.GID = s.GID
	r.Name = s.Name
	r.Info = s.Info
	r.Type = s.Type
	sid := entity.ShowID{
		EID: s.EID, GID: s.GID,
	}
	if s.Type == entity.HTML {
		r.URL = fmt.Sprintf("show/x/%s/index.html", sid.Encode())
	} else if s.Type == entity.EXE {
		r.URL = fmt.Sprintf("show/x/%s/release.zip", sid.Encode())
	}
	r.Thumbnail = s.Thumbnail
	r.Readme = s.Readme
	r.UpdatedAt = s.UpdatedAt
}

func Reccommend(eid uint, uid string, teacher *entity.MUser) error {
	exp, err := GetMExp(eid)
	if err != nil {
		return err
	}
	if !checkMCourseAuth(exp.CID, teacher, entity.Owner) {
		return errors.New("权限不足")
	}
	var mid entity.MExperimentSubmit
	if err := global.GDB.Where("e_id = ? AND uid = ?", eid, uid).
		First(&mid).Error; err != nil {
		return err
	}
	if !mid.Status {
		return errors.New("未提交")
	}
	var sub entity.MSubmission
	if err = global.GDB.Where("e_id = ? AND g_id = ?", eid, mid.GID).
		First(&sub).Error; err != nil {
		return err
	}
	var rec entity.MRecSubmission
	global.GDB.Where("e_id = ? AND g_id = ?", eid, mid.GID).
		First(&rec)
	src := filepath.Join(global.GCoursePath, fmt.Sprintf("%d", sub.EID), mid.GID, "show")
	dist := filepath.Join(global.GShowPath, fmt.Sprintf("%d", sub.EID), mid.GID)
	if rec.EID == 0 {
		sub2rec(&sub, &rec)
		rec.Rec = true
		os.RemoveAll(dist)
		if err = CopyDir(src, dist); err != nil {
			return err
		}
		return global.GDB.Create(&rec).Error
	} else if sub.UpdatedAt.After(rec.UpdatedAt) {
		rec.Rec = true
		os.RemoveAll(dist)
		if err = CopyDir(src, dist); err != nil {
			return err
		}
		rec.Rec = true
		return global.GDB.Save(&rec).Error
	} else {
		// FIXME: if file has been deleted?
		return global.GDB.Model(&entity.MRecSubmission{}).
			Where("e_id = ? AND g_id = ?", eid, mid.GID).
			Update("rec", true).Error
	}
}

func Unrec(eid uint, uid string, teacher *entity.MUser) error {
	exp, err := GetMExp(eid)
	if err != nil {
		return err
	}
	if !checkMCourseAuth(exp.CID, teacher, entity.Owner) {
		return errors.New("权限不足")
	}
	var mid entity.MExperimentSubmit
	if err := global.GDB.Where("e_id = ? AND uid = ?", eid, uid).
		First(&mid).Error; err != nil {
		return err
	}
	if !mid.Status {
		return errors.New("未提交")
	}
	return global.GDB.Model(&entity.MRecSubmission{}).
		Where("e_id = ? AND g_id = ?", eid, mid.GID).
		Update("rec", false).Error
}

func CheckInTeam(uid, gid string, eid uint) bool {
	return global.GDB.
		Where("e_id = ? AND uid = ? AND g_id = ?", eid, uid, gid).
		First(&entity.MExperimentSubmit{}).
		Error == nil
}

func CheckRecommend(gid string, eid uint) bool {
	return global.GDB.
		Where("e_id = ? AND g_id = ? AND rec = ?", eid, gid, true).
		First(&entity.MRecSubmission{}).
		Error == nil
}

// GetRecSubmission 获取推荐的提交信息
func GetRecSubmission(eid uint, uid string, res *entity.SubmissionResp) error {
	var mid entity.MExperimentSubmit
	var sub entity.MRecSubmission
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
	res.UpdatedAt = sub.UpdatedAt.Format(global.TimeTemplateSec)
	res.Name = sub.Name
	res.Info = sub.Info
	res.Type = int(sub.Type)
	res.URL = sub.URL
	res.Readme = sub.Readme
	res.Thumbnail = sub.Thumbnail

	return nil
}
