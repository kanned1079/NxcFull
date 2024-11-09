<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onMounted, ref} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import type {FormInst, FormRules, NotificationType} from 'naive-ui'
import {NIcon, useMessage, useNotification} from 'naive-ui'
import {Wallet as walletIcon} from "@vicons/ionicons5"
import {encodeToBase64, hashPassword} from "@/utils/encryptor";
import instance from "@/axios";

interface ModelType {
  old_password: string
  new_password: string
  new_password_again: string
}

const message = useMessage()

let animated = ref<boolean>(false)

const {t} = useI18n()
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
const apiAddrStore = useApiAddrStore()
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

// let handleClose = () => {
//   // handleCancelSetup2FA()
//   showModal.value = false
// }

let leftTime = ref<number>(0)

let intervalId // 定时器ID

let handleCalTime = () => {
  leftTime.value = 60
  if (intervalId) {
    clearInterval(intervalId) // 清除之前的定时器，防止多次调用重复计时
  }

  intervalId = setInterval(() => {
    if (leftTime.value <= 1) {
      showModal.value = false // 关闭二维码modal
      setTimeout(() => {
        clearInterval(intervalId) // 计时结束时清除定时器
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
          <p class="balance">{{ userInfoStore.thisUser.balance.toFixed(2)}}</p>
          <p class="unit">{{ appInfosStore.appCommonConfig.currency }}</p>
        </div>
        <div class="wallet-bottom">
          <p class="sub">{{ t('userProfile.walletSub') }}</p>
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
              <n-input v-model:value="modelRef.old_password" @keydown.enter.prevent
                       :placeholder="t('userProfile.oldPwdSub')"/>
            </n-form-item>
            <n-form-item path="new_password" :label="t('userProfile.newPwd')">
              <n-input v-model:value="modelRef.new_password" @keydown.enter.prevent
                       :placeholder="t('userProfile.newPwdSub')"/>
            </n-form-item>
            <n-form-item path="new_password_again" :label="t('userProfile.newPwdAgain')">
              <n-input v-model:value="modelRef.new_password_again" @keydown.enter.prevent
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
        <n-p class="title">两步验证2FA</n-p>
        <n-alert :bordered="false" style="margin: 20px" type="info">
          雙重驗證（縮寫為
          2FA）是一個驗證過程，要求透過兩個不同的驗證因素來確立身分。簡而言之，這意味著使用者必須先以兩種不同的方式證明其身分，然後才會被授予存取權限。2FA
          是多重要素驗證的一種形式。
        </n-alert>
        <div class="form">
          <div style="display: flex; flex-direction: row; margin-bottom: 20px; align-items: center">
            <p style="font-size: 0.9rem; font-weight: 500; margin-right: 10px; opacity: 0.9;">两步验证状态</p>
            <n-tag v-if="twoFAEnabled" type="success"> 已启用</n-tag>
            <n-tag v-else type="default"> 未启用</n-tag>
          </div>
          <n-button :disabled="twoFAEnabled" @click="handleSetup2FA" :bordered="false" style="margin-bottom: 20px"
                    type="primary" secondary>
            设置两步验证
          </n-button>
          <n-button :disabled="!twoFAEnabled" @click="handleDisable2FA" :bordered="false"
                    style="margin-bottom: 20px; margin-left: 10px" type="warning" secondary>
            关闭两步验证
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
          <n-button strong style="margin-top: 20px;" type="error" secondary>{{ t('userProfile.deleteBtn') }}</n-button>
        </div>
      </n-card>
    </div>
  </transition>


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
      <p style="opacity: 0.8; margin-bottom: 20px">根据提示在您的验证器上加入</p>
      <div class="qr-modal-body">
        <div class="l-part">
          <n-qr-code
              size="180"
              :value="url"
              type="svg"
          />
        </div>
        <div class="r-part">
          <h3>按照以下步骤以启用2FA</h3>
          <n-p>1. 您的移动设备上需要有一个通用的验证器</n-p>
          <n-p>2. 点击右上角的Scan按钮来扫描左侧的QR Code</n-p>
          <n-p>3. 该QR Code中包含有您的验证信息和唯一密钥，请妥善保存</n-p>
        </div>

      </div>
      <n-hr></n-hr>
      <p style="font-weight: 700; font-size: 1.25rem">STEP2</p>
      <p style="opacity: 0.8; margin-bottom: 20px">为了确保您的验证器能够正常使用 我们需要进行测试</p>
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
          测试
        </n-button>
        <n-button
            @click="handleCancelSetup2FA()"
            size="large"
            class="cancel-btn"
            type="primary"
            secondary
            :bordered="false"
        >
          取消
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


</style>