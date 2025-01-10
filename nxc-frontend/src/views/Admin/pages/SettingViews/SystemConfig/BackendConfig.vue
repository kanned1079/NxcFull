<script setup lang="ts">
// import useThemeStore from "@/stores/useThemeStore"
import useSettingStore from "@/stores/useSettingStore";

// const themeStore = useThemeStore()
const settingStore = useSettingStore()

type BackendModel =
    | "app_sub_description"
    | "why_choose_us_hint"
    | "bilibili_official_link"
    | "youtube_official_link"
    | "instagram_link"
    | "wechat_official_link"
    | "filing_number"
    | "page_suffix";

interface BackendSetting {
  title: string;
  description: string;
  modelValue: BackendModel;
  type: string;
  placeholder?: string;
}

let settingsData: BackendSetting[] = [
  {
    title: "首页描述",
    description: "设置首页的简要描述内容。",
    type: "input",
    placeholder: "请输入首页描述内容",
    modelValue: "app_sub_description"
  },
  {
    title: "为什么选择我们",
    description: "设置关于为什么选择我们的描述。",
    type: "input",
    placeholder: "请输入详细描述",
    modelValue: "why_choose_us_hint"
  },
  {
    title: "Bilibili 官方链接",
    description: "设置 Bilibili 官方账号的链接地址。",
    type: "input",
    placeholder: "https://space.bilibili.com/xxxx",
    modelValue: "bilibili_official_link"
  },
  {
    title: "YouTube 官方链接",
    description: "设置 YouTube 官方账号的链接地址。",
    type: "input",
    placeholder: "https://youtube.com/channel/xxxx",
    modelValue: "youtube_official_link"
  },
  {
    title: "Instagram 官方链接",
    description: "设置 Instagram 官方账号的链接地址。",
    type: "input",
    placeholder: "https://instagram.com/xxxx",
    modelValue: "instagram_link"
  },
  {
    title: "微信公众账号链接",
    description: "设置微信公众账号的链接地址。",
    type: "input",
    placeholder: "请输入微信公众链接",
    modelValue: "wechat_official_link"
  },
  {
    title: "备案号",
    description: "设置站点的备案号。",
    type: "input",
    placeholder: "如：粤ICP备12345678号",
    modelValue: "filing_number"
  },
  {
    title: "站点后缀",
    description: "设置站点名称后缀，用于标题显示。",
    type: "input",
    placeholder: "如：- 你的站点名称",
    modelValue: "page_suffix"
  }
];


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
            v-model:value="settingStore.settings.welcome[item.modelValue]"
            @blur="settingStore.saveOption('welcome', item.modelValue, settingStore.settings.welcome[item.modelValue])"
        ></n-input>
      </span>
      <span v-if="item.type === 'input-number'" class="r-content">
        <n-input-number
            size="large"
            :placeholder="item.placeholder"
            v-model:value.number="settingStore.settings.welcome[item.modelValue]"
            @blur="settingStore.saveOption('welcome', item.modelValue, settingStore.settings.welcome[item.modelValue])"
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