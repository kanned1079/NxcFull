<script setup lang="ts">
import {h, reactive, ref} from "vue"
import useSettingStore from "@/stores/useSettingStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {useI18n} from "vue-i18n"
import instance from "@/axios/index";
import {HourglassOutline as waitIcon, Flask as testIcon} from '@vicons/ionicons5'
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
  message.warning(t('adminViews.systemConfig.sendMail.common.sending') , {
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
    type: 'input',
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
    <n-alert type="warning" style="margin-bottom: 30px" :bordered="false">
      {{ t('adminViews.systemConfig.sendMail.common.warning') }}
    </n-alert>

    <div v-for="setting in smtpSettings" :key="setting.title" class="item">
  <span class="l-content">
    <div class="describe">
      <p class="title">{{ t(setting.title) }}</p>
      <p class="shallow">{{ t(setting.shallow) }}</p>
    </div>
  </span>
      <span class="r-content" v-if="setting.type === 'input'">
    <n-input
        size="large"
        :placeholder="t(setting.placeholder || '')"
        v-model:value="settingStore.settings.sendmail[setting.model]"
        @blur="saveFiled(setting.model, settingStore.settings.sendmail[setting.model])"
    />
  </span>
      <span class="r-content" v-if="setting.type === 'input-number'">
    <n-input-number
        size="large"
        :placeholder="t(setting.placeholder || '')"
        v-model:value.number="settingStore.settings.sendmail[setting.model]"
        @blur="saveFiled(setting.model, settingStore.settings.sendmail[setting.model])"
    />
  </span>
      <span class="r-content" v-if="setting.type === 'select'">
    <n-select
        size="large"
        :options="options"
        :placeholder="t(setting.placeholder || '')"
        v-model:value="settingStore.settings.sendmail[setting.model]"
        @update:value="saveFiled(setting.model, settingStore.settings.sendmail[setting.model])"
    />
  </span>
    </div>

    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">{{ t('adminViews.systemConfig.sendMail.common.sendTestMailTitle') }}</p>
            <p class="shallow">{{ t('adminViews.systemConfig.sendMail.common.sendTestMailShallow') }}</p>
          </div>
        </span>
      <span class="r-content to-right">
        <n-spin :show="show" :delay="100" size="small">
        <n-button @click="sendTestMail" size="large" type="primary" :bordered="false" icon-placement="left">
          {{ t('adminViews.systemConfig.sendMail.common.sendTestMailTo') +' '+ userInfoStore.thisUser.email }}
          <template #icon>
            <n-icon size="16">
               <testIcon />
            </n-icon>
          </template>
        </n-button>
        </n-spin>
        </span>
    </div>

  </n-card>
</template>

<style lang="less" scoped>
.root {
  .warm {
    background-color: #ffefd1;
    height: 50px;
    border: 0;
    line-height: 50px;
    text-align: left;

    p {
      color: #956d34;
      margin-left: 20px;
      font-size: 1rem;
    }
  }
}

.item {
  height: 50px;
  display: flex;
  margin-bottom: 30px;

  .l-content {
    flex: 1;

    .describe {
      .title {
        font-weight: bold;
      }

      .shallow {
        margin-top: 5px;
        opacity: 0.5;
      }
    }
  }

  .r-content {
    margin-left: 30px;
    flex: 0.8;
    justify-content: center;
    line-height: 50px;
  }
}

.to-right {
  text-align: right;
}

</style>