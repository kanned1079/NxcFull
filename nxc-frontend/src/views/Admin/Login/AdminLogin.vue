<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onMounted, onBeforeMount, ref} from 'vue'
import {useRouter, useRoute} from 'vue-router'
import useAppInfosStore from "@/stores/useAppInfosStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";
import {
  type FormInst,
  type FormRules,
  NButton,
  NIcon,
  type NotificationType,
  useMessage,
  useNotification
} from 'naive-ui'
import {encodeToBase64} from '@/utils/encryptor'
import instance from "@/axios";
import {
  ChevronBackOutline as BackIcon,
  ChevronBackOutline as backIcon,
  LogInOutline as loginIcon
} from "@vicons/ionicons5"

const {t} = useI18n()
const notification = useNotification()
const appInfoStore = useAppInfosStore()
const themeStore = useThemeStore()
const message = useMessage()

let animated = ref<boolean>(false)

let notifyErr = (type: NotificationType, msg: string) => {
  notification[type]({
    content: computed(() => t('adminViews.login.card.authFailure')).value,
    meta: msg,
    duration: 2500,
    keepAliveOnHover: true
  })
}

let notifyPass = (type: NotificationType) => {
  notification[type]({
    content: computed(() => t('adminViews.login.card.authPassed')).value,
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
const route = useRoute();
let sourcePath = route.params.path as string || ''
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
     return message.error(computed(() => t('adminViews.login.card.formNotPassed')).value)
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
        case 'incorrect_password':
          return notifyErr('error', computed(() => t('adminViews.login.message.passwordErr')).value)

        case 'user_not_exist':
          return notifyErr('error', computed(() => t('adminViews.login.message.adminNotExist')).value)

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
    backgroundColor: themeStore.enableDarkMode ? '#252525' : '#e3e5e7',
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
  let _title = computed(() => t('adminViews.login.mention.title'))
  let _description = computed(() => t('adminViews.login.mention.description'))
  let _toMainPage = computed(() => t('adminViews.login.card.back'))
  notification.create({
    title: _title.value,
    duration: 10000,
    description: _description.value,
    content: () => {
      return h('div', [
        h(
            NButton,
            {
              text: true,
              type: 'primary',
              style: {textDecoration: 'underline'},
              onClick: () => {
                // 路由跳转
                window.location.replace('/'); // 使用 window.location.replace() 进行页面跳转
              },
            },
            () => _toMainPage.value // 使用函数插槽渲染按钮文本
        ),
      ]);
    },
    // meta: new Date().toLocaleString(),
    onClose: () => {
    },
  });
};

onBeforeMount(() => {
  if (appInfoStore.appCommonConfig.secure_path !== sourcePath.trim()) {
    message.error('请求参数不正确')
    return router.replace("/")
  } else message.success('请求参数正确')
})

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
                content-style="padding: 30px 35px 40px 35px; text-align: center;"
            >
              <div style="display: flex; justify-content: flex-start; margin-bottom: 30px;">
                <n-button
                    type="primary"
                    icon-placement="left"
                    text
                    @click="router.push({path: '/'})"
                >
                  {{ t('adminViews.login.card.back') }}
                  <template #icon>
                    <n-icon>
                      <BackIcon/>
                    </n-icon>
                  </template>
                </n-button>
              </div>
              <p class="title">{{ appInfoStore.appCommonConfig.app_name }}</p>
              <p class="sub-title">
                {{ t('adminViews.login.card.toAdminCenter') }}
              </p>
              <div class="inp" style="margin-top: 40px">
                <n-form
                    ref="userFormInstance"
                    :model="userFormData"
                    :rules="rules"
                    size="large"
                    style="width: 100%"
                    :show-label="true"
                    :show-feedback="false"
                    label-placement="top"
                    label-align="left"
                >
                  <n-form-item path="username" :label="t('adminViews.login.card.emailInputArea.title')">
                    <n-input
                        clearable
                        v-model:value="userFormData.username"
                        type="text"
                        :placeholder="t('adminViews.login.card.emailInputArea.placeholder')"
                        style="opacity: 1"
                        :bordered="true"
                    />
                  </n-form-item>
                  <n-form-item path="password" style="margin-top: 20px" :label="t('adminViews.login.card.passwordInputArea.title')">
                    <n-input
                        v-model:value="userFormData.password"
                        type="password"
                        showPasswordOn="click"
                        :placeholder="t('adminViews.login.card.passwordInputArea.placeholder')"
                        :bordered="true"
                    />
                  </n-form-item>
                  <n-form-item>
                    <n-button
                        secondary
                        type="info"
                        class="login-btn"
                        @click="handleLoginClick"
                        :disabled="!enableLogin"
                        icon-placement="left"
                    >
                      {{ t('adminViews.login.card.login') }}
                      <template #icon>
                        <n-icon>
                          <loginIcon/>
                        </n-icon>
                      </template>
                    </n-button>
                  </n-form-item>
                </n-form>

                <div style="text-align: right; margin-top: 40px">
                  <n-button
                      text
                      type="primary"
                  >
                    {{ t('adminViews.login.card.forgetPassword') }}
                  </n-button>
                </div>
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
  //border-radius: 3px 3px 0 0;

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