<script setup lang="ts">
import AppFrame from "@/views/utils/AppFrame.vue";
import {useI18n} from "vue-i18n";
import {useRoute} from "vue-router"
import {ref, computed, onMounted, onBeforeMount} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import {useMessage} from "naive-ui"
import {handleFetchAllAppDownloadList} from "@/api/user/app";
import {formatDate, formatTimestamp} from "../../../../utils/timeFormat";
import {BatteryHalfOutline as batteryIcon, Cellular as dataIcon, Wifi as wifiIcon} from "@vicons/ionicons5"

const {t} = useI18n()
const route = useRoute()
const message = useMessage()
const themeStore = useThemeStore()
let animated = ref<boolean>(false)
let downloadEnabled = ref<boolean>(false)

interface AppItem {
  id?: number
  platform: string
  link: string
}

let appDownloadList = ref<AppItem[]>([])

let callFetchAllDownloadLink = async () => {
  let data = await handleFetchAllAppDownloadList('en')
  if (data.code === 200) {
    console.log(data)
    downloadEnabled.value = data.download_enabled
    data.download_links.forEach((item: AppItem) => appDownloadList.value.push(item))
  } else {
    message.error('数据获取失败')
  }
}

onBeforeMount(() => {
  themeStore.menuSelected = 'user-app-download'
  themeStore.userPath = route.fullPath
})

onMounted(async () => {
  animated.value = true
  // await callFetchAllDownloadLink()
})

</script>

<script lang="ts">
export default {
  name: 'UserAppDownload',
}
</script>

<template>
<div class="root">
  <transition name="slide-fade">
    <div class="inner-root" v-if="animated">

      <n-alert
        type="warning"
        class="mention"
        v-if="!downloadEnabled"
      >
        很抱歉暂时还没有下载，请您稍后再试，如有问题请提交工单以联系我们的支持服务。
      </n-alert>

      <n-card
          v-if="downloadEnabled"
        hoverable
        :embedded="true"
        :bordered="false"
        v-for="item in appDownloadList"
        :title="item.platform"
        class="platform-card"
      >
        <n-tag
          type="success"
          >
          {{ item.link  }}
        </n-tag>
      </n-card>

<!--      <div style="display: flex; justify-content: center">-->

        <div
          style="width: 100%; display: flex; justify-content: center; margin-top: 30px"
        >
          <n-grid
              cols="1 s:1 m:2"
              responsive="screen"
              x-gap="20px"
              y-gap="20px"
              style="max-width: 1024px"
          >
            <n-grid-item>

              <AppFrame
                  platform="mobile"
              ></AppFrame>
            </n-grid-item>
            <n-grid-item>

              <AppFrame
                  platform="desktop"
              ></AppFrame>
            </n-grid-item>
          </n-grid>

        </div>





    </div>
  </transition>
</div>
</template>

<style scoped lang="less">
.root {
  padding: 20px;
  .inner-root {
    .mention {
      margin-bottom: 20px;
    }
    .platform-card {
      margin-bottom: 20px;
    }
  }
}
</style>