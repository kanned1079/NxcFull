<script setup lang="ts">
// import useThemeStore from "@/stores/useThemeStore"
import useSettingStore from "@/stores/useSettingStore";
import {useI18n} from "vue-i18n";


const {t} = useI18n()
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

interface WelcomeSetting {
  title: string;
  description: string;
  modelValue: BackendModel;
  type: string;
  placeholder?: string;
}

let settingsData: WelcomeSetting[] = [
  {
    title: 'adminViews.systemConfig.welcome.homeDescription.title',
    description: 'adminViews.systemConfig.welcome.homeDescription.description',
    type: 'input',
    placeholder: 'adminViews.systemConfig.welcome.homeDescription.placeholder',
    modelValue: 'app_sub_description',
  },
  {
    title: 'adminViews.systemConfig.welcome.whyChooseUs.title',
    description: 'adminViews.systemConfig.welcome.whyChooseUs.description',
    type: 'input',
    placeholder: 'adminViews.systemConfig.welcome.whyChooseUs.placeholder',
    modelValue: 'why_choose_us_hint',
  },
  {
    title: 'adminViews.systemConfig.welcome.bilibiliLink.title',
    description: 'adminViews.systemConfig.welcome.bilibiliLink.description',
    type: 'input',
    placeholder: 'adminViews.systemConfig.welcome.bilibiliLink.placeholder',
    modelValue: 'bilibili_official_link',
  },
  {
    title: 'adminViews.systemConfig.welcome.youtubeLink.title',
    description: 'adminViews.systemConfig.welcome.youtubeLink.description',
    type: 'input',
    placeholder: 'adminViews.systemConfig.welcome.youtubeLink.placeholder',
    modelValue: 'youtube_official_link',
  },
  {
    title: 'adminViews.systemConfig.welcome.instagramLink.title',
    description: 'adminViews.systemConfig.welcome.instagramLink.description',
    type: 'input',
    placeholder: 'adminViews.systemConfig.welcome.instagramLink.placeholder',
    modelValue: 'instagram_link',
  },
  {
    title: 'adminViews.systemConfig.welcome.wechatLink.title',
    description: 'adminViews.systemConfig.welcome.wechatLink.description',
    type: 'input',
    placeholder: 'adminViews.systemConfig.welcome.wechatLink.placeholder',
    modelValue: 'wechat_official_link',
  },
  {
    title: 'adminViews.systemConfig.welcome.filingNumber.title',
    description: 'adminViews.systemConfig.welcome.filingNumber.description',
    type: 'input',
    placeholder: 'adminViews.systemConfig.welcome.filingNumber.placeholder',
    modelValue: 'filing_number',
  },
  {
    title: 'adminViews.systemConfig.welcome.pageSuffix.title',
    description: 'adminViews.systemConfig.welcome.pageSuffix.description',
    type: 'input',
    placeholder: 'adminViews.systemConfig.welcome.pageSuffix.placeholder',
    modelValue: 'page_suffix',
  },
];


// let settingsData: WelcomeSetting[] = [
//   {
//     title: "首页描述",
//     description: "设置首页的简要描述内容。",
//     type: "input",
//     placeholder: "请输入首页描述内容",
//     modelValue: "app_sub_description"
//   },
//   {
//     title: "为什么选择我们",
//     description: "设置关于为什么选择我们的描述。",
//     type: "input",
//     placeholder: "请输入详细描述",
//     modelValue: "why_choose_us_hint"
//   },
//   {
//     title: "Bilibili 官方链接",
//     description: "设置 Bilibili 官方账号的链接地址。",
//     type: "input",
//     placeholder: "https://space.bilibili.com/xxxx",
//     modelValue: "bilibili_official_link"
//   },
//   {
//     title: "YouTube 官方链接",
//     description: "设置 YouTube 官方账号的链接地址。",
//     type: "input",
//     placeholder: "https://youtube.com/channel/xxxx",
//     modelValue: "youtube_official_link"
//   },
//   {
//     title: "Instagram 官方链接",
//     description: "设置 Instagram 官方账号的链接地址。",
//     type: "input",
//     placeholder: "https://instagram.com/xxxx",
//     modelValue: "instagram_link"
//   },
//   {
//     title: "微信公众账号链接",
//     description: "设置微信公众账号的链接地址。",
//     type: "input",
//     placeholder: "请输入微信公众链接",
//     modelValue: "wechat_official_link"
//   },
//   {
//     title: "备案号",
//     description: "设置站点的备案号。",
//     type: "input",
//     placeholder: "如：粤ICP备12345678号",
//     modelValue: "filing_number"
//   },
//   {
//     title: "站点后缀",
//     description: "设置站点名称后缀，用于标题显示。",
//     type: "input",
//     placeholder: "如：- 你的站点名称",
//     modelValue: "page_suffix"
//   }
// ];


</script>

<script lang="ts">
export default {
  name: 'BackendConfig'
}
</script>

<template>
  <n-card
      class="root"
      :embedded="true"
      :bordered="false"
      :title="t('adminViews.systemConfig.welcome.common.title')"
  >
    <!--    <div v-for="(item, index) in settingsData" :key="index" class="item">-->
    <!--      <span class="l-content">-->
    <!--        <div class="describe">-->
    <!--          <p class="title">{{ item.title }}</p>-->
    <!--          <p class="shallow">{{ item.description }}</p>-->
    <!--        </div>-->
    <!--      </span>-->
    <!--      <span v-if="item.type === 'input'" class="r-content">-->
    <!--        <n-input-->
    <!--            size="large"-->
    <!--            :placeholder="item.placeholder"-->
    <!--            v-model:value="settingStore.settings.welcome[item.modelValue]"-->
    <!--            @blur="settingStore.saveOption('welcome', item.modelValue, settingStore.settings.welcome[item.modelValue])"-->
    <!--        ></n-input>-->
    <!--      </span>-->
    <!--      <span v-if="item.type === 'input-number'" class="r-content">-->
    <!--        <n-input-number-->
    <!--            size="large"-->
    <!--            :placeholder="item.placeholder"-->
    <!--            v-model:value.number="settingStore.settings.welcome[item.modelValue]"-->
    <!--            @blur="settingStore.saveOption('welcome', item.modelValue, settingStore.settings.welcome[item.modelValue])"-->
    <!--        ></n-input-number>-->
    <!--      </span>-->
    <!--    </div>-->

    <div v-for="(item, index) in settingsData" :key="index" class="item">
      <span class="l-content">
       <div class="describe">
          <p class="title">{{ t(item.title) }}</p>
          <p class="shallow">{{ t(item.description) }}</p>
      </div>
  </span>
      <span v-if="item.type === 'input'" class="r-content">
    <n-input
        size="large"
        :placeholder="t(item.placeholder || '')"
        v-model:value="settingStore.settings.welcome[item.modelValue]"
        @blur="settingStore.saveOption('welcome', item.modelValue, settingStore.settings.welcome[item.modelValue])"
    ></n-input>
  </span>
      <span v-if="item.type === 'input-number'" class="r-content">
    <n-input-number
        size="large"
        :placeholder="t(item.placeholder || '')"
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