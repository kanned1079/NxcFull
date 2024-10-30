<script setup lang="ts">
import {onMounted, ref, computed} from "vue"
import {useI18n} from "vue-i18n";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";
import {useMessage} from "naive-ui";
import instance from "@/axios";

const {t} = useI18n()
const message = useMessage()
const userInfoStore = useUserInfoStore()
const themeStore = useThemeStore()

interface TicketItem {
  id: number  // 工单Id
  subject: string // 工单主题
  urgency: number // 紧急程度 1低 2中 3高
  status: number // true工单未关闭 false工单关闭
  created_at: string  // 工单创建时间
  last_reply: string // 最后一次回复
}

let ticketList = ref<TicketItem[]>([])

// 点击打开工单后 单独打开一个聊天窗口
let openTicket = (ticket: TicketItem) => {
  message.info(`查看工单：${ticket.title}`);
  const chatDialogWindow = window.open(
      `http://localhost:5173/user/tickets/chat?user_id=${userInfoStore.thisUser.id}&ticket_id=${ticket.id}&subject=${ticket.subject}&status=${ticket.status}&role=1`,
      '111111',
      'width=480,height=640,resizable=yes,scrollbars=yes',
  );
  chatDialogWindow.addEventListener('onload', () => {
    chatDialogWindow.postMessage({greeting: 'Hello from the parent window!'}, '*'); // * 表示接受来自任何来源的消息
  });

}

let getAllTicket = async () => {
  try {
    let {data} = await instance.get('/api/admin/v1/ticket', {
      params: {
        page: 1,
        size: 10,
      }
    })
    if (data.code === 200) {


    }
  } catch (error: any) {
    console.log(error)
  }
}


onMounted(() => {
  themeStore.menuSelected = 'ticket-mgr'
  themeStore.contentPath = '/admin/dashboard/ticket'

  // getAllMyTickets()
})

</script>

<script lang="ts">
export default {
  name: 'TicketMgr',
}
</script>

<template>
  <div class="root">
    <n-card class="title-card" :embedded="true" :bordered="false" hoverable :title="t('adminTicket.ticketMgr')">
    </n-card>

    <n-card :embedded="true" :bordered="false" hoverable content-style="padding: 0">
      1111
    </n-card>

  </div>
</template>

<style scoped lang="less">
.root {
  margin: 20px;
}
</style>