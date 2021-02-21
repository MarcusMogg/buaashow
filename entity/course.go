package entity

import (
	"gorm.io/gorm"
)

// Term 学期
type Term struct {
	Year   int `json:"year" binding:"required"`
	Season int `json:"season" binding:"required,gte=1,lte=2"`
}

// MTerm 学期
// Season 1 春 2 秋
type MTerm struct {
	gorm.Model `json:"-"`
	Term
}

// MCourse 课程信息
// Name-{Year}年-{Senson}学期
// TODO: More course content
type MCourse struct {
	ID   uint `gorm:"primarykey"`
	TID  uint
	Name string
	Info string
}

// CourseResp 接口返回的课程信息
type CourseResp struct {
	ID   uint   `json:"cid"`
	Name string `json:"name"`
	Info string `json:"info"`
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
