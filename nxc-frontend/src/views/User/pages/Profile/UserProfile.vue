<script setup lang="ts">
import {onMounted, ref} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import {NIcon, useNotification} from 'naive-ui'
import {Wallet as walletIcon} from "@vicons/ionicons5"
import type {FormInst, FormRules, NotificationType} from 'naive-ui'
import {useMessage} from 'naive-ui'
import {hashPassword} from "@/utils/encryptor";
import instance from "@/axios";

interface ModelType {
  old_password: string | null
  new_password: string | null
  new_password_again: string | null
}

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
    let { data } = await instance.post(apiAddrStore.apiAddr.user.verifyOldPassword, {
      email: userInfoStore.thisUser.email,
      old_password: hashPassword(modelRef.value.old_password.trim())
    });
    if (data.code === 200) {
      return data.verified as boolean;
    } else {
      notify('error', '旧密码验证失败', data.msg)
      console.log(data);
      return false;
    }
  } catch (error) {
    notify('error', '旧密码验证失败', error.toString())
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
            let { data } = await instance.post(apiAddrStore.apiAddr.user.applyNewPassword, {
              email: userInfoStore.thisUser.email,
              new_password: hashPassword(modelRef.value.new_password)
            });
            if (data.code === 200 && data.updated as boolean) {
              modelRef.value.old_password = ''
              modelRef.value.new_password = ''
              modelRef.value.new_password_again = ''
              console.log('修改成功', data);
              notify('success', '修改成功', '请使用新密码登录')

            }
          } catch (error) {
            notify('error', '密码修改失败', error.toString())
          }
        } else {
          console.log('旧密码验证失败');
        }
      } else {
        notify('error', '两次密码输入不一致')
        console.log('两次输入密码不一致');
      }
    } else {
      console.log('新密码格式错误');
      notify('error', '新密码格式错误')

    }
  } else {
    console.log('旧密码不能为空');
    notify('error', '旧密码不能为空')
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
    <n-card class="wallet-card" :embedded="true" hoverable content-style="padding: 0">
      <div class="wallet-header">
        <p class="wallet-title">我的钱包</p>
        <n-icon size="30" style="opacity: 0.1">
          <walletIcon/>
        </n-icon>
      </div>
      <div class="wallet-content">
        <p class="balance">30.20</p>
        <p class="unit">USDT</p>
      </div>
      <div class="wallet-bottom">
        <p class="sub">账户余额（仅消费）</p>
      </div>
    </n-card>

    <n-card class="password-alert-card" :embedded="true" hoverable title="修改密码" header-style="font-weight: 300;">
      <div class="form">
        <n-form ref="formRef" :rules="rules" :style="themeStore.menuCollapsed?({width: '100%'}):({width: '60%'})">
          <n-form-item path="old_password" label="旧密码">
            <n-input v-model:value="modelRef.old_password" @keydown.enter.prevent placeholder="请输入旧密码"/>
          </n-form-item>
          <n-form-item path="new_password" label="新密码">
            <n-input v-model:value="modelRef.new_password" @keydown.enter.prevent placeholder="请输入新密码"/>
          </n-form-item>
          <n-form-item path="new_password_again" label="确认密码">
            <n-input v-model:value="modelRef.new_password_again" @keydown.enter.prevent
                     placeholder="请再输入一遍新密码"/>
          </n-form-item>
        </n-form>
        <n-button class="alert-btn" type="primary" @click="saveNewPassword">保存</n-button>
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

      x
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
    .form {
      width: 100%;
    }

    .alert-btn {
      width: 80px;
    }

    //@media screen and (min-width: 769px) {
    //  .form {
    //    width: 50%;
    //  }
    //}

  }


}

.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}


</style>