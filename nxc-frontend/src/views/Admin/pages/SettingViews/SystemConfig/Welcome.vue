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
    type: 'input-area',
    placeholder: 'adminViews.systemConfig.welcome.homeDescription.placeholder',
    modelValue: 'app_sub_description',
  },
  {
    title: 'adminViews.systemConfig.welcome.whyChooseUs.title',
    description: 'adminViews.systemConfig.welcome.whyChooseUs.description',
    type: 'input-area',
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

</script>

<script lang="ts">
export default {
  name: 'BackendConfig'
}
</script>

<template>
  <n-card
      class="root security-panel"
      :embedded="true"
      :bordered="false"
      :title="t('adminViews.systemConfig.welcome.common.title')"
  >
    <n-grid cols="1" responsive="screen" y-gap="16">
      <n-grid-item v-for="(item, index) in settingsData" :key="index" class="grid-item">
        <n-grid cols="1 s:1 m:2 l:2" responsive="screen" align-items="center">
          <!-- 左侧：标题 + 描述 -->
          <n-grid-item>
            <div class="describe">
              <p class="title">{{ t(item.title) }}</p>
              <p class="shallow">{{ t(item.description) }}</p>
            </div>
          </n-grid-item>

          <!-- 右侧：输入框 / 数字输入框 -->
          <n-grid-item>
            <n-input
                v-if="item.type === 'input'"
                size="large"
                :placeholder="t(item.placeholder || '')"
                v-model:value="settingStore.settings.welcome[item.modelValue]"
                @blur="settingStore.saveOption('welcome', item.modelValue, settingStore.settings.welcome[item.modelValue])"
            />
            <n-input
            v-if="item.type === 'input-area'"
            type="textarea"
            :show-count="true"
            size="large"
            rows="2"
            :placeholder="t(item.placeholder || '')"
            v-model:value="settingStore.settings.welcome[item.modelValue]"
            @blur="settingStore.saveOption('welcome', item.modelValue, settingStore.settings.welcome[item.modelValue])"
            >
            </n-input>
            <n-input-number
                v-else-if="item.type === 'input-number'"
                size="large"
                :placeholder="t(item.placeholder || '')"
                v-model:value.number="settingStore.settings.welcome[item.modelValue]"
                @blur="settingStore.saveOption('welcome', item.modelValue, settingStore.settings.welcome[item.modelValue])"
            />
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

.security-panel {
  .grid-item {
    margin-bottom: 16px;
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
}

.to-right {
  text-align: right;
}
</style>