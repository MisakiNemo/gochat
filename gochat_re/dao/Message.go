package dao

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gochat_re/common"
	"gochat_re/global"
	"gochat_re/models"
	"gopkg.in/fatih/set.v0"
	"net/http"
)

func Chat(c *gin.Context) {
	token := c.Request.URL.Query().Get("token")
	claims, err := common.ParseToken(token, "123456")
	if err != nil {
		zap.S().Error("解析token失败")
		common.RespFail(c.Writer, "解析token失败")
		return
	}
	userId, ok := (*claims)["user_id"].(float64)
	if !ok {
		zap.S().Error("解析用户id失败")
		common.RespFail(c.Writer, "解析用户id失败")
		return
	}
	var isvalidate = true
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalidate
		},
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.S().Error("升级websocket连接失败")
	}
	aesKey := (*claims)["aeskey"].(string)
	node := &models.Node{
		Conn:       conn,
		DataQueue:  make(chan []byte, 50),
		Friends:    set.New(set.ThreadSafe),
		PrivateKey: common.GetAESKey(global.PrivateKey, []byte(aesKey)),
	}
	models.RwLocker.Lock()
	models.ClientMap[uint(userId)] = node
	models.RwLocker.Unlock()
	go sendProc(node)

	go recProc(node)
	SendOfflineMessage(uint(userId))
}

func sendProc(node *models.Node) {
	for {
		select {
		case data := <-node.DataQueue:
			encrypteddata, _ := common.EncryptAES(node.PrivateKey, data)
			err := node.Conn.WriteMessage(websocket.TextMessage, encrypteddata)
			if err != nil {
				zap.S().Error("写入消息失败")
				return
			}
			fmt.Println("发送消息成功")

		}
	}

}

func recProc(node *models.Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			zap.S().Error("接收消息失败")
			return
		}
		fmt.Println(string(data))
		decryptedData, _ := common.DecryptAES(node.PrivateKey, data)
		var msg models.Message
		err = json.Unmarshal(decryptedData, &msg)
		if err != nil {
			zap.S().Error("解析消息失败", err)
			return
		}
		msg.UpdateTime = msg.CreateTime
		targetNode, ok := models.ClientMap[msg.Receiver]
		if !ok {
			err := SaveMessage(&msg)
			if err != nil {
				zap.S().Error("保存消息失败")
				return
			}
			fmt.Println("对方不在线，消息已存储")
			continue
		}
		targetNode.DataQueue <- decryptedData
		fmt.Println("消息发送成功")

	}
}

func SaveMessage(msg *models.Message) error {
	fmt.Println(msg)
	err := global.DB.Create(msg).Error
	if err != nil {
		zap.S().Error("保存消息失败", err)
		return err
	}
	return nil
}

func getOfflineMessage(userId uint) ([]models.Message, error) {
	var msgs []models.Message
	err := global.DB.Where("receiver = ?", userId).Find(&msgs).Error
	if err != nil {
		return nil, err
	}
	if len(msgs) == 0 {
		return nil, nil
	}
	return msgs, nil
}

func SendOfflineMessage(userId uint) {
	msgs, err := getOfflineMessage(userId)
	if err != nil {
		zap.S().Error("获取离线消息失败")
		return
	}
	models.RwLocker.RLock()
	node, ok := models.ClientMap[userId]
	models.RwLocker.RUnlock()
	if !ok {
		zap.S().Error("用户不在线")
		return
	}
	for _, msg := range msgs {
		data, err := json.Marshal(msg)
		if err != nil {
			zap.S().Error("序列化消息失败")
			return
		}
		node.DataQueue <- data
	}
	zap.S().Info("离线消息发送成功")
}

func BindAESKey(userId uint, AESkey []byte) {
	models.RwLocker.RLock()
	_, ok := models.ClientMap[userId]
	models.RwLocker.RUnlock()
	if !ok {
		zap.S().Error("用户不在线")
		return
	}

}
