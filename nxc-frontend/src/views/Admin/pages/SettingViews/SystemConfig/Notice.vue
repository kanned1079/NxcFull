<script setup lang="ts">
// import {defineComponent} from 'vue'
// import useThemeStore from '@/stores/useThemeStore'
import useSettingStore from "@/stores/useSettingStore";
import {useI18n} from "vue-i18n"
// const themeStore = useThemeStore()
const {t} = useI18n()
const settingStore = useSettingStore()

type NoticeModel =
    | "notice_name"
    | "bark_host"
    | "bark_group";

interface NoticeSetting {
  title: string;
  shallow: string;
  model: NoticeModel;
  placeholder?: string;
}

const noticeSettings: NoticeSetting[] = [
  {
    title: 'adminViews.systemConfig.notice.displayName.title',
    shallow: 'adminViews.systemConfig.notice.displayName.shallow',
    placeholder: 'adminViews.systemConfig.notice.displayName.placeholder',
    model: 'notice_name',
  },
  {
    title: 'adminViews.systemConfig.notice.barkEndpoint.title',
    shallow: 'adminViews.systemConfig.notice.barkEndpoint.shallow',
    placeholder: 'adminViews.systemConfig.notice.barkEndpoint.placeholder',
    model: 'bark_host',
  },
  {
    title: 'adminViews.systemConfig.notice.barkGroup.title',
    shallow: 'adminViews.systemConfig.notice.barkGroup.shallow',
    placeholder: 'adminViews.systemConfig.notice.barkGroup.placeholder',
    model: 'bark_group',
  },
];


// let noticeSettings: NoticeSetting[] = [
//   {
//     title: "显示名称",
//     shallow: "仅用于前端页面显示",
//     model: "notice_name",
//     placeholder: "通用通知接口1"
//   },
//   {
//     title: "Bark 接入点",
//     shallow: "Bark 服务器后端 API 地址",
//     model: "bark_host",
//     placeholder: "https://<ip>:<port>/<secret-key>"
//   },
//   {
//     title: "Bark 群组",
//     shallow: "客户端显示的群组名称",
//     model: "bark_group",
//     placeholder: "web"
//   }
// ]

</script>

<script lang="ts">
export default {
  name: 'Notice'
}
</script>

<template>
  <n-card
      hoverable
      :embedded="true"
      class="root"
      :title="t('adminViews.systemConfig.notice.common.title')"
      :bordered="false"
  >


    <!--    <div v-for="setting in noticeSettings" :key="setting.title" class="item">-->
    <!--      <span class="l-content">-->
    <!--        <div class="describe">-->
    <!--          <p class="title">{{ setting.title }}</p>-->
    <!--          <p class="shallow">{{ setting.shallow }}</p>-->
    <!--        </div>-->
    <!--      </span>-->
    <!--      <span class="r-content">-->
    <!--        <n-input-->
    <!--            type="text"-->
    <!--            placeholder="{{ setting.placeholder }}"-->
    <!--            size="large"-->
    <!--            v-model:value="settingStore.settings.notice[setting.model]"-->
    <!--            @blur="settingStore.saveOption('notice', setting.model, settingStore.settings.notice[setting.model])"-->
    <!--        />-->
    <!--      </span>-->
    <!--    </div>-->

    <div v-for="setting in noticeSettings" :key="setting.title" class="item">
  <span class="l-content">
    <div class="describe">
      <p class="title">{{ t(setting.title) }}</p>
      <p class="shallow">{{ t(setting.shallow) }}</p>
    </div>
  </span>
      <span class="r-content">
    <n-input
        type="text"
        :placeholder="t(setting.placeholder || '')"
        size="large"
        v-model:value="settingStore.settings.notice[setting.model]"
        @blur="settingStore.saveOption('notice', setting.model, settingStore.settings.notice[setting.model])"
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

//.n-card {
//  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
//  border: 0;
//}
</style>