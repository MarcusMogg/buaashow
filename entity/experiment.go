package entity

import (
	"time"
)

// MExperiment 数据库实现信息表
type MExperiment struct {
	ID        uint `gorm:"primarykey"`
	CID       uint
	Name      string
	Info      string
	BeginTime time.Time
	EndTime   time.Time
	Resources string
}

// MExperimentSubmit 学生作业提交状态
// if GID == EID,  then he can modify the group member.
// Anyone in the group can modify the submission
// TODO : 用户提交的作业内容
type MExperimentSubmit struct {
	EID uint   `gorm:"primarykey"`
	UID string `gorm:"primarykey"`
	GID string
}

// MSubmission 学生作业提交
type MSubmission struct {
	EID       uint   `gorm:"primarykey"`
	GID       string `gorm:"primarykey"`
	Name      string
	Info      string
	Type      string
	URL       string
	Readme    string
	UpdatedAt time.Time
}

// ExperimentResponse 完整的实验信息
type ExperimentResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Info string `json:"info"`
	// course info
	CourseID    uint   `json:"cid"`
	CourseName  string `json:"cname"`
	TeacherName string `json:"teacher"`
	// YYYY-MM-DD HH-MM-SS
	BeginTime string `json:"begin"`
	EndTime   string `json:"end"`

	Resources []string `json:"resources"`
}

// ExperimentReq 创建 or 修改实验
type ExperimentReq struct {
	Name string `json:"name"`
	Info string `json:"info"`
	// YYYY-MM-DD HH-MM-SS
	BeginTime string `json:"begin"`
	EndTime   string `json:"end"`

	Resources []string `json:"resources"`
}
