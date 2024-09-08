<script setup lang="ts" name="Site">
import {reactive} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import useSettingStore from "@/stores/useSettingStore"

const themeStore = useThemeStore();
const settingStore = useSettingStore();

let handleSaveOptions = () => {
  settingStore.saveSetting()
}


let siteInfo = reactive({
  // app_name: '',
  // app_description:'',
  // app_url: '',
  // force_https: false,
  // logo_url: '',
  // subscribe_url: '',
  // tos_url: '',
  // stop_register: false,
  // trial_time: '',
  // trial_subscribe: '',
  subscribe_list: [
    {
      label: 'Everybody\'s Got Something to Hide Except Me and My Monkey',
      value: 'song0',
      disabled: true
    },
    {
      label: 'Drive My Car',
      value: 'song1'
    },
    {
      label: 'Norwegian Wood',
      value: 'song2'
    },],
  // currency: '',
  // currency_symbol: '',

})

</script>

<template>
  <div class="root">
    <n-card :embedded="true" class="security-panel" title="站点">
      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">站点名称</p>
            <p class="shallow">用于显示需要站点名称的地方。</p>
          </div>
        </span>
        <span class="r-content">
          <n-input
              v-model:value="settingStore.settings.site.app_name"
              type="text"
              placeholder="请输入站点名称"
              size="large"
              @blur="settingStore.saveOption('site','app_name', settingStore.settings.site.app_name)"
          />
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">站点描述</p>
            <p class="shallow">用于显示需要站点描述的地方。</p>
          </div>
        </span>
        <span class="r-content">
          <n-input
              v-model:value="settingStore.settings.site.app_description"
              type="text"
              placeholder="请输入站点描述"
              size="large"
              @blur="settingStore.saveOption('site','app_description', settingStore.settings.site.app_description)"
          />
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">站点网址</p>
            <p class="shallow">当前网站最新网址，将会在邮件等需要用于网址处体现。</p>
          </div>
        </span>
        <span class="r-content">
          <n-input
              v-model:value="settingStore.settings.site.app_url"
              type="text"
              placeholder="请输入站点网址"
              size="large"
              @blur="settingStore.saveOption('site','app_url', settingStore.settings.site.app_url)"
          />
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">强制HTTPS</p>
            <p class="shallow">当站点没有使用HTTPS，CDN或反代开启强制HTTPS时需要开启。</p>
          </div>
        </span>
        <span class="r-content" style="text-align: right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.site.force_https"
              @update:value="settingStore.saveOption('site','force_https', settingStore.settings.site.force_https)"
          />
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">LOGO</p>
            <p class="shallow">用于显示需要LOGO的地方。</p>
          </div>
        </span>
        <span class="r-content">
          <n-input
              v-model:value="settingStore.settings.site.logo_url"
              type="text"
              placeholder="请输入logo的url地址"
              size="large"
              @blur="settingStore.saveOption('site','logo_url', settingStore.settings.site.logo_url)"
          />
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">订阅URL</p>
            <p class="shallow">用于订阅所使用，留空则为站点URL。如需多个订阅URL随机获取请使用逗号进行分割。</p>
          </div>
        </span>
        <span class="r-content">
          <n-input
              v-model:value="settingStore.settings.site.subscribe_url"
              type="text"
              placeholder="请输入订阅url"
              size="large"
              @blur="settingStore.saveOption('site','subscribe_url', settingStore.settings.site.subscribe_url)"
          />
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">用户条款(TOS)URL</p>
            <p class="shallow">用于跳转到用户条款(TOS)</p>
          </div>
        </span>
        <span class="r-content">
          <n-input
              v-model:value="settingStore.settings.site.tos_url"
              type="text"
              placeholder="请输入用户条款地址"
              size="large"
              @blur="settingStore.saveOption('site','tos_url', settingStore.settings.site.tos_url)"
          />
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">停止新用户注册</p>
            <p class="shallow">开启后任何人都将无法进行注册。</p>
          </div>
        </span>
        <span class="r-content" style="text-align: right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.site.stop_register"
              @update:value="settingStore.saveOption('site','stop_register', settingStore.settings.site.stop_register)"
          />
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">注册试用</p>
            <p class="shallow">选择需要试用的订阅，如果没有选项请先前往订阅管理添加。</p>
          </div>
        </span>
        <span class="r-content">
          <n-select
              v-model:value="settingStore.settings.site.trial_subscribe"
              size="large"
              :options="siteInfo.subscribe_list"
              @update:value="settingStore.saveOption('site','trial_subscribe', settingStore.settings.site.trial_subscribe)"
          />
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">试用时间(小时)</p>
            <p class="shallow">新用户注册时订阅试用事件。</p>
          </div>
        </span>
        <span class="r-content">
          <n-input-number
              v-model:value.number="settingStore.settings.site.trial_time"
              type="text"
              placeholder="请输入"
              size="large"
              @update:value="settingStore.saveOption('site','trial_time', settingStore.settings.site.trial_time)"
          />
        </span>
      </div>

      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">货币单位</p>
            <p class="shallow">仅用于展示使用，更改后系统中所有的货币单位都将发生变更。</p>
          </div>
        </span>
        <span class="r-content">
          <n-input
              v-model:value="settingStore.settings.site.currency"
              type="text"
              placeholder="CNY"
              size="large"
              @update:value="settingStore.saveOption('site','currency', settingStore.settings.site.currency)"
          />
        </span>
      </div>

      <div class="item" style="margin-bottom: 0">
        <span class="l-content">
          <div class="describe">
            <p class="title">货币符号</p>
            <p class="shallow">仅用于展示使用，更改后系统中所有的货币单位都将发生变更。</p>
          </div>
        </span>
        <span class="r-content">
          <n-input v-model:value="settingStore.settings.site.currency_symbol" type="text" placeholder="¥" size="large"/>
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

.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}
</style>