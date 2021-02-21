package entity

import "time"

// MUser 数据库用户表
type MUser struct {
	Account   string `gorm:"not null;unique;primarykey"`
	CreatedAt time.Time
	Name      string
	Email     string
	Password  string `gorm:"not null" json:"-"`
	Role      Role
}

// Role 用户身份
type Role int

const (
	// Student 学生
	Student Role = iota + 1
	// Teacher 老师
	Teacher
	// Admin 管理员
	Admin
)

// UserInfoRes common user info
type UserInfoRes struct {
	Account string `json:"id"`
	Role    int    `json:"role"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}
