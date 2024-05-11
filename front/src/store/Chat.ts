import {defineStore} from "pinia";
import {Ref, ref} from "vue";
import {ChatRecord, MessagesData} from "@/api/message";
import {UserData} from "@/api/user";
import websocketService from "@/api/chat/type/WebsocketService";


export const useChatStore = defineStore("Chat",
    () => {
        const chatList = ref<ChatRecord[]>([]);
        const key:Ref<string> = ref<string>("");

        const initChatList = (list: UserData[]) => {
            list.forEach((item) => {
                chatList.value.push({
                    friendId: item.id,
                    MessagesData: []
                })
            })
        }

        const sendMessage = (message: MessagesData) => {
            chatList.value.forEach((item) => {
                if (item.friendId === message.receiver) {
                    item.MessagesData.push(message);
                }
            })
            console.log("消息发送成功", message)
            console.log("消息发送成功",chatList)
            websocketService.sendMessage(JSON.stringify(message)
            )
        }


        const receiveMessage = (message: MessagesData) => {
             let msg:MessagesData = {
                sender: message.sender,
                receiver: message.receiver,
                content: message.content,
                create_time: message.create_time
            }
            chatList.value.forEach((item) => {
                if (item.friendId === message.sender) {
                    item.MessagesData.push(msg);
                    console.log("消息接收",item.MessagesData)
                }
            })
        }


        const getChatRecord = (friendId: number) => {
            const messages= chatList.value.find((item) => item.friendId === friendId)
            if (messages) {
                return messages.MessagesData
            }
            return []
        }


        return {
            key,
            chatList,
            initChatList,
            sendMessage,
            receiveMessage,
            getChatRecord
        }

    }
)



