<script setup lang="ts">
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
              <n-card
                  hoverable
                  :embedded="true"
                  :bordered="false"
                  content-style="padding: 0;"
                  style="background: linear-gradient(to bottom right, rgb(211, 149, 155), rgb(191, 230, 186)); height: 500px;"
              >
                <div style="width: 100%; margin-top: 100px; display: flex; flex-direction: row; justify-content: center">
                  <n-card style="width: 300px; height: 400px; border-radius: 10px 10px 0 0; opacity: 0.9; backdrop-filter: blur(10px)" content-style="padding: 0;" hoverable>
                      <div style="display: flex; flex-direction: row; justify-content: space-between; align-items: center; margin-top: 10px">
                        <div style="opacity: 0.9">
                          <p style="font-size: 0.8rem; margin: 0 0 0 15px;">10:13</p>
                        </div>

                        <div style="margin-right: 15px; opacity: 0.6; display: flex; gap: 5px" >
<!--                          <n-icon size="15">-->
<!--                            <dataIcon />-->
<!--                          </n-icon>-->
                          <div style="width: 16px; display: flex; align-items: center"><dataIcon /></div>
                          <div style="width: 16px; display: flex; align-items: center"><wifiIcon /></div>
                          <div style="width: 20px; display: flex; align-items: center"><batteryIcon /></div>

                        </div>
<!--                        <div style="margin-right: 10px; align-content: center">-->
<!--                          -->
<!--                        </div>-->

                      </div>
                    <div style="margin-top: 30px; padding: 20px">
                      <n-skeleton :width="100" :height="20"></n-skeleton>
                      <n-skeleton :width="150" :height="12" style="margin-top: 10px"></n-skeleton>
                      <n-skeleton :height="10" style="margin-top: 50px"></n-skeleton>
                      <n-skeleton :height="10" style="margin-top: 10px"></n-skeleton>
                      <n-skeleton :height="10" style="margin-top: 10px"></n-skeleton>
                      <n-skeleton :height="10" style="margin-top: 10px"></n-skeleton>
                    </div>
                  </n-card>
                </div>
              </n-card>
            </n-grid-item>
            <n-grid-item>
              <n-card
                  hoverable
                  :embedded="true"
                  :bordered="false"
                  content-style="padding: 0;"
                  style="background: linear-gradient(to bottom right, rgb(185, 147, 214), rgb(140, 166, 219)); height: 500px;"
              >
                <div style="width: 100%; margin-top: 100px; display: flex; flex-direction: row; justify-content: center">
                  <n-card style="width: 400px; border-radius: 10px"  content-style="padding: 0;">
                    <div style="display: flex; flex-direction: row; justify-content: space-between; align-items: center; margin-top: 10px">
                      <div style="opacity: 0.9; display: flex; margin-left: 15px">
<!--                        <p style="font-size: 0.8rem; margin: 0 0 0 15px;">10:13</p>-->
                        <div style="width: 10px; height: 10px; border-radius: 50%; background-color: #ff6666; margin-right: 5px"></div>
                        <div style="width: 10px; height: 10px; border-radius: 50%; background-color: #ffcc66; margin-right: 5px"></div>
                        <div style="width: 10px; height: 10px; border-radius: 50%; background-color: #66cc66; margin-right: 5px"></div>

                      </div>

                      <n-flex style="margin-right: 15px; opacity: 0.6" >
<!--                       <div style="width: 10px; height: 10px; border-radius: 50%; background-color: rgba(51,51,51,0.7)"></div>-->
                      </n-flex>
                      <!--                        <div style="margin-right: 10px; align-content: center">-->
                      <!--                          -->
                      <!--                        </div>-->

                    </div>
                    <div style="margin-top: 30px; padding: 20px">
                      <n-skeleton :width="100" :height="20"></n-skeleton>
                      <n-skeleton :width="150" :height="12" style="margin-top: 10px"></n-skeleton>
                      <n-skeleton :height="10" style="margin-top: 50px"></n-skeleton>
                      <n-skeleton :height="10" style="margin-top: 10px"></n-skeleton>
                      <n-skeleton :height="10" style="margin-top: 10px"></n-skeleton>
                      <n-skeleton :height="10" style="margin-top: 10px"></n-skeleton>
                    </div>
                  </n-card>
                </div>
              </n-card>
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