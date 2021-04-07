package entity

import "time"

// Term 学期
type Term struct {
	TID   uint   `json:"tid"`
	TName string `json:"tname" binding:"required"`
	Begin string `json:"tbegin"`
	End   string `json:"tend"`
}

// MTerm 学期
type MTerm struct {
	ID    uint `gorm:"primarykey"`
	TName string
	Begin time.Time
	End   time.Time
}

// MCourseName 限定课程名称，用于展示
type MCourseName struct {
	Name string `gorm:"primarykey"`
}

// MCourse 课程信息
// Name-{Year}年-{Senson}学期
// TODO: More course content
type MCourse struct {
	ID   uint `gorm:"primarykey"`
	TID  uint
	Name string
	Info string
	// Teacher Account
	Teacher string
}

// CourseResp 接口返回的课程信息
type CourseResp struct {
	ID          uint   `json:"cid"`
	Name        string `json:"name"`
	Info        string `json:"info"`
	Teacher     string `json:"teacher"`
	TeacherName string `json:"teacher_name"`
	Term
}

// RCourseStudent 课程与学生关联表
type RCourseStudent struct {
	CourseID uint   `gorm:"primarykey"`
	UserID   string `gorm:"primarykey"`
	Auth     CourseAuth
}

// CourseAuth 用户对课程的权限
type CourseAuth int

const (
	// Member 学生 可以查看作业，提交作业
	Member CourseAuth = iota + 1
	// Manager TA权限 可以发布作业、下载作业
	Manager
	// Owner 管理员 可以赋予权限，删除课程
	Owner
)
