package entity

import (
	"time"

	"gorm.io/gorm"
)

// Term 学期
// Season 0 春 1 夏 2 秋
type Term struct {
	Year   int `json:"year"`
	Season int `json:"season"`
}

// Course 课程信息
// Name-{Year}年-{Senson}学期
// TODO: More course content
type Course struct {
	gorm.Model
	Name string
	Info string
	Term
}

// CourseStudents 课程与学生关联表
type CourseStudents struct {
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
