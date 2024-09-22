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
import {hashPassword} from "@/utils/encryptor";
import instance from "@/axios";

interface ModelType {
  old_password: string | null
  new_password: string | null
  new_password_again: string | null
}

const {t} = useI18n()
const formRef = ref<FormInst | null>(null)
const message = useMessage()
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

let notify = (type: NotificationType, title: string, meta: string) => {
  notification[type]({
    content: title,
    meta: meta,
    duration: 2500,
    keepAliveOnHover: true
  })
}

let verifyOldPassword = async (): Promise<boolean> => {
  try {
    let {data} = await instance.post(apiAddrStore.apiAddr.user.verifyOldPassword, {
      email: userInfoStore.thisUser.email,
      old_password: hashPassword(modelRef.value.old_password.trim())
    });
    if (data.code === 200) {
      return data.verified as boolean;
    } else {
      notify('error', t('userProfile.oldPwdVerifiedFailure'), data.msg)
      console.log(data);
      return false;
    }
  } catch (error) {
    notify('error', t('userProfile.oldPwdVerifiedFailure'), error.toString())
    console.log(error.toString());
    return false;
  }
}

let saveNewPassword = async () => {
  if (modelRef.value.old_password.trim() !== '') {
    if (modelRef.value.new_password.trim() !== '' && modelRef.value.new_password_again.trim() !== '') {
      if (modelRef.value.new_password.trim() === modelRef.value.new_password_again.trim()) {
        console.log('验证ok');
        if (await verifyOldPassword()) { // 等待验证旧密码的结果
          try {
            let {data} = await instance.post(apiAddrStore.apiAddr.user.applyNewPassword, {
              email: userInfoStore.thisUser.email,
              new_password: hashPassword(modelRef.value.new_password)
            });
            if (data.code === 200 && data.updated as boolean) {
              modelRef.value.old_password = ''
              modelRef.value.new_password = ''
              modelRef.value.new_password_again = ''
              console.log('修改成功', data);
              notify('success', t('userProfile.alertSuccess'), t('userProfile.alertSuccessSub'))

            }
          } catch (error) {
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


onMounted(() => {
  console.log('挂载个人中心')
  themeStore.menuSelected = 'user-profile'
  themeStore.userPath = '/dashboard/profile'
})

</script>

<script lang="ts">
export default {
  name: 'UserProfile',
}
</script>

<template>
  <div class="root">
    <n-card
        class="wallet-card"
        :embedded="true"
        hoverable
        content-style="padding: 0"
    >
      <div class="wallet-header">
        <p class="wallet-title">{{ t('userProfile.myWallet') }}</p>
        <n-icon size="30" style="opacity: 0.1">
          <walletIcon/>
        </n-icon>
      </div>
      <div class="wallet-content">
        <p class="balance">30.20</p>
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
    >
      <n-p class="title">{{ t('userProfile.alertPwd') }}</n-p>
      <div class="form">
        <n-form ref="formRef" :rules="rules" :style="themeStore.menuCollapsed?({width: '100%'}):({width: '60%'})">
          <n-form-item path="old_password" :label="t('userProfile.oldPwd')">
            <n-input v-model:value="modelRef.old_password" @keydown.enter.prevent :placeholder="t('userProfile.oldPwdSub')"/>
          </n-form-item>
          <n-form-item path="new_password" :label="t('userProfile.newPwd')">
            <n-input v-model:value="modelRef.new_password" @keydown.enter.prevent :placeholder="t('userProfile.newPwdSub')"/>
          </n-form-item>
          <n-form-item path="new_password_again" :label="t('userProfile.newPwdAgain')">
            <n-input v-model:value="modelRef.new_password_again" @keydown.enter.prevent
                     :placeholder="t('userProfile.newPwdAgainSub')"/>
          </n-form-item>
        </n-form>
        <n-button class="alert-btn" type="primary" @click="saveNewPassword">{{ t('userProfile.saveBtn') }}</n-button>
      </div>
    </n-card>

    <n-card
        class="notify-card"
        :embedded="true"
        hoverable
        content-style="padding: 0"
    >
      <n-p class="title">{{ t('userProfile.notify') }}</n-p>
      <div class="form">
        <n-form ref="formRef" :rules="rules" :style="themeStore.menuCollapsed?({width: '100%'}):({width: '60%'})">
          <n-form-item path="old_password" :label="t('userProfile.enableNotify')">
            <n-switch />
          </n-form-item>
        </n-form>
      </div>

    </n-card>

    <n-card
        class="del-card"
        :embedded="true"
        hoverable
        content-style="padding: 0"
    >
      <n-p class="title">{{ t('userProfile.deleteAccount') }}</n-p>
      <div class="form">
        <n-alert type="warning">
          {{ t('userProfile.deleteAccountSub') }}
        </n-alert>
        <n-button strong style="margin-top: 20px; color: #e3e4e7" type="error">{{ t('userProfile.deleteBtn') }}</n-button>
      </div>
    </n-card>


  </div>
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
      background-color: rgba(216,216,216,0.1);
      padding: 10px 20px 10px 20px;
      border-radius: 3px 3px 0 0;
    }
    .form {
      margin: 20px;
    }

    .alert-btn {
      width: 80px;
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
      background-color: rgba(216,216,216,0.1);
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
      background-color: rgba(216,216,216,0.1);
      padding: 10px 20px 10px 20px;
      border-radius: 3px 3px 0 0;
    }
    .form {
      margin: 20px;
    }
    margin-bottom: 20px;
  }


}

.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}


</style>