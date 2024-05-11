<script setup lang="ts">
import {MessagesData} from "@/api/message/type/messages";
import {useUserStore} from "@/store/User";
import {handleTime} from "@/utils/handleMessage";
import {UserData} from "@/api/user";
import {ref} from "vue";

const userStore = useUserStore()
const scrollbar = ref(null)


const Pronps = defineProps({
  messages: {
    type: Array as () => MessagesData[],
    required: true,
    default: () => []
  },
  friendData: {
    type: Object as () => UserData,
    required: true,
    default: () => {
    }
  }
})


</script>

<template>
  <el-scrollbar id="messageBox" ref="scrollbar" noresize="noresize">
    <div v-for="message in messages">
      <div v-if="userStore.isUser(message.sender)" class="user">
        <div class="messageBubbles">
          {{ message.content }}
          <div class="time">{{ handleTime(message.create_time) }}</div>
        </div>
        <el-avatar :src="userStore.userInfoData.avatar"></el-avatar>
      </div>
      <div v-else>
        <el-avatar :src="friendData.avatar" class="other"></el-avatar>
        <div class="messageBubbles">
          {{ message.content }}
          <div class="time">{{ handleTime(message.create_time) }}</div>
        </div>
      </div>

    </div>

  </el-scrollbar>
</template>

<style scoped lang="scss">
.el-scrollbar {
  display: block;
  padding-top: 20px;
  width: 100%;

  .user {
    text-align: right;
    margin-right: 20px;


  }

  .messageBubbles {
    font-size: 20px;
    display: inline-block;
    padding: 10px 20px;
    margin: 10px;
    background-color: #f0f0f0;
    border-radius: 20px;
    max-width: 60%;
    overflow-wrap: break-word;
    text-align: left;


    .time {
      text-align: right;
      color: #a9a9a9;
    }


  }

}


</style>