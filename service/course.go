package service

import (
	"buaashow/entity"
	"buaashow/global"
	"buaashow/utils"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// table m_courses;
// table r_course_students;

// checkMCourseAuth 检查用户是否有操作table m_courses的权限
func checkMCourseAuth(cid uint, user *entity.MUser, auth entity.CourseAuth) bool {
	if cid == 0 {
		return false
	}
	if user.Role == entity.Admin {
		return true
	}
	var relation entity.RCourseStudent
	if err := global.GDB.Where("course_id = ? and user_id = ?", cid, user.Account).First(&relation).Error; err != nil {
		return false
	}
	return relation.Auth >= auth
}

// CreateCourse 创建课程，并创建教师与课程之间的关联
func CreateCourse(course *entity.MCourse, user *entity.MUser) (*entity.CourseResp, error) {
	relation := entity.RCourseStudent{
		UserID: user.Account,
		Auth:   entity.Owner,
	}
	var t entity.MTerm
	var cname string
	if err := global.GDB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("id = ?", course.TID).First(&t)
		if result.Error != nil {
			return errors.New("该学期不存在")
		}
		result = tx.Model(&entity.MCourseName{}).
			Where("id = ?", course.CID).
			Select("name").First(&cname)
		if result.Error != nil {
			return errors.New("课程名称不存在")
		}
		if err := tx.Create(course).Error; err != nil {
			zap.S().Debug(err)
			return err
		}
		relation.CourseID = course.ID
		err := tx.Create(&relation).Error
		if err != nil {
			zap.S().Debug(err)
		}
		return err
	}); err != nil {
		return nil, err
	}
	return &entity.CourseResp{
		ID:      course.ID,
		Name:    cname,
		Info:    course.Info,
		Teacher: course.Teacher,
		Term: entity.Term{
			TID:   course.TID,
			TName: t.TName,
			Begin: t.Begin.Format(global.TimeTemplateDay),
			End:   t.End.Format(global.TimeTemplateDay),
		},
	}, nil
}

// GetMyCourses 获取与用户相关的课程信息
func GetMyCourses(user *entity.MUser) []*entity.CourseResp {
	var res []*entity.CourseResp
	db := global.GDB
	db = db.Model(&entity.MCourse{}).
		Select(`m_courses.id,m_course_names.name,m_courses.info,m_courses.teacher,
			m_users.name as teacher_name,
			m_courses.t_id, 
			m_terms.t_name,
			date_format(m_terms.begin,'%Y-%m-%d') as begin,
			date_format(m_terms.end,'%Y-%m-%d') as end`).
		Joins("LEFT JOIN m_terms ON m_courses.t_id = m_terms.ID").
		Joins("LEFT JOIN m_users ON m_courses.teacher = m_users.account").
		Joins("LEFT JOIN m_course_names ON m_courses.c_id = m_course_names.id")
	if user.Role == entity.Admin {
		db.Find(&res)
	} else {
		db.Joins("LEFT JOIN r_course_students ON r_course_students.course_id = m_courses.id").
			Where("r_course_students.user_id = ?", user.Account).
			Find(&res)
	}
	return res
}

// GetCourseInfoByID 查询某个课程的信息
func GetCourseInfoByID(id uint) (*entity.CourseResp, error) {
	var res entity.CourseResp
	result := global.GDB.Model(&entity.MCourse{}).
		Select(`m_courses.id,m_course_names.name,m_courses.info,m_courses.teacher,
				m_users.name as teacher_name,
				m_courses.t_id, 
				m_terms.t_name,
				date_format(m_terms.begin,'%Y-%m-%d') as begin,
				date_format(m_terms.end,'%Y-%m-%d') as end`).
		Joins("LEFT JOIN m_terms ON m_courses.t_id = m_terms.ID").
		Joins("LEFT JOIN m_users ON m_courses.teacher = m_users.account").
		Joins("LEFT JOIN m_course_names ON m_courses.c_id = m_course_names.id").
		Where("m_courses.id = ?", id).
		First(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return &res, nil
}

// CreateStudentsToCourse 向课程里添加学生，如果学生账号不存在，则创建学生
func CreateStudentsToCourse(accounts []string, cid uint, user *entity.MUser) (fails []string, err error) {
	if !checkMCourseAuth(cid, user, entity.Manager) {
		err = errors.New("权限不足")
		return
	}
	var eds []uint
	global.GDB.Model(&entity.MExperiment{}).Select("id").
		Where("c_id = ?", cid).Scan(&eds)
	fails = make([]string, 0)
	basePwd := utils.AesEncrypt("666666")
	for i := range accounts {
		err = global.GDB.Transaction(func(tx *gorm.DB) error {
			var user entity.MUser
			result := tx.Where(entity.MUser{Account: accounts[i]}).
				Attrs(entity.MUser{
					Password: basePwd,
					Role:     entity.Student,
				}).
				FirstOrCreate(&user)
			if result.Error != nil {
				return result.Error
			}
			for _, id := range eds {
				tx.Create(&entity.MExperimentSubmit{
					EID:    id,
					UID:    accounts[i],
					GID:    accounts[i],
					Status: false,
				})
			}
			return tx.Create(&entity.RCourseStudent{
				CourseID: cid,
				UserID:   user.Account,
				Auth:     entity.Member,
			}).Error

		})
		if err != nil {
			fails = append(fails, accounts[i])
		}
	}
	err = nil
	return

}

// GetStudentsInCourse 获取课程里的学生信息
func GetStudentsInCourse(cid uint) []entity.UserInfoRes {
	var res []entity.UserInfoRes
	global.GDB.Model(&entity.RCourseStudent{}).
		Select("m_users.account,m_users.role,m_users.name,m_users.email").
		Joins("INNER JOIN m_users ON r_course_students.user_id = m_users.account").
		Where("r_course_students.course_id = ?  AND r_course_students.auth = ?", cid, entity.Member).
		Find(&res)
	return res
}

// DeleteStudent 删除学生
func DeleteStudent(cid uint, uid string, user *entity.MUser) error {
	if !checkMCourseAuth(cid, user, entity.Manager) {
		return errors.New("权限不足")
	}
	var relation entity.RCourseStudent
	if err := global.GDB.Where("course_id = ? and user_id = ?", cid, uid).First(&relation).Error; err != nil {
		return errors.New("无此人")
	}
	if relation.Auth >= entity.Owner {
		return errors.New("不能删除老师")
	}
	var eds []uint
	global.GDB.Model(&entity.MExperiment{}).Select("id").
		Where("c_id = ?", cid).Scan(&eds)
	for _, id := range eds {
		global.GDB.Delete(&entity.MExperimentSubmit{
			EID: id,
			UID: uid,
		})
		global.GDB.Delete(&entity.MSubmission{
			EID: id,
			GID: uid,
		})
	}
	return global.GDB.Delete(&entity.RCourseStudent{
		CourseID: cid,
		UserID:   uid,
	}).Error
}

// DeleteAllStudents 清空学生
func DeleteAllStudents(cid uint, user *entity.MUser) error {
	if !checkMCourseAuth(cid, user, entity.Owner) {
		return errors.New("权限不足")
	}

	var eds []uint
	global.GDB.Model(&entity.MExperiment{}).Select("id").
		Where("c_id = ?", cid).Scan(&eds)
	for _, id := range eds {
		global.GDB.Delete(&entity.MExperimentSubmit{
			EID: id,
		})
		global.GDB.Delete(&entity.MSubmission{
			EID: id,
		})
	}
	return global.GDB.Where("course_id = ? AND auth != ?", cid, entity.Owner).
		Delete(&entity.RCourseStudent{}).Error
}

// DeleteCourse 删除课程
func DeleteCourse(cid uint, user *entity.MUser) error {
	if !checkMCourseAuth(cid, user, entity.Owner) {
		return errors.New("权限不足")
	}
	var eds []uint
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&entity.MCourse{
			ID: cid,
		}).Error
		if err != nil {
			return err
		}
		global.GDB.Model(&entity.MExperiment{}).Select("id").
			Where("c_id = ?", cid).Scan(&eds)
		for _, id := range eds {
			global.GDB.Delete(&entity.MExperiment{
				ID: id,
			})
			global.GDB.Delete(&entity.MExperimentSubmit{
				EID: id,
			})
			global.GDB.Delete(&entity.MSubmission{
				EID: id,
			})
		}
		return tx.Where("course_id = ?", cid).Delete(&entity.RCourseStudent{}).Error
	})
}
