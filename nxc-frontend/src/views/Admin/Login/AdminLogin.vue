<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, onMounted, ref} from 'vue'
import {useRouter} from 'vue-router'
import useAppInfosStore from "@/stores/useAppInfosStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";
import {type FormInst, type FormRules, type InputProps, NButton, NIcon,  type NotificationType} from 'naive-ui'
import {useMessage, useNotification} from 'naive-ui'
import {encodeToBase64} from '@/utils/encryptor'
import instance from "@/axios";
import {ChevronBackOutline as BackIcon} from "@vicons/ionicons5"
import {
  ChevronBackOutline as backIcon,
  LogInOutline as loginIcon,
  LogoGithub as githubIcon,
  LogoMicrosoft as microsoftIcon,
  ReloadOutline as reloadIcon,
  Home as homeIcon,
} from '@vicons/ionicons5'

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

interface AdminForm {
  username: string,
  password: string,
}

const userInfoStore = useUserInfoStore();
const router = useRouter();
let userFormInstance = ref<FormInst | null>(null)
let userFormData = ref<AdminForm>({
  username: '',
  password: '',
})
const rules: FormRules = {
  username: {
    required: true,
    type: "email",
    trigger: ["blur", "change"],
  },
  password: {
    required: true,
    type: "string",
    trigger: ["blur", "change"],
  }
}

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

let bindUserInfo = (data: DataWithAuth) => {
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

let handleLoginClick = async (e: MouseEvent) => {
  e.preventDefault()
  await userFormInstance.value?.validate((errors) => {
    if (errors) {
      message.error('表单验证不通过')
    } else {
      submitLogin()
    }
  })

}

let submitLogin = async () => {
  enableLogin.value = false
  console.log('登陆信息', userFormData.value)
  try {
    // let hashedPwd =  hashPassword(password.value.trim())
    let {data} = await instance.post('http://localhost:8081/api/admin/v1/login', {
      email: userFormData.value.username,
      password: encodeToBase64(userFormData.value.password),
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
      await router.push({path: '/admin/dashboard/summary'});
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
  } catch (error) {
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

let showStartupNotification = () => {
  notification.create({
    title: '提示',
    description: '此页面是管理员页面，只有管理员才能访问，如果您没有管理员权限或意外来到此处，请点击下面的链接返回主页。',
    content: () => {
      return h('div', [
        h(
            NButton,
            {
              text: true,
              type: 'primary',
              style: { textDecoration: 'underline' },
              onClick: () => {
                // 路由跳转
                window.location.replace('/'); // 使用 window.location.replace() 进行页面跳转
              },
            },
            () => '返回主页' // 使用函数插槽渲染按钮文本
        ),
      ]);
    },
    // meta: new Date().toLocaleString(),
    onClose: () => {
    },
  });
};

onMounted(() => {
  console.log('AdminLogin挂载')
  // userInfoStore.isAuthed = false
  // sessionStorage.setItem('isAuthed', JSON.stringify(false))
  console.log('6666', JSON.parse(sessionStorage.getItem('isAuthed') as string))
  console.log('store.isAuthed: ', userInfoStore.isAuthed)

  if (JSON.parse(sessionStorage.getItem('isAuthed') as string)) {
    console.log('to dashboard')
    router.push({path: '/admin/dashboard/summary'})
  }
  animated.value = true
  showStartupNotification()

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
          <div v-if="animated" class="card-all-root">
            <n-card
                class="layer-up"
                :embedded="true"
                :bordered="false"
                content-style="padding: 30px 35px 45px 35px; text-align: center;"
            >
              <div style="display: flex; justify-content: flex-start; margin-bottom: 30px;">
                <n-button
                  type="primary"
                  icon-placement="left"
                  text
                  @click="router.push({path: '/'})"
                >
                  返回首页
                  <template #icon>
                    <n-icon><BackIcon /></n-icon>
                  </template>
                </n-button>
              </div>
              <p class="title">{{ appInfoStore.appCommonConfig.app_name }}</p>
              <p class="sub-title">登陆到管理中心</p>
              <div class="inp">
                <n-form
                    ref="userFormInstance"
                    :model="userFormData"
                    :rules="rules"
                    size="large"
                    style="width: 100%"
                    :show-label="false"
                    :show-feedback="false"
                >
                  <n-form-item path="username">
                    <n-input
                        v-model:value="userFormData.username"
                        type="text"
                        placeholder="邮箱"
                        style="opacity: 1"
                        :bordered="true"
                    />
                  </n-form-item>
                  <n-form-item path="password" style="margin-top: 20px">
                    <n-input
                        v-model:value="userFormData.password"
                        type="password"
                        placeholder="密码"
                        :bordered="true"
                    />
                  </n-form-item>
<!--                  <n-divider style="margin: 10px 0 10px 0 !important;">-->
<!--                    <p style="opacity: 0.2">·</p>-->
<!--                  </n-divider>-->

                  <n-form-item>
                    <n-button
                        secondary
                        type="info"
                        class="login-btn"
                        @click="handleLoginClick"
                        :disabled="!enableLogin"
                        icon-placement="left"
                    >
                      登入
                      <template #icon>
                        <n-icon><loginIcon/></n-icon>
                      </template>
                    </n-button>
                  </n-form-item>
                </n-form>



<!--                <p style="opacity: 0.9">{{ appInfoStore.appCommonConfig.app_description }}</p>-->

<!--                <div>-->
<!--                  <n-h3>提示</n-h3>-->
<!--                  <p style="font-size: 1rem">此页面是管理员页面，只有管理员才能访问。如果您没有管理员权限或意外来到此处，请点击-->
<!--                    <n-button-->
<!--                        text-->
<!--                        type="primary"-->
<!--                        style="text-decoration: underline"-->
<!--                        @click="router.replace('/')"-->
<!--                    >-->
<!--                      这里-->
<!--                    </n-button>-->
<!--                    返回主页。-->
<!--                  </p>-->
<!--                </div>-->

              </div>

            </n-card>
            <n-card
                class="layer-down"
                :bordered="false"
                :embedded="true"
                content-style="padding: 0"
            >
              <div class="layer-down-inner">
                <n-button
                    tertiary
                    type="default"
                    class="reset-pwd-btn"
                    @click="handleForgetPassword"
                >
                  忘记密码
<!--                  <template #icon>-->
<!--                    <n-icon size="14"><reloadIcon /></n-icon>-->
<!--                  </template>-->
                </n-button>
              </div>
            </n-card>
          </div>
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
  backdrop-filter: blur(4px);
  border-radius: 3px 3px 0 0;

  .title {
    font-size: 2rem;
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

@media (max-width: 768px) {
  .layer-up {
    width: 360px;

    .title {
      font-size: 1.3rem;
      opacity: 1;
    }
  }
}

.layer-down {
  padding: 0;
  border-radius: 0 0 3px 3px;
  opacity: 0.9;

}

.card-all-root {
  transition: ease 300ms;
  border-radius: 3px;

  .layer-down-inner {
    text-align: center;

    .reset-pwd-btn {
      height: 52px;
      width: 100%;
      border-radius: 0 0 3px 3px;
    }
  }
}

.card-all-root:hover {
  box-shadow: 3px 3px 10px rgba(25, 25, 25, 0.2);
}
</style>