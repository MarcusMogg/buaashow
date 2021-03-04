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
func checkMCourseAuth(cid uint, uid string, auth entity.CourseAuth) bool {
	var relation entity.RCourseStudent
	if err := global.GDB.Where("course_id = ? and user_id = ?", cid, uid).First(&relation).Error; err != nil {
		return false
	}
	return relation.Auth >= auth
}

// CreateCourse 创建课程，并创建教师与课程之间的关联
func CreateCourse(term *entity.Term, course *entity.MCourse, user *entity.MUser) error {
	relation := entity.RCourseStudent{
		UserID: user.Account,
		Auth:   entity.Owner,
	}
	t := &entity.MTerm{
		Term: (*term),
	}
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("year = ? and season = ?", t.Year, t.Season).
			First(t)
		if result.Error != nil {
			return errors.New("该学期不存在")
		}
		course.TID = t.ID
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
	})
}

// GetMyCourses 获取与用户相关的课程信息
func GetMyCourses(user *entity.MUser) []entity.CourseResp {
	var res []entity.CourseResp
	global.GDB.Model(&entity.MCourse{}).
		Select("m_courses.id,m_courses.name,m_courses.info,m_terms.year,m_terms.season").
		Joins("INNER JOIN r_course_students ON r_course_students.course_id = m_courses.id").
		Joins("INNER JOIN m_terms ON m_courses.t_id = m_terms.ID").
		Where("r_course_students.user_id = ?", user.Account).
		Find(&res)
	return res
}

// GetCourseInfoByID 查询某个课程的信息
func GetCourseInfoByID(id uint) (*entity.CourseResp, error) {
	var res entity.CourseResp
	result := global.GDB.Model(&entity.MCourse{}).
		Select("m_courses.id,m_courses.name,m_courses.info,m_terms.year,m_terms.season").
		Joins("INNER JOIN m_terms ON m_courses.t_id = m_terms.ID").
		Where("m_courses.id = ?", id).
		First(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return &res, nil
}

// CreateStudentsToCourse 向课程里添加学生，如果学生账号不存在，则创建学生
func CreateStudentsToCourse(accounts []string, cid uint, uid string) (fails []string, err error) {
	if !checkMCourseAuth(cid, uid, entity.Manager) {
		err = errors.New("权限不足")
		return
	}
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
		Where("r_course_students.course_id = ?", cid).
		Find(&res)
	return res
}

// DeleteStudent 删除学生
func DeleteStudent(cid uint, uid string, user *entity.MUser) error {
	if !checkMCourseAuth(cid, user.Account, entity.Manager) {
		return errors.New("权限不足")
	}
	if uid == user.Account {
		return errors.New("不能删除老师")
	}
	return global.GDB.Delete(&entity.RCourseStudent{
		CourseID: cid,
		UserID:   uid,
	}).Error
}

// DeleteCourse 删除课程
func DeleteCourse(cid uint, user *entity.MUser) error {
	if !checkMCourseAuth(cid, user.Account, entity.Owner) {
		return errors.New("权限不足")
	}

	return global.GDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&entity.MCourse{
			ID: cid,
		}).Error
		if err != nil {
			return err
		}
		return tx.Where("course_id = ?", cid).Delete(&entity.RCourseStudent{}).Error
	})
}
