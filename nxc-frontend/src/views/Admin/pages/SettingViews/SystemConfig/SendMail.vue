<script setup lang="ts">
import {h, reactive, ref} from "vue"
// import useThemeStore from '@/stores/useThemeStore'
import useSettingStore from "@/stores/useSettingStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import instance from "@/axios/index";
// import {makeNotify} from "@/utils/notify"
import {HourglassOutline as waitIcon} from '@vicons/ionicons5'
import {NIcon, type NotificationType, useDialog, useMessage, useNotification} from 'naive-ui'

const notification = useNotification()
const userInfoStore = useUserInfoStore()
const apiAddrStore = useApiAddrStore();
// const themeStore = useThemeStore()
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
  message.warning('发送邮件中', {
    icon: () => h(NIcon, null, {default: () => h(waitIcon)})
  })
  try {
    // let {data} = await instance.post(apiAddrStore.apiAddr.admin.sendTestMail, {
    let {data} = await instance.post('/api/admin/v1/mail/test', {
      email: userInfoStore.thisUser.email
    })
    show.value = false
    if (data.code === 200) {
      // makeNotify('success', '发送邮件成功', '请查收该管理员邮箱')
      console.log(data)
      let {info} = data
      dialog.info({
        title: '成功',
        content: `收信地址: ${userInfoStore.thisUser.email}\t\n
                  发信服务器: ${info.host}
                  发信端口: ${info.port}
                  发信加密方式: ${info.use_ssl ? 'SSL' : 'TLS'}
                  发信用户名: ${info.username}`
      })
    } else {
      makeNotify('error', '发送邮件失败', data.msg.toString())

    }
  } catch (error) {
    console.log(error)
  }


}

// interface FormItem {
//   title: string;
//   shallow: string;
//   model: keyof SendmailSettings;
//   type: 'input' | 'input-number' | 'select';
//   placeholder: string;
//   options?: Array<{ label: string; value: string }>;
// }

interface SendmailSettings {
  email_host: string;
  email_port: number;
  email_encryption: string;
  email_username: string;
  email_password: string;
  email_from_address: string;
  email_template: string;
}

let options = reactive([
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

let smtpSettings: SmtpSetting[] = [
  {
    title: "SMTP 服务器地址",
    shallow: "由邮件服务商提供的服务地址",
    model: "email_host",
    type: "input"
  },
  {
    title: "SMTP 服务端口",
    shallow: "常见的端口有 25, 465, 587",
    model: "email_port",
    type: "input-number"
  },
  {
    title: "SMTP 加密方式",
    shallow: "465 端口加密方式一般为 SSL，587 端口加密方式一般为 TLS",
    model: "email_encryption",
    type: "input"
  },
  {
    title: "SMTP 账号",
    shallow: "由邮件服务商提供的账号",
    model: "email_username",
    type: "input"
  },
  {
    title: "SMTP 密码",
    shallow: "由邮件服务商提供的密码",
    model: "email_password",
    type: "input"
  },
  {
    title: "发件地址",
    shallow: "由邮件服务商提供的发件地址",
    model: "email_from_address",
    type: "input"
  },
  {
    title: "邮件模版",
    shallow: "你可以在文档查看如何自定义邮件模板",
    model: "email_template",
    type: "select"
  }
]

</script>

<script lang="ts">
export default {
  name: 'SendMail'
}
</script>

<template>

  <n-card class="root" :embedded="true" title="邮件设置" :bordered="false">
    <n-alert type="warning" style="margin-bottom: 30px" :bordered="false">
      如果你更改了本页配置，需要对队列服务进行重启。&nbsp;另外本页配置优先级高于.env中邮件配置。
    </n-alert>


    <div v-for="setting in smtpSettings" :key="setting.title" class="item">
    <span class="l-content">
      <div class="describe">
        <p class="title">{{ setting.title }}</p>
        <p class="shallow">{{ setting.shallow }}</p>
      </div>
    </span>
      <span class="r-content" v-if="setting.type === 'input'">
      <n-input
          size="large"
          placeholder="请输入"
          v-model:value="settingStore.settings.sendmail[setting.model]"
          @blur="settingStore.saveOption('sendmail', setting.model, settingStore.settings.sendmail[setting.model])"
      />
    </span>
      <span class="r-content" v-if="setting.type === 'input-number'">
      <n-input-number
          size="large"
          placeholder="请输入"
          v-model:value.number="settingStore.settings.sendmail[setting.model]"
          @blur="settingStore.saveOption('sendmail', setting.model, settingStore.settings.sendmail[setting.model])"
      />
    </span>
      <span class="r-content" v-if="setting.type === 'select'">
      <n-select
          size="large"
          :options="options"
          v-model:value="settingStore.settings.sendmail[setting.model]"
          @update:value="settingStore.saveOption('sendmail', setting.model, settingStore.settings.sendmail[setting.model])"
      />
    </span>
    </div>


    <!--        <div class="item">-->
    <!--            <span class="l-content">-->
    <!--              <div class="describe">-->
    <!--                <p class="title">SMTP服务器地址</p>-->
    <!--                <p class="shallow">由邮件服务商提供的服务地址</p>-->
    <!--              </div>-->
    <!--            </span>-->
    <!--          <span class="r-content">-->
    <!--              <n-input-->
    <!--                  size="large"-->
    <!--                  placeholder="请输入"-->
    <!--                  v-model:value="settingStore.settings.sendmail.email_host"-->
    <!--                  @blur="settingStore.saveOption('sendmail','email_host', settingStore.settings.sendmail.email_host)"-->
    <!--              ></n-input>-->
    <!--            </span>-->
    <!--        </div>-->

    <!--        <div class="item">-->
    <!--            <span class="l-content">-->
    <!--              <div class="describe">-->
    <!--                <p class="title">SMTP服务端口</p>-->
    <!--                <p class="shallow">常见的端口有25, 465, 587</p>-->
    <!--              </div>-->
    <!--            </span>-->
    <!--          <span class="r-content">-->
    <!--              <n-input-number-->
    <!--                  size="large"-->
    <!--                  placeholder="请输入"-->
    <!--                  v-model:value.number="settingStore.settings.sendmail.email_port"-->
    <!--                  @blur="settingStore.saveOption('sendmail','email_port', settingStore.settings.sendmail.email_port)"-->
    <!--              ></n-input-number>-->
    <!--            </span>-->
    <!--        </div>-->

    <!--        <div class="item">-->
    <!--            <span class="l-content">-->
    <!--              <div class="describe">-->
    <!--                <p class="title">SMTP加密方式</p>-->
    <!--                <p class="shallow">465端口加密方式一般为SSL，587端口加密方式一般为TLS</p>-->
    <!--              </div>-->
    <!--            </span>-->
    <!--          <span class="r-content">-->
    <!--              <n-input-->
    <!--                  size="large"-->
    <!--                  placeholder="请输入"-->
    <!--                  v-model:value="settingStore.settings.sendmail.email_encryption"-->
    <!--                  @blur="settingStore.saveOption('sendmail','email_encryption', settingStore.settings.sendmail.email_encryption)"-->
    <!--              ></n-input>-->
    <!--            </span>-->
    <!--        </div>-->

    <!--        <div class="item">-->
    <!--            <span class="l-content">-->
    <!--              <div class="describe">-->
    <!--                <p class="title">SMTP账号</p>-->
    <!--                <p class="shallow">由邮件服务商提供的账号</p>-->
    <!--              </div>-->
    <!--            </span>-->
    <!--          <span class="r-content">-->
    <!--              <n-input-->
    <!--                  size="large"-->
    <!--                  placeholder="请输入"-->
    <!--                  v-model:value="settingStore.settings.sendmail.email_username"-->
    <!--                  @blur="settingStore.saveOption('sendmail','email_username', settingStore.settings.sendmail.email_username)"-->
    <!--              ></n-input>-->
    <!--            </span>-->
    <!--        </div>-->

    <!--        <div class="item">-->
    <!--            <span class="l-content">-->
    <!--              <div class="describe">-->
    <!--                <p class="title">SMTP密码</p>-->
    <!--                <p class="shallow">由邮件服务商提供的密码</p>-->
    <!--              </div>-->
    <!--            </span>-->
    <!--          <span class="r-content">-->
    <!--              <n-input-->
    <!--                  size="large"-->
    <!--                  placeholder="请输入"-->
    <!--                  v-model:value="settingStore.settings.sendmail.email_password"-->
    <!--                  @blur="settingStore.saveOption('sendmail','email_password', settingStore.settings.sendmail.email_password)"-->
    <!--              ></n-input>-->
    <!--            </span>-->
    <!--        </div>-->

    <!--        <div class="item">-->
    <!--            <span class="l-content">-->
    <!--              <div class="describe">-->
    <!--                <p class="title">发件地址</p>-->
    <!--                <p class="shallow">由邮件服务商提供的发件地址</p>-->
    <!--              </div>-->
    <!--            </span>-->
    <!--          <span class="r-content">-->
    <!--              <n-input-->
    <!--                  size="large"-->
    <!--                  placeholder="请输入"-->
    <!--                  v-model:value="settingStore.settings.sendmail.email_from_address"-->
    <!--                  @blur="settingStore.saveOption('sendmail','email_from_address', settingStore.settings.sendmail.email_from_address)"-->
    <!--              ></n-input>-->
    <!--            </span>-->
    <!--        </div>-->

    <!--        <div class="item">-->
    <!--            <span class="l-content">-->
    <!--              <div class="describe">-->
    <!--                <p class="title">邮件模版</p>-->
    <!--                <p class="shallow">你可以在文档查看如何自定义邮件模板</p>-->
    <!--              </div>-->
    <!--            </span>-->
    <!--          <span class="r-content">-->
    <!--            <n-select-->
    <!--                size="large"-->
    <!--                :options="options"-->
    <!--                v-model:value="settingStore.settings.sendmail.email_template"-->
    <!--                @update:value="settingStore.saveOption('sendmail','email_template', settingStore.settings.sendmail.email_template)"-->
    <!--            />-->
    <!--            </span>-->
    <!--        </div>-->

    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">发送测试邮件</p>
            <p class="shallow">邮件将会发送到当前登陆用户邮箱</p>
          </div>
        </span>
      <span class="r-content to-right">
        <n-spin :show="show" :delay="100" size="small">
        <n-button @click="sendTestMail" size="large" type="primary">发送测试邮件</n-button>
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

//.n-card {
//  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
//  border: 0;
//}
</style>