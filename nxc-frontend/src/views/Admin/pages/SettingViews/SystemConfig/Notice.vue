<script setup lang="ts" name="Notice">
import {defineComponent} from 'vue'
import useThemeStore from '@/stores/useThemeStore'
import useSettingStore from "@/stores/useSettingStore";
const themeStore = useThemeStore()
const settingStore = useSettingStore()

</script>

<template>
  <n-card :embedded="true" class="root" title="通知">
<!--    <n-alert type="warning" style="margin-bottom: 30px" :bordered="false">-->
<!--      用于自有客户端(APP)的版本管理及更新-->
<!--    </n-alert>-->
    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">显示名称</p>
            <p class="shallow">仅用于前端页面显示</p>
          </div>
        </span>
      <span class="r-content">
          <n-input
              type="text"
              placeholder="通用通知接口1"
              size="large"
              v-model:value="settingStore.settings.notice.notice_name"
              @blur="settingStore.saveOption('notice','notice_name', settingStore.settings.notice.notice_name)"
          />
        </span>
    </div>

    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">Bark接入点</p>
            <p class="shallow">Bark服务器后端API地址</p>
          </div>
        </span>
      <span class="r-content">
          <n-input
              type="text"
              placeholder="https://<ip>:<port>/<secret-key>"
              size="large"
              v-model:value="settingStore.settings.notice.bark_host"
              @blur="settingStore.saveOption('notice','bark_host', settingStore.settings.notice.bark_host)"
          />
        </span>
    </div>

    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">Bark群组</p>
            <p class="shallow">客户端显示的群组名称</p>
          </div>
        </span>
      <span class="r-content">
          <n-input
              type="text"
              placeholder="web"
              size="large"
              v-model:value="settingStore.settings.notice.bark_group"
              @blur="settingStore.saveOption('notice','bark_group', settingStore.settings.notice.bark_group)"
          />
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