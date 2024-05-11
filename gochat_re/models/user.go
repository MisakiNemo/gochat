package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID         uint           `gorm:"primarykey"`
	CreateTime time.Time      `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time      `gorm:"column:update_time"`
	DeleteAt   gorm.DeletedAt `gorm:"column:delete_at;index"`
}

type User struct {
	Model
	Name     string `json:"name" form:"name" gorm:"column:name"`
	Password string `json:"password" form:"password" gorm:"column:password" `
	Avatar   string `json:"avatar" form:"avatar" gorm:"column:avatar"`
}

type UserResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Status bool   `json:"status"`
}

func GetUserResponse(user User) UserResponse {
	return UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Avatar: user.Avatar,
		Status: IsOnline(user.ID),
	}
}

func (table *User) UserTableName() string {
	return "User"
}
