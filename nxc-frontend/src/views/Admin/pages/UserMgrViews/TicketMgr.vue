<script setup lang="ts">
import {computed, onMounted, ref} from "vue"
import {useI18n} from "vue-i18n";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";
import {NButton, NTag, useMessage} from "naive-ui";
import {formatDate} from "@/utils/timeFormat";
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

let animated = ref<boolean>(false)

let ticketList = ref<TicketItem[]>([])
let pendingTicketList = ref<TicketItem[]>([])

// 点击打开工单后 单独打开一个聊天窗口
let openTicket = (ticket: TicketItem) => {
  message.info(`查看工单：${ticket.title}`);
  const chatDialogWindow = window.open(
      `http://localhost:5173/user/tickets/chat?user_id=${userInfoStore.thisUser.id}&ticket_id=${ticket.id}&subject=${ticket.subject}&status=${ticket.status}&role=0`,
      '111111',
      'width=480,height=640,resizable=yes,scrollbars=yes',
  );
  chatDialogWindow.addEventListener('onload', () => {
    chatDialogWindow.postMessage({greeting: 'Hello from the parent window!'}, '*'); // * 表示接受来自任何来源的消息
  });

}

// ----------------------

// 定义表格的列
const columns = ref([
  {title: computed(() => t('userTickets.ticketId')).value, key: 'id'},
  {title: computed(() => t('userTickets.ticketSubject')).value, key: 'subject'},
  {
    title: computed(() => t('userTickets.ticketUrgency')).value, key: 'urgency', render(row: TicketItem) {
      let level = '';
      switch (row.urgency) {
        case 1:
          level = computed(() => t('userTickets.ticketUrgencyLevel.low')).value
          break;
        case 2:
          level = computed(() => t('userTickets.ticketUrgencyLevel.med')).value
          break;
        case 3:
          level = computed(() => t('userTickets.ticketUrgencyLevel.high')).value
          break;
      }
      return h(NTag, {
        type: row.urgency === 3 ? 'error' : row.urgency === 2 ? 'warning' : 'success',
        bordered: false,
      }, {default: () => level});
    }
  },
  {
    title: computed(() => t('userTickets.ticketStatus')).value, key: 'status', render(row: TicketItem) {
      return h(NTag, {
        type: row.status !== 204 ? 'info' : 'default',
        bordered: false,
      }, {default: () => row.status !== 204 ? computed(() => t('userTickets.ticketActive')).value : computed(() => t('userTickets.ticketInActive')).value});
    }
  },
  {
    title: computed(() => t('userTickets.ticketCreatedAt')), key: 'created_at', render(row: TicketItem) {
      return formatDate(row.created_at)
    }
  },
  {
    title: computed(() => t('userTickets.lastResponse')), key: 'last_response', render(row: TicketItem) {
      return row.last_reply ? formatDate(row.last_reply) : computed(() => t('userTickets.noReply')).value
    }
  },
  {
    title: computed(() => t('userTickets.operate')).value, key: 'actions', fixed: 'right', render(row: TicketItem) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          secondary: true,
          bordered: false,
          disabled: false,
          style: {marginLeft: '10px'},
          onClick: () => openTicket(row)
        }, {
          default: () => computed(() => t('userTickets.checkTicket')).value
        }),
        h(NButton, {
          size: 'small',
          type: 'error',
          secondary: true,
          disabled: row.status === 204,
          style: {marginLeft: '10px'},
          onClick: () => closeTicket(row.id)
        }, {
          default: () => computed(() => t('userTickets.closeTicket')).value
        })
      ]);
    }
  }
])

// 定义表格的列
const pendingColumns = ref([
  {title: computed(() => t('userTickets.ticketId')).value, key: 'id'},
  {title: computed(() => t('userTickets.ticketSubject')).value, key: 'subject'},
  {
    title: computed(() => t('userTickets.ticketUrgency')).value, key: 'urgency', render(row: TicketItem) {
      let level = '';
      switch (row.urgency) {
        case 1:
          level = computed(() => t('userTickets.ticketUrgencyLevel.low')).value
          break;
        case 2:
          level = computed(() => t('userTickets.ticketUrgencyLevel.med')).value
          break;
        case 3:
          level = computed(() => t('userTickets.ticketUrgencyLevel.high')).value
          break;
      }
      return h(NTag, {
        type: row.urgency === 3 ? 'error' : row.urgency === 2 ? 'warning' : 'success',
        bordered: false,
      }, {default: () => level});
    }
  },
  {
    title: computed(() => t('userTickets.ticketStatus')).value, key: 'status', render(row: TicketItem) {
      return h(NTag, {
        type: row.status ? 'info' : 'default',
        bordered: false,
      }, {default: () => row.status ? computed(() => t('userTickets.ticketActive')).value : computed(() => t('userTickets.ticketInActive')).value});
    }
  },
  {
    title: computed(() => t('userTickets.ticketCreatedAt')), key: 'created_at', render(row: TicketItem) {
      return formatDate(row.created_at)
    }
  },
  {
    title: computed(() => t('userTickets.lastResponse')), key: 'last_response', render(row: TicketItem) {
      return row.last_reply ? formatDate(row.last_reply) : computed(() => t('userTickets.noReply')).value
    }
  },
  {
    title: computed(() => t('userTickets.operate')).value, key: 'actions', fixed: 'right', render(row: TicketItem) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          secondary: true,
          bordered: false,
          disabled: false,
          style: {marginLeft: '10px'},
          onClick: () => openTicket(row)
        }, {
          default: () => computed(() => t('userTickets.checkTicket')).value
        }),
        h(NButton, {
          size: 'small',
          type: 'error',
          secondary: true,
          disabled: row.status === 204,
          style: {marginLeft: '10px'},
          onClick: () => closeTicket(row.id)
        }, {
          default: () => computed(() => t('userTickets.closeTicket')).value
        })
      ]);
    }
  }
])


let getAllTicket = async () => {
  try {
    animated.value = false
    let {data} = await instance.get('/api/admin/v1/ticket', {
      params: {
        page: 1,
        size: 10,
      }
    })
    if (data.code === 200) {
      ticketList.value = []
      pendingTicketList.value = []
      data.tickets.forEach((ticket: TicketItem) => ticketList.value.push(ticket))
      data.pending_tickets.forEach((ticket: TicketItem) => pendingTicketList.value.push(ticket))
      animated.value = true
    }
  } catch (error: any) {
    console.log(error)
  }
}

let closeTicket = async (ticket_id: number) => {
  try {
    let {data} = await instance.delete('/api/user/v1/ticket', {
      params: {
        user_id: userInfoStore.thisUser.id,
        ticket_id,
      }
    })
    if (data.code === 200 && data.closed) {
      message.success(computed(() => t('userTickets.ticketCloseSuccess')).value)
      await getAllTicket()
    } else {
      message.error(computed(() => t('userTickets.ticketCloseFailure')).value + data.msg as string || '')
    }
  } catch (error: any) {
    console.log(error)
  }
}


onMounted(async () => {
  themeStore.menuSelected = 'ticket-mgr'
  themeStore.contentPath = '/admin/dashboard/ticket'

  await getAllTicket()
  animated.value = true

  // getAllMyTickets()
})

</script>

<script lang="ts">
export default {
  name: 'TicketMgr',
}
</script>

<template>
  <div style="padding: 20px 20px 0 20px">
    <n-card class="title-card" :embedded="true" :bordered="false" hoverable :title="t('adminTicket.ticketMgr')">
    </n-card>
  </div>
  <transition name="slide-fade">
    <div class="root" v-if="animated">



      <n-card style="margin-top: 20px" :embedded="true" hoverable content-style="padding: 0;" :bordered="false"
              title="待处理工单">
        <n-data-table
            class="table"
            :columns="pendingColumns"
            :data="pendingTicketList"
            :pagination="false"
            :bordered="true"
            style=""
            :scroll-x="800"
        />
      </n-card>

      <n-card
          style="margin-top: 20px"
          :embedded="true"
          :bordered="false"
          hoverable
          content-style="padding: 0"
          title="已完成工单"
      >
        <n-data-table
            class="table"
            :columns="columns"
            :data="ticketList"
            :pagination="false"
            :bordered="true"
            style=""
            :scroll-x="800"
        />
      </n-card>

    </div>
  </transition>

</template>

<style scoped lang="less">
.root {
  margin: 20px;
}
</style>