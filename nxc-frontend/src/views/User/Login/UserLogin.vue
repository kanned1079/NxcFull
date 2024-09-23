<script setup lang="ts">
import {computed, onMounted, ref} from 'vue'
import {useRouter} from 'vue-router'
import useSiteInfo from "@/stores/siteInfo";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";
import type {NotificationType} from 'naive-ui'
import {FormInst, useMessage, useNotification} from 'naive-ui'
import instance from '@/axios/index'
import {GlobeOutline as LanguageIcon} from '@vicons/ionicons5'
import useApiAddrStore from "@/stores/useApiAddrStore";
import {hashPassword} from "@/utils/encryptor";

const apiAddrStore = useApiAddrStore();
const notification = useNotification()
const themeStore = useThemeStore()
const message = useMessage()
const siteInfo = useSiteInfo();
const userInfoStore = useUserInfoStore();
const router = useRouter();

const formRef = ref<FormInst | null>(null)

let formValue = ref({
  user: {
    email: '',
    password: '',
  }
})
let rules = {
  user: {
    email: {
      required: true,
      message: '邮件地址是必填的',
      trigger: 'blur'
    },
    password: {
      required: true,
      message: '你还没有输入密码',
      trigger: 'blur'
    }
  }
}

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


let enableLogin = ref<boolean>(true)
let clickedCount = ref<number>(0)

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

let bindUserInfo = (data: DataWithAuth) => {
  console.log('绑定数据', data)
  userInfoStore.isAuthed = data.isAuthed
  let {user_data} = data

  userInfoStore.thisUser.id = user_data.id
  userInfoStore.thisUser.inviteUserId = user_data.invite_user_id
  userInfoStore.thisUser.name = user_data.name
  userInfoStore.thisUser.email = user_data.email
  userInfoStore.thisUser.isAdmin = user_data.is_admin
  userInfoStore.thisUser.isStaff = user_data.is_staff
  userInfoStore.thisUser.balance = user_data.balance
  userInfoStore.thisUser.lastLogin = user_data.last_login ? user_data.last_login.toString() : ''
  userInfoStore.thisUser.lastLoginIp = user_data.last_login_ip
  // userInfoStore.thisUser.licenseActive = user_data.license_active
  // userInfoStore.thisUser.licenseExpiration = user_data.license_expiration
  // userInfoStore.thisUser.licenseId = user_data.license_id
  console.log('绑定结束')
}

let enterDashboard = (data) => {

}


let handleLogin = async () => {
  console.log('用户数据')
  console.log(formValue.value.user.email, formValue.value.user.password)
  try {
    let {data} = await instance.post(apiAddrStore.apiAddr.user.userLogin, {
      email: formValue.value.user.email,
      // password: encodeToBase64(formValue.value.user.password),
      password: hashPassword(formValue.value.user.password),
      role: 'user',
    })
    console.log(data)
    if (data.code === 200 && data.isAuthed === true) {
      // 验证通过 保存token
      sessionStorage.setItem('token', data.token)
      // sessionStorage.setItem('isAuthed', JSON.stringify(true))
      notifyPass('success');
      console.log('绑定用户数据到pinia')
      bindUserInfo(data)
      console.log("结束")
      console.log('跳转到用户仪表板')
      await router.push({path: '/dashboard'});
    } else {
      enableLogin.value = true
      switch (data.msg) {
        case 'incorrect_password': {
          notifyErr('error', '密码不正确')
          enableLogin.value = true
          break
        }
        case 'user_not_exist': {
          notifyErr('error', '用户不存在 请注册')
          enableLogin.value = true
          break
        }
      }
    }
  } catch (error) {
    console.log(error)
    // notifyErr('error', error.toString())
    enableLogin.value = true
  }


}

let toRegister = () => {
  console.log('跳转到注册页面')
  router.push({
    path: '/register',
  })
}

// 处理忘记密码
let handleForgetPassword = () => {
  message.info('处理忘记密码')
}

let backgroundStyle = computed(() => ({
  backgroundSize: 'cover', // 或者 'contain' 根据你需要的效果选择
  // backgroundImage: `url(${themeStore.backgroundUrl})`,
  background: `linear-gradient(to right, rgb(190, 147, 197), rgb(123, 198, 204))`,


}))

onMounted(() => {
  console.log('挂载普通用户登录')
})

let handleValidateClick = () => {
  // enableLogin.value = false
  clickedCount.value += 1
  console.log('clickedCount: ', clickedCount.value)
  if (clickedCount.value == 5) {
    enableLogin.value = false
    setTimeout(() => {
      enableLogin.value = true
      clickedCount.value = 0
    }, 60000)
    return
  }
  if (formValue.value.user.email === '' || formValue.value.user.password === '') {
    notifyErr('error', '邮箱或密码不能为空')
    return
  }
  console.log('验证完成 调用登录')
  handleLogin()
  // if (formValue.value.user.password === '') {
  //   notifyErr('error', '密码不能为空')
  //   return false
  // }
}


</script>


<script lang="ts">
export default {
  name: "UserLogin",
}
</script>

<template>

  <n-layout style="width: 100%; height: 100vh;" justify="center" :vertical="true" align="center"
            :style="backgroundStyle">
    <n-flex justify="center" :vertical="true" align="center" style="gap: 0">
      <n-card class="layer-up" :embedded="true" hoverable>
        <p class="title">{{ siteInfo.siteName }}</p>
        <p class="sub-title">常州站点</p>
        <div class="inp">

          <n-form
              ref="formRef"
              :model="formValue"
              :rules="rules"
          >
            <n-form-item path="user.email" :show-feedback="false">
              <n-input v-model:value="formValue.user.email" placeholder="邮箱地址" size="large"
                       style="user-select: none"/>
            </n-form-item>

            <n-form-item path="user.password" :show-feedback="false">
              <n-input type="password" v-model:value="formValue.user.password" placeholder="密码" size="large"/>
            </n-form-item>
          </n-form>
        </div>
        <n-button secondary type="info" class="login-btn" size="large" @click="handleValidateClick"
                  :disabled="!enableLogin">
          登入
        </n-button>

      </n-card>
      <n-card class="layer-down" content-style="padding: 0;">
        <div class="bottom-root">
          <div class="l-con">
            <n-button text @click="toRegister">注册</n-button>
            <n-divider vertical/>
            <n-button text @click="handleForgetPassword">忘记密码</n-button>
          </div>

          <div class="r-con">
            <n-button text>
              <n-icon style="margin-right: 5px;">
                <LanguageIcon/>
              </n-icon>
              语言
            </n-button>
          </div>
        </div>

      </n-card>

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
  height: auto;
  border: 0;
  border-radius: 5px 5px 0 0;
  margin-bottom: 0;

  .title {
    margin-top: 20px;
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
    margin-top: 30px;
    width: 90%;
    margin-bottom: 30px;
  }

  .register-btn {
    margin-top: 30px;
    width: 90%;
  }

}

.bottom-panel {
  height: 50px;
  width: 480px;
  border-radius: 0 0 5px 5px;
  padding: 0;
  margin: 0;
}

.layer-down {
  width: 480px;
  height: 50px;
  border-radius: 0 0 5px 5px;

  .bottom-root {
    border-radius: 0 0 5px 5px;
    height: 100%;
    display: flex;
    backdrop-filter: blur(10px);
    background-color: rgba(225, 225, 225, 0.2);
    justify-content: space-between;

    .l-con {
      line-height: 50px;
      margin-left: 20px;
      opacity: 0.8;
    }

    .r-con {
      line-height: 50px;
      margin-right: 20px;

    }
  }
}

.n-card {
  //background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  //background-color: v-bind('themeStore.getTheme.globeTheme.loginCardBgColor');
  border: 0;
}

//element.style {
//  gap: 0;
//}

</style>
