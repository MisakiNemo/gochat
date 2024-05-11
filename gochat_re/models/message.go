package models

import (
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"sync"
	"time"
)

type Message struct {
	Model
	Content  string `json:"content" gorm:"column:content"`
	Sender   uint   `json:"sender" gorm:"column:sender"`
	Receiver uint   `json:"receiver" gor:"column:receiver"`
}

type MessageResponse struct {
	Content    string    `json:"content"`
	Sender     string    `json:"sender"`
	Receiver   string    `json:"receiver"`
	CreateTime time.Time `json:"create_time"`
}

func (table *Message) MessageTableName() string {
	return "Message"
}

type Node struct {
	Conn       *websocket.Conn
	Addr       string
	DataQueue  chan []byte
	Friends    set.Interface
	PrivateKey []byte
}

var ClientMap map[uint]*Node = make(map[uint]*Node)

var RwLocker sync.RWMutex

func IsOnline(userId uint) bool {
	RwLocker.RLock()
	_, ok := ClientMap[userId]
	RwLocker.RUnlock()
	return ok
}
