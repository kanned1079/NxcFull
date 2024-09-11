<script setup lang="ts" name="UserRegister">
import {computed, onMounted, ref} from 'vue'
import {useRouter} from 'vue-router'
import useSiteInfo from "@/stores/siteInfo";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";
import type {NotificationType} from 'naive-ui'
import {FormInst, useMessage, useNotification} from 'naive-ui'
import instance from '@/axios/index'

const notification = useNotification()
const themeStore = useThemeStore()
const message = useMessage()
const siteInfo = useSiteInfo();
const userInfoStore = useUserInfoStore();
const router = useRouter();

const formRef = ref<FormInst | null>(null)

interface uniConfig{
  app_description: string
  app_url: string
  email_whitelist_suffix: boolean
  is_email_verify: boolean
  is_invite_force: boolean
  is_recaptcha: boolean
  logo: string
  recaptcha_site_key: string
  tos_url: string
}

let universalConfig = reactive<uniConfig>({
  app_description: '',
  app_url: '',
  email_whitelist_suffix: false,
  is_email_verify: false,
  is_invite_force: false,
  is_recaptcha: false,
  logo: '',
  recaptcha_site_key: '',
  tos_url: '',
})

let formValue = ref({
  user: {
    email: '',
    verify_code: '',
    password: '',
    password_confirmation: '',
    invest_code: '',
    agreement: false,
  }
})
let rules = {
  user: {
    email: {
      required: true,
      trigger: 'blur'
    },
    verify_code: {
      required: true,
      trigger: 'blur'
    },
    password: {
      required: true,
      trigger: 'blur'
    },
    password_confirmation: {
      required: true,
      trigger: 'blur'
    },
    invest_code: {
      required: true,
      trigger: 'blur'
    }

  }
}

let notifyErr = (type: NotificationType, msg: string) => {
  notification[type]({
    content: '注册',
    meta: msg,
    duration: 2500,
    keepAliveOnHover: true
  })
}

let notifyPass = (type: NotificationType) => {
  notification[type]({
    content: '注册成功返回登录',
    // meta: '1111',
    duration: 2500,
    keepAliveOnHover: true
  })
}

let agreementChecked = ref<boolean>(true)
let verify = ref<boolean>(false)
// pass
let enableRegister = computed(() => '')
let enableLoginViaForm = computed(() => agreementChecked.value)
let clickedCount = ref<number>(0)
// pass

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

let sendVerifyCode = async () => {
  console.log("发送邮箱验证")
  let {data} = instance.post('/api/v1/')

}

// 返回登录界面
let backToLogin = () => {
  console.log('返回登录')
  router.push({
    path: '/login',
  })
}

// 处理注册
let handleRegister = async () => {
  console.log('用户注册信息', formRef.value)
  try {
    let {data} = instance.post('/api/v1/user-register', formValue)
    console.log(data)
    if (data.code === 200) {
      notifyPass('注册成功')
    } else {
      notifyErr(data.msg)
    }
  }catch (error) {
    notifyErr(error)
  }

}


let backgroundStyle = computed(() => ({
  backgroundSize: 'cover', // 或者 'contain' 根据你需要的效果选择
  // backgroundImage: `url(${themeStore.backgroundUrl})`,
  background: `linear-gradient(to right, rgb(190, 147, 197), rgb(123, 198, 204))`,

}))

// 获取通用的设置
let getConfig = async () => {
  let {data} = instance.get('/api/v1/get-universal-config')
  console.log('通用设置', data)
}

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
  handleRegister()
  // if (formValue.value.user.password === '') {
  //   notifyErr('error', '密码不能为空')
  //   return false
  // }
}


</script>

<template>

  <n-layout style="width: 100%; height: 100vh;" justify="center" :vertical="true" align="center"
            :style="backgroundStyle">
    <n-flex justify="center" :vertical="true" align="center" style="gap: 0">
      <n-card class="layer-up" :embedded="true">
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

            <n-form-item path="user.verify_code" :show-feedback="false">
              <n-input v-model:value="formValue.user.verify_code" placeholder="邮箱验证码" size="large"/>
              <n-button @click="sendVerifyCode" secondary style="margin-left: 20px" size="large" type="primary">
                发送验证码
              </n-button>
            </n-form-item>

            <n-form-item path="user.password" :show-feedback="false">
              <n-input v-model:value="formValue.user.password" placeholder="密码" size="large"/>
            </n-form-item>

            <n-form-item path="user.password_confirmation" :show-feedback="false">
              <n-input v-model:value="formValue.user.password_confirmation" placeholder="确认密码" size="large"/>
            </n-form-item>

            <n-form-item path="user.invest_code" :show-feedback="false">
              <n-input v-model:value="formValue.user.invest_code" placeholder="邀请码（选填）" size="large"/>
            </n-form-item>

            <n-form-item path="user.agreement" :show-feedback="false">
              <n-checkbox v-model:checked="formValue.user.agreement"></n-checkbox>
              <p style="font-weight: bold; opacity: 0.8; margin-left: 8px">我已经阅读并同意</p>
              <n-button text type="primary" style="font-weight: bold; opacity: 0.9; margin-left: 5px">用户条款
              </n-button>
            </n-form-item>
          </n-form>
        </div>
        <n-button secondary type="info" class="login-btn" size="large" @click="handleValidateClick"
                  :disabled="(!enableLogin) && enableLoginViaForm">
          注册
        </n-button>

      </n-card>
            <n-card class="layer-down" content-style="padding: 0;">
              <div class="bottom-root">
                <div class="l-con">
                  <n-button text @click="backToLogin">返回登录</n-button>
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
  height: 620px;
  border: 0;
  border-radius: 5px 5px 0 0;
  margin-bottom: 0;

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
    margin-top: 30px;
    width: 90%;
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
  background-color: v-bind('themeStore.getTheme.globeTheme.loginCardBgColor');
  border: 0;
}

//element.style {
//  gap: 0;
//}

</style>
