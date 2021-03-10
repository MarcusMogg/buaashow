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
	EID    uint   `gorm:"primarykey"`
	UID    string `gorm:"primarykey"`
	GID    string
	Status bool
}

// SummitType 允许提交的作品类型
type SummitType int

const (
	// HTML 静态网页
	HTML SummitType = iota + 1
	// EXE 可执行文件
	EXE
	// URL 外部链接
	URL
)

// MSubmission 学生作业提交
type MSubmission struct {
	EID       uint   `gorm:"primarykey"`
	GID       string `gorm:"primarykey"`
	Name      string
	Info      string
	Type      SummitType
	URL       string
	OldURL    string
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

// SubmissionResp 作业信息
type SubmissionResp struct {
	Status    bool     `json:"status"`
	Groups    []string `json:"groups"`
	UpdatedAt string   `json:"time"`
	Name      string   `json:"name"`
	Info      string   `json:"info"`
	Type      int      `json:"type"`
	URL       string   `json:"url"`
	Readme    string   `json:"readme"`
}
