<script setup lang="ts">
import {ref} from "vue";
import {registerService, loginService} from "@/api/user";
import {useRouter} from "vue-router";
import type {RegisterResponseData} from "@/api/user";
import {setToken} from "@/utils/request";
//注册表单
const router = useRouter()
const FormRef = ref(null);
const isLogin = ref(true);
const registerBox = ref(null)
const loginBox = ref(null)
const Form = ref({
  username: "",
  password: "",
  //验证用户输入的数据
  validate: {
    username: [
      {required: true, message: "请输入用户名", trigger: "blur"},
      {min: 3, max: 10, message: "长度在 3 到 10 个字符", trigger: "blur"},
    ],
    password: [
      {required: true, message: "请输入密码", trigger: "blur"},
      {min: 6, max: 20, message: "长度在 6 到 20 个字符", trigger: "blur"},
    ],
  },

});
const changeLogin = () => {
  isLogin.value = !isLogin.value
  Form.value.username = ""
  Form.value.password = ""
}

const register = async () => {
  try {
    await FormRef.value.validate();
    const registerData = {
      name: Form.value.username,
      password: Form.value.password
    }
    await registerService(registerData).then((res) => {
      let data: RegisterResponseData = res.data;
      console.log(data)
      if (data.code == 0) {
        ElMessage.success(data.msg)
        console.log(data.data.access_token)
        console.log(data.data.access_expire)
        setToken(data.data)
        router.push("/index")
      } else {
        ElMessage.error({
          message: data.msg,
          type: "error",
        });
      }

    }).catch((err) => {
      console.log(err)
    })
  } catch (error) {
    console.log(error)
  }
}
const login = async () => {
  try {
    await FormRef.value.validate();
    const loginData = {
      name: Form.value.username,
      password: Form.value.password
    }
    await loginService(loginData).then((res) => {
      let data: RegisterResponseData = res.data;
      console.log(data)
      if (data.code == 0) {
        ElMessage.success(data.msg)
        console.log(data.data.access_token)
        console.log(data.data.access_expire)
        setToken(data.data)
        router.push("/index")
      } else {
        ElMessage.error({
          message: data.msg,
          type: "error",
        });
      }

    }).catch((err) => {
      console.log(err)
    })

  } catch (error) {
    console.log(error)
  }
}



</script>

<template>
  <div v-if="isLogin" ref="loginBox">
    <el-form :model="Form" :rules="Form.validate" ref="FormRef" label-width="80px">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="Form.username"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input type="password" v-model="Form.password"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="login">登录</el-button>
      </el-form-item>
    </el-form>
    <el-button @click="changeLogin">没有账户，去注册</el-button>
  </div>

  <div class="box" v-else ref="registerBox">
    <el-form :model="Form" :rules="Form.validate" ref="FormRef" label-width="80px">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="Form.username"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input type="password" v-model="Form.password"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="register">注册</el-button>
      </el-form-item>
    </el-form>
    <el-button @click="changeLogin">已有账户，去登录</el-button>
  </div>

</template>

<style scoped>

</style>