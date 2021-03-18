package service

// table m_users;

import (
	"buaashow/entity"
	"buaashow/global"
	"buaashow/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

//Register 用户注册
func Register(u *entity.MUser) error {
	u.Password = utils.AesEncrypt(u.Password)
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(u).Error
	})
}

//Login 用户登录
func Login(u *entity.MUser) bool {
	result := global.GDB.Where("account = ? AND password = ?", u.Account, utils.AesEncrypt(u.Password)).First(u)
	return result.Error == nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(account string) (*entity.MUser, error) {
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

//GetUserInfoList 查询用户列表
func GetUserInfoList(page, pageSize int, account string) (int, []*entity.UserInfoRes, error) {
	var res []*entity.UserInfoRes
	offset := (page - 1) * pageSize
	var tot int64
	err := global.GDB.Model(&entity.MUser{}).
		Where("account LIKE ? AND role != ?", fmt.Sprintf("%%%s%%", account), entity.Admin).
		Select("account,role,name,email").
		Count(&tot).Offset(offset).Limit(pageSize).Find(&res).Error
	return int(tot), res, err
}

// DeleteUser 删除用户
func DeleteUser(account string) error {
	return global.GDB.Delete(&entity.MUser{
		Account: account,
	}).Error
}
