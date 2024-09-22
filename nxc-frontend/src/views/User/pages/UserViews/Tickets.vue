<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {ref, onMounted} from "vue"
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";

const {t} = useI18n();
const userInfoStore = useUserInfoStore();
const themeStore = useThemeStore();

let dialogActive = ref<boolean>(false)

// 表格的字段名
// 主题 工单级别 工单状态 创建时间最后回复 操作（查看工单/关闭工单）
interface TicketItem {
  id: number  // 工单Id
  title: string // 工单主题
  urgency: number // 紧急程度 1低 2中 3高
  status: boolean // true工单未关闭 false工单关闭
  created_at: string  // 工单创建时间
  last_response: string // 最后一次回复
}

// 点击打开工单后 单独打开一个聊天窗口
let openTicket = () => {

}

// ticketList 将从服务器获取
let ticketList = ref<TicketItem[]>([])


onMounted(() => {
  themeStore.menuSelected = ''
  themeStore.userPath = '/dashboard/tickets'
})

</script>

<script lang="ts">
export default {
  name: 'Tickets',
}
</script>

<template>
<div class="root">
  <n-card class="tickets-history" :embedded="true" hoverable :title="t('tickets.userTickets')">
    <n-button type="primary" @click="dialogActive = true">{{ t('tickets.addTicket') }}</n-button>
<!--    表格在这里-->
  </n-card>
</div>
</template>

<style scoped lang="less">
.root {
  padding: 20px;
}

.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}
</style>