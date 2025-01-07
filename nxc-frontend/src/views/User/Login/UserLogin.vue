<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, onMounted, ref} from 'vue'
import {useRouter} from 'vue-router'
import useSiteInfo from "@/stores/siteInfo";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import useThemeStore from "@/stores/useThemeStore";
import type {NotificationType} from 'naive-ui'
import {type FormInst, type FormRules, useMessage, useNotification} from 'naive-ui'
import instance from '@/axios/index'
import {
  ChevronBackOutline as backIcon,
  LogInOutline as loginIcon,
  LogoApple as appleIcon,
  LogoGithub as githubIcon,
  LogoGoogle as googleIcon,
  LogoMicrosoft as microsoftIcon,
} from '@vicons/ionicons5'
import useApiAddrStore from "@/stores/useApiAddrStore";
import {encodeToBase64} from "@/utils/encryptor";

const {t} = useI18n()
const appInfoStore = useAppInfosStore()
const apiAddrStore = useApiAddrStore();
const notification = useNotification()
const themeStore = useThemeStore()
const message = useMessage()
const siteInfo = useSiteInfo();
const userInfoStore = useUserInfoStore();
const router = useRouter();

const loginFormRef = ref<FormInst | null>(null)

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
  backgroundColor: '#f2fafd'
})

let textLeftColor = computed(() => themeStore.enableDarkMode ? {
  color: '#fff',
} : {})

let formValue = ref({
  user: {
    email: '',
    password: '',
    two_fa_code: '',
  }
})

let rules: FormRules = {
  user: {
    email: {
      required: true,
      type: 'email',
      message: '邮件地址是必填的',
      trigger: ['blur', 'input'],
    },
    password: {
      required: true,
      message: '你还没有输入密码',
      trigger: 'blur'
    },
    two_fa_code: {
      required: false,
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
  invite_code: string;
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
  userInfoStore.thisUser.inviteCode = user_data.invite_code
  userInfoStore.thisUser.name = user_data.name
  userInfoStore.thisUser.email = user_data.email
  userInfoStore.thisUser.isAdmin = user_data.is_admin
  userInfoStore.thisUser.isStaff = user_data.is_staff
  userInfoStore.thisUser.balance = user_data.balance
  userInfoStore.thisUser.lastLogin = user_data.last_login ? user_data.last_login.toString() : ''
  userInfoStore.thisUser.lastLoginIp = user_data.last_login_ip
  console.log('绑定结束')
}

let callLogin = async () => {
  console.log('用户数据')
  console.log(formValue.value.user.email, formValue.value.user.password)
  try {
    // let hashedPwd = await hashPassword(formValue.value.user.password)
    // console.log('测试用hashedPwd ', hashedPwd)
    let {data} = await instance.post("http://localhost:8081/api/user/v1/login", {
      email: formValue.value.user.email.trim(),
      password: encodeToBase64(formValue.value.user.password.trim() as string),
      two_fa_code: formValue.value.user.two_fa_code.trim(),
      // password: hashedPwd,
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
      notifyErr("error", data.msg)
    }
  } catch (error) {
    console.log(error)
    // notifyErr('error', error.toString())
    enableLogin.value = true
  }
}

// 处理忘记密码
let handleForgetPassword = () => {
  message.info('处理忘记密码')
  // TODO: 忘记密码步骤不完整
}

let backgroundStyle = computed(() => ({
  backgroundSize: 'cover', // 或者 'contain' 根据你需要的效果选择
  // backgroundImage: `url(${themeStore.backgroundUrl})`,
  background: `linear-gradient(to right, rgb(190, 147, 197), rgb(123, 198, 204))`,

}))

let loginClick = async () => {
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
  await loginFormRef.value?.validate((errors) =>{
    if (errors) { // 表单验证不通过
      message.error('表单不完整')
    } else {
      // message.info('ok11111')
      callLogin() // 表单验证完成后提交登陆请求
    }
  })
}

// let handleValidateClick = () => {
//   enableLogin.value = false
//   clickedCount.value += 1
//   console.log('clickedCount: ', clickedCount.value)
//   if (clickedCount.value == 5) {
//     enableLogin.value = false
//     setTimeout(() => {
//       enableLogin.value = true
//       clickedCount.value = 0
//     }, 60000)
//     return
//   }
//   if (formValue.value.user.email === '' || formValue.value.user.password === '') {
//     notifyErr('error', '邮箱或密码不能为空')
//     return
//   }
//   console.log('验证完成 调用登录')
//   handleLogin()
//   // if (formValue.value.user.password === '') {
//   //   notifyErr('error', '密码不能为空')
//   //   return false
//   // }
// }

onMounted(() => {
  console.log('挂载普通用户登录')
  setTimeout(() => animated.value.leftAnimated = true, 100)
  setTimeout(() => animated.value.rightAnimated = true, 150)
})



</script>


<script lang="ts">
export default {
  name: "UserLogin",
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
      <transition name="slide-login-fade">
        <div v-if="animated.leftAnimated" class="l-container">
          <p :style="textLeftColor" class="title">{{ appInfoStore.appCommonConfig.app_name }}</p>
          <p :style="textLeftColor" class="sub">{{ appInfoStore.appCommonConfig.app_description }}</p>
        </div>
      </transition>

    </div>
    <div class="r-content" :style="bgColor">
      <div class="r-content-color-cover"></div>
      <div class="r-content-color1"></div>
      <div class="r-content-color2"></div>
      <div class="r-content-color3"></div>
      <transition name="slide-login-fade">
        <n-card v-if="animated.rightAnimated" class="login-card" :embedded="false" :bordered="false">
          <n-button class="back" text :bordered="false" @click="router.replace({path: '/'})">
            <n-icon style="margin-right: 5px" size="20">
              <backIcon/>
            </n-icon>
            {{ t('userLogin.backHomePage') }}
          </n-button>
          <p class="login-title">{{ t('userLogin.loginToContinue') }}</p>
          <n-form
              :rules="rules"
              :model="formValue"
              :show-label="false"
              :show-feedback="false"
              size="large"
              ref="loginFormRef"
          >
            <n-form-item path="user.email">
              <n-input
                  :style="placeholderBgColor"
                  :bordered="true"
                  type="text"
                  class="username"
                  :placeholder="t('userLogin.email')"
                  v-model:value="formValue.user.email"
              />
            </n-form-item>
            <n-form-item path="user.password">
              <n-input
                  :style="placeholderBgColor"
                  :bordered="true"
                  type="password"
                  class="password"
                  :placeholder="t('userLogin.password')"
                  v-model:value="formValue.user.password"
              />
            </n-form-item>
            <n-form-item path="user.two_fa_code">
              <n-input
                  type="text"
                  :style="placeholderBgColor"
                  :bordered="true"
                  class="password"
                  maxlength="6"
                  show-count
                  :placeholder="'如果您启用了两步验证（2FA）'"
                  v-model:value="formValue.user.two_fa_code"
              />
            </n-form-item>
          </n-form>


          <div class="no-account">
            <p class="no-account-prefix">{{ t('userLogin.haveNoAccount') }}</p>
            <n-button
                class="no-account-btn"
                text
                @click="router.push({path: '/register'})">{{
                t('userLogin.reg')
              }}
            </n-button>
          </div>
          <n-button
              :bordered="false"
              type="primary"
              class="login-btn"
              size="large"
              @click="loginClick"
              :disabled="!enableLogin"
          >
            <n-icon style="margin-right: 10px" size="25">
              <loginIcon/>
            </n-icon>
            {{ t('userLogin.login') }}
          </n-button>
          <n-divider>
            <p style="opacity: 0.5;  font-weight: 400; font-size: 0.8rem">{{ t('userLogin.otherMethods') }}</p>
          </n-divider>
          <n-button
              :bordered="false"
              class="other-login-method github"
              type="primary"
              @click=" "
              size="large"
          >
            <n-icon style="margin-right: 10px;" size="25">
              <githubIcon/>
            </n-icon>
            {{ t('userLogin.github') }}
          </n-button>
          <n-button :bordered="false" class="other-login-method apple" type="primary">
            <n-icon style="margin-right: 10px;" size="25">
              <appleIcon/>
            </n-icon>
            {{ t('userLogin.apple') }}&nbsp;
          </n-button>
          <n-button :bordered="false" class="other-login-method google" type="primary">
            <n-icon style="margin-right: 10px;" size="25">
              <googleIcon/>
            </n-icon>
            {{ t('userLogin.google') }}
          </n-button>
          <n-button
              :bordered="false"
              class="other-login-method microsoft"
              type="primary"
              size="large"
          >
            <n-icon style="margin-right: 10px;" size="23">
              <microsoftIcon/>
            </n-icon>
            {{ t('userLogin.microsoft') }}
          </n-button>
        </n-card>
      </transition>
      <!--      <n-card class="login-card" :embedded="false" :bordered="false">-->
      <!--        <n-button class="back" text :bordered="false" @click="router.back()">-->
      <!--          <n-icon style="margin-right: 5px" size="20">-->
      <!--            <backIcon/>-->
      <!--          </n-icon>-->
      <!--          {{ t('userLogin.backHomePage') }}-->
      <!--        </n-button>-->
      <!--        <p class="login-title">{{ t('userLogin.loginToContinue') }}</p>-->
      <!--        <n-input-->
      <!--            size="large"-->
      <!--            :style="placeholderBgColor"-->
      <!--            :bordered="false"-->
      <!--            type="text"-->
      <!--            class="username"-->
      <!--            :placeholder="t('userLogin.email')" v-model:value="formValue.user.email"-->
      <!--        ></n-input>-->
      <!--        <n-input-->
      <!--            size="large"-->
      <!--            :style="placeholderBgColor"-->
      <!--            :bordered="false"-->
      <!--            type="password"-->
      <!--            class="password"-->
      <!--            :placeholder="t('userLogin.password')"-->
      <!--            v-model:value="formValue.user.password"-->
      <!--        ></n-input>-->

      <!--        <n-input-->
      <!--            size="large"-->
      <!--            type="text"-->
      <!--            :style="placeholderBgColor"-->
      <!--            :bordered="false"-->
      <!--            class="password"-->
      <!--            maxlength="6"-->
      <!--            show-count-->
      <!--            :placeholder="'如果您启用了两步验证（2FA）'"-->
      <!--            v-model:value="formValue.user.two_fa_code"-->
      <!--        ></n-input>-->
      <!--        <div class="no-account">-->
      <!--          <p class="no-account-prefix">{{ t('userLogin.haveNoAccount') }}</p>-->
      <!--          <n-button-->
      <!--              class="no-account-btn"-->
      <!--              text-->
      <!--              @click="router.push({path: '/register'})">{{-->
      <!--              t('userLogin.reg')-->
      <!--            }}-->
      <!--          </n-button>-->
      <!--        </div>-->
      <!--        <n-button-->
      <!--            :bordered="false"-->
      <!--            type="primary"-->
      <!--            class="login-btn"-->
      <!--            size="large"-->
      <!--            @click="handleValidateClick"-->
      <!--            :disabled="!enableLogin"-->
      <!--        >-->
      <!--          <n-icon style="margin-right: 10px" size="25">-->
      <!--            <loginIcon/>-->
      <!--          </n-icon>-->
      <!--          {{ t('userLogin.login') }}-->
      <!--        </n-button>-->
      <!--        <n-divider>-->
      <!--          <p style="opacity: 0.5;  font-weight: 400; font-size: 0.8rem">{{ t('userLogin.otherMethods') }}</p>-->
      <!--        </n-divider>-->
      <!--        <n-button-->
      <!--            :bordered="false"-->
      <!--            class="other-login-method github"-->
      <!--            type="primary"-->
      <!--            @click=" "-->
      <!--            size="large"-->
      <!--        >-->
      <!--          <n-icon style="margin-right: 10px;" size="25">-->
      <!--            <githubIcon/>-->
      <!--          </n-icon>-->
      <!--          {{ t('userLogin.github') }}-->
      <!--        </n-button>-->
      <!--        &lt;!&ndash;        <n-button :bordered="false" class="other-login-method apple" type="primary">&ndash;&gt;-->
      <!--        &lt;!&ndash;          <n-icon style="margin-right: 10px;" size="25">&ndash;&gt;-->
      <!--        &lt;!&ndash;            <appleIcon/>&ndash;&gt;-->
      <!--        &lt;!&ndash;          </n-icon>&ndash;&gt;-->
      <!--        &lt;!&ndash;          {{ t('userLogin.apple') }}&nbsp;&ndash;&gt;-->
      <!--        &lt;!&ndash;        </n-button>&ndash;&gt;-->
      <!--        &lt;!&ndash;        <n-button :bordered="false" class="other-login-method google" type="primary">&ndash;&gt;-->
      <!--        &lt;!&ndash;          <n-icon style="margin-right: 10px;" size="25">&ndash;&gt;-->
      <!--        &lt;!&ndash;            <googleIcon/>&ndash;&gt;-->
      <!--        &lt;!&ndash;          </n-icon>&ndash;&gt;-->
      <!--        &lt;!&ndash;          {{ t('userLogin.google') }}&ndash;&gt;-->
      <!--        &lt;!&ndash;        </n-button>&ndash;&gt;-->
      <!--        <n-button-->
      <!--            :bordered="false"-->
      <!--            class="other-login-method microsoft"-->
      <!--            type="primary"-->
      <!--            size="large"-->
      <!--        >-->
      <!--          <n-icon style="margin-right: 10px;" size="23">-->
      <!--            <microsoftIcon/>-->
      <!--          </n-icon>-->
      <!--          {{ t('userLogin.microsoft') }}-->
      <!--        </n-button>-->
      <!--      </n-card>-->
    </div>
  </div>
</template>

<style scoped lang="less">

.slide-login-fade-enter-active {
  transition: all 500ms ease;
}

.slide-login-fade-leave-active {
  transition: all 500ms ease;
}

.slide-login-fade-enter-from,
.slide-login-fade-leave-to {
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
    width: 65%;

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

    .username {
      margin-top: 20px;
      border-radius: 3px;
      align-items: center;
      //background-color: #f2fafd;
      //box-shadow: rgba(242,251,254,0.5) 0 1px 10px 0;
    }

    .password {
      margin-top: 30px;
      border-radius: 3px;
      align-items: center;
      //background-color: #f2fafd;
    }

    .no-account {
      margin-top: 10px;
      display: flex;
      flex-direction: row;
      justify-content: right;

      .no-account-prefix {
        font-size: 1rem;
        margin-right: 3px;
        opacity: 0.8;
      }

      .no-account-btn {
        opacity: 0.5;
        font-size: 1rem;
      }
    }

    .login-btn {
      margin-top: 50px;
      width: 100%;
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
      //height: 45px;
      margin-bottom: 20px;
      transition: ease 200ms;
    }

    .other-login-method:hover {
      transform: translateX(0) translateY(-2px);
    }

    .github {
      color: white;
      background-color: #3a4754;
    }

    .github:hover {
      box-shadow: #3a4754 0 5px 20px 0;
    }

    .google {
      color: white;
      background-color: #e84d40;
    }

    .google:hover {
      box-shadow: #e84d40 0 5px 20px 0;
    }

    .apple {
      color: white;
      background-color: #000;
    }

    .apple:hover {
      box-shadow: #000 0 5px 20px 0;
    }

    .microsoft {
      color: white;
      background-color: #00a4ef;
    }

    .microsoft:hover {
      box-shadow: #00a4ef 0 5px 20px 0;
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

//@media (max-width: 900px) {
//  .r-content {
//    //background: linear-gradient(to right, rgb(217, 167, 199), rgb(255, 252, 220));
//  }
//}

</style>