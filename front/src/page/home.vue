<script setup lang="ts">
import {onMounted, onUnmounted, provide} from "vue";
import websocketService from "@/api/chat/type/WebsocketService";
import FriendList from "@/components/Aside/FriendList.vue";
import Searchbox from "@/components/Aside/searchbox.vue";
import {getToken} from "@/utils/request";
import {useUserStore} from "@/store/User";
import {useChatStore} from "@/store/Chat";

const userStore = useUserStore()
const chatStore = useChatStore()

onMounted(async () => {
    await userStore.updateFriendList()
    await userStore.updateUserInfo()
    let friendList = userStore.friendList
    chatStore.initChatList(friendList)
    conne()
    userStore.friendList.forEach((item) => {
      console.log(item)
    })
})

const conne = () => {
  websocketService.connect("ws://localhost:8080/v1/chat/ws?token=" + getToken())
}

</script>

<template>
  <div id="mainbox">
    <el-container id="ctn">
      <el-aside id="left">
        <Searchbox></Searchbox>
        <FriendList></FriendList>
      </el-aside>

      <el-main>
        <router-view></router-view>
      </el-main>

    </el-container>


  </div>


</template>

<style scoped lang="scss">

#mainbox {
  //设置边框
  border: 1px solid #e9e9e9;
  border-radius: 20px;
  height: 80%;
  width: 60%;
  position: absolute;
  top: 10%;
  left: 20%;


  #ctn {
    height: 100%;


    #left {
      width: 15vw;
      border-right: 0.1px solid #e9e9e9;
      background-color: white;
    }

    .el-main {
      padding-top: 0;
      padding-bottom: 0;
    }


  }


}


</style>