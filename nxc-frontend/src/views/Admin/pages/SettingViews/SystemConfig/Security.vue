<script setup lang="ts">
import useThemeStore from "@/stores/useThemeStore";
import useSettingStore from "@/stores/useSettingStore";

const themeStore = useThemeStore();
const settingStore = useSettingStore();

// let saveSetting = () => {
//   settingStore.saveSetting()
// }

// const optionsAll = [
//   {
//     title: 'Verify',
//     describe: '开启后将会强制要求用户进行邮箱验证。',
//     placeHolder: '',
//     inputType: 'switch',
//     data: settingStore.settings.security.email_verify,
//     dataType: 'string'
//   },
//   {
//     title: '禁止使用Gmail多别名',
//     describe: '开启后Gmail多别名将无法注册。',
//     placeHolder: '请输入',
//     inputType: 'input',
//     data: settingStore.settings.security.email_gmail_limit_enable,
//     dataType: 'string'
//   },
//
// ]

type SecurityModel =
    | "email_verify"
    | "email_gmail_limit_enable"
    | "safe_mode_enable"
    | "secure_path"
    | "email_whitelist_enable"
    | "recaptcha_enable"
    | "recaptcha_site_key"
    | "ip_register_limit_enable"
    | "ip_register_limit_times"
    | "ip_register_lock_time";

interface SecuritySetting {
  title: string;
  // shallow: string;
  description: string
  modelValue: SecurityModel;
  type: string;
  placeholder?: string;
}

let securitySettingsData: SecuritySetting[] = [
  {
    title: "邮箱验证",
    description: "开启后将会强制要求用户进行邮箱验证。",
    type: "switch",
    modelValue: "email_verify"
  },
  {
    title: "禁止使用Gmail多别名",
    description: "开启后Gmail多别名将无法注册。",
    type: "switch",
    modelValue: "email_gmail_limit_enable"
  },
  {
    title: "安全模式",
    description: "开启后除了站点URL以外的绑定本站点的域名访问都将会被403。",
    type: "switch",
    modelValue: "safe_mode_enable"
  },
  {
    title: "后台路径",
    description: "后台管理路径，修改后将会改变原有的admin路径。",
    type: "input",
    placeholder: "https://x.com/logo.jpeg",
    modelValue: "secure_path"
  },
  {
    title: "邮箱后缀白名单",
    description: "开启后在名单中的邮箱后缀才允许进行注册。",
    type: "switch",
    modelValue: "email_whitelist_enable"
  },
  {
    title: "防机器人",
    description: "开启后将会使用Google reCAPTCHA防止机器人。",
    type: "switch",
    modelValue: "recaptcha_enable"
  },
  {
    title: 'hCaptcha SiteKey',
    description: '该SiteKey用于请求hCaptcha服务器来标识网站编号',
    type: "input",
    placeholder: 'a3ca066c-0ea0-42fe-bcd2-23f4ab48d528',
    modelValue: 'recaptcha_site_key',
  },
  {
    title: "IP注册限制",
    description: "开启后如果IP注册账户达到规则要求将会被限制注册，请注意IP判断可能因为CDN或前置代理导致问题。",
    type: "switch",
    modelValue: "ip_register_limit_enable"
  },
  {
    title: "次数",
    description: "达到注册次数后开启惩罚。",
    type: "input-number",
    placeholder: "请输入",
    modelValue: "ip_register_limit_times"
  },
  {
    title: "惩罚时间(分钟)",
    description: "需要等待惩罚时间过后才可以再次注册。",
    type: "input-number",
    placeholder: "请输入",
    modelValue: "ip_register_lock_time"
  }
]

</script>

<script lang="ts">
export default {
  name: 'Security'
}
</script>

<template>
  <div class="root">
    <n-card :embedded="true" class="security-panel" title="安全设置" :bordered="false">

<!--      <div class="item">-->
<!--        <span class="l-content">-->
<!--          <div class="describe">-->
<!--            <p class="title">邮箱验证</p>-->
<!--            <p class="shallow">开启后将会强制要求用户进行邮箱验证。</p>-->
<!--          </div>-->
<!--        </span>-->
<!--        <span class="r-content to-right">-->
<!--          <n-switch-->
<!--              size="medium"-->
<!--              v-model:value="settingStore.settings.security.email_verify"-->
<!--              @update:value="settingStore.saveOption('security','email_verify', settingStore.settings.security.email_verify)"-->
<!--          ></n-switch>-->
<!--        </span>-->
<!--      </div>-->

<!--      <div class="item">-->
<!--        <span class="l-content">-->
<!--          <div class="describe">-->
<!--            <p class="title">禁止使用Gmail多别名</p>-->
<!--            <p class="shallow">开启后Gmail多别名将无法注册。</p>-->
<!--          </div>-->
<!--        </span>-->
<!--        <span class="r-content to-right">-->
<!--          <n-switch-->
<!--              size="medium"-->
<!--              v-model:value="settingStore.settings.security.email_gmail_limit_enable"-->
<!--              @update:value="settingStore.saveOption('security','email_gmail_limit_enable', settingStore.settings.security.email_gmail_limit_enable)"-->
<!--          ></n-switch>-->
<!--        </span>-->
<!--      </div>-->

<!--      <div class="item">-->
<!--        <span class="l-content">-->
<!--          <div class="describe">-->
<!--            <p class="title">安全模式</p>-->
<!--            <p class="shallow">开启后除了站点URL以外的绑定本站点的域名访问都将会被403。</p>-->
<!--          </div>-->
<!--        </span>-->
<!--        <span class="r-content to-right">-->
<!--          <n-switch-->
<!--              size="medium"-->
<!--              v-model:value="settingStore.settings.security.safe_mode_enable"-->
<!--              @update:value="settingStore.saveOption('security','safe_mode_enable', settingStore.settings.security.safe_mode_enable)"-->
<!--          ></n-switch>-->
<!--        </span>-->
<!--      </div>-->

<!--      <div class="item">-->
<!--        <span class="l-content">-->
<!--          <div class="describe">-->
<!--            <p class="title">后台路径</p>-->
<!--            <p class="shallow">后台管理路径，修改后将会改变原有的admin路径。</p>-->
<!--          </div>-->
<!--        </span>-->
<!--        <span class="r-content">-->
<!--          <n-input-->
<!--              size="large"-->
<!--              placeholder="https://x.com/logo.jpeg"-->
<!--              v-model:value="settingStore.settings.security.secure_path"-->
<!--              @blur="settingStore.saveOption('security','secure_path', settingStore.settings.security.secure_path)"-->
<!--          ></n-input>-->
<!--        </span>-->
<!--      </div>-->

<!--      <div class="item">-->
<!--        <span class="l-content">-->
<!--          <div class="describe">-->
<!--            <p class="title">邮箱后缀白名单</p>-->
<!--            <p class="shallow">开启后在名单中的邮箱后缀才允许进行注册。</p>-->
<!--          </div>-->
<!--        </span>-->
<!--        <span class="r-content to-right">-->
<!--          <n-switch-->
<!--              size="medium"-->
<!--              v-model:value="settingStore.settings.security.email_whitelist_enable"-->
<!--              @update:value="settingStore.saveOption('security','email_whitelist_enable', settingStore.settings.security.email_whitelist_enable)"-->
<!--          ></n-switch>-->
<!--        </span>-->
<!--      </div>-->

<!--      <div class="item">-->
<!--        <span class="l-content">-->
<!--          <div class="describe">-->
<!--            <p class="title">防机器人</p>-->
<!--            <p class="shallow">开启后将会使用Google reCAPTCHA防止机器人。</p>-->
<!--          </div>-->
<!--        </span>-->
<!--        <span class="r-content to-right">-->
<!--          <n-switch-->
<!--              size="medium"-->
<!--              v-model:value="settingStore.settings.security.recaptcha_enable"-->
<!--              @update:value="settingStore.saveOption('security','recaptcha_enable', settingStore.settings.security.recaptcha_enable)"-->
<!--          ></n-switch>-->
<!--        </span>-->
<!--      </div>-->

<!--      <div class="item">-->
<!--        <span class="l-content">-->
<!--          <div class="describe">-->
<!--            <p class="title">IP注册限制</p>-->
<!--            <p class="shallow">开启后如果IP注册账户达到规则要求将会被限制注册，请注意IP判断可能因为CDN或前置代理导致问题。</p>-->
<!--          </div>-->
<!--        </span>-->
<!--        <span class="r-content to-right">-->
<!--          <n-switch-->
<!--              size="medium"-->
<!--              v-model:value="settingStore.settings.security.ip_register_limit_enable"-->
<!--              @update:value="settingStore.saveOption('security','ip_register_limit_enable', settingStore.settings.security.ip_register_limit_enable)"-->
<!--          ></n-switch>-->
<!--        </span>-->
<!--      </div>-->

<!--      <div class="item">-->
<!--        <span class="l-content">-->
<!--          <div class="describe">-->
<!--            <p class="title">次数</p>-->
<!--            <p class="shallow">达到注册次数后开启惩罚。</p>-->
<!--          </div>-->
<!--        </span>-->
<!--        <span class="r-content">-->
<!--          <n-input-number-->
<!--              size="large"-->
<!--              placeholder="请输入"-->
<!--              v-model:value.number="settingStore.settings.security.ip_register_limit_times"-->
<!--              @blur="settingStore.saveOption('security','ip_register_limit_times', settingStore.settings.security.ip_register_limit_times)"-->
<!--          ></n-input-number>-->
<!--        </span>-->
<!--      </div>-->

<!--      <div class="item">-->
<!--        <span class="l-content">-->
<!--          <div class="describe">-->
<!--            <p class="title">惩罚时间(分钟)</p>-->
<!--            <p class="shallow">需要等待惩罚时间过后才可以再次注册。</p>-->
<!--          </div>-->
<!--        </span>-->
<!--        <span class="r-content">-->
<!--          <n-input-number-->
<!--              size="large"-->
<!--              placeholder="请输入"-->
<!--              v-model:value.number="settingStore.settings.security.ip_register_lock_time"-->
<!--              @blur="settingStore.saveOption('security','ip_register_lock_time', settingStore.settings.security.ip_register_lock_time)"-->
<!--          ></n-input-number>-->
<!--        </span>-->
<!--      </div>-->



      <div v-for="(item, index) in securitySettingsData" :key="index" class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">{{item.title}}</p>
            <p class="shallow">{{item.description}}</p>
          </div>
        </span>
        <span v-if="item.type === 'switch'" class="r-content to-right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.security[item.modelValue]"
              @update:value="settingStore.saveOption('security', item.modelValue, settingStore.settings.security[item.modelValue])"
          ></n-switch>
        </span>
        <span v-if="item.type === 'input'" class="r-content">
          <n-input
              size="large"
              :placeholder="item.placeholder"
              v-model:value="settingStore.settings.security[item.modelValue]"
              @blur="settingStore.saveOption('security', item.modelValue, settingStore.settings.security[item.modelValue])"
          ></n-input>
        </span>
        <span v-if="item.type === 'input-number'" class="r-content">
          <n-input-number
              size="large"
              :placeholder="item.placeholder"
              v-model:value.number="settingStore.settings.security[item.modelValue]"
              @blur="settingStore.saveOption('security', item.modelValue, settingStore.settings.security[item.modelValue])"
          ></n-input-number>
        </span>
      </div>



    </n-card>

  </div>
</template>

<style lang="less" scoped>
.root {
  min-width: 900px;

  .security-panel {
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
  }
}
.to-right {
  text-align: right;
}

//
//.n-card {
//  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
//  border: 0;
//}
</style>