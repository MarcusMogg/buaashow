package service

import (
	"errors"
	"buaashow/global"
	"buaashow/model/entity"
	"buaashow/utils"

	"gorm.io/gorm"
)

//Register 用户注册
func Register(u *entity.MUser) error {
	u.Password = utils.AesEncrypt(u.Password)
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("user_name = ?", u.UserName).First(u)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return tx.Create(u).Error
		}
		return errors.New("用户名已注册")
	})
}

//Login 用户登录
func Login(u *entity.MUser) bool {
	result := global.GDB.Where("user_name = ? AND password = ?", u.UserName, utils.AesEncrypt(u.Password)).First(u)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// GetUserInfoByID 获取用户信息
func GetUserInfoByID(id uint) (*entity.MUser, error) {
	var u entity.MUser
	err := global.GDB.First(&u, id).Error
	return &u, err
}

// UpdateUser 修改用户信息
func UpdateUser(user *entity.MUser) error {
	return global.GDB.Save(user).Error
}
