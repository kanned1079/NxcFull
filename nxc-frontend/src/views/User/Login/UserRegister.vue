<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, onMounted, ref} from 'vue'
import {useRouter} from 'vue-router'
import useThemeStore from "@/stores/useThemeStore";

import type {NotificationType} from 'naive-ui'
import {type FormInst, useMessage, useNotification, type FormItemRule} from 'naive-ui'
import {ChevronBackOutline as backIcon,} from "@vicons/ionicons5"

import authInstance from '@/axios/authInstance'
import useAppInfosStore from "@/stores/useAppInfosStore";
import {hashPassword} from "@/utils/encryptor";
import instance from "@/axios";

const {t} = useI18n()
const appInfosStore = useAppInfosStore()
const notification = useNotification()
const themeStore = useThemeStore()
const message = useMessage()
const router = useRouter();

const formRef = ref<FormInst | null>(null)

let verifyCodePassed = ref<boolean>(false)

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
let enableSendCode = ref<boolean>(true)
let waitSendMail = ref<number>(0)

let formValue = ref({
  user: {
    email: '',
    verify_code: '',
    password: '',
    password_confirmation: '',
    invite_user_id: ''
  }
})

// 该选项如果为true则代表禁用gmail多别名
// appInfosStore.registerPageConfig.email_gmail_limit_enable

// 该选项如果为true则代表需要进行邮箱验证
//appInfosStore.registerPageConfig.is_email_verify

let rules = {
  user: {
    // email: {
    //   required: true,
    //   message: '邮箱不能为空',
    //   type: 'email',
    //   trigger: 'blur',
    //   validator: async (rule: FormItemRule, value: string) => {
    //     if (!value) {
    //       return new Error('请输入邮箱');
    //     }
    //
    //     // 检查邮箱格式
    //     const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    //     if (!emailRegex.test(value)) {
    //       return new Error('请输入有效的邮箱地址');
    //     }
    //
    //     // 如果启用了 Gmail 多别名限制
    //     if (appInfosStore.registerPageConfig.email_gmail_limit_enable) {
    //       // 提取邮箱用户名和域名部分
    //       const [localPart, domain] = value.split('@');
    //
    //       // 检查是否为 Gmail 或 Googlemail
    //       const isGmail = domain.toLowerCase() === 'gmail.com';
    //       const isGooglemail = domain.toLowerCase() === 'googlemail.com';
    //       if (isGmail || isGooglemail) {
    //         // 禁止使用 `+` 别名
    //         if (localPart.includes('+')) {
    //           return new Error('禁止使用 Gmail 中的 "+" 别名');
    //         }
    //
    //         // 禁止使用 `.` 作为用户名中的无效分隔符
    //         const normalizedLocalPart = localPart.replace(/\./g, '');
    //         if (normalizedLocalPart !== localPart) {
    //           return new Error('禁止使用 "." 分隔符生成的 Gmail 别名');
    //         }
    //
    //         // 禁止大小写变化（统一小写）
    //         if (localPart !== localPart.toLowerCase()) {
    //           return new Error('禁止通过大小写变换生成 Gmail 别名');
    //         }
    //
    //         // 禁止 `googlemail.com` 作为 Gmail 的别名
    //         if (isGooglemail) {
    //           return new Error('禁止使用 googlemail.com 作为 Gmail 的别名');
    //         }
    //       }
    //     }
    //
    //     return true; // 验证通过
    //   }
    // },
    email: {
      required: true,
      message: '邮箱不能为空',
      type: 'email',
      trigger: 'blur',
      validator: async (rule: FormItemRule, value: string) => {
        if (!value) {
          return new Error('请输入邮箱');
        }

        // 调用 validateEmail 函数进行验证
        const isValid = validateEmail(value);
        if (!isValid) {
          return new Error('邮箱验证未通过，请输入有效邮箱');
        }

        return true; // 验证通过
      },
    },
    verify_code: {
      required: appInfosStore.registerPageConfig.is_email_verify,
      message: '验证码不能为空',
      trigger: 'blur',
      validator: async (rule: FormItemRule, value: string) => {
        // 验证码框是否显示
        if (!appInfosStore.registerPageConfig.is_email_verify) {
          return true; // 不需要验证码时直接通过
        }
        if (!value) {
          return new Error('请输入验证码');
        }
        const codeRegex = /^\d{6}$/;
        if (!codeRegex.test(value)) {
          return new Error('验证码格式不正确');
        }
        //
        // if (!verifyCodePassed.value) {
        //   return new Error('验证码验证未通过，请点击验证按钮完成验证');
        // }

        return true; // 验证通过
      }
    },
    password: {
      required: true,
      message: '密码不能为空',
      trigger: 'blur',
      validator: async (rule: FormItemRule, value: any) => {
        if (!value) {
          return new Error('请输入密码');
        }
        // 校验密码长度
        if (value.length < 6) {
          return new Error('密码长度不能少于6位');
        }
        return true;
      }
    },
    password_confirmation: {
      required: true,
      message: '请确认密码',
      trigger: 'blur',
      validator: async (rule: FormItemRule, value: any) => {
        if (!value) {
          return new Error('请输入确认密码');
        }
        // 校验与密码是否一致
        if (value !== formValue.value.user.password) {
          return new Error('两次密码输入不一致');
        }
        return true;
      }
    },
    invite_user_id: {
      required: false,
    },
  }
};

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

let validateEmail = (email: string): boolean => {
  // 检查邮箱格式的正则
  const emailRegex = /^[0-9a-zA-Z_.-]+@[0-9a-zA-Z_.-]+(\.[a-zA-Z]+){1,2}$/;
  if (!emailRegex.test(email)) {
    message.info('格式无效')
    return false; // 格式无效
  }
  // 提取邮箱用户名和域名部分
  const [localPart, domain] = email.split('@');
  const isGmail = domain.toLowerCase() === 'gmail.com';
  const isGooglemail = domain.toLowerCase() === 'googlemail.com';

  // 如果是 Gmail 或 Googlemail，且启用了多别名限制
  if ((isGmail || isGooglemail) && appInfosStore.registerPageConfig.email_gmail_limit_enable) {
    message.info('判断多别名')
    // 禁止使用 `+` 别名
    if (localPart.includes('+')) {
      return false;
    }

    // 禁止使用 `.` 分隔符（去掉 `.` 后检查是否改变）
    const normalizedLocalPart = localPart.replace(/\./g, '');
    if (normalizedLocalPart !== localPart) {
      return false;
    }

    // 禁止大小写变换（统一小写）
    if (localPart !== localPart.toLowerCase()) {
      return false;
    }

    // 禁止使用 `googlemail.com`
    if (isGooglemail) {
      return false;
    }
  }
  message.success('ok')
  return true; // 邮箱有效
};

// 发送邮箱验证码
let sendVerifyCode = async () => {
  console.log(formValue.value.user.email.trim())
  if (formValue.value.user.email === '' || !validateEmail(formValue.value.user.email.trim())) {
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
      // 验证成功后则为true
      verifyCodePassed.value = data.passed as boolean
      console.log('verifyCode 验证码正确返回true')
      console.log(data)
    } else {
      message.error('验证码错误')
    }
  } catch (error) {
    console.log(error)
    return false
  }
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

onMounted(async () => {
  console.log('挂载普通用户注册')

  // 检查是否是邀请链接 如果是则解析邀请码填入
  const queryParams = new URLSearchParams(window.location.search);
  const code = queryParams.get('code');
  code?formValue.value.user.invite_user_id = code:null

  setTimeout(() => animated.value.leftAnimated = true, 100)
  setTimeout(() => animated.value.rightAnimated = true, 150)
})

// let handleValidateClick = async () => {
//   clickedCount.value += 1
//   console.log('clickedCount: ', clickedCount.value)
//   if (clickedCount.value == 5) {
//     enableRegister.value = false
//     setTimeout(() => {
//       enableRegister.value = true
//       clickedCount.value = 0
//     }, 60000)
//     return
//   }
//
//   if (formValue.value.user.email === '' || formValue.value.user.password === '' || formValue.value.user.password_confirmation === '') {
//     notifyErr('error', computed(() => t('userRegister.reg')), computed(() => t('userRegister.infoIncomplete')))
//     return
//   } else {  // 有效信息完整
//     if (formValue.value.user.password !== formValue.value.user.password_confirmation) {
//       notifyErr('error', computed(() => t('userRegister.reg')), computed(() => t('userRegister.pwdIncorrect')))
//       return
//     }
//     if (appInfosStore.registerPageConfig.is_email_verify) { // 如果启用了邮箱验证
//       console.log('开始验证邮箱')
//       if (await verifyCode()) {
//         // 验证码正确
//         console.log('验证码正确')
//         await handleRegister()
//       } else {
//         notifyErr('error', computed(() => t('userRegister.failure')), computed(() => t('userRegister.verifyCodeErr')))
//         return
//       }
//
//     } else {
//       // 没有启用邮箱验证
//       await handleRegister()
//     }
//   }
//   // await handleRegister()
// }

let handleValidateClick = async () => {
  clickedCount.value += 1;
  console.log('clickedCount: ', clickedCount.value);

  // 限制连续点击注册按钮的次数
  if (clickedCount.value === 5) {
    enableRegister.value = false;
    setTimeout(() => {
      enableRegister.value = true;
      clickedCount.value = 0;
    }, 60000);
    return;
  }

  // 表单验证
  if (!formRef.value) {
    message.error('Form reference is not defined');
    return;
  }

  try {
    // 验证表单
    message.info('验证表单')
    await formRef.value.validate();

    // await verifyCode(formValue.value.user.verify_code)
    // if (await verifyCode()) {
    //   message.success('ok')
    // } else {
    //   message.error('验证码错误')
    //   return
    // }

    await verifyCode()

    // 如果启用了邮箱验证但未通过验证码验证（应通过规则控制）
    if (
        appInfosStore.registerPageConfig.is_email_verify &&
        !verifyCodePassed.value
    ) {
      notifyErr(
          'error',
          t('userRegister.failure'),
          t('userRegister.verifyCodeErr')
      );
      return;
    }



    // 执行注册逻辑
    await handleRegister();
    message.info('注册成功');
  } catch (errors) {
    message.error('表单验证未通过:' + errors);
    // 验证失败由 Naive UI 的反馈机制展示
  }
};

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

<!--          <n-form-->
<!--              ref="formRef"-->
<!--              :model="formValue"-->
<!--              :rules="rules">-->
<!--            <n-form-item-->
<!--                path="user.email"-->
<!--                :show-feedback="false"-->
<!--            >-->
<!--              <n-input-->
<!--                  v-model:value="formValue.user.email"-->
<!--                  :placeholder="t('userRegister.email')"-->
<!--                  size="large"-->
<!--                  :bordered="false"-->
<!--                  :style="placeholderBgColor"-->

<!--              />-->
<!--            </n-form-item>-->
<!--            <n-form-item-->
<!--                path="user.verify_code"-->
<!--                :show-feedback="false"-->
<!--                v-if="appInfosStore.registerPageConfig.is_email_verify"-->
<!--            >-->
<!--              <n-input-->
<!--                  v-model:value.number="formValue.user.verify_code"-->
<!--                  :placeholder="t('userRegister.verifyCode')"-->
<!--                  size="large"-->
<!--                  class="input-general"-->
<!--                  :bordered="false"-->
<!--                  :style="placeholderBgColor"-->

<!--              />-->
<!--              <n-button-->
<!--                  :disabled="!enableSendCode"-->
<!--                  secondary type="primary"-->
<!--                  @click="sendVerifyCode"-->
<!--                  style="margin-left: 12px;"-->
<!--                  size="large">-->
<!--                {{ t('userRegister.sendVerifyCode') }}-->
<!--                {{ waitSendMail !== 0 ? ` (${waitSendMail})` : '' }}-->
<!--              </n-button>-->
<!--            </n-form-item>-->
<!--            <n-form-item-->
<!--                path="user.password"-->
<!--                :show-feedback="false"-->
<!--            >-->
<!--              <n-input-->
<!--                  type="password"-->
<!--                  v-model:value="formValue.user.password"-->
<!--                  :placeholder="t('userRegister.pwd')"-->
<!--                  size="large"-->
<!--                  :bordered="false"-->
<!--                  :style="placeholderBgColor"-->

<!--              />-->
<!--            </n-form-item>-->
<!--            <n-form-item-->
<!--                path="user.password_confirmation"-->
<!--                :show-feedback="false"-->
<!--            >-->
<!--              <n-input-->
<!--                  type="password"-->
<!--                  v-model:value="formValue.user.password_confirmation"-->
<!--                  :placeholder="t('userRegister.pwdAgain')"-->
<!--                  size="large"-->
<!--                  class="input-general"-->
<!--                  :bordered="false"-->
<!--                  :style="placeholderBgColor"-->

<!--              />-->
<!--            </n-form-item>-->
<!--            <n-form-item-->
<!--                path="user.invest_code"-->
<!--                :show-feedback="false"-->
<!--                v-if="appInfosStore.registerPageConfig.is_invite_force"-->
<!--            >-->
<!--              <n-input-->
<!--                  v-model:value.number="formValue.user.invite_user_id"-->
<!--                  :placeholder="t('userRegister.inviteCode')"-->
<!--                  size="large"-->
<!--                  class="input-general"-->
<!--                  :bordered="false"-->
<!--                  :style="placeholderBgColor"-->

<!--              />-->
<!--            </n-form-item>-->

<!--            <n-form-item>-->
<!--              <n-checkbox v-model:checked="agreementChecked"></n-checkbox>-->
<!--              <p style="font-weight: bold; opacity: 0.8; margin-left: 8px">{{ t('userRegister.agreement') }}</p>-->
<!--              <n-button text type="primary" style="font-weight: bold; opacity: 0.9; margin-left: 3px">-->
<!--                {{ t('userRegister.terminalUserAgreement') }}-->
<!--              </n-button>-->
<!--            </n-form-item>-->

<!--            <n-form-item>-->
<!--              <n-button :bordered="false" type="primary" class="login-btn" size="large" @click="handleValidateClick"-->
<!--                        :disabled="!regBtnEnabled">-->
<!--                {{ t('userRegister.reg') }}-->
<!--              </n-button>-->
<!--            </n-form-item>-->

<!--          </n-form>-->

          <n-form
              ref="formRef"
              :model="formValue"
              :rules="rules"
          >
            <!-- 邮箱输入框 -->
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

            <!-- 验证码输入框 -->
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
                  secondary
                  type="primary"
                  @click="sendVerifyCode"
                  style="margin-left: 12px;"
                  size="large"
              >
                {{ t('userRegister.sendVerifyCode') }}
                {{ waitSendMail !== 0 ? ` (${waitSendMail})` : '' }}
              </n-button>
            </n-form-item>

            <!-- 密码输入框 -->
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

            <!-- 确认密码输入框 -->
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

            <!-- 邀请码输入框 -->
            <n-form-item
                path="user.invite_user_id"
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

            <!-- 用户协议 -->
            <n-form-item>
              <n-checkbox v-model:checked="agreementChecked"></n-checkbox>
              <p style="font-weight: bold; opacity: 0.8; margin-left: 8px">{{ t('userRegister.agreement') }}</p>
              <n-button
                  text
                  type="primary"
                  style="font-weight: bold; opacity: 0.9; margin-left: 3px"
              >
                {{ t('userRegister.terminalUserAgreement') }}
              </n-button>
            </n-form-item>

            <!-- 注册按钮 -->
            <n-form-item>
              <n-button
                  :bordered="false"
                  type="primary"
                  class="login-btn"
                  size="large"
                  @click="handleValidateClick"
                  :disabled="!regBtnEnabled"
              >
                {{ t('userRegister.reg') }}
              </n-button>
            </n-form-item>
          </n-form>
<!--          <n-form>-->
<!--            <n-form-item>-->
<!--              <n-checkbox v-model:checked="agreementChecked"></n-checkbox>-->
<!--              <p style="font-weight: bold; opacity: 0.8; margin-left: 8px">{{ t('userRegister.agreement') }}</p>-->
<!--              <n-button text type="primary" style="font-weight: bold; opacity: 0.9; margin-left: 3px">-->
<!--                {{ t('userRegister.terminalUserAgreement') }}-->
<!--              </n-button>-->
<!--            </n-form-item>-->
<!--          </n-form>-->

<!--          <n-button :bordered="false" type="primary" class="login-btn" size="large" @click="handleValidateClick"-->
<!--                    :disabled="!regBtnEnabled">-->
<!--            {{ t('userRegister.reg') }}-->
<!--          </n-button>-->
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