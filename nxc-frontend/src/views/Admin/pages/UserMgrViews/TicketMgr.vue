<script setup lang="ts">
import {computed, h, onMounted, onBeforeMount, ref} from "vue"
import {useI18n} from "vue-i18n";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import useThemeStore from "@/stores/useThemeStore";
import {NButton, NTag, useMessage, type DataTableColumns} from "naive-ui";
import {formatDate} from "@/utils/timeFormat";
import instance from "@/axios";
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue";

const {t} = useI18n()
const i18nPrefix = 'adminViews.adminTicket'
const message = useMessage()
const userInfoStore = useUserInfoStore()
const appInfoStore = useAppInfosStore()
const themeStore = useThemeStore()

interface TicketItem {
  id: number  // 工单Id
  subject: string // 工单主题
  urgency: number // 紧急程度 1低 2中 3高
  status: number // true工单未关闭 false工单关闭
  created_at: string  // 工单创建时间
  last_reply: string // 最后一次回复
}

let pageCountPending = ref<number>(10)
let pageCountAll = ref<number>(10)

let dataSizePending = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

let dataSizeAll = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})


let animated = ref<boolean>(false)

let ticketList = ref<TicketItem[]>([])
let pendingTicketList = ref<TicketItem[]>([])

// 点击打开工单后 单独打开一个聊天窗口
let openTicket = (ticket: TicketItem) => {
  // message.info(`查看工单：${ticket.subject}`);
  const chatDialogWindow = window.open(
      `${appInfoStore.appCommonConfig.app_url}/user/tickets/chat?user_id=${userInfoStore.thisUser.id}&ticket_id=${ticket.id}&subject=${ticket.subject}&status=${ticket.status}&role=0`,
      '111111',
      'width=480,height=640,resizable=yes,scrollbars=yes',
  );
  chatDialogWindow?.addEventListener('onload', () => {
    chatDialogWindow?.postMessage({greeting: 'Hello from the parent window!'}, '*'); // * 表示接受来自任何来源的消息
  });

}

// 定义表格的列
const columns = computed<DataTableColumns<TicketItem>>( () => [
  {
    title: t('userTickets.ticketId'),
    key: 'id'
  },
  {
    title: t('userTickets.ticketSubject'),
    key: 'subject'
  },
  {
    title: t('userTickets.ticketUrgency'),
    key: 'urgency',
    render(row: TicketItem) {
      let level = '';
      switch (row.urgency) {
        case 1:
          level = t('userTickets.ticketUrgencyLevel.low')
          break;
        case 2:
          level = t('userTickets.ticketUrgencyLevel.med')
          break;
        case 3:
          level = t('userTickets.ticketUrgencyLevel.high')
          break;
      }
      return h(NTag, {
        type: row.urgency === 3 ? 'error' : row.urgency === 2 ? 'warning' : 'success',
        bordered: false,
      }, {default: () => level});
    }
  },
  {
    title: t('userTickets.ticketStatus'),
    key: 'status',
    render(row: TicketItem) {
      return h(NTag, {
        type: row.status !== 204 ? 'info' : 'default',
        bordered: false,
      }, {default: () => row.status !== 204 ? t('userTickets.ticketActive') : t('userTickets.ticketInActive')});
    }
  },
  {
    title: t('userTickets.ticketCreatedAt'),
    key: 'created_at',
    render(row: TicketItem) {
      return formatDate(row.created_at)
    }
  },
  {
    title: t('userTickets.lastResponse'),
    key: 'last_response',
    render(row: TicketItem) {
      return row.last_reply ? formatDate(row.last_reply) : t('userTickets.noReply')
    }
  },
  {
    title: t('userTickets.operate'),
    key: 'actions', fixed: 'right',
    render(row: TicketItem) {
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
          default: () => t('userTickets.checkTicket')
        }),
        h(NButton, {
          size: 'small',
          type: 'error',
          secondary: true,
          disabled: row.status === 204,
          style: {marginLeft: '10px'},
          onClick: () => closeTicket(row.id)
        }, {
          default: () => t('userTickets.closeTicket')
        })
      ]);
    }
  }
])

// 定义表格的列
// const pendingColumns = ref([
//   {
//     title: computed(() => t('userTickets.ticketId')).value,
//     key: 'id'
//   },
//   {
//     title: computed(() => t('userTickets.ticketSubject')).value,
//     key: 'subject'
//   },
//   {
//     title: computed(() => t('userTickets.ticketUrgency')).value,
//     key: 'urgency',
//     render(row: TicketItem) {
//       let level = '';
//       switch (row.urgency) {
//         case 1:
//           level = computed(() => t('userTickets.ticketUrgencyLevel.low')).value
//           break;
//         case 2:
//           level = computed(() => t('userTickets.ticketUrgencyLevel.med')).value
//           break;
//         case 3:
//           level = computed(() => t('userTickets.ticketUrgencyLevel.high')).value
//           break;
//       }
//       return h(NTag, {
//         type: row.urgency === 3 ? 'error' : row.urgency === 2 ? 'warning' : 'success',
//         bordered: false,
//       }, {default: () => level});
//     }
//   },
//   {
//     title: computed(() => t('userTickets.ticketStatus')).value,
//     key: 'status',
//     render(row: TicketItem) {
//       return h(NTag, {
//         type: row.status ? 'info' : 'default',
//         bordered: false,
//       }, {default: () => row.status ? computed(() => t('userTickets.ticketActive')).value : computed(() => t('userTickets.ticketInActive')).value});
//     }
//   },
//   {
//     title: computed(() => t('userTickets.ticketCreatedAt')),
//     key: 'created_at',
//     render(row: TicketItem) {
//       return formatDate(row.created_at)
//     }
//   },
//   {
//     title: computed(() => t('userTickets.lastResponse')),
//     key: 'last_response',
//     render(row: TicketItem) {
//       return row.last_reply ? formatDate(row.last_reply) : computed(() => t('userTickets.noReply')).value
//     }
//   },
//   {
//     title: computed(() => t('userTickets.operate')).value,
//     key: 'actions',
//     fixed: 'right',
//     render(row: TicketItem) {
//       return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
//         h(NButton, {
//           size: 'small',
//           type: 'primary',
//           secondary: true,
//           bordered: false,
//           disabled: false,
//           style: {marginLeft: '10px'},
//           onClick: () => openTicket(row)
//         }, {
//           default: () => computed(() => t('userTickets.checkTicket')).value
//         }),
//         h(NButton, {
//           size: 'small',
//           type: 'error',
//           secondary: true,
//           disabled: row.status === 204,
//           style: {marginLeft: '10px'},
//           onClick: () => closeTicket(row.id)
//         }, {
//           default: () => computed(() => t('userTickets.closeTicket')).value
//         })
//       ]);
//     }
//   }
// ])


let getAllTicket = async () => {
  try {
    animated.value = false
    let {data} = await instance.get('/api/admin/v1/ticket', {
      params: {
        pending_page: dataSizePending.value.page,
        pending_size: dataSizePending.value.pageSize,
        finished_page: dataSizeAll.value.page,
        finished_size: dataSizeAll.value.pageSize,
      }
    })
    if (data.code === 200) {
      ticketList.value = []
      pendingTicketList.value = []
      data.tickets.forEach((ticket: TicketItem) => ticketList.value.push(ticket))
      pageCountAll.value = data.finished_page_count ? data.finished_page_count : 1
      data.pending_tickets.forEach((ticket: TicketItem) => pendingTicketList.value.push(ticket))
      pageCountPending.value = data.pending_page_count ? data.pending_page_count : 1
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

onBeforeMount(() => {
  themeStore.menuSelected = 'ticket-mgr'
  themeStore.breadcrumb = t(`${i18nPrefix}.ticketMgr`)
})

onMounted(async () => {
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
<!--  <div style="padding: 20px 20px 0 20px">-->
<!--    <n-card class="title-card" :embedded="true" :bordered="false" hoverable :title="t('adminTicket.ticketMgr')">-->
<!--    </n-card>-->
<!--  </div>-->

  <PageHead
      :title="t(`${i18nPrefix}.ticketMgr`)"
      :description="t(`${i18nPrefix}.description`)"
  />

  <transition name="slide-fade">
    <div class="root" v-if="animated">

      <n-card style="margin-top: 20px" :embedded="true" hoverable content-style="padding: 0;" :bordered="false"
              :title="t(`${i18nPrefix}.pendingTicket`)">
        <n-data-table
            class="table"
            size="medium"
            striped
            :columns="columns"
            :data="pendingTicketList"
            :pagination="false"
            :bordered="true"
            style=""
            :scroll-x="800"
        />
      </n-card>

<!--      <div style="margin: 20px 0 40px 0; display: flex; flex-direction: row; justify-content: right;">-->
<!--        <n-pagination-->
<!--            size="medium"-->
<!--            v-model:page.number="dataSizePending.page"-->
<!--            :page-count="pageCountPending"-->
<!--            @update:page="animated=false; getAllTicket()"-->
<!--        />-->
<!--        <n-select-->
<!--            style="width: 160px; margin-left: 20px"-->
<!--            v-model:value.number="dataSizePending.pageSize"-->
<!--            size="small"-->
<!--            :options="dataCountOptions"-->
<!--            :remote="true"-->
<!--            @update:value="animated=false; dataSizePending.page = 1; getAllTicket()"-->
<!--        />-->
<!--      </div>-->

      <DataTableSuffix
          v-model:data-size="dataSizePending"
          v-model:page-count="pageCountPending"
          v-model:animated="animated"
          :update-data="getAllTicket"
      />

      <n-card
          style="margin-top: 20px"
          :embedded="true"
          :bordered="false"
          hoverable
          content-style="padding: 0"
          :title="t(`${i18nPrefix}.finishedTicket`)"
      >
        <n-data-table
            class="table"
            size="medium"
            striped
            :columns="columns"
            :data="ticketList"
            :pagination="false"
            :bordered="true"
            style=""
            :scroll-x="800"
        />
      </n-card>

<!--      <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">-->
<!--        <n-pagination-->
<!--            size="medium"-->
<!--            v-model:page.number="dataSizeAll.page"-->
<!--            :page-count="pageCountAll"-->
<!--            @update:page="animated=false; getAllTicket()"-->
<!--        />-->
<!--        <n-select-->
<!--            style="width: 160px; margin-left: 20px"-->
<!--            v-model:value.number="dataSizeAll.pageSize"-->
<!--            size="small"-->
<!--            :options="dataCountOptions"-->
<!--            :remote="true"-->
<!--            @update:value="animated=false; dataSizeAll.page = 1; getAllTicket()"-->
<!--        />-->
<!--      </div>-->

      <DataTableSuffix
          v-model:data-size="dataSizeAll"
          v-model:page-count="pageCountAll"
          v-model:animated="animated"
          :update-data="getAllTicket"
      />

    </div>
  </transition>

</template>

<style scoped lang="less">
.root {
  margin: 20px;
}
</style>