package dao

import (
	"go.uber.org/zap"
	"gochat_re/global"
	"gochat_re/models"
)

func GetUserList() ([]*models.User, error) {
	var list []*models.User
	if tx := global.DB.Find(&list); tx.RowsAffected == 0 {
		zap.S().Error("无法获取用户列表")
		return nil, tx.Error
	}
	return list, nil
}

func GetUserByUserid(userId uint) (*models.User, error) {
	var user = models.User{}
	if tx := global.DB.Where("id = ?", userId).First(&user); tx.RowsAffected == 0 {

		return nil, tx.Error
	}
	return &user, nil
}

func GetUserByName(name string) (*models.User, error) {
	user := models.User{}
	if tx := global.DB.Where("name = ?", name).First(&user); tx.RowsAffected == 0 {
		zap.S().Error("无法获取用户列表")
		return nil, tx.Error
	}
	return &user, nil
}

func GetUserByNameAndPassword(user models.User) (*models.User, error) {
	if tx := global.DB.Where("name = ? AND password = ?", user.Name, user.Password).First(&user); tx.RowsAffected == 0 {
		zap.S().Error("用户名或密码错误")
		return nil, tx.Error
	}
	return &user, nil
}

func CreateUser(user models.User) (*models.User, error) {
	if tx := global.DB.Create(&user); tx.RowsAffected == 0 {
		zap.S().Error("创建用户失败")
		return nil, tx.Error
	}
	return &user, nil
}

func UpdateUser(user models.User) (*models.User, error) {
	tx := global.DB.Model(&user).Updates(models.User{
		Name:     user.Name,
		Password: user.Password,
		Avatar:   user.Avatar,
	})
	if tx.RowsAffected == 0 {
		zap.S().Error("更新用户信息失败")
		return nil, tx.Error
	}
	return &user, nil
}

func DeleteUser(user models.User) error {
	if tx := global.DB.Delete(&user); tx.RowsAffected == 0 {
		zap.S().Error("删除用户失败")
		return tx.Error
	}
	return nil
}
