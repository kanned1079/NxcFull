<script setup lang="ts">
import {onMounted, onBeforeMount, ref} from "vue";
import {useRouter} from "vue-router";
import useUserInfoStore from "@/stores/useUserInfoStore";
import Settings from '@/views/Admin/pages/RootViews/dashboard/Settings.vue'
import useThemeStore from "@/stores/useThemeStore";
import Income from '@/views/Admin/pages/RootViews/dashboard/Income.vue'
import IncomeChart from "@/views/Admin/pages/RootViews/dashboard/IncomeChart.vue";
import instance from "@/axios";
import {ChevronForwardOutline as toRightIcon,} from "@vicons/ionicons5"

const userInfoStore = useUserInfoStore();
const themeStore = useThemeStore();
const router = useRouter();

let animated = ref<boolean>(false)

let haveOpenTickets = ref<boolean>(false)
let openingTicketCount = ref<number>(0)

let checkIsUserHaveOpenTickets = async () => {
  try {
    let {data} = await instance.get('/api/user/v1/ticket/check', {
      params: {
        user_id: userInfoStore.thisUser.id,
      }
    })
    if (data.code === 200) {
      // thisNotices.value = data.notices
      // data.notices.forEach((notice: Notice) => thisNotices.value.push(notice))
      haveOpenTickets.value = data.exist || false
      openingTicketCount.value = data.ticket_count || 0
    }
  } catch (error) {
    console.log(error)
  }
}



onBeforeMount(async () => {
  // await getAppOverview()
  await checkIsUserHaveOpenTickets()
})

onMounted(() => {
  themeStore.contentPath = '/admin/dashboard/summary'
  themeStore.menuSelected = 'dashboard'

  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'Summary'
}
</script>

<template>
  <transition name="slide-fade">
    <div class="root" v-if="animated">

      <n-collapse-transition :show="haveOpenTickets">
        <n-alert
            :bordered="false"
            style="margin: 20px 20px 0 20px"
            title=""
            type="warning"
        >
          <div style="display: flex; flex-direction: row; align-items: center">
            {{ `您有${openingTicketCount}条待处理的工单` }}
            <n-button
                style="margin-left: 5px"
                text
                type="warning"
                @click="router.push({path: '/admin/dashboard/ticket'})"
            >
              去查看
              <n-icon>
                <toRightIcon/>
              </n-icon>
            </n-button>
          </div>
        </n-alert>
      </n-collapse-transition>



      <div class="setting-panel">
        <Settings></Settings>
      </div>
      <div class="income-chart">
        <IncomeChart></IncomeChart>
      </div>
    </div>
  </transition>
  <div style="height: 3rem"></div>
</template>

<style lang="less" scoped>
.root {
  width: 100%;

  .setting-panel {
    display: flex;
    align-items: center;
    line-height: 160px;
    justify-content: center;
  }

  .income-panel {
    display: flex;
  }

  .income-chart {
    display: flex;
    align-items: center;
    padding: 0;
    justify-content: center;
  }
}
//.n-card {
//  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
//  border: 0;
//}
</style>