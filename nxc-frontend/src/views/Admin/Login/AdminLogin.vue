<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from 'vue'
import {useRouter} from 'vue-router'
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
import {
  ChevronBackOutline as BackIcon,
  ChevronBackOutline as backIcon,
  LogInOutline as loginIcon
} from "@vicons/ionicons5"
import {handleAdminLoginFunc} from "@/api/admin/auth";

const {t} = useI18n()
const notification = useNotification()
const appInfoStore = useAppInfosStore()
const themeStore = useThemeStore()
const message = useMessage()

let animated = ref<boolean>(false)

let notifyErr = (type: NotificationType, msg: string) => {
  notification[type]({
    content: computed(() => t('adminViews.login.message.authFailure')).value,
    meta: msg,
    duration: 2500,
    keepAliveOnHover: true
  })
}

let notifyPass = (type: NotificationType) => {
  notification[type]({
    content: computed(() => t('adminViews.login.message.authPassed')).value,
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
// const route = useRoute();
// let sourcePath = route.params.path as string || ''
let thisSecurePath = ref<string | null>(null)
let rememberSecurePath = ref<boolean>(false)
let showInputSecurePathModal = ref<boolean>(true)
let checkSecureBtnDisabled = computed(() => !thisSecurePath.value)
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
  last_login: string;
  last_login_ip: string
}

let bindUserInfo = (data: DataWithAuth) => {
  userInfoStore.thisUser.token = data.token
  userInfoStore.isAuthed = data.isAuthed
  let {user_data} = data
  userInfoStore.thisUser.id = user_data.id
  userInfoStore.thisUser.inviteUserId = user_data.invite_user_id
  userInfoStore.thisUser.name = user_data.name
  userInfoStore.thisUser.email = user_data.email
  userInfoStore.thisUser.isAdmin = user_data.is_admin
  userInfoStore.thisUser.isStaff = user_data.is_staff
  userInfoStore.thisUser.balance = user_data.balance
  userInfoStore.thisUser.lastLogin = user_data.last_login
  userInfoStore.thisUser.lastLoginIp = user_data.last_login_ip
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

const submitLogin = async () => {
  enableLogin.value = false
  let data = await handleAdminLoginFunc(userFormData.value.username, userFormData.value.password, 'admin')
  // console.log(data)
  if (data.code === 200 && data.isAuthed === true) {
    if (data.user_data.is_admin || data.user_data.is_staff) {
      // 验证通过 保存token
      if (!data.token) return message.error(t('userLogin.tokenNotExist'))
      sessionStorage.setItem('token', data.token)
      // 保存验证状态
      userInfoStore.setAndSaveAuthStatus(true)
      notifyPass('success');
      bindUserInfo(data)
      console.log(userInfoStore.thisUser)
      setTimeout(async () => await router.push({path: '/admin/dashboard/summary'}), 1000)
    } else {
      message.error(t('adminViews.login.message.noPrivilege'))
    }
  } else {
    console.error('failure admin login')
    enableLogin.value = true
    switch (data.code) {
      case 401:
        return notifyErr('error', t('adminViews.login.message.passwordErr'))
      case 404:
        return notifyErr('error', t('adminViews.login.message.adminNotExist'))
      default:
        return notifyErr('error', t('adminViews.login.message.otherErr.'))
    }
  }
  // } catch (error) {
  //   console.log(error)
  // }
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
    duration: 3000,
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

let checkSecurePath = () => {
  if (
      thisSecurePath.value
      && thisSecurePath.value === appInfoStore.appCommonConfig.secure_path
  ) {
    showInputSecurePathModal.value = false
    message.success(t('adminViews.login.message.pathCheckPassed'))
    sessionStorage.setItem('secure_path', JSON.stringify(thisSecurePath.value))
  } else {
    message.error(t('adminViews.login.message.pathCheckFailure'))
  }
}

let changeRememberSecurePath = () => {
  console.log('changeRememberSecurePath')
  if (rememberSecurePath.value) {
    message.info(t('adminViews.login.message.rememberSecureMention'))
    localStorage.setItem('secure_path', JSON.stringify(thisSecurePath.value))
  } else {
    localStorage.removeItem('secure_path')
  }
}

onBeforeMount(() => {
  let path: string | null = null;
  try {
    // 优先从 sessionStorage 获取 secure_path，否则从 localStorage 获取
    let rawPath = sessionStorage.getItem('secure_path') || localStorage.getItem('secure_path');
    path = rawPath ? JSON.parse(rawPath) : null;
  } catch (e: any) {
    console.error('Error parsing secure_path:', e);
    path = null;
  }

  if (path && path === appInfoStore.appCommonConfig.secure_path) {
    thisSecurePath.value = path
    showInputSecurePathModal.value = false;
  }
  if (localStorage.getItem('secure_path')) rememberSecurePath.value = true
});
onMounted(() => {
  if (JSON.parse(sessionStorage.getItem('isAuthed') as string)) {
    console.log('to dashboard')
    setTimeout(() => router.push({path: '/admin/dashboard/summary'}), 1000)
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
                    size="medium"
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
                  <n-form-item path="password" style="margin-top: 20px"
                               :label="t('adminViews.login.card.passwordInputArea.title')">
                    <n-input
                        v-model:value="userFormData.password"
                        type="password"
                        showPasswordOn="click"
                        :placeholder="t('adminViews.login.card.passwordInputArea.placeholder')"
                        :bordered="true"
                    />
                  </n-form-item>
                  <n-form-item>
                    <n-checkbox v-model:checked="rememberSecurePath" @change="changeRememberSecurePath">
                      {{ t('adminViews.login.secureCard.rememberPath') }}
                    </n-checkbox>
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

      <n-modal
          v-model:show="showInputSecurePathModal"
          preset="card"
          :title="t('adminViews.login.secureCard.title')"
          size="medium"
          style="width: 400px"
          :bordered="false"
          :closable="false"
          :mask-closable="false"
      >
        <n-p style="margin-top: 10px">
          {{ t('adminViews.login.secureCard.hint') }}
        </n-p>
        <n-form-item
            :label="t('adminViews.login.secureCard.securePath')"
            :show-require-mark="true"
            style="margin-top: 20px"
            :show-feedback="false"
            :rule="{trigger: ['blur', 'input'],validator:() => thisSecurePath!==''} as FormRules"
        >
          <n-input
              :placeholder="t('adminViews.login.secureCard.placeholder')"
              v-model:value="thisSecurePath"
          />
        </n-form-item>

        <n-form-item
            :show-require-mark="true"
        >
          <n-button
              type="primary"
              secondary
              :bordered="false"
              style="width: 100%"
              :disabled="checkSecureBtnDisabled"
              @click="checkSecurePath"
          >
            {{ t('modal.confirm') }}
          </n-button>
        </n-form-item>

      </n-modal>

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