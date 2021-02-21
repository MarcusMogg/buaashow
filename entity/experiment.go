package entity

import "time"

// MExperiment 数据库实现信息表
type MExperiment struct {
	ID        uint `gorm:"primarykey"`
	CID       uint
	Name      string
	Info      string
	BeginTime time.Time
	EndTime   time.Time
}

// MExperimentResource 参考资源
type MExperimentResource struct {
	EID  uint
	Path string
}

// MExperimentSubmit 学生作业提交状态
// if GID == EID,  then he can modify the group member.
// Anyone in the group can modify the submission
// TODO : 用户提交的作业内容
type MExperimentSubmit struct {
	EID    uint   `gorm:"primarykey"`
	UID    string `gorm:"primarykey"`
	GID    uint
	Status bool
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
