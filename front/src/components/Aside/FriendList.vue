<script setup lang="ts">
import FriendBox from "@/components/Aside/Friend.vue";
import {useRouter} from "vue-router";
import {UserData} from "@/api/user";
import {useUserStore} from "@/store/User";

const userStore = useUserStore()
const router = useRouter()
import {onMounted, ref, watch} from "vue";

const activeIndex = ref("idx")
const handleClick = (index: number) => {
  router.push(`/chat/${index}`)
  activeIndex.value = index.toString()
}
const friendDataList = ref<UserData[]>([])
onMounted(() => {
  friendDataList.value = userStore.friendList
})
watch(() => userStore.friendList, (newvalue) => {
  friendDataList.value = newvalue
})


</script>

<template>
  <el-menu :default-active="activeIndex">
    <el-menu-item v-for="friend in friendDataList" :key="friend.id" :index="friend.id.toString()"
                  @click="handleClick(friend.id)">
      <FriendBox :friendData="friend"></FriendBox>
    </el-menu-item>
  </el-menu>

</template>
<style scoped lang="scss">

.el-menu-item {
  height: 10%;
}


.el-menu-item:hover {
  background-color: #d3d3d3;
}


.el-menu-item.is-active {
  background-color: #DDE4FF;
}


</style>