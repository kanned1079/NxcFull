<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, onMounted, ref} from 'vue'
import {useRouter} from 'vue-router'
// import useSiteInfo from "@/stores/siteInfo";
// import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";

import type {NotificationType} from 'naive-ui'
import {type FormInst, useMessage, useNotification} from 'naive-ui'
import {ChevronBackOutline as backIcon,} from "@vicons/ionicons5"

import authInstance from '@/axios/authInstance'
import useAppInfosStore from "@/stores/useAppInfosStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import {hashPassword} from "@/utils/encryptor";
import instance from "@/axios";

const {t} = useI18n()
const apiAddrStore = useApiAddrStore();
const appInfosStore = useAppInfosStore()
const notification = useNotification()
const themeStore = useThemeStore()
const message = useMessage()
// const siteInfo = useSiteInfo();
// const userInfoStore = useUserInfoStore();
const router = useRouter();

const formRef = ref<FormInst | null>(null)

let animated = ref<{
  leftAnimated: boolean
  rightAnimated: boolean
}>({
  leftAnimated: false,
  rightAnimated: false
})

let bgColor = computed(() => themeStore.enableDarkMode ? {
  backgroundColor: 'rgba(40, 41, 41, 1)',
} : {
  backgroundColor: '#fff'
})

let coverBgColor = computed(() => themeStore.enableDarkMode ? {
  backgroundColor: 'rgba(40, 40, 40, 0.2)',
} : {
  backgroundColor: 'rgba(255, 255, 255, 0.0)',
})

let placeholderBgColor = computed(() => themeStore.enableDarkMode ? {} : {
  backgroundColor: '#f2fafd',
  height: '45px',
})

let textLeftColor = computed(() => themeStore.enableDarkMode ? {
  color: '#fff',
} : {})


let agreementChecked = ref<boolean>(false)  // 同意按钮是否被选中
let enableRegister = ref<boolean>(true) // 由检查逻辑控制
let regBtnEnabled = computed(() => agreementChecked.value && enableRegister.value)
let clickedCount = ref<number>(0) // 注册按钮点击次数
// pass
let enableSendCode = ref<boolean>(true)
let waitSendMail = ref<number>(0)

// 通用配置移动到pinia中

// let userIp = ref({
//   ip_address: '',
//   position: '',
// })

let formValue = ref({
  user: {
    email: '',
    verify_code: '',
    password: '',
    password_confirmation: '',
    invite_user_id: ''
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
    invite_user_id: {
      required: true,
      trigger: 'blur'
    }

  }
}

let notifyErr = (type: NotificationType, title: any, msg: any) => {
  notification[type]({
    content: title,
    meta: msg,
    duration: 2500,
    keepAliveOnHover: true
  })
}

let notifyPass = (type: NotificationType, title: any, msg: any) => {
  notification[type]({
    content: title,
    meta: msg,
    duration: 2500,
    keepAliveOnHover: true
  })
}


// interface DataWithAuth {
//   isAuthed: boolean;
//   user_data: UserData;
// }

// interface UserData {
//   id: number;
//   invite_user_id: number;
//   name: string;
//   email: string;
//   is_admin: boolean;
//   is_staff: boolean;
//   balance: number;
//   last_login: Date;
//   last_login_ip: string;
// }

// const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

let validateEmail = (email: string): boolean => {
  const emailRegex = /^[0-9a-zA-Z_.-]+[@][0-9a-zA-Z_.-]+([.][a-zA-Z]+){1,2}$/;
  return !emailRegex.test(email);
};

// 发送邮箱验证码
let sendVerifyCode = async () => {
  console.log(formValue.value.user.email.trim())
  if (formValue.value.user.email === '' || validateEmail(formValue.value.user.email.trim())) {
    notifyErr('error', '注册', computed(() => t('userRegister.invalidEmailFormat')))
    return
  } else {
    console.log("邮件校验通过 发送邮箱验证")
    try {
      let {data} = await authInstance.post('/api/user/v1/mail/register/get', {
        email: formValue.value.user.email
      })
      if (data.code === 200) {
        enableSendCode.value = false
        startWait(60)
        notifyPass('success', '注册', computed(() => t('userRegister.sendSuccess')))
      } else {
        notifyErr('error', '注册', data.error.toString())
      }
    } catch (error) {
      // notifyErr('error', error.toString())
      console.log(error)
    }

  }
}

// pass
let verifyCode = async () => {
  console.log('验证验证码是否正确')
  try {
    let {data} = await authInstance.post('/api/user/v1/mail/register/verify', {
      email: formValue.value.user.email,
      code: formValue.value.user.verify_code
    })
    if (data.code === 200 && data.passed) {
      console.log('verifyCode 验证码正确返回true')
      console.log(data)
      return true
    } else {
      return false
    }
  } catch (error) {
    // notifyErr('error', error.toString())
    console.log(error)
    return false
  }
  // return false
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
  try {
    let hashedPassword = await hashPassword(formValue.value.user.password.trim() as string)
    let {data} = await instance.post('/api/user/v1/register/register', {
      email: formValue.value.user.email,
      password: hashedPassword,
      invite_user_id: formValue.value.user.invite_user_id,
    })
    if (data.code === 200 && data.is_registered) {
      console.log("注册成功")
      notifyPass('success', computed(() => t('userRegister.reg')), computed(() => t('userRegister.regSuccess')))
      setTimeout(async () => {
        await router.push({
          path: '/login'
        })
      }, 1000)
    } else {
      notifyPass('error', computed(() => t('userRegister.regFailure')), data.msg)
    }
  } catch (error) {
    notifyPass('error', computed(() => t('userRegister.unknowErr')), error)
  }
}


let backgroundStyle = computed(() => ({
  backgroundSize: 'cover', // 或者 'contain' 根据你需要的效果选择
  // backgroundImage: `url(${themeStore.backgroundUrl})`,
  // background: `linear-gradient(to right, rgb(190, 147, 197), rgb(123, 198, 204))`,
  backgroundImage: `url(https://ikanned.com:24444/d/Upload/NXC/pexels-mads-thomsen-2739013-2.jpg)`,
  // backgroundImage: `linear-gradient(to top, #fad0c4 0%, #fad0c4 1%, #ffd1ff 100%)`

}))

// 获取通用的设置
// let getConfig = async () => {
//   let {data} = authInstance.get('/api/v1/get-universal-config')
//   console.log('通用设置', data)
// }

onMounted(async () => {
  // await getUserIpAddress()
  console.log('挂载普通用户注册')

  setTimeout(() => animated.value.leftAnimated = true, 100)
  setTimeout(() => animated.value.rightAnimated = true, 150)
})

let handleValidateClick = async () => {
  // enableLogin.value = false
  clickedCount.value += 1
  console.log('clickedCount: ', clickedCount.value)
  if (clickedCount.value == 5) {
    enableRegister.value = false
    setTimeout(() => {
      enableRegister.value = true
      clickedCount.value = 0
    }, 60000)
    return
  }

  if (formValue.value.user.email === '' || formValue.value.user.password === '' || formValue.value.user.password_confirmation === '') {
    notifyErr('error', computed(() => t('userRegister.reg')), computed(() => t('userRegister.infoIncomplete')))
    return
  } else {  // 有效信息完整
    if (formValue.value.user.password !== formValue.value.user.password_confirmation) {
      notifyErr('error', computed(() => t('userRegister.reg')), computed(() => t('userRegister.pwdIncorrect')))
      return
    }
    if (appInfosStore.registerPageConfig.is_email_verify) { // 如果启用了邮箱验证
      console.log('开始验证邮箱')
      if (await verifyCode()) {
        // 验证码正确
        console.log('验证码正确')
        await handleRegister()
        // handleRegister()  // 注册并写入数据库
        // try {
        //   let {data} = await instance.post(apiAddrStore.apiAddr.user.newUserRegister, {
        //     email: formValue.value.user.email,
        //     password: hashPassword(formValue.value.user.password),
        //     invite_user_id: formValue.value.user.invite_user_id,
        //   })
        //   if (data.code === 200 && data.is_registered) {
        //     console.log("注册成功")
        //     notifyPass('success', computed(() => t('userRegister.reg')), computed(() => t('userRegister.regSuccess')))
        //     setTimeout(async () => {
        //       await router.push({
        //         path: '/login'
        //       })
        //     }, 1000)
        //   } else {
        //     notifyPass('error', computed(() => t('userRegister.regFailure')), data.msg)
        //   }
        // } catch (error) {
        //   notifyPass('error', computed(() => t('userRegister.unknowErr')), error)
        // }
        // return;
      } else {
        notifyErr('error', computed(() => t('userRegister.failure')), computed(() => t('userRegister.verifyCodeErr')))
        return
      }

    } else {
      // 没有启用邮箱验证
      await handleRegister()
    }
  }
  // await handleRegister()
}

let startWait = (sec: number) => {
  let countdown = sec;
  let intervalId = setInterval(async () => {
    if (countdown >= 0) {
      waitSendMail.value = countdown;
      countdown--;
    } else {
      clearInterval(intervalId);
      enableSendCode.value = true;
    }
  }, 1000);
};

// let getUserIpAddress = async () => {
//   try {
//     let {data} = await authInstance.get('https://2024.ipchaxun.com/')
//     userIp.value.ip_address = data.ip.toString()
//     // userIp.value.position = `${response.data.data[0]} + ${response.data.data[1]} + ${response.data.data[2]}`
//     message.info('你的ip地址为' + userIp.value.ip_address)
//   } catch (error) {
//     console.log(error)
//   }
//
// }

</script>

<script lang="ts">
export default {
  name: 'UserRegister',
}
</script>

<template>
  <div class="root">
    <div class="l-content" :style="bgColor">
      <div class="l-content-top" :style="coverBgColor"></div>
      <div class="l-content-color1"></div>
      <div class="l-content-color2"></div>
      <div class="l-content-color3"></div>
      <div class="l-content-color4"></div>
      <transition name="slide-register-fade">
        <div v-if="animated.leftAnimated" class="l-container">
          <p :style="textLeftColor" class="title">{{ appInfosStore.appCommonConfig.app_name }}</p>
          <p :style="textLeftColor" class="sub">{{ appInfosStore.appCommonConfig.app_description }}</p>
        </div>
      </transition>

    </div>
    <div class="r-content" :style="bgColor">
      <div class="r-content-color-cover"></div>
      <div class="r-content-color1"></div>
      <div class="r-content-color2"></div>
      <div class="r-content-color3"></div>
      <transition name="slide-register-fade">
        <n-card v-if="animated.rightAnimated" class="login-card" :embedded="false" :bordered="false">
          <n-button class="back" text :bordered="false" @click="router.replace({path: '/welcome'})">
            <n-icon style="margin-right: 5px" size="20">
              <backIcon/>
            </n-icon>
            {{ t('userRegister.backHomePage') }}
          </n-button>
          <p class="login-title">{{ t('userRegister.newAccount') }}</p>

          <n-form
              ref="formRef"
              :model="formValue"
              :rules="rules">
            <n-form-item
                path="user.email"
                :show-feedback="false"
            >
              <n-input
                  v-model:value="formValue.user.email"
                  :placeholder="t('userRegister.email')"
                  size="large"
                  :bordered="false"
                  :style="placeholderBgColor"

              />
            </n-form-item>
            <n-form-item
                path="user.verify_code"
                :show-feedback="false"
                v-if="appInfosStore.registerPageConfig.is_email_verify"
            >
              <n-input
                  v-model:value.number="formValue.user.verify_code"
                  :placeholder="t('userRegister.verifyCode')"
                  size="large"
                  class="input-general"
                  :bordered="false"
                  :style="placeholderBgColor"

              />
              <n-button
                  :disabled="!enableSendCode"
                  secondary type="primary"
                  @click="sendVerifyCode"
                  style="margin-left: 12px;"
                  size="large">
                {{ t('userRegister.sendVerifyCode') }}
                {{ waitSendMail !== 0 ? ` (${waitSendMail})` : '' }}
              </n-button>
            </n-form-item>
            <n-form-item
                path="user.password"
                :show-feedback="false"
            >
              <n-input
                  type="password"
                  v-model:value="formValue.user.password"
                  :placeholder="t('userRegister.pwd')"
                  size="large"
                  :bordered="false"
                  :style="placeholderBgColor"

              />
            </n-form-item>
            <n-form-item
                path="user.password_confirmation"
                :show-feedback="false"
            >
              <n-input
                  type="password"
                  v-model:value="formValue.user.password_confirmation"
                  :placeholder="t('userRegister.pwdAgain')"
                  size="large"
                  class="input-general"
                  :bordered="false"
                  :style="placeholderBgColor"

              />
            </n-form-item>
            <n-form-item
                path="user.invest_code"
                :show-feedback="false"
                v-if="appInfosStore.registerPageConfig.is_invite_force"
            >
              <n-input
                  v-model:value.number="formValue.user.invite_user_id"
                  :placeholder="t('userRegister.inviteCode')"
                  size="large"
                  class="input-general"
                  :bordered="false"
                  :style="placeholderBgColor"

              />
            </n-form-item>
          </n-form>
          <n-form>
            <n-form-item>
              <n-checkbox v-model:checked="agreementChecked"></n-checkbox>
              <p style="font-weight: bold; opacity: 0.8; margin-left: 8px">{{ t('userRegister.agreement') }}</p>
              <n-button text type="primary" style="font-weight: bold; opacity: 0.9; margin-left: 3px">
                {{ t('userRegister.terminalUserAgreement') }}
              </n-button>
            </n-form-item>
          </n-form>

          <n-button :bordered="false" type="primary" class="login-btn" size="large" @click="handleValidateClick"
                    :disabled="!regBtnEnabled">
            {{ t('userRegister.reg') }}
          </n-button>
        </n-card>
      </transition>
    </div>
  </div>
</template>

<style scoped lang="less">
.slide-register-fade-enter-active {
  transition: all 500ms ease;
}

.slide-register-fade-leave-active {
  transition: all 500ms ease;
}

.slide-register-fade-enter-from,
.slide-register-fade-leave-to {
  transform: translateY(150px);
  opacity: 0;
}

.root {
  height: 100vh;
  display: flex;
  position: relative;
}

.l-content {
  flex: 1;
  position: relative;
  overflow: hidden;
  //background-color: #292929;

  .l-container {
    position: absolute;
    z-index: 1001;
    top: 35%;
    left: 15%;
    //background-color: #4cae4c;
    width: 70%;
    transition: ease 300ms;


    .title {
      font-size: 2.5rem;
      font-weight: 700;
      opacity: 0.8;
      //color: #205387;
      //color: white;
    }

    .sub {
      //color: #136685;
      margin-top: 20px;
      font-size: 1.2rem;
      opacity: 0.5;

      //color: white;
    }
  }

  .l-container:hover {
    transform: translateX(0) translateY(-10px);
  }

  .l-content-top {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    //background-color: rgba(255, 255, 255, 0.0);
    backdrop-filter: blur(90px);
    -webkit-backdrop-filter: blur(100px);
    z-index: 1000;
  }

  .l-content-color1 {
    width: 500px;
    height: 500px;
    border-radius: 50%;
    background-color: rgba(225, 107, 140, 0.2);
    position: absolute;
    //top: calc(70% - 400px / 2);
    //left: calc(30% - 400px / 2);
    //bottom: calc(100% - 100px);
    left: -250px;
    bottom: -250px;
    z-index: 1;
  }

  .l-content-color2 {
    width: 150px;
    height: 150px;
    border-radius: 50%;
    //background-color: rgba(84, 234, 255, 0);
    background: linear-gradient(to right, rgb(129, 129, 221), rgb(142, 177, 255), rgb(131, 255, 244));
    position: absolute;
    top: calc(30% - 150px / 2);
    left: calc(30% - 150px / 2);
    z-index: 2;
  }

  .l-content-color3 {
    width: 200px;
    height: 200px;
    border-radius: 50%;
    //background-color: rgba(220, 135, 227, 0);
    background: linear-gradient(to right, rgb(142, 255, 22), rgb(100, 179, 244));
    position: absolute;
    top: calc(60% - 400px / 2);
    left: calc(50% - 400px / 2);
    z-index: 3;
  }

  .l-content-color4 {
    width: 200px;
    height: 200px;
    border-radius: 50%;
    //background-color: rgba(220, 135, 227, 0);
    background: linear-gradient(to right, rgb(217, 167, 199), rgb(255, 252, 220));
    position: absolute;
    top: calc(50% - 200px / 2);
    left: calc(50% - 200px / 2);
    z-index: 3;
  }
}

@media (max-width: 900px) {
  .l-content {
    display: flex;
    flex: 0;
  }
}

.r-content {
  flex: 1;
  //background-color: #c8d9eb;
  //position: relative;
  display: flex;
  flex-direction: row;

  overflow: hidden;
  justify-content: center;
  align-items: center;

  .login-card {
    z-index: 2000;
    display: flex;
    background-color: rgba(218, 144, 144, 0.0);
    //background-color: #66afe9;
    width: 75%;

    .back {
      //background-color: #31739f;
      font-size: 1rem;
      font-weight: 600;
      margin-bottom: 20px;
      opacity: 0.7;
    }


    .login-title {
      font-size: 2rem;
      font-weight: 500;
      margin-bottom: 10px;
      opacity: 0.9;
    }

    .login-btn {
      margin-top: 50px;
      width: 100%;
      height: 45px;
      color: white;
      background: linear-gradient(to right, rgba(167, 112, 239, 0.5), rgba(207, 139, 243, 0.9), rgba(253, 185, 155, 0.5));
      box-shadow: rgba(139, 129, 195, 0.5) 3px 3px 200px 0;
      margin-bottom: 20px;
      transition: ease 300ms;
    }

    .login-btn:hover {
      transform: translateX(0) translateY(-3px);
    }

    .other-login-method {
      width: 100%;
      height: 45px;
      margin-bottom: 20px;
      transition: ease 200ms;
    }

    .other-login-method:hover {
      transform: translateX(0) translateY(-3px);
    }

    .github {
      color: white;
      background-color: #3a4754;
      box-shadow: #3a4754 1px 1px 5px 0;
    }

    .google {
      color: white;
      background-color: #3382f0;
      box-shadow: #3382f0 1px 1px 5px 0;
    }

    .apple {
      color: white;
      background-color: #000;
      box-shadow: #000 1px 1px 5px 0;
    }
  }

  @media (max-width: 900px) {
    .login-card {
      width: 70%;
    }
  }
  @media (max-width: 768px) {
    .login-card {
      width: 90%;
    }
  }
}

@media (max-width: 900px) {
  .r-content {
    //background: linear-gradient(to right, rgb(217, 167, 199), rgb(255, 252, 220));

  }
}

</style>