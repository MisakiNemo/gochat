<script setup lang="ts">
import {ref, Ref,defineProps} from "vue";
//导入好友数据类型
import {FriendData} from "@/api/user";
const props =defineProps({
  friendData: {
    type: Object as () => FriendData,
    required: true
  }

})


</script>

<template>
  <div id="box">
    <el-avatar :src="friendData.avatar"></el-avatar>
    <div id="detail">
      <div id="name">{{ friendData.name }}</div>
      <div id="status">
        <div v-if="friendData.status">在线</div>
        <div v-else>离线</div>
      </div>
    </div>
  </div>


</template>

<style scoped lang="scss">
#box {
  height: 90px;
  padding-left: 10px;
  display: flex;
  background-color: transparent;
  align-items: center;

  .el-avatar {
    width: 70px;
    height: 70px;
    margin-right: 20px;


  }

  #detail {
    display: flex;
    text-align: justify;
    font-size: 20px;


    #status{
      background-color: red;
      position: absolute;
      right: 20px;
    }


  }


}


</style>