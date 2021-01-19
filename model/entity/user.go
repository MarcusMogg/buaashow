package entity

import "gorm.io/gorm"

// MUser 数据库用户字段
type MUser struct {
	gorm.Model
	Account  string `gorm:"not null;unique"`
	Email    string
	Password string `gorm:"not null" json:"-"`
	Role     Role
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
