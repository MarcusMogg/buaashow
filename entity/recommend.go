package entity

import "time"

// MRecSubmission 学生作业提交
type MRecSubmission struct {
	EID       uint   `gorm:"primarykey"`
	GID       string `gorm:"primarykey"`
	Name      string
	Info      string
	Type      SummitType
	URL       string
	Thumbnail string
	Readme    string
	Rec       bool
	UpAt      time.Time
}
