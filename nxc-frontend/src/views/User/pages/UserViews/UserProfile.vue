<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onMounted, ref, computed} from "vue";
import {useRouter} from "vue-router";
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import type {FormInst, FormRules, NotificationType} from 'naive-ui'
import {NIcon, useMessage, useNotification} from 'naive-ui'
import {
  Wallet as walletIcon,
  ChevronForwardOutline as toRightIcon,
  InformationCircle as infoIcon,
} from "@vicons/ionicons5"
import {encodeToBase64, hashPassword} from "@/utils/encryptor";
import instance from "@/axios";

interface ModelType {
  old_password: string
  new_password: string
  new_password_again: string
}

let confirmDeleteAccountEmailInput = ref<string>('')
let loadingDeleteSpin = ref<boolean>(false)

const message = useMessage()

let animated = ref<boolean>(false)

let showDeleteMyModal = ref<boolean>(false)

const {t} = useI18n()
const router = useRouter()
const formRef = ref<FormInst | null>(null)
// const message = useMessage()
const modelRef = ref<ModelType>({
  old_password: '',
  new_password: '',
  new_password_again: ''
})

const rules: FormRules = {
  old_password: {
    required: false,
    trigger: "blur"
  },
  new_password: {
    required: false,
    trigger: "blur"
  },
  new_password_again: {
    required: false,
    trigger: "blur"
  },
}


const notification = useNotification()
const themeStore = useThemeStore();
const userInfoStore = useUserInfoStore();
const appInfosStore = useAppInfosStore();

let notify = (type: NotificationType, title: string, meta?: string) => {
  notification[type]({
    content: title,
    meta: meta,
    duration: 2500,
    keepAliveOnHover: true
  })
}

// verifyOldPassword 验证旧密码是否正确
let verifyOldPassword = async (): Promise<boolean> => {
  try {
    let {data} = await instance.post('http://localhost:8081/api/user/v1/auth/passcode/verify', {
      email: userInfoStore.thisUser.email,
      // old_password: hashPassword(modelRef.value.old_password.trim() as string)
      old_password: encodeToBase64(modelRef.value.old_password.trim() as string)
    });
    if (data.code === 200) {
      return data.verified as boolean;
    } else {
      notify('error', t('userProfile.oldPwdVerifiedFailure'), data.msg)
      console.log(data);
      return false;
    }
  } catch (error: any) {
    notify('error', t('userProfile.oldPwdVerifiedFailure'), error.toString())
    console.log(error.toString());
    return false;
  }
}

// saveNewPassword 保存新的密码
let saveNewPassword = async () => {
  if (modelRef.value.old_password.trim() !== '') {
    if (modelRef.value.new_password.trim() !== '' && modelRef.value.new_password_again.trim() !== '') {
      if (modelRef.value.new_password.trim() === modelRef.value.new_password_again.trim()) {
        console.log('验证ok');
        if (await verifyOldPassword()) { // 等待验证旧密码的结果
          try {
            let hashedNewPassword = await hashPassword(modelRef.value.new_password.trim() as string)
            let {data} = await instance.post('http://localhost:8081/api/user/v1/auth/passcode/update', {
              email: userInfoStore.thisUser.email,
              // new_password: hashPassword(modelRef.value.new_password)
              new_password: hashedNewPassword
            });
            if (data.code === 200 && data.updated as boolean) {
              modelRef.value.old_password = ''
              modelRef.value.new_password = ''
              modelRef.value.new_password_again = ''
              console.log('修改成功', data);
              notify('success', t('userProfile.alertSuccess'), t('userProfile.alertSuccessSub'))

            }
          } catch (error: any) {
            notify('error', t('userProfile.alertFailure'), error.toString())
          }
        } else {
          console.log('旧密码验证失败');
        }
      } else {
        notify('error', t('userProfile.pwdNotMatch'))
        console.log('两次输入密码不一致');
      }
    } else {
      console.log('新密码格式错误');
      notify('error', t('userProfile.errorPwdFormat'))

    }
  } else {
    console.log('旧密码不能为空');
    notify('error', t('userProfile.oldPwdNotNull'))
  }
}

let url = ref<string>('')
let showModal = ref<boolean>(false)

// handleSetup2FA 创建一个新的2FA新建请求
let handleSetup2FA = async () => {
  try {
    let {data} = await instance.post('/api/user/v1/auth/2fa/setup', {
      id: userInfoStore.thisUser.id,
      email: userInfoStore.thisUser.email,
    })
    if (data.code === 200) {
      console.log(data)
      url.value = data.url
      showModal.value = true
      handleCalTime()
    }
  } catch (error: any) {
    console.log(error)
  }
}

let twoFaCode = ref<string>('')

// handleTest2FA 测试2FA可用性 如果可用则写入进数据库
let handleTest2FA = async () => {
  try {
    let {data} = await instance.post('/api/user/v1/auth/2fa/setup/test', {
      id: userInfoStore.thisUser.id,
      email: userInfoStore.thisUser.email,
      two_fa_code: twoFaCode.value.trim(),
    })
    if (data.code === 200) {
      console.log('绑定2fa成功', data)
      showModal.value = false
      message.success('绑定两步验证成功')
    } else {
      message.error(data.msg)
    }
  } catch (error: any) {
    console.log(error)
  }
}

// handleCancelSetup2FA 删除redis中的新建临时2FA数据
let handleCancelSetup2FA = async () => {
  try {
    let {data} = await instance.delete('/api/user/v1/auth/2fa/setup/cancel', {
      params: {
        // id: userInfoStore.thisUser.id,
        email: userInfoStore.thisUser.email,
      }
    })
    if (data.code === 200) {
      console.log('取消2fa成功', data)
      message.info('已取消')
    }
    showModal.value = false
  } catch (error: any) {
    console.log(error)
    message.info('错误' + error)
  }
}

let twoFAEnabled = ref<boolean>(false)
// handleGet2FAStatus 测试2FA可用性 如果可用则写入进数据库
let handleGet2FAStatus = async () => {
  try {
    let {data} = await instance.get('/api/user/v1/auth/2fa/status', {
      params: {
        id: userInfoStore.thisUser.id,
      }
    })
    if (data.code === 200) {
      twoFAEnabled.value = data.enabled as boolean
    } else {
      message.error(data.msg)
    }
  } catch (error: any) {
    message.info('错误' + error)
    console.log(error)
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

let handleDisable2FA = async () => {
  try {
    let {data} = await instance.delete('/api/user/v1/auth/2fa/disable', {
      params: {
        // id: userInfoStore.thisUser.id,
        email: userInfoStore.thisUser.email,
      }
    })
    if (data.code === 200) {
      console.log('关闭2fa成功', data)
      message.info('已关闭2fa验证')
      // 再次查询2fa状态
      await handleGet2FAStatus()
    }
  } catch (error: any) {
    console.log(error)
    message.info('错误' + error)
  }
}

let handleClickToTopUp = () => {
  message.info('To top up page')
  router.push({
    path: '/dashboard/topup'
  })
}

let handleDeleteMyAccount = async () => {
  try {
    let {data} = await instance.delete('/api/user/v1/user/delete', {
      params: {
        user_id: userInfoStore.thisUser.id,
        confirmed: true,
      },
    })
    if (data.code === 200 && data.deleted) {
      message.success(computed(() => t('userProfile.deletedSuccessMsg')).value)
      setTimeout(() => {
        userInfoStore.logout()
      }, 1200)
    } else {
      message.error(computed(() => t('userProfile.deleteErrOccur')).value + ' ' + data.msg || '')
    }
  } catch (err: any) {
    console.log(err + '')
  }
}

let isEmailCurrent = computed(() => confirmDeleteAccountEmailInput.value === userInfoStore.thisUser.email)

let confirmedDelete = () => {
  console.log('确认删除')
  loadingDeleteSpin.value = true
  setTimeout(() => {
    loadingDeleteSpin.value = false
    showDeleteMyModal.value = false
    handleDeleteMyAccount()
  }, 1500)
}

onMounted(async () => {
  console.log('挂载个人中心')
  themeStore.menuSelected = 'user-profile'
  themeStore.userPath = '/dashboard/profile'

  //
  let isUpdated = await userInfoStore.updateUserInfo()
  console.log(isUpdated)
  !isUpdated ? notify('error', '失败', '个人信息更新失败以下当前数据可能不是最新的') : null
  await handleGet2FAStatus()

  animated.value = true

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
          <n-form ref="formRef" :rules="rules" :style="themeStore.menuCollapsed?({width: '100%'}):({width: '60%'})">
            <n-form-item path="old_password" :label="t('userProfile.oldPwd')">
              <n-input type="password" v-model:value="modelRef.old_password" @keydown.enter.prevent
                       :placeholder="t('userProfile.oldPwdSub')"/>
            </n-form-item>
            <n-form-item path="new_password" :label="t('userProfile.newPwd')">
              <n-input type="password" v-model:value="modelRef.new_password" @keydown.enter.prevent
                       :placeholder="t('userProfile.newPwdSub')"/>
            </n-form-item>
            <n-form-item path="new_password_again" :label="t('userProfile.newPwdAgain')">
              <n-input type="password" v-model:value="modelRef.new_password_again" @keydown.enter.prevent
                       :placeholder="t('userProfile.newPwdAgainSub')"/>
            </n-form-item>
          </n-form>
          <n-button class="alert-btn" type="primary" secondary @click="saveNewPassword" :bordered="false">
            {{ t('userProfile.saveBtn') }}
          </n-button>
        </div>
      </n-card>

      <n-card
          class="notify-card"
          :embedded="true"
          hoverable
          content-style="padding: 0"
          :bordered="false"
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
            <p style="font-size: 0.9rem; font-weight: 500; margin-right: 10px; opacity: 0.9;">{{ t('userProfile.faAuthStatus') }}</p>
            <div
                :style="twoFAEnabled?{backgroundColor: '#86c166'}:{backgroundColor: 'rgba(25,25,25,0.5)'}"
                style="width: 6px; height: 6px; border-radius: 50%; margin-right: 5px"
            ></div>
            <p v-if="twoFAEnabled" style="color: #86c166">{{ t('userProfile.faEnabled') }}</p>
            <p v-else style="color: rgba(25,25,25,0.5)">{{ t('userProfile.faNotEnabled') }}</p>
          </div>
          <n-button
              secondary
              type="primary"
              :disabled="twoFAEnabled"
              @click="handleSetup2FA"
              :bordered="false"
              style="margin-bottom: 20px"
          >
            {{ t('userProfile.setup2Fa') }}
          </n-button>
          <n-button
              secondary
              :disabled="!twoFAEnabled"
              @click="handleDisable2FA"
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
        注销账号
      </div>
    </template>

    <div class="delete-my-modal-header">
      <div class="delete-my-modal-header-inner">
        <n-ul>
          <n-li>
            账号注销是一个不可逆的操作。一旦您确认删除，您将永久性地失去该账号的访问权限，这意味着您将无法再登录，且与此账号相关的所有数据，包括但不限于您的个人信息、历史记录、收藏内容、购买记录等都将全部无法访问。
          </n-li>
          <n-li>
            如果您在我们的平台上有正在进行的业务，如未完成的订单、正在参与的活动、订阅服务等，这些都将随着账号删除而终止或取消，可能会给您带来相应的损失。同时，您与其他用户之间通过本平台建立的联系、互动信息等也都将不复存在。
          </n-li>
          <n-li>
            请再次确认您的决定，如果您还有任何疑虑或问题，欢迎联系我们的客服，我们将竭诚为您解答。若您仍然希望删除账号，请点击
            “确认删除” 按钮。
          </n-li>
        </n-ul>
      </div>
    </div>
    <div class="delete-my-modal-hr"></div>
    <n-spin :show="loadingDeleteSpin" size="small">
      <div class="delete-my-modal-body">
        <div class="delete-my-modal-body-title">
          <p style="font-size: 1rem; font-weight: 600; margin-right: 4px">*</p>
          <p class="delete-my-modal-body-title-hex1">{{ '输入 "' }}</p>
          <p class="delete-my-modal-body-title-hex">{{ userInfoStore.thisUser.email }}</p>
          <p class="delete-my-modal-body-title-hex2">{{ '" 以继续。' }}</p>
        </div>

        <n-input

            v-model:value="confirmDeleteAccountEmailInput"
            :placeholder="''"
            size="medium"
        >

          <!--        oncut="return false"-->
          <!--        oncopy="return false"-->
          <!--        onpaste="return false"-->

        </n-input>


        <n-button
            type="error"
            secondary
            style="margin-top: 10px; width: 100%"
            :disabled="!isEmailCurrent"
            @click="confirmedDelete"
        >
          确认删除
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
            @click="handleTest2FA"
            size="large"
            class="test-btn"
            type="primary"
            :bordered="false"
        >
          {{ t('userProfile.setup2FaModal.test2Fa') }}
        </n-button>
        <n-button
            @click="handleCancelSetup2FA()"
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

</template>

<style scoped lang="less">
.root {
  margin: 20px;

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

    margin-bottom: 20px;
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
    }

    .alert-btn {
      width: 100px;
    }

    //@media screen and (min-width: 769px) {
    //  .form {
    //    width: 50%;
    //  }
    //}
    margin-bottom: 20px;

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

    margin-bottom: 20px;
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

    margin-bottom: 20px;
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


</style>