<script setup lang="ts">
import {useRoute} from 'vue-router';
import {watch, ref, Ref, onMounted, watchEffect, computed} from "vue";
import {UserData} from "@/api/user";
import MessageBox from "@/components/Chat/MessageBox.vue";
import {MessagesData} from "@/api/message/type/messages";
import {useUserStore} from "@/store/User";
import {useChatStore} from "@/store/Chat";

const ChatStore = useChatStore()
const UserStore = useUserStore()
const content = ref<string>("")
const route = useRoute();
const friendId: Ref<number | null> = ref(null)
const messageBox = ref(null)
const messages: Ref<MessagesData[]> = ref([])

onMounted(() => {
  friendId.value = Number(route.params.friendId)
  changeFriendData(Number(friendId.value))
  changeChatMessages(Number(friendId.value))
})

const validate = () => {
  return content.value.trim() === ""
}

const sendMessage = () => {
  if (validate()) {
    ElMessage.error("发送内容不能为空")
    content.value = ""
    return
  }
  let message: MessagesData = {
    content: content.value,
    receiver: friendId.value,
    sender: UserStore.userInfoData.id,
    create_time: new Date()
  }
  console.log(message)
  ChatStore.sendMessage(message)
  content.value = ""
}

const changeFriendData = (friendId: number) => {
  console.log(friendId)
  friendData.value = UserStore.getFriendData(friendId)
}

const changeChatMessages = (friendId: number) => {
  messages.value=ChatStore.getChatRecord(friendId)
}


const friendData: Ref<UserData> = ref({
  id: -1,
  name: "",
  avatar: "",
  status: false
})

watch(() => route.params.friendId, (newvalue) => {
  friendId.value = Number(newvalue)
  changeFriendData(Number(newvalue))
  changeChatMessages(Number(newvalue))
})



watchEffect(() => {
      messages.value = ChatStore.getChatRecord(friendId.value)
    }
)


</script>

<template>
  <div id="chatbox">
    <div id="header">
      <el-avatar :src="friendData.avatar"></el-avatar>
      <div id="name">{{ friendData.name }}</div>
    </div>
    <div id="messageBox">
      <MessageBox :messages="messages" :friendData="friendData" ref="messagesBox"></MessageBox>

    </div>
    <div id="inputBox">
      <textarea v-model="content" placeholder="请输入内容" @keyup.enter="sendMessage"></textarea>
      <el-button @click="sendMessage">发送</el-button>
    </div>

  </div>

</template>

<style scoped lang="scss">
#chatbox {
  width: 100%;
  height: 100%;
  background-color: white;
  top: 0;
  margin-left: -20px;


  #header {
    height: 10%;
    width: 100%;
    display: flex;
    border-bottom: 0.1px solid #e9e9e9;
    align-items: center;
    font-size: 20px;
  }

  #messageBox {
    width: 100%;
    height: 70%;

  }

  #inputBox {
    border-top: 2px solid #d3d3d3;
    height: 20%;


    textarea {
      resize: none;
      width: 100%;
      height: 70%;
      border: none;
      font-size: 20px;
      font-family: "Microsoft YaHei";

    }

    textarea:focus {
      outline: none;
      border: none;
    }


    .el-button {
      display: block;
      position: absolute;
      width: 100px;
      right: 50px;
      margin-top: 8px;
    }

  }
}

</style>
