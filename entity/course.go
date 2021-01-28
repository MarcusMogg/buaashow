package entity

import (
	"time"

	"gorm.io/gorm"
)

// Term 学期
// Season 0 春 1 夏 2 秋
type Term struct {
	Year   int `json:"year" binding:"gte=2020"`
	Season int `json:"season" binding:"gte=0,lte=2"`
}

// MCourse 课程信息
// Name-{Year}年-{Senson}学期
// TODO: More course content
type MCourse struct {
	gorm.Model
	Name string
	Info string
	Term
}

// CourseResp 接口返回的课程信息
type CourseResp struct {
	ID   uint   `json:"cid"`
	Name string `json:"name" binding:"require,gte=4,lte=32"`
	Info string `json:"info"`
	Term
}

// RCourseStudent 课程与学生关联表
type RCourseStudent struct {
	CourseID  uint `gorm:"primarykey"`
	UserID    uint `gorm:"primarykey"`
	Auth      CourseAuth
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
