<script setup lang="ts" name="BackendConfig">
import {defineComponent} from 'vue'
import useThemeStore from "@/stores/useThemeStore"
import useSettingStore from "@/stores/useSettingStore";
const themeStore = useThemeStore()
const settingStore = useSettingStore()
let handok = () => {

}
</script>

<template>
  <n-card class="root" :embedded="true" title="后端设置">
    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">通讯密钥</p>
            <p class="shallow">与节点通讯的密钥，以便数据不会被他人获取。</p>
          </div>
        </span>
      <span class="r-content">
          <n-input size="large" placeholder="b5#*(PG*(3tgf4e^!$P$Gp$G789" @blur="handok" v-model:value="settingStore.settings.server.server_token"></n-input>
        </span>
    </div>

    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">节点拉取动作轮询间隔</p>
            <p class="shallow">节点从面板获取数据的间隔频率。</p>
          </div>
        </span>
      <span class="r-content">
          <n-input-number size="large" placeholder="https://x.com/logo.jpeg" @blur="handok" v-model:value.number="settingStore.settings.server.server_pull_interval"></n-input-number>
        </span>
    </div>

    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">节点推送动作轮询间隔</p>
            <p class="shallow">节点推送数据到面板的间隔频率。</p>
          </div>
        </span>
      <span class="r-content">
          <n-input-number size="large" placeholder="https://x.com/logo.jpeg" @blur="handok" v-model:value.number="settingStore.settings.server.server_push_interval"></n-input-number>
        </span>
    </div>
  </n-card>
</template>

<style lang="less" scoped>
.root {
  border: 0;
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
</style>