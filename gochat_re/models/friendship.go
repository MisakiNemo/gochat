package models

type Friendship struct {
	Model
	UserID   uint `json:"user_id" form:"user_id" gorm:"column:user_id"`
	FriendID uint `json:"friend_id" form:"friend_id" gorm:"column:friend_id"`
}

type FriendshipResponse struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Avtar  string `json:"avtar"`
	Status bool   `json:"status"`
}

func (table *Friendship) FriendshipTableName() string {
	return "Friendship"
}
