package dao

import (
	"fmt"
	"go.uber.org/zap"
	"gochat_re/global"
	"gochat_re/models"
)

func GetFriendList(userId uint) (list []models.Friendship, _ error) {
	if tx := global.DB.Where("user_id = ?", userId).Find(&list); tx.RowsAffected == 0 {
		zap.S().Info("查询好友失败或未查询到好友")
		return nil, tx.Error
	}
	fmt.Printf("list:%v", list)
	return list, nil
}

func CreateFriendShip(userId uint, friendId uint) (models.Friendship, error) {

	friendShip := models.Friendship{
		UserID:   userId,
		FriendID: friendId,
	}
	if tx := global.DB.Create(&friendShip); tx.RowsAffected == 0 {
		zap.S().Error("创建好友关系失败")
		return models.Friendship{}, tx.Error
	}
	return friendShip, nil
}

func GetFriendShipByUserIdAndFriendId(userId uint, friendId uint) (models.Friendship, error) {
	var friendShip models.Friendship
	if tx := global.DB.Where("userId = ? and friendId = ?", userId, friendId).First(&friendShip); tx.RowsAffected == 0 {
		zap.S().Error("不存在该好友关系")
		return models.Friendship{}, nil
	}
	return friendShip, nil
}

func DeleteFriendShipByUserIdAndFriendId(userId uint, friendId uint) error {
	if tx := global.DB.Where("userId = ? and friendId  = ?", userId, friendId); tx.RowsAffected == 0 {
		zap.S().Error("删除好友关系失败")
		return tx.Error
	}
	return nil
}
