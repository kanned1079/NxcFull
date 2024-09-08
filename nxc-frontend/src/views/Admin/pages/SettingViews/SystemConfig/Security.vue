<script setup lang="ts" name="Security">
import useThemeStore from "@/stores/useThemeStore";
import useSettingStore from "@/stores/useSettingStore";

const themeStore = useThemeStore();
const settingStore = useSettingStore();

let saveSetting = () => {
  settingStore.saveSetting()
}

const optionsAll = [
  {
    title: 'Verify',
    describe: '开启后将会强制要求用户进行邮箱验证。',
    placeHolder: '',
    inputType: 'switch',
    data: settingStore.settings.security.email_verify,
    dataType: 'string'
  },
  {
    title: '禁止使用Gmail多别名',
    describe: '开启后Gmail多别名将无法注册。',
    placeHolder: '请输入',
    inputType: 'input',
    data: settingStore.settings.security.email_gmail_limit_enable,
    dataType: 'string'
  },

]
</script>

<template>
  <div class="root">

<!--    <n-card v-for="item in options" :key="item.title" :embedded="true" class="security-panel" style="background-color: #e3e5e7" title="安全设置">-->
<!--      <div v-for="item in options" :key="item.title" class="item" style="background-color: #e3e5e7">-->
<!--        <span class="l-content">-->
<!--          <div class="describe">-->
<!--            <p class="title">{{ item.title }}</p>-->
<!--            <p class="shallow">{{ item.describe }}</p>-->
<!--          </div>-->
<!--        </span>-->
<!--        <span class="r-content to-right">-->
<!--          <n-switch v-if="item.inputType='switch'" size="medium" @update:value="saveSetting" :v-model="settingStore.settings.security.email_verify"></n-switch>-->
<!--          <n-input v-if="item.inputType='input' && item.dataType='string'" size="large"  placeholder="请输入" @blur="saveSetting" v-model:value.number="settingStore.settings.security.ip_register_limit_times"></n-input>-->

<!--        </span>-->
<!--      </div>-->

    <n-card :embedded="true" class="security-panel" title="安全设置">



<!--      <div v-for="i in optionsAll" :key="i.title" class="item" style="background-color: #e3e5e7">-->
<!--        <span class="l-content">-->
<!--          <div class="describe">-->
<!--            <p class="title">{{ i.title }}</p>-->
<!--            <p class="shallow">{{ i.describe }}</p>-->
<!--          </div>-->
<!--        </span>-->
<!--        <span class="r-content to-right">-->
<!--          <n-switch v-if="i.inputType=='switch'" size="medium" @update:value="saveSetting" v-model:value="i.data"></n-switch>-->
<!--          <n-input v-if="i.inputType=='input' && i.dataType=='string'" size="large"  placeholder="请输入" @blur="saveSetting" v-model:value.number="i.data"></n-input>-->
<!--        </span>-->
<!--      </div>-->




      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">邮箱验证</p>
            <p class="shallow">开启后将会强制要求用户进行邮箱验证。</p>
          </div>
        </span>
        <span class="r-content to-right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.security.email_verify"
              @update:value="settingStore.saveOption('security','email_verify', settingStore.settings.security.email_verify)"
          ></n-switch>
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">禁止使用Gmail多别名</p>
            <p class="shallow">开启后Gmail多别名将无法注册。</p>
          </div>
        </span>
        <span class="r-content to-right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.security.email_gmail_limit_enable"
              @update:value="settingStore.saveOption('security','email_gmail_limit_enable', settingStore.settings.security.email_gmail_limit_enable)"
          ></n-switch>
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">安全模式</p>
            <p class="shallow">开启后除了站点URL以外的绑定本站点的域名访问都将会被403。</p>
          </div>
        </span>
        <span class="r-content to-right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.security.safe_mode_enable"
              @update:value="settingStore.saveOption('security','safe_mode_enable', settingStore.settings.security.safe_mode_enable)"
          ></n-switch>
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">后台路径</p>
            <p class="shallow">后台管理路径，修改后将会改变原有的admin路径。</p>
          </div>
        </span>
        <span class="r-content">
          <n-input
              size="large"
              placeholder="https://x.com/logo.jpeg"
              v-model:value="settingStore.settings.security.secure_path"
              @blur="settingStore.saveOption('security','secure_path', settingStore.settings.security.secure_path)"
          ></n-input>
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">邮箱后缀白名单</p>
            <p class="shallow">开启后在名单中的邮箱后缀才允许进行注册。</p>
          </div>
        </span>
        <span class="r-content to-right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.security.email_whitelist_enable"
              @update:value="settingStore.saveOption('security','email_whitelist_enable', settingStore.settings.security.email_whitelist_enable)"
          ></n-switch>
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">防机器人</p>
            <p class="shallow">开启后将会使用Google reCAPTCHA防止机器人。</p>
          </div>
        </span>
        <span class="r-content to-right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.security.recaptcha_enable"
              @update:value="settingStore.saveOption('security','recaptcha_enable', settingStore.settings.security.recaptcha_enable)"
          ></n-switch>
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">IP注册限制</p>
            <p class="shallow">开启后如果IP注册账户达到规则要求将会被限制注册，请注意IP判断可能因为CDN或前置代理导致问题。</p>
          </div>
        </span>
        <span class="r-content to-right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.security.ip_register_limit_enable"
              @update:value="settingStore.saveOption('security','ip_register_limit_enable', settingStore.settings.security.ip_register_limit_enable)"
          ></n-switch>
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">次数</p>
            <p class="shallow">达到注册次数后开启惩罚。</p>
          </div>
        </span>
        <span class="r-content">
          <n-input-number
              size="large"
              placeholder="请输入"
              v-model:value.number="settingStore.settings.security.ip_register_limit_times"
              @blur="settingStore.saveOption('security','ip_register_limit_times', settingStore.settings.security.ip_register_limit_times)"
          ></n-input-number>
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">惩罚时间(分钟)</p>
            <p class="shallow">需要等待惩罚时间过后才可以再次注册。</p>
          </div>
        </span>
        <span class="r-content">
          <n-input-number
              size="large"
              placeholder="请输入"
              v-model:value.number="settingStore.settings.security.ip_register_lock_time"
              @blur="settingStore.saveOption('security','ip_register_lock_time', settingStore.settings.security.ip_register_lock_time)"
          ></n-input-number>
        </span>
      </div>

<!--      <div class="item">-->
<!--        <span class="l-content">-->
<!--          <div class="describe">-->
<!--            <p class="title">防爆破限制</p>-->
<!--            <p class="shallow">开启后如果该账户尝试登陆失败次数过多将会被限制。</p>-->
<!--          </div>-->
<!--        </span>-->
<!--        <span class="r-content to-right">-->
<!--          <n-switch size="medium" @update:value="saveSetting"></n-switch>-->
<!--        </span>-->
<!--      </div>-->

<!--      <div class="item">-->
<!--        <span class="l-content">-->
<!--          <div class="describe">-->
<!--            <p class="title">次数</p>-->
<!--            <p class="shallow">达到失败次数后开启惩罚。</p>-->
<!--          </div>-->
<!--        </span>-->
<!--        <span class="r-content">-->
<!--          <n-input size="large"  placeholder="请输入" @blur="saveSetting"></n-input>-->
<!--        </span>-->
<!--      </div>-->

<!--      <div class="item" style="margin-bottom: 0">-->
<!--        <span class="l-content">-->
<!--          <div class="describe">-->
<!--            <p class="title">惩罚时间(分钟)</p>-->
<!--            <p class="shallow">需要等待惩罚时间过后才可以再次登陆。</p>-->
<!--          </div>-->
<!--        </span>-->
<!--        <span class="r-content">-->
<!--          <n-input size="large"  placeholder="请输入" @blur="saveSetting"></n-input>-->
<!--        </span>-->
<!--      </div>-->


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


.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}
</style>