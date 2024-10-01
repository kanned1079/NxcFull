<script setup lang="ts" name="BackendConfig">
import {defineComponent} from 'vue'
// import useThemeStore from "@/stores/useThemeStore"
import useSettingStore from "@/stores/useSettingStore";

// const themeStore = useThemeStore()
const settingStore = useSettingStore()
// let handok = () => {
//
// }

type BackendModel =
    | "server_token"
    | "server_pull_interval"
    | "server_push_interval";

interface BackendSetting {
  title: string;
  description: string;
  modelValue: BackendModel;
  type: string;
  placeholder?: string;
}

let settingsData: BackendSetting[] = [
  {
    title: "通讯密钥",
    description: "与节点通讯的密钥，以便数据不会被他人获取。",
    type: "input",
    placeholder: "b5#*(PG*(3tgf4e^!$P$Gp$G789",
    modelValue: "server_token"
  },
  {
    title: "节点拉取动作轮询间隔",
    description: "节点从面板获取数据的间隔频率。",
    type: "input-number",
    placeholder: "https://x.com/logo.jpeg",
    modelValue: "server_pull_interval"
  },
  {
    title: "节点推送动作轮询间隔",
    description: "节点推送数据到面板的间隔频率。",
    type: "input-number",
    placeholder: "https://x.com/logo.jpeg",
    modelValue: "server_push_interval"
  }
]


</script>

<script lang="ts">
export default {
  name: 'BackendConfig'
}
</script>

<template>
  <n-card class="root" :embedded="true" title="后端设置">
    <!--    <div class="item">-->
    <!--        <span class="l-content">-->
    <!--          <div class="describe">-->
    <!--            <p class="title">通讯密钥</p>-->
    <!--            <p class="shallow">与节点通讯的密钥，以便数据不会被他人获取。</p>-->
    <!--          </div>-->
    <!--        </span>-->
    <!--      <span class="r-content">-->
    <!--          <n-input-->
    <!--              size="large"-->
    <!--              placeholder="b5#*(PG*(3tgf4e^!$P$Gp$G789"-->
    <!--              v-model:value="settingStore.settings.server.server_token"-->
    <!--              @blur="settingStore.saveOption('server','server_token', settingStore.settings.server.server_token)"-->
    <!--          ></n-input>-->
    <!--        </span>-->
    <!--    </div>-->

    <!--    <div class="item">-->
    <!--        <span class="l-content">-->
    <!--          <div class="describe">-->
    <!--            <p class="title">节点拉取动作轮询间隔</p>-->
    <!--            <p class="shallow">节点从面板获取数据的间隔频率。</p>-->
    <!--          </div>-->
    <!--        </span>-->
    <!--      <span class="r-content">-->
    <!--          <n-input-number-->
    <!--              size="large"-->
    <!--              placeholder="https://x.com/logo.jpeg"-->
    <!--              v-model:value.number="settingStore.settings.server.server_pull_interval"-->
    <!--              @blur="settingStore.saveOption('server','server_pull_interval', settingStore.settings.server.server_pull_interval)"-->
    <!--          ></n-input-number>-->
    <!--        </span>-->
    <!--    </div>-->

    <!--    <div class="item">-->
    <!--        <span class="l-content">-->
    <!--          <div class="describe">-->
    <!--            <p class="title">节点推送动作轮询间隔</p>-->
    <!--            <p class="shallow">节点推送数据到面板的间隔频率。</p>-->
    <!--          </div>-->
    <!--        </span>-->
    <!--      <span class="r-content">-->
    <!--          <n-input-number-->
    <!--              size="large"-->
    <!--              placeholder="https://x.com/logo.jpeg"-->
    <!--              v-model:value.number="settingStore.settings.server.server_push_interval"-->
    <!--              @blur="settingStore.saveOption('server','server_push_interval', settingStore.settings.server.server_push_interval)"-->
    <!--          ></n-input-number>-->
    <!--        </span>-->
    <!--    </div>-->


    <div v-for="(item, index) in settingsData" :key="index" class="item">
      <span class="l-content">
        <div class="describe">
          <p class="title">{{ item.title }}</p>
          <p class="shallow">{{ item.description }}</p>
        </div>
      </span>
      <span v-if="item.type === 'input'" class="r-content">
        <n-input
            size="large"
            :placeholder="item.placeholder"
            v-model:value="settingStore.settings.server[item.modelValue]"
            @blur="settingStore.saveOption('server', item.modelValue, settingStore.settings.server[item.modelValue])"
        ></n-input>
      </span>
      <span v-if="item.type === 'input-number'" class="r-content">
        <n-input-number
            size="large"
            :placeholder="item.placeholder"
            v-model:value.number="settingStore.settings.server[item.modelValue]"
            @blur="settingStore.saveOption('server', item.modelValue, settingStore.settings.server[item.modelValue])"
        ></n-input-number>
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