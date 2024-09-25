<script setup lang="ts">
import {computed, onMounted, ref} from 'vue'
import {useRouter} from 'vue-router'
// import useSiteInfo from "@/stores/siteInfo";
// import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";
import type {NotificationType} from 'naive-ui'
import {FormInst, useMessage, useNotification} from 'naive-ui'
import {GlobeOutline as LanguageIcon} from '@vicons/ionicons5'

import authInstance from '@/axios/authInstance'
import useAppInfosStore from "@/stores/useAppInfosStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import {hashPassword} from "@/utils/encryptor";
import instance from "@/axios";

const apiAddrStore = useApiAddrStore();
const appInfosStore = useAppInfosStore()
const notification = useNotification()
const themeStore = useThemeStore()
const message = useMessage()
// const siteInfo = useSiteInfo();
// const userInfoStore = useUserInfoStore();
const router = useRouter();

const formRef = ref<FormInst | null>(null)


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

let notifyErr = (type: NotificationType, title: string, msg: string) => {
  notification[type]({
    content: title,
    meta: msg,
    duration: 2500,
    keepAliveOnHover: true
  })
}

let notifyPass = (type: NotificationType, title: string, msg: string) => {
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
    notifyErr('error', '注册', '邮件格式不合法')
    return
  } else {
    console.log("邮件校验通过 发送邮箱验证")
    try {
      let {data} = await authInstance.post(apiAddrStore.apiAddr.user.sendAuthMail, {
        email: formValue.value.user.email
      })
      if (data.code === 200) {
        enableSendCode.value = false
        startWait(60)
        notifyPass('success', '注册', '邮件发送成功 请查看邮箱')
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
    let {data} = await authInstance.post(apiAddrStore.apiAddr.user.verifyAuthMail, {
      email: formValue.value.user.email,
      code: formValue.value.user.verify_code
    })
    if (data.code === 200) {
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
  console.log('用户注册信息', formRef.value)
  try {
    let {data} = authInstance.post(apiAddrStore.apiAddr.user.newUserRegister, {
      email: formValue.value.user.email,
      password: formValue.value.user.password,
      invite_user_id: formValue.value.user.invite_user_id,
    })
    console.log(data)
    if (data.code === 200) {
      notifyPass('success', '注册', '注册成功')
    } else {
      notifyErr('error', '注册', data.msg.toString())
    }
  } catch (error) {
    // notifyErr('error',error.toString())
    console.log(error)
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
    notifyErr('error', '注册', '缺少有效信息')
    return
  } else {  // 有效信息完整
    if (formValue.value.user.password !== formValue.value.user.password_confirmation) {
      notifyErr('error', '注册', '密码不一致')
      return
    }
    if (appInfosStore.registerPageConfig.is_email_verify) { // 如果启用了邮箱验证
      console.log('开始验证邮箱')
      if (await verifyCode()) {
        // 验证码正确
        console.log('验证码正确')
        // handleRegister()  // 注册并写入数据库
        try {
          let {data} = await instance.post(apiAddrStore.apiAddr.user.newUserRegister, {
            email: formValue.value.user.email,
            password: hashPassword(formValue.value.user.password),
            invite_user_id: formValue.value.user.invite_user_id,
          })
          if (data.code === 200) {
            console.log("注册成功")
            notifyPass('success', '注册', '注册成功 返回登录')
            setTimeout(async () => {
              await router.push({
                path: '/login'
              })
            }, 1000)
          } else {
            notifyPass('error', '注册失败', data.msg)

          }
        } catch (error) {
          notifyPass('error', '未知错误', error.toString())
        }
        return;
      } else {
        notifyErr('error', '注册', '验证码错误')
        return
      }

    }
  }
  await handleRegister()
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

  <n-layout style="width: 100%; height: 100vh;" justify="center" :vertical="true" align="center"
            :style="backgroundStyle">
    <n-flex justify="center" :vertical="true" align="center" style="gap: 0">
      <n-card class="layer-up" :embedded="true">
        <p class="title">{{ appInfosStore.registerPageConfig.app_name }}</p>
        <p class="sub-title">{{ appInfosStore.registerPageConfig.app_description }}</p>
        <div class="inp">
          <n-form
              ref="formRef"
              :model="formValue"
              :rules="rules">
            <n-form-item path="user.email" :show-feedback="false">
              <n-input v-model:value="formValue.user.email" placeholder="邮箱地址" size="large"
                       style="user-select: none"/>
            </n-form-item>
            <n-form-item path="user.verify_code" :show-feedback="false"
                         v-if="appInfosStore.registerPageConfig.is_email_verify">
              <n-input v-model:value.number="formValue.user.verify_code" placeholder="邮箱验证码" size="large"/>
              <n-button :disabled="!enableSendCode" secondary type="primary" @click="sendVerifyCode"
                        style="margin-left: 12px" size="large">
                发送验证码{{ waitSendMail !== 0 ? ` (${waitSendMail})` : '' }}
              </n-button>
            </n-form-item>
            <n-form-item path="user.password" :show-feedback="false">
              <n-input type="password" v-model:value="formValue.user.password" placeholder="密码" size="large"/>
            </n-form-item>
            <n-form-item path="user.password_confirmation" :show-feedback="false">
              <n-input type="password" v-model:value="formValue.user.password_confirmation" placeholder="确认密码"
                       size="large"/>
            </n-form-item>
            <n-form-item path="user.invest_code" :show-feedback="false"
                         v-if="appInfosStore.registerPageConfig.is_invite_force">
              <n-input v-model:value.number="formValue.user.invite_user_id" placeholder="邀请码（选填）" size="large"/>
            </n-form-item>
          </n-form>
          <n-form>
            <n-form-item>
              <n-checkbox v-model:checked="agreementChecked"></n-checkbox>
              <p style="font-weight: bold; opacity: 0.8; margin-left: 8px">我已经阅读并同意</p>
              <n-button text type="primary" style="font-weight: bold; opacity: 0.9; margin-left: 3px">用户条款
              </n-button>
            </n-form-item>
          </n-form>
        </div>
        <n-button type="primary" class="login-btn" size="large" @click="handleValidateClick" :disabled="!regBtnEnabled">
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
              <n-button text>语言</n-button>
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
  border: 0;
  border-radius: 5px 5px 0 0;
  margin-bottom: 0;

  .title {
    margin-top: 20px;
    font-size: 30px;
  }

  .sub-title {
    font-size: 1rem;
    opacity: 0.8;
  }

  .inp {
    margin-top: 5px;
    text-align: left;
    width: 95%;
  }

  .login-btn {
    margin-top: 10px;
    width: 95%;
  }

  .register-btn {
    margin-top: 30px;
    width: 95%;
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
    backdrop-filter: blur(20px);
    //background-color: rgba(225, 225, 225, 0.0);
    background-color: rgba(202, 202, 202, 0.1);
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
  backdrop-filter: blur(50px);
  //background-color: v-bind('themeStore.getTheme.globeTheme.loginCardBgColor');
  //border: 0;

}

//element.style {
//  gap: 0;
//}

</style>
