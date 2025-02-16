<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onMounted, onBeforeMount, ref} from "vue";
import {useRouter} from "vue-router";
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import type {FormInst, FormItemRule, FormRules, NotificationType} from 'naive-ui'
import {NIcon, type UploadCustomRequestOptions, useMessage, useNotification} from 'naive-ui'
import {
  Add as createIcon,
  ChevronForwardOutline as toRightIcon,
  InformationCircle as infoIcon,
  Wallet as walletIcon,
} from "@vicons/ionicons5"
// import {encodeToBase64, hashPassword} from "@/utils/encryptor";
import {
  handleAlterMyName,
  handleCancelSetup2FA,
  handleDeleteMyAccount,
  handleDisable2FA,
  handleGet2FAStatus,
  handleGetUserAvatar,
  handleSetup2FA,
  handleTest2FA,
  handleUploadAvatar,
  saveNewPassword,
  verifyOldPassword
} from "@/api/user/profile";
import {compressAvatar} from "@/utils/compress"
// import instance from "@/axios";

const notification = useNotification()
const themeStore = useThemeStore();
const userInfoStore = useUserInfoStore();
const appInfosStore = useAppInfosStore();

// import {handleTest2FA, handleCancelSetup2FA, handleSetup2FA, handleGet2FAStatus, verifyOldPassword, saveNewPassword} from "@/api/user/profile";

interface ModelType {
  old_password: string
  new_password: string
  new_password_again: string
}

let confirmDeleteAccountEmailInput = ref<string>('')
let loadingDeleteSpin = ref<boolean>(false)
const message = useMessage()
let animated = ref<boolean>(false)
let showAlterAvatarModal = ref<boolean>(false)
let showDeleteMyModal = ref<boolean>(false)
let showAlterNameModal = ref<boolean>(false)
// let uploadSuccess = ref<boolean>(false)
const {t} = useI18n()
const router = useRouter()
const updatePwdForm = ref<FormInst | null>(null)
// const message = useMessage()
const newName = ref<string>('')
const modelRef = ref<ModelType>({
  old_password: '',
  new_password: '',
  new_password_again: ''
})

let failReasons = ref<string[]>([]);

let newNameIsValid = ref<boolean>(false);
const newNameRule: FormItemRule = {
  required: true,
  trigger: ['blur', 'input'],
  validator() {
    let val = newName.value.trim();

    // 检查：
    // 1. 不能包含空格
    // 2. 不能是纯数字
    // 3. 不能以数字开头
    // 4. 其他字符（包括字母、中文、日文等）都允许

    // 正则表达式：
    // 1. ^[^\d\s] 表示开头不能是数字和空格
    // 2. [^\s]+$ 表示整个字符串中不能有空格
    // 3. isNaN(Number(val)) 确保不是纯数字
    const isValid = /^[^\d\s][^\s]*$/.test(val) // 不允许空格和数字开头
        && !/^\d+$/.test(val)  // 确保不是纯数字
        && !/\s/.test(val);  // 不允许包含空格

    if (val && val.length >= 3 && isValid) {
      newNameIsValid.value = true;
      return true;  // 返回true表示验证通过
    } else {
      newNameIsValid.value = false;
      return new Error(t('userProfile.secondary.modify.input.spaceIsNotAllowed'));  // 返回错误
    }
  },
};

const rules: FormRules = {
  old_password: {
    required: true,
    trigger: 'blur',
    validator(rule: FormItemRule, value: string) {
      if (value) return true
      else {
        failReasons.value.push(t('userProfile.resetPassword.previousPasswordRequire'))
        return new Error('Please input previous password first')
      }
    }
  },
  new_password: {
    required: true,
    trigger: 'blur',
    validator(rule: FormItemRule, value: string) {
      let newPassword = value.trim()
      // 檢查密碼是否為空
      if (!newPassword) {
        failReasons.value.push(t('userProfile.resetPassword.passwordRequire')); // 繁體中文
        return new Error('Password cannot be empty');
      }

      // 檢查新密碼是否與舊密碼相同
      if (modelRef.value.old_password === newPassword) {
        failReasons.value.push(t('userProfile.resetPassword.passwordConflict')); // 繁體中文
        return new Error('New password must be different from previous password');
      }

      // Check password length
      if (newPassword.length < 10) {
        failReasons.value.push(t('userProfile.resetPassword.passwordLengthRequire')); // 繁體中文
        return new Error("Password length is less than 12 characters");
      }

      // Check if password contains at least 3 character types
      const hasUpperCase = /[A-Z]/.test(newPassword);  // Uppercase letter
      const hasLowerCase = /[a-z]/.test(newPassword);  // Lowercase letter
      const hasNumber = /\d/.test(newPassword);        // Number
      const hasSpecialChar = /[!@#$%^&*(),.?":{}|<>]/.test(newPassword);  // Special character

      const typeCount = [hasUpperCase, hasLowerCase, hasNumber, hasSpecialChar].filter(Boolean).length;

      // Check if password meets at least 3 types
      if (typeCount < 3) {
        failReasons.value.push(t('userProfile.resetPassword.passwordComplexRequire')); // 繁體中文
        return new Error("Password must contain at least three types of characters: uppercase, lowercase, digits, or special characters");
      }

      return true;
    }
  },
  new_password_again: {
    required: true,
    trigger: 'blur',
    validator: (rule: FormItemRule, value: string) => {
      let newPasswordAgain = value.trim()
      if (!newPasswordAgain) {
        failReasons.value.push(t('userProfile.resetPassword.passwordAgainRequire'))
        return new Error('Please input new password again')
      }
      if (newPasswordAgain !== modelRef.value.new_password) {
        failReasons.value.push(t('userProfile.resetPassword.passwordAgainNotMatch'))
        return new Error('Password not match')
      }
      return true
    }
  },
}

let showFailReasons = () => {
  let _failReasons = failReasons.value
  notification.create({
    title: computed(() => t("userRegister.form.checkForm")).value, // 通知标题
    duration: 10000,
    description: () => {
      // 使用 div 来渲染多个错误信息
      return h('div', [
        _failReasons.map((reason: string, index: number) => {
          return h('p', {key: index}, `${index + 1} ${reason}`); // 渲染每个错误信息
        }),
      ]);
    },
    // meta: new Date().toLocaleString(), // 可以加入时间戳或其他额外信息
    onClose: () => {
      // 关闭通知时的回调
      console.log('Fail reasons notification closed');
    },
  });
};


let notify = (type: NotificationType, title: string, meta?: string) => {
  notification[type]({
    content: title,
    meta: meta,
    duration: 2500,
    keepAliveOnHover: true
  })
}

let saveNewPasswordClick = async (e: MouseEvent) => {
  e.preventDefault()
  failReasons.value = []
  await updatePwdForm.value?.validate((errors) => {
    if (errors) return showFailReasons()
  })
  if (await verifyOldPassword(userInfoStore.thisUser.email, modelRef.value.old_password.trim())) { // 等待验证旧密码的结果
    let data = await saveNewPassword(userInfoStore.thisUser.email, modelRef.value.new_password.trim())
    if (data && data.code === 200 && data.updated as boolean) {
      modelRef.value.old_password = ''
      modelRef.value.new_password = ''
      modelRef.value.new_password_again = ''
      notify('success', t('userProfile.alertSuccess'), t('userProfile.alertSuccessSub'))
    } else {
      message.error(t('userProfile.resetPassword.updatePasswordFailure') + data.msg || '')
    }
  } else {
    message.error(t('userProfile.resetPassword.previousPasswordVerifiedFailure'))
  }


}

let url = ref<string>('')
let showModal = ref<boolean>(false)

// handleSetup2FA 创建一个新的2FA新建请求
let callHandleSetup2FA = async () => {
  // try {
  //   let {data} = await instance.post('/api/user/v1/auth/2fa/setup', {
  //     id: userInfoStore.thisUser.id,
  //     email: userInfoStore.thisUser.email,
  //   })
  let data = await handleSetup2FA(userInfoStore.thisUser.id, userInfoStore.thisUser.email)
  if (data && data.code === 200) {
    console.log(data)
    url.value = data.url
    showModal.value = true
    handleCalTime()
  } else {
    message.error('err: ' + data.msg || '')
  }
  // } catch (error: any) {
  //   console.log(error)
  // }
}

let twoFaCode = ref<string>('')

// handleTest2FA 测试2FA可用性 如果可用则写入进数据库
let callHandleTest2FA = async () => {
  let data = await handleTest2FA(userInfoStore.thisUser.id, userInfoStore.thisUser.email, twoFaCode.value.trim())
  if (data && data.code === 200) {
    console.log('绑定2fa成功', data)
    showModal.value = false
    message.success('绑定两步验证成功')
    await callHandleGet2FAStatus()
  } else {
    message.error(data.msg)
  }
}

// handleCancelSetup2FA 删除redis中的新建临时2FA数据
let callHandleCancelSetup2FA = async () => {
  let data = await handleCancelSetup2FA(userInfoStore.thisUser.email)
  if (data && data.code === 200) {
    message.info(t('userProfile.disable2FaCancelled'))
  } else {
    message.error(t('userProfile.unknownErr') + data.msg || '')
  }
  showModal.value = false
}

let twoFAEnabled = ref<boolean>(false)
// handleGet2FAStatus 测试2FA可用性 如果可用则写入进数据库
let callHandleGet2FAStatus = async () => {
  let data = await handleGet2FAStatus(userInfoStore.thisUser.id)
  if (data && data.code === 200) {
    twoFAEnabled.value = data.enabled as boolean
  } else {
    message.error(data.msg)
  }
}

let leftTime = ref<number>(0)

let intervalId = ref<number | null>(null) // 定时器ID

let handleCalTime = () => {
  leftTime.value = 60
  if (intervalId.value) {
    intervalId.value ? clearInterval(intervalId.value) : null // 清除之前的定时器，防止多次调用重复计时
  }

  intervalId.value = setInterval(() => {
    if (leftTime.value <= 1) {
      showModal.value = false // 关闭二维码modal
      setTimeout(() => {
        intervalId.value ? clearInterval(intervalId.value) : null // 计时结束时清除定时器
        return
      }, 1)
    }
    leftTime.value -= 1
  }, 1000)
}

let callHandleDisable2FA = async () => {
  let data = await handleDisable2FA(userInfoStore.thisUser.email)
  if (data && data.code === 200) {
    message.info(t('userProfile.closed2FaHint'))
    // 再次查询2fa状态
    // await handleGet2FAStatus()
    await callHandleGet2FAStatus()
  } else {
    message.error(data.msg || '')
  }
}

let handleClickToTopUp = () => {
  message.info('To top up page')
  router.push({
    path: '/dashboard/topup'
  })
}

let callHandleDeleteMyAccount = async () => {
  let data = await handleDeleteMyAccount(userInfoStore.thisUser.id, true)
  if (data && data.code === 200 && data.deleted) {
    message.success(computed(() => t('userProfile.deletedSuccessMsg')).value)
    setTimeout(() => {
      userInfoStore.logout()
    }, 1200)
  } else {
    message.error(computed(() => t('userProfile.deleteErrOccur')).value + ' ' + data.msg || '')
  }
}

let isEmailCurrent = computed(() => confirmDeleteAccountEmailInput.value === userInfoStore.thisUser.email)

let confirmedDelete = () => {
  loadingDeleteSpin.value = true
  setTimeout(() => {
    loadingDeleteSpin.value = false
    showDeleteMyModal.value = false
    // handleDeleteMyAccount()
    callHandleDeleteMyAccount()
  }, 1500)
}

let uploadStatus = ref<string>('')
let callCustomUploadAvatar = async ({
                                      file,
                                      onFinish,
                                      onError,
                                      onProgress,
                                    }: UploadCustomRequestOptions) => {
  try {
    onProgress({percent: 0})
    const compressedFile = await compressAvatar(file.file as File, 160);
    onProgress({percent: 30})

    let data = await handleUploadAvatar(userInfoStore.thisUser.id, file.name, compressedFile)
    if (data && data.code === 200) {
      message.success(t('userProfile.secondary.mention.alterAvatarSuccess'))
      onProgress({percent: 95})
      setTimeout(() => {
        onFinish()
        callGetUserAvatarByUserId()
      }, 1000)
    } else {
      onError()
      message.error(t('userProfile.secondary.mention.alterAvatarSuccess') + data.msg || '')
    }
  } catch (error: any) {
    console.error(error);
  }
}

let avatarData = ref<any>()
let callGetUserAvatarByUserId = async () => {
  let data = await handleGetUserAvatar(userInfoStore.thisUser.id)
  if (data.code === 404) return message.warning('no avatar')
  else if (data && data.code === 200) {
    avatarData.value = `data:image/jpeg;base64,${data.file_data}`
    // data.file_data 是头像的 Base64 编码，所以将其拼接成 data:image/jpeg;base64,xxx 的格式，以 Base64 URL 方式存入 avatarData.value。
  } else {
    message.error(t('userProfile.secondary.mention.fetchAvatarErr'))
  }
  showAlterAvatarModal.value = false
}

let getAvatarFilter = computed(() => themeStore.enableDarkMode ? ({filter: 'brightness(60%)'}) : null)

const nameClick = () => {
  if (userInfoStore.thisUser.name) {
    newName.value = userInfoStore.thisUser.name
  }
  showAlterNameModal.value = true
}

const callAlterUsername = async () => {
  if (newNameIsValid.value) {
  } else return message.warning(t('userProfile.secondary.mention.newNameIsNotValid'))
  let data = await handleAlterMyName(userInfoStore.thisUser.id, newName.value)
  if (data && data.code === 200) {
    message.success(t('userProfile.secondary.mention.alterNameSuccess'))
    let isUpdated = await userInfoStore.updateUserInfo()
    console.log(isUpdated)
    !isUpdated ? notify('error', t('userProfile.failure'), t('userProfile.notLatestHint')) : null
  } else {
    message.error(t('userProfile.secondary.mention.alterNameErr'))
  }
}

onBeforeMount(() => {
  themeStore.breadcrumb = t('userProfile.title')
  themeStore.menuSelected = 'user-profile'


})

onMounted(async () => {
  themeStore.userPath = '/dashboard/profile'

  //
  let isUpdated = await userInfoStore.updateUserInfo()
  console.log(isUpdated)
  !isUpdated ? notify('error', t('userProfile.failure'), t('userProfile.notLatestHint')) : null
  // await handleGet2FAStatus()
  await callHandleGet2FAStatus()

  animated.value = true

  await callGetUserAvatarByUserId()

})

</script>

<script lang="ts">
export default {
  name: 'UserProfile',
}
</script>

<template>
  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
          content-style="padding: 0"
          class="secondary-card"
      >
        <div class="secondary-card-inner">
          <n-p class="title">{{ t('userProfile.secondary.title') }}</n-p>

          <div class="secondary-show">
            <div class="secondary-l">
              <n-avatar
                  v-if="avatarData"
                  :src="avatarData"
                  :size="60"
                  @click="showAlterAvatarModal=true"
                  :style="getAvatarFilter"
              />
              <div
                  v-else
                  class="no-avatar-div"
                  @click="showAlterAvatarModal=true"
              >
                <n-icon size="30" style="opacity: 0.5">
                  <createIcon/>
                </n-icon>
                <p class="drag-icon-hint">{{ t('userProfile.secondary.modify.uploadIconHint') }}</p>
              </div>
            </div>
            <div class="secondary-r">
              <div class="secondary-r-up">
                <n-button
                    text
                    type="primary"
                    class="secondary-name"
                    @click="nameClick"
                >
                  {{ userInfoStore.thisUser.name || t('userProfile.secondary.mention.click2SetName') }}
                </n-button>
              </div>
              <div class="secondary-r-below">
                <p class="secondary-email">{{ userInfoStore.thisUser.email }}</p>
              </div>
            </div>
          </div>

          <p
              style="font-size: 0.7rem; opacity: 0.3"
          >{{ t('userProfile.secondary.modify.alterShallow') }}</p>

        </div>

      </n-card>

      <n-card
          class="wallet-card"
          :embedded="true"
          hoverable
          content-style="padding: 0"
          :bordered="false"
      >
        <div class="wallet-header">
          <p class="wallet-title">{{ t('userProfile.myWallet') }}</p>
          <n-icon size="30" style="opacity: 0.1">
            <walletIcon/>
          </n-icon>
        </div>
        <div class="wallet-content">
          <p class="balance">{{ userInfoStore.thisUser.balance.toFixed(2) }}</p>
          <p class="unit">{{ appInfosStore.appCommonConfig.currency }}</p>
        </div>
        <div class="wallet-bottom">
          <p class="sub">{{ t('userProfile.walletSub') }}</p>
          <n-button type="primary" quaternary @click="handleClickToTopUp">
            {{ t('userProfile.toTopUp') }}
            <n-icon size="18" style="margin-left: 6px">
              <toRightIcon/>
            </n-icon>
          </n-button>
        </div>
      </n-card>

      <n-card
          class="password-alert-card"
          :embedded="true"
          hoverable
          content-style="padding: 0"
          :bordered="false"
      >
        <n-p class="title">{{ t('userProfile.alertPwd') }}</n-p>
        <div class="form">
          <n-form
              ref="updatePwdForm"
              :model="modelRef"
              :rules="rules"
              :style="themeStore.menuCollapsed?({width: '100%'}):({width: '60%'})"
              :show-feedback="false"
              :show-require-mark="true"
          >
            <n-form-item path="old_password" :label="t('userProfile.oldPwd')" class="form-item">
              <n-input
                  type="password"
                  showPasswordOn="click"
                  v-model:value="modelRef.old_password"
                  @keydown.enter.prevent
                  :placeholder="t('userProfile.oldPwdSub')"/>
            </n-form-item>
            <n-form-item path="new_password" :label="t('userProfile.newPwd')" class="form-item">
              <n-input
                  type="password"
                  showPasswordOn="click"
                  v-model:value="modelRef.new_password"
                  @keydown.enter.prevent
                  :placeholder="t('userProfile.newPwdSub')"/>
            </n-form-item>
            <n-form-item path="new_password_again" :label="t('userProfile.newPwdAgain')" class="form-item">
              <n-input
                  type="password"
                  showPasswordOn="click"
                  v-model:value="modelRef.new_password_again"
                  @keydown.enter.prevent
                  :placeholder="t('userProfile.newPwdAgainSub')"/>
            </n-form-item>
            <n-form-item>
              <n-button
                  class="alert-btn"
                  type="primary"
                  secondary
                  @click="saveNewPasswordClick"
                  :bordered="false">
                {{ t('userProfile.saveBtn') }}
              </n-button>
            </n-form-item>
          </n-form>

        </div>
      </n-card>

      <n-card
          class="notify-card"
          :embedded="true"
          hoverable
          content-style="padding: 0"
          :bordered="false"
          v-if="false"
      >
        <n-p class="title">{{ t('userProfile.notify') }}</n-p>
        <div class="form">
          <n-form ref="formRef" :rules="rules" :style="themeStore.menuCollapsed?({width: '100%'}):({width: '60%'})">
            <n-form-item path="old_password" :label="t('userProfile.enableNotify')">
              <n-switch/>
            </n-form-item>
          </n-form>
        </div>
      </n-card>

      <n-card
          class="two-fa-card"
          :embedded="true"
          hoverable
          content-style="padding: 0"
          :bordered="false"
      >
        <n-p class="title">{{ t('userProfile.faAuth') }}</n-p>
        <n-alert :bordered="false" style="margin: 20px" type="info">
          {{ t('userProfile.faAuthHint') }}
        </n-alert>
        <div class="form">
          <div style="display: flex; flex-direction: row; margin-bottom: 20px; align-items: center">
            <p style="font-size: 0.9rem; font-weight: 500; margin-right: 10px; opacity: 0.9;">
              {{ t('userProfile.faAuthStatus') }}</p>
            <div
                :style="twoFAEnabled?{backgroundColor: '#86c166'}:{backgroundColor: themeStore.enableDarkMode?'rgba(225, 225, 225, 0.8)':'rgba(25,25,25,0.5)'}"
                style="width: 6px; height: 6px; border-radius: 50%; margin-right: 5px"
            ></div>
            <p v-if="twoFAEnabled" style="color: #86c166">{{ t('userProfile.faEnabled') }}</p>
            <p v-else :style="{color: themeStore.enableDarkMode?'rgba(225, 225, 225, 0.8)':'rgba(25,25,25,0.5)'}">
              {{ t('userProfile.faNotEnabled') }}</p>
          </div>
          <n-button
              secondary
              type="primary"
              :disabled="twoFAEnabled"
              @click="callHandleSetup2FA"
              :bordered="false"
              style="margin-bottom: 20px"
          >
            {{ t('userProfile.setup2Fa') }}
          </n-button>
          <n-button
              secondary
              :disabled="!twoFAEnabled"
              @click="callHandleDisable2FA"
              :bordered="false"
              style="margin-bottom: 20px; margin-left: 10px" type="warning"
          >
            {{ t('userProfile.disable2Fa') }}
          </n-button>
        </div>
      </n-card>

      <n-card
          class="del-card"
          :embedded="true"
          hoverable
          content-style="padding: 0"
          :bordered="false"
      >
        <n-p class="title">{{ t('userProfile.deleteAccount') }}</n-p>
        <div class="form">
          <n-alert type="warning" :bordered="false">
            {{ t('userProfile.deleteAccountSub') }}
          </n-alert>
          <n-button
              strong
              secondary
              type="error"
              style="margin-top: 20px;"
              @click="showDeleteMyModal=true"
          >
            {{ t('userProfile.deleteBtn') }}
          </n-button>
        </div>
      </n-card>
    </div>
  </transition>

  <n-modal
      v-model:show="showDeleteMyModal"
      class="delete-my-modal-card"
      style="width: 450px"
      preset="card"
      type="warning"
      :bordered="false"
      content-style="padding: 0"
  >
    <template #header>
      <div style="display: flex; flex-direction: row; align-items: center;">
        <n-icon style="margin-right: 5px" size="23">
          <infoIcon/>
        </n-icon>
        {{ t('userProfile.deleteMyAccountModal.title') }}
      </div>
    </template>

    <div class="delete-my-modal-header">
      <div class="delete-my-modal-header-inner">
        <n-ul>
          <n-li>{{ t('userProfile.deleteMyAccountModal.contentLine1') }}</n-li>
          <n-li>{{ t('userProfile.deleteMyAccountModal.contentLine2') }}</n-li>
          <n-li>{{ t('userProfile.deleteMyAccountModal.contentLine3') }}</n-li>
        </n-ul>
      </div>
    </div>
    <div class="delete-my-modal-hr"></div>
    <n-spin :show="loadingDeleteSpin" size="small">
      <div class="delete-my-modal-body">
        <div class="delete-my-modal-body-title">
          <p style="font-size: 1rem; font-weight: 600; margin-right: 4px">*</p>
          <p class="delete-my-modal-body-title-hex1">{{ t('userProfile.deleteMyAccountModal.inputHint1') }}</p>
          <p class="delete-my-modal-body-title-hex">{{ userInfoStore.thisUser.email }}</p>
          <p class="delete-my-modal-body-title-hex2">{{ t('userProfile.deleteMyAccountModal.inputHint2') }}</p>
        </div>

        <n-input

            v-model:value="confirmDeleteAccountEmailInput"
            :placeholder="''"
            size="medium"
        />
        <n-button
            type="error"
            secondary
            style="margin-top: 10px; width: 100%"
            :disabled="!isEmailCurrent"
            @click="confirmedDelete"
        >
          {{ t('userProfile.deleteMyAccountModal.confirmDelete') }}
        </n-button>


      </div>
    </n-spin>
    <template #footer>

    </template>
  </n-modal>


  <n-modal
      v-model:show="showModal"
      auto-focus
      @after-hide=""
  >
    <n-card
        class="qr-modal-root"
        :bordered="false"
        size="huge"
    >
      <template #header-extra></template>
      <!--      <n-h4 style="opacity: 0.8">按照以下步骤以设置验证器</n-h4>-->

      <p style="font-weight: 700; font-size: 1.25rem">STEP1</p>
      <p style="opacity: 0.8; margin-bottom: 20px">{{ t('userProfile.setup2FaModal.followStep') }}</p>
      <div class="qr-modal-body">
        <div class="l-part">
          <n-qr-code
              :size="180"
              :value="url"
              type="svg"
          />
        </div>
        <div class="r-part">
          <h3>{{ t('userProfile.setup2FaModal.step1Title') }}</h3>
          <n-p>{{ t('userProfile.setup2FaModal.step1Context1') }}</n-p>
          <n-p>{{ t('userProfile.setup2FaModal.step1Context2') }}</n-p>
          <n-p>{{ t('userProfile.setup2FaModal.step1Context3') }}</n-p>
        </div>

      </div>
      <n-hr></n-hr>
      <p style="font-weight: 700; font-size: 1.25rem">STEP2</p>
      <p style="opacity: 0.8; margin-bottom: 20px">{{ t('userProfile.setup2FaModal.step2Context1') }}</p>
      <div class="code-test-body">
        <n-input
            class="code-inp"
            :placeholder="leftTime + 's'"
            maxlength="6"
            show-count
            size="large"
            v-model:value="twoFaCode"
        >
        </n-input>
        <n-button
            @click="callHandleTest2FA"
            size="large"
            class="test-btn"
            type="primary"
            :bordered="false"
        >
          {{ t('userProfile.setup2FaModal.test2Fa') }}
        </n-button>
        <n-button
            @click="callHandleCancelSetup2FA()"
            size="large"
            class="cancel-btn"
            type="primary"
            secondary
            :bordered="false"
        >
          {{ t('userProfile.setup2FaModal.cancel') }}
        </n-button>
      </div>
    </n-card>
  </n-modal>

  <n-modal
      v-model:show="showAlterAvatarModal"
      auto-focus
      @after-hide=""
  >
    <n-card
        class="alter-avatar-modal"
        :bordered="false"
        size="medium"
        :title="t('userProfile.secondary.modify.alterAvatar')"
    >
      <div class="alter-avatar-inner">
        <p class="alter-avatar-title">{{ t('userProfile.secondary.mention.uploadImageHint') }}</p>
        <n-upload
            :max="1"
            :custom-request="callCustomUploadAvatar"
            accept="image/*"
            :status="uploadStatus"
            class="upload-board"
        >
          <n-upload-dragger>
            <n-text style="font-size: 0.9rem">
              {{ t('userProfile.secondary.mention.click2Upload') }}
            </n-text>
            <n-p depth="3" style="margin: 8px 0 0 0; font-size: 0.7rem">
              {{ t('userProfile.secondary.mention.uploadWarn') }}
            </n-p>
          </n-upload-dragger>
        </n-upload>

        <p class="alter-avatar-hint-upper">{{ t('userProfile.secondary.mention.imageRequire.title') }}</p>
        <p
            class="alter-avatar-hint"
            v-for="(item, index) in ['p1', 'p2', 'p3']"
            :key="item"
        >{{ `${index + 1}. ` + t(`userProfile.secondary.mention.imageRequire.${item}`) }}</p>
        <!--       <p>{{ t('userProfile.secondary.mention.imageRequire.p2') }}</p>-->
        <!--       <p>{{ t('userProfile.secondary.mention.imageRequire.p3') }}</p>-->
      </div>
    </n-card>
  </n-modal>


  <n-modal
      :title="t('userProfile.secondary.modify.alterName')"
      v-model:show="showAlterNameModal"
      preset="dialog"
      :positive-text="t('modal.confirm')"
      :negative-text="t('modal.cancel')"
      style="width: 400px;"
      @positive-click="callAlterUsername"
      @negative-click="message.warning('negative')"
      :show-icon="false"
  >

    <div style="margin-top: 30px">
      <n-form-item
          :label="t('userProfile.secondary.modify.input.label')"
          label-placement="top"
          :rule="newNameRule"
      >
        <n-input
            v-model:value="newName"
            :placeholder="t('userProfile.secondary.modify.input.placeholder')"
            show-count
            clearable
        />
      </n-form-item>

      <n-li
          style="opacity: 0.7; margin-top: 4px"
          v-for="i in 3"
          :key="i"
      >{{ t(`userProfile.secondary.modify.input.require.p${i}`) }}
      </n-li>
    </div>

  </n-modal>

</template>

<style scoped lang="less">
.root {
  margin: 20px;

  .secondary-card {
    margin-bottom: 14px;

    .secondary-card-inner {
      display: flex;
      flex-direction: column;
      justify-content: flex-start;
      background-color: rgba(216, 216, 216, 0.0);
      padding: 10px 20px 10px 20px;

      .title {
        font-size: 1rem;
        font-weight: normal;
        opacity: 0.6;
      }

      .secondary-show {
        display: flex;
        flex-direction: row;
        background-color: rgba(216, 216, 216, 0.0);
        margin: 10px 0;
        justify-content: flex-start;
        align-items: center;

        .secondary-l {
          .no-avatar-div {
            width: 60px;
            height: 60px;
            border-radius: 3px;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            /* border: rgba(51,51,51,0.3) 1px solid; */

            .drag-icon-hint {
              font-size: 0.6rem;
              margin-top: 2px;
              opacity: 0.5;
            }
          }
        }

        .secondary-r {
          height: 100%;
          display: flex;
          flex-direction: column;
          justify-content: flex-start;
          margin-left: 20px;

          .secondary-r-up {
            align-items: center;

            .secondary-name {
              font-size: 1.1rem;
              font-weight: bold;
              opacity: 0.8;
            }
          }

          .secondary-r-below {
            .secondary-email {
              font-size: 1rem;
              opacity: 0.6;
            }
          }
        }
      }
    }
  }

  .wallet-card {
    .wallet-header {
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      background-color: rgba(216, 216, 216, 0.0);
      padding: 10px 20px 10px 20px;

      .wallet-title {
        font-size: 1rem;
        font-weight: normal;
        opacity: 0.6;
      }
    }

    .wallet-content {
      display: flex;
      flex-direction: row;
      align-items: baseline;
      padding: 0 20px 0 20px;

      .balance {
        font-size: 3.25rem;
        font-weight: 200;
        font-family: Inter, -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Helvetica Neue, Arial, Noto Sans, Liberation Sans, sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol, Noto Color Emoji;

      }

      .unit {
        font-size: 1rem;
        margin-left: 15px;
        font-weight: 300;
        font-family: Inter, -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Helvetica Neue, Arial, Noto Sans, Liberation Sans, sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol, Noto Color Emoji;
        opacity: 0.6;
      }
    }

    .wallet-bottom {
      margin: 0 20px 20px 20px;
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: center;

      .sub {
        opacity: 0.6;
      }
    }

    margin-bottom: 15px;
  }

  .password-alert-card {
    .title {
      font-size: 1.1rem;
      font-weight: 400;
      background-color: rgba(216, 216, 216, 0.1);
      padding: 10px 20px 10px 20px;
      border-radius: 3px 3px 0 0;
    }

    .form {
      margin: 20px;

      .form-item {
        margin-top: 20px;
      }
    }

    .alert-btn {
      width: auto;
    }

    //@media screen and (min-width: 769px) {
    //  .form {
    //    width: 50%;
    //  }
    //}
    margin-bottom: 15px;

  }

  .notify-card {
    .title {
      font-size: 1.1rem;
      font-weight: 400;
      background-color: rgba(216, 216, 216, 0.1);
      padding: 10px 20px 10px 20px;
      border-radius: 3px 3px 0 0;
    }

    .form {
      margin: 20px 20px 0 20px;
    }

    margin-bottom: 20px;
  }

  .two-fa-card {
    .title {
      font-size: 1.1rem;
      font-weight: 400;
      background-color: rgba(216, 216, 216, 0.1);
      padding: 10px 20px 10px 20px;
      border-radius: 3px 3px 0 0;
    }

    .form {
      margin: 20px 20px 0 20px;

    }

    margin-bottom: 15px;
  }

  .del-card {
    .title {
      font-size: 1.1rem;
      font-weight: 400;
      background-color: rgba(216, 216, 216, 0.1);
      padding: 10px 20px 10px 20px;
      border-radius: 3px 3px 0 0;
    }

    .form {
      margin: 20px;
    }

    margin-bottom: 15px;
  }


}

//.n-card {
//  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
//  border: 0;
//}

.qr-modal-root {
  width: 720px;

  .qr-modal-body {
    display: flex;
    flex-direction: row;
    //background-color: #4cae4c;

    .l-part {
      flex: 2;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      //background-color: #66afe9;
    }

    .r-part {
      flex: 4;
      margin-left: 40px;
      align-items: center;
    }
  }

  @media (max-width: 1000px) {
    .qr-modal-body {
      display: flex;
      width: 100%;
      flex-direction: column;
      justify-content: center;

      .l-part {
        width: 100%;
        //background-color: #4cae4c;
      }

      .r-part {
        flex: 1;
        margin: 30px 0 20px 0;
        //background-color: #66afe9;
      }
    }

  }
}

@media (max-width: 1000px) {
  .qr-modal-root {
    width: 60%;
  }
}

@media (max-width: 768px) {
  .qr-modal-root {
    width: 90%;
  }
}

.code-test-body {
  width: 80%;
  display: flex;
  flex-direction: row;
  margin-bottom: 10px;

  .code-inp {
    flex: 4;
    margin-right: 20px;
  }

  .test-btn {
    flex: 1;
    margin-right: 20px;
  }

  .cancel-btn {
    flex: 1;
  }
}

@media (max-width: 1000px) {
  .code-test-body {
    width: 100%;
    display: flex;
    flex-direction: row;
    margin-bottom: 10px;

    .code-inp {
      flex: 4;
      margin-right: 20px;
    }

    .test-btn {
      flex: 1;
      margin-right: 20px;
    }

    .cancel-btn {
      flex: 1;
    }
  }
}

@media (max-width: 768px) {
  .code-test-body {
    flex-direction: column;
    margin-bottom: 10px;
    height: auto;

    .code-inp {
      flex: auto;
    }

    .test-btn {
      flex: auto;
      margin-right: 0;
      margin-top: 10px;
    }

    .cancel-btn {
      flex: auto;
      margin-top: 10px;
    }
  }
}


.delete-my-modal-card {

  .delete-my-modal-header {
    padding: 0 20px;
    display: flex; // 确保整体居中
    align-items: center;

    .delete-my-modal-header-inner {
      display: flex;
      flex-direction: column;
      font-size: 0.9rem;
      padding-right: 5px;
    }

  }

  .delete-my-modal-hr {
    background-color: #252525;
    height: 1px;
    opacity: 0.1;
    margin: 20px;
  }

  .delete-my-modal-body {
    padding: 0 20px 30px 20px;

    .delete-my-modal-body-title {
      display: flex;
      flex-direction: row;
      align-items: center;
      align-content: center;
      line-height: 1;
      font-size: 0.9rem;
      margin: 20px 0 10px 0;

      .delete-my-modal-body-title-hex1 {
        margin-right: 1px;
      }

      .delete-my-modal-body-title-hex {
        opacity: 0.8;
        font-style: italic;
      }

      .delete-my-modal-body-title-hex2 {
        margin-left: 1px;
      }
    }
  }
}

.alter-avatar-modal {
  width: 400px;

  .alter-avatar-inner {
    padding-bottom: 30px;

    .alter-avatar-title {
      margin-bottom: 20px;
    }

    .upload-board {
      margin-bottom: 20px;
    }

    .alter-avatar-hint-upper {
      font-weight: bold;
      font-size: 1.1rem;
      margin: 20px 0 10px 0;
      opacity: 0.8;
    }

    .alter-avatar-hint {
      opacity: 0.6;
      margin-top: 4px;
    }
  }
}

</style>