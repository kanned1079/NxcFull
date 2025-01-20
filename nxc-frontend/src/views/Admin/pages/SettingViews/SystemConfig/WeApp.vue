<script setup lang="ts">
import useSettingStore from "@/stores/useSettingStore";
import {useI18n} from "vue-i18n"
import {useMessage} from "naive-ui"

const {t} = useI18n()
const settingStore = useSettingStore();
// const themeStore = useThemeStore()
const message = useMessage()

type AppModel =
    | "download_enabled"
    | "win_download"
    | "osx_download"
    | "android_download";

interface AppSetting {
  type: string
  title: string;
  shallow: string;
  model: AppModel;
  placeholder?: string;
}

// const appDownloadSettings: AppSetting[] = [
//   {
//     title: "Windows",
//     shallow: "Windows 端版本号及下载地址。",
//     model: "win_download",
//     placeholder: "https://xxxx.com/xxx.exe"
//   },
//   {
//     title: "macOS",
//     shallow: "macOS 端版本号及下载地址",
//     model: "osx_download",
//     placeholder: "https://xxxx.com/xxx.dmg"
//   },
//   {
//     title: "Android",
//     shallow: "Android 端版本号及下载地址",
//     model: "android_download",
//     placeholder: "https://xxxx.com/xxx.apk"
//   }
// ]

const appDownloadSettings: AppSetting[] = [
  {
    type: 'switch',
    title: 'adminViews.systemConfig.appDownload.enabled.title',
    shallow: 'adminViews.systemConfig.appDownload.enabled.shallow',
    model: 'download_enabled'
  },
  {
    type: 'input',
    title: 'adminViews.systemConfig.appDownload.windows.title',
    shallow: 'adminViews.systemConfig.appDownload.windows.shallow',
    placeholder: 'adminViews.systemConfig.appDownload.windows.placeholder',
    model: 'win_download',
  },
  {
    type: 'input',
    title: 'adminViews.systemConfig.appDownload.macos.title',
    shallow: 'adminViews.systemConfig.appDownload.macos.shallow',
    placeholder: 'adminViews.systemConfig.appDownload.macos.placeholder',
    model: 'osx_download',
  },
  {
    type: 'input',
    title: 'adminViews.systemConfig.appDownload.android.title',
    shallow: 'adminViews.systemConfig.appDownload.android.shallow',
    placeholder: 'adminViews.systemConfig.appDownload.android.placeholder',
    model: 'android_download',
  },
];

const saveFiled = async (k: string, v: any) => {
  let updateResponse = await settingStore.saveOption('my_app', k, v)
  if (updateResponse.code === 200)
    message.success(t('adminViews.systemConfig.common.success'))
  else
    message.error(t('adminViews.systemConfig.common.err') + updateResponse.msg || '')
}

</script>

<script lang="ts">
export default {
  name: 'WeApp'
}
</script>

<template>
  <n-card
      :embedded="true"
      class="root"
      :title="t('adminViews.systemConfig.appDownload.common.title')"
      :bordered="false"
  >
    <n-alert type="info" style="margin-bottom: 30px" :bordered="false">
      {{ t('adminViews.systemConfig.appDownload.common.hint') }}
    </n-alert>

    <div v-for="setting in appDownloadSettings" :key="setting.title" class="item">
      <span class="l-content">
        <div class="describe">
          <p class="title">{{ t(setting.title) }}</p>
          <p class="shallow">{{ t(setting.shallow) }}</p>
        </div>
     </span>

      <span v-if="setting.type === 'switch'" class="r-content to-right">
      <n-switch
          size="medium"
          v-model:value="settingStore.settings.my_app[setting.model]"
          @update:value="saveFiled(setting.model, settingStore.settings.my_app[setting.model])"
      ></n-switch>
      </span>

      <span v-if="setting.type === 'input'" class="r-content">
      <n-input
          type="text"
          :placeholder="t(setting.placeholder || '')"
          size="large"
          v-model:value="settingStore.settings.my_app[setting.model]"
          @blur="saveFiled(setting.model, settingStore.settings.my_app[setting.model])"
      />
  </span>
    </div>


    <!--    <div v-for="setting in appDownloadSettings" :key="setting.title" class="item">-->
    <!--    <span class="l-content">-->
    <!--      <div class="describe">-->
    <!--        <p class="title">{{ setting.title }}</p>-->
    <!--        <p class="shallow">{{ setting.shallow }}</p>-->
    <!--      </div>-->
    <!--    </span>-->
    <!--      <span class="r-content">-->
    <!--      <n-input-->
    <!--          type="text"-->
    <!--          placeholder="{{ setting.placeholder }}"-->
    <!--          size="large"-->
    <!--          v-model:value="settingStore.settings.my_app[setting.model]"-->
    <!--          @blur="settingStore.saveOption('my_app', setting.model, settingStore.settings.my_app[setting.model])"-->
    <!--      />-->
    <!--    </span>-->
    <!--    </div>-->

  </n-card>
</template>

<style lang="less" scoped>
.root {
  min-width: 900px;

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

.to-right {
  text-align: right;
}

//.n-card {
//  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
//  border: 0;
//}
</style>