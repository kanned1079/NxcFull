<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { ref } from "vue";
import { useMessage } from "naive-ui";
import useThemeStore from "@/stores/useThemeStore";
import useSettingStore from "@/stores/useSettingStore";

const { t } = useI18n();
const message = useMessage();
const themeStore = useThemeStore();
const settingStore = useSettingStore();

interface ColorItem {
  label: string,
  value: string,
  disabled: boolean,
}

const options = ref<ColorItem[]>([
  { label: t('adminViews.systemConfig.frontend.themePropColor.default'), value: "defaultDay", disabled: false },
  { label: t('adminViews.systemConfig.frontend.themePropColor.darkBlueDay'), value: "darkBlueDay", disabled: false },
  { label: t('adminViews.systemConfig.frontend.themePropColor.milkGreenDay'), value: "milkGreenDay", disabled: false },
  { label: t('adminViews.systemConfig.frontend.themePropColor.bambooGreen'), value: "bambooGreen", disabled: false },
  { label: t('adminViews.systemConfig.frontend.themePropColor.mistyPine'), value: "mistyPine", disabled: false },
  { label: t('adminViews.systemConfig.frontend.themePropColor.glacierBlue'), value: "glacierBlue", disabled: false },
]);

type FrontendModel =
    | "frontend_theme_sidebar"
    | "frontend_theme_header"
    | "frontend_theme"
    | "frontend_background_url";

interface FrontendSetting {
  title: string;
  shallow: string;
  model: FrontendModel;
  placeholder?: string;
  type: string
}

// 配置表单项
const settings: FrontendSetting[] = [
  {
    title: 'adminViews.systemConfig.frontend.sidebarStyle.title',
    shallow: 'adminViews.systemConfig.frontend.sidebarStyle.shallow',
    type: 'switch',
    model: 'frontend_theme_sidebar',
  },
  {
    title: 'adminViews.systemConfig.frontend.headerStyle.title',
    shallow: 'adminViews.systemConfig.frontend.headerStyle.shallow',
    type: 'switch',
    model: 'frontend_theme_header',
  },
  {
    title: 'adminViews.systemConfig.frontend.themeColor.title',
    shallow: 'adminViews.systemConfig.frontend.themeColor.shallow',
    type: 'select',
    model: 'frontend_theme',
  },
  {
    title: 'adminViews.systemConfig.frontend.background.title',
    shallow: 'adminViews.systemConfig.frontend.background.shallow',
    type: 'input',
    model: 'frontend_background_url',
    placeholder: 'adminViews.systemConfig.frontend.background.placeholder',
  },
];

// TODO: 完善全球化 日文英文翻譯 保存函數統一使用saveField()方法

const saveFiled = async (k: string, v: any) => {
  let updateResponse = await settingStore.saveOption('frontend', k, v)
  if (updateResponse.code === 200)
    message.success(t('adminViews.systemConfig.common.success'))
  else
    message.error(t('adminViews.systemConfig.common.err') + updateResponse.msg || '')
}

</script>

<script lang="ts">
export default {
  name: 'Personalization',
}
</script>

<template>
  <div class="root">
    <n-card
        :embedded="true"
        class="security-panel"
        :title="t('adminViews.systemConfig.frontend.common.title')"
        :bordered="false"
    >
      <div v-for="setting in settings" :key="setting.model" class="item">
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
              v-model:value="settingStore.settings.frontend[setting.model]"
              @blur="saveFiled(setting.model, settingStore.settings.frontend[setting.model])"
          />
        </span>

        <span class="r-content" v-if="setting.type === 'input-number'">
          <n-input-number
              size="large"
              :placeholder="t(setting.placeholder || '')"
              v-model:value.number="settingStore.settings.frontend[setting.model]"
              @blur="saveFiled(setting.model, settingStore.settings.frontend[setting.model])"
          />
        </span>

        <span class="r-content" v-if="setting.type === 'select'">
          <n-select
              size="large"
              :options="options"
              :placeholder="t(setting.placeholder || '')"
              v-model:value="settingStore.settings.frontend[setting.model]"
              @update:value="saveFiled(setting.model, settingStore.settings.frontend[setting.model])"
          />
        </span>

        <span class="r-content to-right" v-if="setting.type === 'switch'">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.frontend[setting.model]"
              @update:value="saveFiled(setting.model, settingStore.settings.frontend[setting.model])"
          />
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


</style>