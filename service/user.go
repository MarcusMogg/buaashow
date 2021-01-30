package service

// table m_users;

import (
	"buaashow/entity"
	"buaashow/global"
	"buaashow/utils"
	"errors"

	"gorm.io/gorm"
)

//Register 用户注册
func Register(u *entity.MUser) error {
	u.Password = utils.AesEncrypt(u.Password)
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("account = ?", u.Account).First(u)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return tx.Create(u).Error
		}
		return errors.New("用户名已注册")
	})
}

//Login 用户登录
func Login(u *entity.MUser) bool {
	result := global.GDB.Where("account = ? AND password = ?", u.Account, utils.AesEncrypt(u.Password)).First(u)
	return result.Error == nil
}

// GetUserInfoByID 获取用户信息
func GetUserInfoByID(id uint) (*entity.MUser, error) {
	var u entity.MUser
	err := global.GDB.First(&u, id).Error
	return &u, err
}

// GetUserInfoByAccount 获取用户信息
func GetUserInfoByAccount(account string) (*entity.MUser, error) {
	var u entity.MUser
	err := global.GDB.Where("account = ?", account).First(&u).Error
	return &u, err
}

// UpdateUser 修改用户信息
func UpdateUser(user *entity.MUser) error {
	return global.GDB.Save(user).Error
}

// UpdateEmail 修改邮箱
func UpdateEmail(user *entity.MUser, email string) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&user).Update("email", email).Error
	})
}

// UpdatePassword 修改密码
func UpdatePassword(user *entity.MUser, oldPassword, newPassword string) error {
	oldPassword = utils.AesEncrypt(oldPassword)
	newPassword = utils.AesEncrypt(newPassword)
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		result := global.GDB.Where("account = ? AND password = ?", user.Account, oldPassword).First(user)
		if result.Error != nil {
			return errors.New("密码错误")
		}
		return tx.Model(user).Update("password", newPassword).Error
	})
}
