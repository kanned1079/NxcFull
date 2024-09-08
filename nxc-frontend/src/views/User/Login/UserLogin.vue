<script setup lang="ts" name="UserLogin">
import {computed, onMounted, ref} from 'vue'
import {useRouter} from 'vue-router'
import useSiteInfo from "@/stores/siteInfo";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";
import type {NotificationType} from 'naive-ui'
import {useNotification} from 'naive-ui'
import axios from 'axios';

const notification = useNotification()
const themeStore = useThemeStore()
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
let enableLogin = ref<boolean>(true)

// encodeToBase64 将密码进行base64加密
let encodeToBase64 = (str: string): string => {
  const utf8Encode = new TextEncoder();
  const encoded = utf8Encode.encode(str);
  const base64Encoded = btoa(String.fromCharCode(...encoded));
  return base64Encoded;
}

interface DataWithAuth {
  isAuthed: boolean;
  user_data: UserData;
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
  last_login_ip: string;
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
}



let handleLogin = async () => {
  enableLogin.value = false
  try {
    let { data } = await axios.post('http://localhost:8080/api/admin/login', {
      email: username.value,
      password: encodeToBase64(password.value),
    })
    console.log(data)
    if (data.code === 200 && data.isAuthed === true) {
      // 验证通过 保存token
      sessionStorage.setItem('token', data.token)
      // sessionStorage.setItem('isAuthed', JSON.stringify(true))
      notifyPass('success');
      await bindUserInfo(data)
      await router.push({ path: '/admin/dashboard' });
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

}

let backgroundStyle = computed(() => ({
  backgroundSize: 'cover', // 或者 'contain' 根据你需要的效果选择
  backgroundImage: `url(${themeStore.backgroundUrl})`
}))

onMounted(() => {
  // console.log('UserLogin挂载')
  // sessionStorage.setItem('isAuthed', JSON.stringify(false))
  //
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

<template>

  <n-layout style="width: 100%; height: 100vh;" justify="center" :vertical="true" align="center" :style="backgroundStyle">
    <n-flex justify="center" :vertical="true" align="center">
      <n-card class="layer-up" :embedded="true">
        <p class="title">{{ siteInfo.siteName }}</p>
        <p class="sub-title">常州站点</p>
        <div class="inp">
          <n-input secondary v-model:value="username" type="text" placeholder="邮箱" size="large"/>
          <n-input v-model:value="password" type="password" placeholder="密码" size="large" style="margin-top: 20px"/>
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
  //background-image:
  //background-repeat:no-repeat;background-size:cover;background-attachment:fixed;background-position-x:center;
}

.layer-up {
  width: 480px;
  height: 400px;
  border: 0;

  .title {
    font-size: 30px;
  }

  .sub-title {
    font-size: 13px;
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

.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}

</style>
