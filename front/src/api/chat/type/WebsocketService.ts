import {useChatStore} from "@/store/Chat";
import {encryptAES, encryptAESKeyWithPublicKey, generateAESKey} from "@/utils/crypto.js";
import {MessagesData} from "@/api/message";

const chatStore = useChatStore();

class WebsocketService {
    private ws: WebSocket | null = null;

    public connect(url: string) {
        this.ws = new WebSocket(url);

        this.ws.onopen = () => {
            console.log('connected');
        }
        this.ws.onmessage =   async (event) => {
            const message = JSON.parse(event.data);
            console.log("message", message.content);
            if (message.type === 1) {
                chatStore.key = generateAESKey();
                console.log("key", chatStore.key);
                encryptAESKeyWithPublicKey(message.content,chatStore.key).then((res) => {
                    console.log("res", res);
                const msg: MessagesData = {
                    sender: 0,
                    receiver: 0,
                    content: res,
                    type: 2,
                    create_time: new Date()
                }
                this.sendMessage(JSON.stringify(msg));
            })
            }
            chatStore.receiveMessage(message);

        }
        this.ws.onclose = () => {

        }
        this.ws.onerror = (error) => {
            console.log(error);
        }
    }

    sendMessage(message: string) {
        if (this.ws) {
            this.ws.send(message)
        }
    }

    public disconnect() {
        if (this.ws) {
            this.ws.close();
        }
    }


}


export default new WebsocketService();