<script setup lang="ts">
import {computed, onMounted, ref} from 'vue'
import {useRouter} from 'vue-router'
import useSiteInfo from "@/stores/siteInfo";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";
import type {NotificationType} from 'naive-ui'
import {useNotification} from 'naive-ui'
import { useMessage } from 'naive-ui'
import {hashPassword, encodeToBase64} from '@/utils/encryptor'
import axios from 'axios';
import useApiAddrStore from "@/stores/useApiAddrStore";
import instance from "@/axios";

const apiAddrStore = useApiAddrStore();

const notification = useNotification()
const themeStore = useThemeStore()
const message = useMessage()
let notifyErr = (type: NotificationType, msg: string) => {
  notification[type]({
    content: '登陆失败',
    meta: msg,
    duration: 2500,
    keepAliveOnHover: true
  })
}

let notifyPass = (type: NotificationType) => {
  notification[type]({
    content: '验证通过',
    // meta: '1111',
    duration: 2500,
    keepAliveOnHover: true
  })
}

const siteInfo = useSiteInfo();
const userInfoStore = useUserInfoStore();
const router = useRouter();
let username = ref<string>('')
let password = ref<string>('')
let usernameInputStatus = ref()
let passwordInputStatus = ref()
let enableLogin = ref<boolean>(true)

// encodeToBase64 将密码进行base64加密
// 移动到utils中
// let encodeToBase64 = (str: string): string => {
//   const utf8Encode = new TextEncoder();
//   const encoded = utf8Encode.encode(str);
//   const base64Encoded = btoa(String.fromCharCode(...encoded));
//   return base64Encoded;
// }

interface DataWithAuth {
  isAuthed: boolean;
  user_data: UserData;
  token: string;
}

interface UserData {
  id: number;
  invite_user_id: number;
  name: string;
  email: string;
  is_admin: boolean;
  is_staff: boolean;
  balance: number;
  last_login: Date;
  last_login_ip: string
}

let bindUserInfo =  (data: DataWithAuth) => {
  userInfoStore.isAuthed = data.isAuthed
  let {user_data} = data

  userInfoStore.thisUser.id = user_data.id
  userInfoStore.thisUser.inviteUserId = user_data.invite_user_id
  userInfoStore.thisUser.name = user_data.name
  userInfoStore.thisUser.email = user_data.email
  userInfoStore.thisUser.isAdmin = user_data.is_admin
  userInfoStore.thisUser.isStaff = user_data.is_staff
  userInfoStore.thisUser.balance = user_data.balance
  userInfoStore.thisUser.lastLogin = user_data.last_login.toString()
  userInfoStore.thisUser.lastLoginIp = user_data.last_login_ip
  userInfoStore.thisUser.token = data.token
  // userInfoStore.thisUser = {
  //   id: user_data.id,
  //   inviteUserId: user_data.invite_user_id,
  //   name: user_data.name,
  //   email: user_data.email,
  //   isAdmin: user_data.is_admin,
  //   isStaff: user_data.is_staff,
  //   balance: user_data.balance,
  //   lastLogin: user_data.last_login.toString(),
  //   lastLoginIp: user_data.last_login_ip,
  //   token: data.token,
  // };
  console.log('isAdmin', user_data.is_admin)
  console.log('admin/login: ', userInfoStore.thisUser.isAdmin)
}

// let ifIsNull = () => {
//   if (username.value === '' || password.value === '') {
//     notifyErr('error', '不能为空')
//     return false
//   } else {
//     return true
//   }
// }

let setInputStatus = (status: boolean) => {
  return !status?'error':null
}

let handleLogin = async () => {
  if (username.value === '' || password.value === '') {
    notifyErr('error', '不能为空')
    return
  }
  enableLogin.value = false
  console.log('管理员密码hash: ', hashPassword(password.value))
  try {
    // let hashedPwd =  hashPassword(password.value.trim())
    let { data } = await instance.post('http://localhost:8081/api/admin/v1/login', {
      email: username.value,

      password: encodeToBase64(password.value.trim()),
      // password: encodeToBase64(password.value),
      // password: hashPassword(password.value),
      // password: hashedPwd,
      role: 'admin', // 限制权限
    })
    console.log(data)
    if (data.code === 200 && data.isAuthed === true) {
      // 验证通过 保存token
      sessionStorage.setItem('token', data.token)
      // 保存验证状态
      userInfoStore.setAndSaveAuthStatus(true)
      // userInfoStore.isAuthed = true // 鉴权通过
      // sessionStorage.setItem('isAuthed', JSON.stringify(true))
      notifyPass('success');
      bindUserInfo(data)
      console.log(userInfoStore.thisUser)
      await router.push({ path: '/admin/dashboard/summary' });
    } else {
      enableLogin.value = true
      switch (data.msg) {
        case 'incorrect_password': {
          notifyErr('error', '密码不正确')
          break
        }
        case 'user_not_exist': {
          notifyErr('error', '用户不存在 请注册')
          break
        }
      }
    }
  }
  catch (error) {
    console.log(error)
  }
}

let handleFrogetPassword = () => {
  noticeInfo()
}

let noticeInfo = () => {
  message.info('自己想办法')
}

let backgroundStyle = computed(() => ({
  backgroundSize: 'cover', // 或者 'contain' 根据你需要的效果选择
  // backgroundImage: `url(${themeStore.backgroundUrl})`,
  background: `linear-gradient(to right, rgb(167, 112, 239), rgb(207, 139, 243), rgb(253, 185, 155))`
}))



onMounted(() => {
  console.log('AdminLogin挂载')
  // userInfoStore.isAuthed = false
  // sessionStorage.setItem('isAuthed', JSON.stringify(false))
  console.log('6666', JSON.parse(sessionStorage.getItem('isAuthed') as string))
  console.log('store.isAuthed: ', userInfoStore.isAuthed)

  if (JSON.parse(sessionStorage.getItem('isAuthed') as string)) {
    console.log('to dashboard')
    router.push({path:'/admin/dashboard/summary'})
  }

  // if (sessionStorage.getItem('isAuthed') != null) {
  //   if (JSON.parse(sessionStorage.getItem('isAuthed') as string) == true) {
  //     setTimeout(() => {
  //       notifyPass('success')
  //       router.push({
  //         path: '/admin/dashboard'
  //       })
  //     }, 500)
  //
  //   }
  // }
})
</script>

<script lang="ts">
export default {
  name: 'AdminLogin',
}
</script>

<template>

  <n-layout style="width: 100%; height: 100vh;" justify="center" :vertical="true" align="center" :style="backgroundStyle">
    <n-flex justify="center" :vertical="true" align="center">
      <n-card class="layer-up" :embedded="true" hoverable :bordered="false">
        <p class="title">{{ siteInfo.siteName }}</p>
        <p class="sub-title">登陆到管理中心</p>
        <div class="inp">
          <n-input
              secondary
              v-model:value="username"
              type="text" placeholder="邮箱"
              size="large"
              style="opacity: 0.8"
              :bordered="true"
              :status="usernameInputStatus"
              @blur="!username?usernameInputStatus = 'error':null"
          />
          <n-input
              v-model:value="password"
              type="password"
              placeholder="密码"
              size="large"
              style="margin-top: 20px; opacity: 0.8"
              :bordered="true"
              :status="passwordInputStatus"
              @blur="!password?passwordInputStatus = 'error':null"
          />
        </div>
        <n-button secondary type="info" class="login-btn" size="large" @click="handleLogin" :disabled="!enableLogin">
          登入
        </n-button>
        <n-button strong tertiary type="warning" size="large" class="login-btn" @click="handleFrogetPassword">
          忘记密码
        </n-button>
      </n-card>
      <!--      <n-card class="layer-down">-->

      <!--      </n-card>-->
    </n-flex>
  </n-layout>


</template>

<style lang="less" scoped>

.n-flex {
  height: 100vh;
  //background-color: rgba(255, 255, 255, 0.001);
  //background-color: #66afe9;
  //background-image:
  //background-repeat:no-repeat;background-size:cover;background-attachment:fixed;background-position-x:center;
}

.layer-up {
  width: 480px;
  border: 0;
  padding-bottom: 20px;
  backdrop-filter: blur(4px);
  .title {
    font-size: 30px;
    opacity: 1;
  }
  .sub-title {
    font-size: 13px;
    opacity: 0.7;
  }

  .inp {
    margin-top: 30px;
    text-align: left;
    width: 90%;
  }

  .login-btn {
    margin-top: 20px;
    width: 90%;
  }

}

.layer-down {
  background-color: #2c3e50;
  width: 480px;
  height: 40px;
}

//.n-card {
//  background-color: v-bind('themeStore.getTheme.globeTheme.loginCardBgColor');
//  //background-color: rgba(0, 0, 0, 0.7);
//  backdrop-filter: blur(50px);
//  border: 0;
//}

</style>
