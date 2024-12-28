<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, onMounted, ref} from 'vue'
import {useRouter} from 'vue-router'
import useAppInfosStore from "@/stores/useAppInfosStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";
import type {NotificationType} from 'naive-ui'
import {useNotification} from 'naive-ui'
import { useMessage } from 'naive-ui'
import {hashPassword, encodeToBase64} from '@/utils/encryptor'
import instance from "@/axios";

const {t} = useI18n()
const notification = useNotification()
const appInfoStore = useAppInfosStore()
const themeStore = useThemeStore()
const message = useMessage()

let animated = ref<boolean>(false)

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

const userInfoStore = useUserInfoStore();
const router = useRouter();
let username = ref<string>('')
let password = ref<string>('')
let usernameInputStatus = ref()
let passwordInputStatus = ref()
let enableLogin = ref<boolean>(true)


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
  console.log('isAdmin', user_data.is_admin)
  console.log('admin/login: ', userInfoStore.thisUser.isAdmin)
}

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
      role: 'admin', // 限制权限
    })
    console.log(data)
    if (data.code === 200 && data.isAuthed === true) {
      // 验证通过 保存token
      sessionStorage.setItem('token', data.token)
      // 保存验证状态
      userInfoStore.setAndSaveAuthStatus(true)
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

let handleForgetPassword = () => {
  noticeInfo()
}

let noticeInfo = () => {
  message.info('...')
  // TODO 待实现
}

let backgroundStyle = computed(() => {
  const baseStyle = {
    backgroundSize: 'cover', // 或者 'contain'，根据需要调整
    backgroundImage: `url(${appInfoStore.appCommonConfig.frontend_background_url})`,
  };

  // 如果启用了深色模式，调整亮度和对比度
  if (themeStore.enableDarkMode) {
    return {
      ...baseStyle,
      filter: 'brightness(0.5) contrast(1)', // 亮度为0.5，对比度为0.8
    };
  }

  return baseStyle;
});


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
  animated.value = true

})

</script>

<script lang="ts">
export default {
  name: 'AdminLogin',
}
</script>

<template>
  <div class="background-container">
    <!-- 背景层 -->
    <div class="background-layer" :style="backgroundStyle"></div>

    <!-- 内容层 -->
    <div
        class="content-layer"
        style="width: 100%; height: 100vh;"
    >
      <n-flex justify="center" :vertical="true" align="center">
        <transition name="slide-fade">
          <n-card
              v-if="animated"
              class="layer-up"
              :embedded="true"
              hoverable
              :bordered="false"
              style="text-align: center"
              content-style="padding: 40px"
          >
            <p class="title">{{ appInfoStore.appCommonConfig.app_name }}</p>
            <p class="sub-title">登陆到管理中心</p>
            <div class="inp">
              <n-input
                  secondary
                  v-model:value="username"
                  type="text"
                  placeholder="邮箱"
                  size="large"
                  style="opacity: 0.8"
                  :bordered="true"
                  :status="usernameInputStatus"
                  @blur="!username ? (usernameInputStatus = 'error') : null"
              />
              <n-input
                  v-model:value="password"
                  type="password"
                  placeholder="密码"
                  size="large"
                  style="margin-top: 20px; opacity: 0.8"
                  :bordered="true"
                  :status="passwordInputStatus"
                  @blur="!password ? (passwordInputStatus = 'error') : null"
              />
            </div>
            <n-button
                secondary
                type="info"
                class="login-btn"
                size="large"
                @click="handleLogin"
                :disabled="!enableLogin"
            >
              登入
            </n-button>
            <n-button
                strong
                tertiary
                type="warning"
                size="large"
                class="login-btn"
                @click="handleForgetPassword"
            >
              忘记密码
            </n-button>
          </n-card>
        </transition>
      </n-flex>
    </div>
  </div>
</template>

<style lang="less" scoped>
.background-container {
  position: relative;
  width: 100%;
  height: 100vh;
  overflow: hidden;
}

.background-layer {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
  z-index: 0;
}

.content-layer {
  position: relative;
  z-index: 1; /* 确保内容层在背景层之上 */

}

.n-flex {
  height: 100vh;
}

.layer-up {
  filter: brightness(1) contrast(1) !important;
  z-index: 1000;
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
    width: 100%;
  }

  .login-btn {
    margin-top: 20px;
    width: 100%;
  }
}

.layer-down {
  background-color: #2c3e50;
  width: 480px;
  height: 40px;
}
</style>