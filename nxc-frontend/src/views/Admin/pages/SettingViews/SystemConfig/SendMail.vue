<script setup lang="ts">
import {h, ref} from "vue"
import useSettingStore from "@/stores/useSettingStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {useI18n} from "vue-i18n"
import instance from "@/axios/index";
import {Flask as testIcon, HourglassOutline as waitIcon} from '@vicons/ionicons5'
import {NIcon, type NotificationType, useDialog, useMessage, useNotification} from 'naive-ui'

const {t} = useI18n()
const notification = useNotification()
const userInfoStore = useUserInfoStore()
const settingStore = useSettingStore()
const message = useMessage()
const dialog = useDialog()

let show = ref<boolean>(false)

let makeNotify = (type: NotificationType, title: string, msg: string) => {
  notification[type]({
    content: title,
    meta: msg,
    duration: 2500,
    keepAliveOnHover: true
  })
}

let sendTestMail = async () => {
  show.value = true
  message.warning(t('adminViews.systemConfig.sendMail.common.sending'), {
    icon: () => h(NIcon, null, {default: () => h(waitIcon)})
  })
  try {
    let {data} = await instance.post('/api/admin/v1/mail/test', {
      email: userInfoStore.thisUser.email
    })
    show.value = false
    if (data.code === 200) {
      let {info} = data
      dialog.success({
        title: t('adminViews.systemConfig.sendMail.common.success'),
        content: () =>
            h('div', {style: {marginTop: '15px'}}, [
              h('p', {style: {marginBottom: '5px'}}, `${t('adminViews.systemConfig.sendMail.common.receiverAddr')}: ${userInfoStore.thisUser.email}`),
              h('p', {style: {marginBottom: '5px'}}, `${t('adminViews.systemConfig.sendMail.common.senderHost')}: ${info.host}`),
              h('p', {style: {marginBottom: '5px'}}, `${t('adminViews.systemConfig.sendMail.common.senderPort')}: ${info.port}`),
              h('p', {style: {marginBottom: '5px'}}, `${t('adminViews.systemConfig.sendMail.common.senderEncrypt')}: ${info.use_ssl ? 'SSL' : 'TLS'}`),
              h('p', {style: {marginBottom: '5px'}}, `${t('adminViews.systemConfig.sendMail.common.senderUsername')}: ${info.username}`),
            ])
      });
    } else {
      makeNotify('error', t('adminViews.systemConfig.sendMail.common.sendErr'), data.msg.toString() || '')
    }
  } catch (error) {
    console.log(error)
  }
}

let options = ref([
  {
    label: 'default',
    value: 'default',
    disabled: false,
  },
  {
    label: 'classic',
    value: 'classic',
    disabled: false,
  },
])

type SmtpModel =
    | "email_host"
    | "email_port"
    | "email_encryption"
    | "email_username"
    | "email_password"
    | "email_from_address"
    | "email_template";

interface SmtpSetting {
  title: string;
  shallow: string;
  model: SmtpModel;
  placeholder?: string;
  type: string;
}

const smtpSettings: SmtpSetting[] = [
  {
    title: 'adminViews.systemConfig.sendMail.smtpServerAddress.title',
    shallow: 'adminViews.systemConfig.sendMail.smtpServerAddress.shallow',
    placeholder: 'adminViews.systemConfig.sendMail.smtpServerAddress.placeholder',
    model: 'email_host',
    type: 'input',
  },
  {
    title: 'adminViews.systemConfig.sendMail.smtpServerPort.title',
    shallow: 'adminViews.systemConfig.sendMail.smtpServerPort.shallow',
    placeholder: 'adminViews.systemConfig.sendMail.smtpServerPort.placeholder',
    model: 'email_port',
    type: 'input-number',
  },
  {
    title: 'adminViews.systemConfig.sendMail.smtpEncryption.title',
    shallow: 'adminViews.systemConfig.sendMail.smtpEncryption.shallow',
    placeholder: 'adminViews.systemConfig.sendMail.smtpEncryption.placeholder',
    model: 'email_encryption',
    type: 'input',
  },
  {
    title: 'adminViews.systemConfig.sendMail.smtpAccount.title',
    shallow: 'adminViews.systemConfig.sendMail.smtpAccount.shallow',
    placeholder: 'adminViews.systemConfig.sendMail.smtpAccount.placeholder',
    model: 'email_username',
    type: 'input',
  },
  {
    title: 'adminViews.systemConfig.sendMail.smtpPassword.title',
    shallow: 'adminViews.systemConfig.sendMail.smtpPassword.shallow',
    placeholder: 'adminViews.systemConfig.sendMail.smtpPassword.placeholder',
    model: 'email_password',
    type: 'input-pwd',
  },
  {
    title: 'adminViews.systemConfig.sendMail.senderAddress.title',
    shallow: 'adminViews.systemConfig.sendMail.senderAddress.shallow',
    placeholder: 'adminViews.systemConfig.sendMail.senderAddress.placeholder',
    model: 'email_from_address',
    type: 'input',
  },
  {
    title: 'adminViews.systemConfig.sendMail.emailTemplate.title',
    shallow: 'adminViews.systemConfig.sendMail.emailTemplate.shallow',
    placeholder: 'adminViews.systemConfig.sendMail.emailTemplate.placeholder',
    model: 'email_template',
    type: 'select',
  },
];

const saveFiled = async (k: string, v: any) => {
  let updateResponse = await settingStore.saveOption('sendmail', k, v)
  if (updateResponse.code === 200)
    message.success(t('adminViews.systemConfig.common.success'))
  else
    message.error(t('adminViews.systemConfig.common.err') + updateResponse.msg || '')
}

</script>

<script lang="ts">
export default {
  name: 'SendMail'
}
</script>

<template>
  <n-card
      class="root"
      :embedded="true"
      :title="t('adminViews.systemConfig.sendMail.common.title')"
      :bordered="false"
  >
    <n-alert type="warning" class="alert">
      {{ t('adminViews.systemConfig.sendMail.common.warning') }}
    </n-alert>

    <n-grid cols="1" responsive="screen" y-gap="16">
      <n-grid-item v-for="setting in smtpSettings" :key="setting.title" class="grid-item">
        <n-grid cols="1 s:1 m:2 l:2" responsive="screen" align-items="center">
          <!-- 左侧：标题 + 描述 -->
          <n-grid-item>
            <div class="describe">
              <p class="title">{{ t(setting.title) }}</p>
              <p class="shallow">{{ t(setting.shallow) }}</p>
            </div>
          </n-grid-item>

          <!-- 右侧：输入框 / 选择框 -->
          <n-grid-item>
            <n-input
                v-if="setting.type === 'input'"
                size="large"
                :placeholder="t(setting.placeholder || '')"
                v-model:value="settingStore.settings.sendmail[setting.model]"
                @blur="saveFiled(setting.model, settingStore.settings.sendmail[setting.model])"
            />
            <n-input
                v-if="setting.type === 'input-pwd'"
                type="password"
                show-password-on="click"
                size="large"
                :placeholder="t(setting.placeholder || '')"
                v-model:value="settingStore.settings.sendmail[setting.model]"
                @blur="saveFiled(setting.model, settingStore.settings.sendmail[setting.model])"
            />
            <n-input-number
                v-else-if="setting.type === 'input-number'"
                size="large"
                :placeholder="t(setting.placeholder || '')"
                v-model:value.number="settingStore.settings.sendmail[setting.model]"
                @blur="saveFiled(setting.model, settingStore.settings.sendmail[setting.model])"
            />
            <n-select
                v-else-if="setting.type === 'select'"
                size="large"
                :options="options"
                :placeholder="t(setting.placeholder || '')"
                v-model:value="settingStore.settings.sendmail[setting.model]"
                @update:value="saveFiled(setting.model, settingStore.settings.sendmail[setting.model])"
            />
          </n-grid-item>
        </n-grid>
      </n-grid-item>

      <!-- 测试邮件部分 -->
      <n-grid-item class="grid-item">
        <n-grid cols="1 s:1 m:2 l:2" responsive="screen" align-items="center">
          <n-grid-item>
            <div class="describe">
              <p class="title">{{ t('adminViews.systemConfig.sendMail.common.sendTestMailTitle') }}</p>
              <p class="shallow">{{ t('adminViews.systemConfig.sendMail.common.sendTestMailShallow') }}</p>
            </div>
          </n-grid-item>

          <n-grid-item class="to-right">
            <n-spin :show="show" :delay="100" size="small">
              <n-button @click="sendTestMail" size="large" type="primary" :bordered="false" icon-placement="left">
                {{ t('adminViews.systemConfig.sendMail.common.sendTestMailTo') + ' ' + userInfoStore.thisUser.email }}
                <template #icon>
                  <n-icon size="16">
                    <testIcon/>
                  </n-icon>
                </template>
              </n-button>
            </n-spin>
          </n-grid-item>
        </n-grid>
      </n-grid-item>
    </n-grid>

  </n-card>
</template>

<style lang="less" scoped>
.root {
  margin: 0 auto;
}

.alert {
  margin-bottom: 16px;
}

.grid-item {
  margin-bottom: 16px; // 统一所有配置项间距
}

.describe {
  .title {
    font-weight: bold;
  }

  .shallow {
    margin-top: 5px;
    opacity: 0.5;
  }
}

.to-right {
  text-align: right;
}
</style>