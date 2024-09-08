<script setup lang="ts" name="WeApp">
import {defineComponent} from 'vue'
import useThemeStore from '@/stores/useThemeStore'
import useSettingStore from "@/stores/useSettingStore";
const settingStore = useSettingStore();
const themeStore = useThemeStore()

</script>

<template>
  <n-card :embedded="true" class="root" title="APP">
    <n-alert type="info" style="margin-bottom: 30px" :bordered="false">
      用于自有客户端(APP)的版本管理及更新
    </n-alert>
    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">Windows</p>
            <p class="shallow">Windows端版本号及下载地址。</p>
          </div>
        </span>
      <span class="r-content">
          <n-input type="text" placeholder="https://xxxx.com/xxx.exe" size="large" v-model:value="settingStore.settings.myapp.win_download"/>
        </span>
    </div>

    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">macOS</p>
            <p class="shallow">macOS端版本号及下载地址</p>
          </div>
        </span>
      <span class="r-content">
          <n-input type="text" placeholder="https://xxxx.com/xxx.dmg" size="large" v-model:value="settingStore.settings.myapp.osx_download"/>
        </span>
    </div>

    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">Android</p>
            <p class="shallow">Android端版本号及下载地址</p>
          </div>
        </span>
      <span class="r-content">
          <n-input type="text" placeholder="https://xxxx.com/xxx.apk" size="large" v-model:value="settingStore.settings.myapp.android_download"/>
        </span>
    </div>

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
.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}
</style>