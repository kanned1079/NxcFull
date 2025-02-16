<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {ref} from "vue";
import {useMessage} from "naive-ui";
import useThemeStore from "@/stores/useThemeStore";
import useSettingStore from "@/stores/useSettingStore";
import {handleFetchRootRuntimeEnvConfig} from "@/api/common/env"

const {t} = useI18n();
const message = useMessage();
const themeStore = useThemeStore();
const settingStore = useSettingStore();

interface ColorItem {
  label: string,
  value: string,
  disabled: boolean,
}

const options = ref<ColorItem[]>([
  {label: t('adminViews.systemConfig.frontend.themePropColor.default'), value: "defaultDay", disabled: false},
  {label: t('adminViews.systemConfig.frontend.themePropColor.darkBlueDay'), value: "darkBlueDay", disabled: false},
  {label: t('adminViews.systemConfig.frontend.themePropColor.milkGreenDay'), value: "milkGreenDay", disabled: false},
  {label: t('adminViews.systemConfig.frontend.themePropColor.bambooGreen'), value: "bambooGreen", disabled: false},
  {label: t('adminViews.systemConfig.frontend.themePropColor.mistyPine'), value: "mistyPine", disabled: false},
  {label: t('adminViews.systemConfig.frontend.themePropColor.glacierBlue'), value: "glacierBlue", disabled: false},
  // grayTheme
  {label: t('adminViews.systemConfig.frontend.themePropColor.grayTheme'), value: "grayTheme", disabled: false},

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

const saveFiled = async (k: string, v: any) => {
  let updateResponse = await settingStore.saveOption('frontend', k, v)
  if (updateResponse.code === 200)
    message.success(t('adminViews.systemConfig.common.success'))
  else
    message.error(t('adminViews.systemConfig.common.err') + updateResponse.msg || '')
}

// const refreshTheme = () => location.reload()

const refreshTheme = async (fieldName: string) => {
  if (fieldName === 'frontend_theme') {
    await handleFetchRootRuntimeEnvConfig()
    // setTimeout(() => themeStore.setThemeFromSetting(), 1000)
    location.reload() // 直接刷新页面
  }
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
      <n-grid cols="1" responsive="screen" y-gap="16">
        <n-grid-item v-for="setting in settings" :key="setting.model" class="grid-item">
          <n-grid cols="1 s:1 m:2 l:2" responsive="screen" align-items="center">
            <!-- 左侧：标题 + 描述 -->
            <n-grid-item>
              <div class="describe">
                <p class="title">{{ t(setting.title) }}</p>
                <p class="shallow">{{ t(setting.shallow) }}</p>
              </div>
            </n-grid-item>

            <!-- 右侧：输入框 / 选择框 / 开关 -->
            <n-grid-item>
              <n-input
                  v-if="setting.type === 'input'"
                  size="large"
                  :placeholder="t(setting.placeholder || '')"
                  v-model:value="settingStore.settings.frontend[setting.model]"
                  @blur="saveFiled(setting.model, settingStore.settings.frontend[setting.model])"
              />
              <n-input-number
                  v-else-if="setting.type === 'input-number'"
                  size="large"
                  :placeholder="t(setting.placeholder || '')"
                  v-model:value.number="settingStore.settings.frontend[setting.model]"
                  @blur="saveFiled(setting.model, settingStore.settings.frontend[setting.model])"
              />
              <n-select
                  v-else-if="setting.type === 'select'"
                  size="large"
                  :options="options"
                  :placeholder="t(setting.placeholder || '')"
                  v-model:value="settingStore.settings.frontend[setting.model]"
                  @update:value="saveFiled(setting.model, settingStore.settings.frontend[setting.model]); refreshTheme(setting.model)"
              />
              <div v-else-if="setting.type === 'switch'" class="to-right">
                <n-switch
                    size="medium"
                    v-model:value="settingStore.settings.frontend[setting.model]"
                    @update:value="saveFiled(setting.model, settingStore.settings.frontend[setting.model])"
                />
              </div>
            </n-grid-item>
          </n-grid>
        </n-grid-item>
      </n-grid>
    </n-card>
  </div>
</template>

<style lang="less" scoped>
.root {
  margin: 0 auto;

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
}

.to-right {
  text-align: right;
}
</style>