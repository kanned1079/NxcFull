<script setup lang="ts">
import {computed, h, onMounted, onBeforeMount, ref} from "vue"
import {useI18n} from "vue-i18n";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import useThemeStore from "@/stores/useThemeStore";
import {NButton, NTag, useMessage, type DataTableColumns, type CheckboxProps, useDialog} from "naive-ui";
import {formatDate} from "@/utils/timeFormat";
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue";
import useTablePagination from "@/hooks/useTablePagination";
import {handleCloseTicketById, handleGetAllTicket} from "@/api/admin/ticket";
import {handleDeleteGroup} from "@/api/admin/groups";
// import {ArrowUpOutline} from "@vicons/ionicons5"

const {t} = useI18n()
const i18nPrefix = 'adminViews.adminTicket'
const message = useMessage()
const dialog = useDialog()
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

let typeChooseCheckVal = ref<any[]>(['pending'])
// typeChooseCheckVal.value.push('pending')

const [dataSizePending, pageCountPending] = useTablePagination();
const [dataSizeAll, pageCountAll] = useTablePagination()
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
const columns = computed<DataTableColumns<TicketItem>>(() => [
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
          onClick: () => closeTicket(row)
        }, {
          default: () => t('userTickets.closeTicket')
        })
      ]);
    }
  }
])


let getAllTicket = async () => {
  let data = await handleGetAllTicket(
      dataSizePending.value.page,
      dataSizePending.value.pageSize,
      dataSizeAll.value.page,
      dataSizeAll.value.pageSize,
  )
  if (data && data.code === 200) {
    ticketList.value = []
    pendingTicketList.value = []
    if (data.tickets && data.pending_tickets) {
      data.tickets.forEach((ticket: TicketItem) => ticketList.value.push(ticket))
      pageCountAll.value = data.finished_page_count ? data.finished_page_count : 0
      data.pending_tickets.forEach((ticket: TicketItem) => pendingTicketList.value.push(ticket))
      pageCountPending.value = data.pending_page_count ? data.pending_page_count : 0
    }
  } else {
    message.error(t('adminViews.common.fetchDataFailure'))
  }
  animated.value = true
}

let closeTicket = async (row: TicketItem) => {
  let runCloseTask = async () => {
    animated.value = false
    let data = await handleCloseTicketById(userInfoStore.thisUser.id, row.id)
    if (data && data.code === 200 && data.closed) {
      message.success(t('userTickets.ticketCloseSuccess'))
      await getAllTicket()
    } else {
      message.error(t('userTickets.ticketCloseFailure') + data.msg || '')
    }
  }
  dialog.warning({
    title: t('adminViews.adminTicket.mention.title'),
    content: () => {
      return h('div', {}, [
        h('p', {style: {fontWeight: 'bold', fontSize: '1rem', opacity: '0.8'}}, row.subject),
        h('p', {style: {marginTop: '4px'}}, t('adminViews.adminTicket.mention.content'))
      ])
    },
    positiveText: t('adminViews.common.confirm'),
    negativeText: t('adminViews.common.cancel'),
    showIcon: false,
    actionStyle: {
      marginTop: '20px',
    },
    contentStyle: {
      marginTop: '20px',
    },
    onPositiveClick: async () => {
      await runCloseTask()
    }
  })
}

const triggerTypeCheck = () => {
  animated.value = false
  setTimeout(() => animated.value = true, 20)
}

// let showPendingTickets = computed(() => {
//   animated.value = true
//   return typeChooseCheckVal.value.includes('pending');
// });

onBeforeMount(() => {
  themeStore.menuSelected = 'ticket-mgr'
  themeStore.breadcrumb = `${i18nPrefix}.ticketMgr`
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
  <PageHead
      :title="t(`${i18nPrefix}.ticketMgr`)"
      :description="t(`${i18nPrefix}.description`)"
  >
    <n-form-item
        :show-label="false"
        path="checkboxGroupValue"
        :show-feedback="false"
        :show-require-mark="true"
    >
      <n-checkbox-group v-model:value="typeChooseCheckVal" @update:value="triggerTypeCheck">
        <n-space>
          <n-checkbox :value="'pending'">
            {{ t('adminViews.adminTicket.pendingTicket') }}
          </n-checkbox>
          <n-checkbox :value="'finished'">
            {{ t('adminViews.adminTicket.finishedTicket') }}
          </n-checkbox>
        </n-space>
      </n-checkbox-group>
    </n-form-item>

  </PageHead>
  <transition name="slide-fade">
    <div class="root" v-if="animated">

      <div v-show="typeChooseCheckVal.includes('pending')">
        <PageHead
            :title="t(`${i18nPrefix}.pendingTicket`)"
            :description="t('adminViews.adminTicket.type.pendingDescription')"
            style="margin: 0 !important;"
            secondary
        />
        <n-card
            :embedded="true"
            hoverable
            content-style="padding: 0;"
            :bordered="false"
        >
          <n-data-table
              class="table"
              size="medium"
              striped
              :columns="columns"
              :data="pendingTicketList"
              :pagination="false"
              :bordered="true"
              style=""
              :scroll-x="1200"
          />
        </n-card>
        <DataTableSuffix
            v-model:data-size="dataSizePending"
            v-model:page-count="pageCountPending"
            v-model:animated="animated"
            :update-data="getAllTicket"
        />
      </div>

      <div v-show="typeChooseCheckVal.includes('finished')">
        <PageHead
            :title="t(`${i18nPrefix}.finishedTicket`)"
            :description="t('adminViews.adminTicket.type.finishedDescription')"
            style="margin: 0 !important;"
            secondary
        />
        <n-card
            :embedded="true"
            :bordered="false"
            hoverable
            content-style="padding: 0"
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
              :scroll-x="1200"
          />
        </n-card>
        <DataTableSuffix
            v-model:data-size="dataSizeAll"
            v-model:page-count="pageCountAll"
            v-model:animated="animated"
            :update-data="getAllTicket"
        />
      </div>

      <div class="none-select" v-if="typeChooseCheckVal.length === 0">
        <!--        <n-h3>-->
        <!--          {{ t('adminViews.adminTicket.chooseOneNecessary') }}-->
        <!--        </n-h3>-->
        <!--        <n-icon size="20"><ArrowUpOutline /></n-icon>-->
        <n-alert
            :title="t('adminViews.adminTicket.chooseOneNecessary')"
            type="info"
            :bordered="!themeStore.enableDarkMode"
          >

        </n-alert>
      </div>

    </div>


  </transition>

</template>

<style scoped lang="less">
.root {
  margin: 20px;

  .ticket-title {
    font-size: 1.1rem;
    font-weight: bold;
    margin: 10px 10px 10px 4px;
  }

  .none-select {
  }
}
</style>