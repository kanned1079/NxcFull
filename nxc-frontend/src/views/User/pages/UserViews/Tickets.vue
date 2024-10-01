<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onMounted, ref, h} from "vue"
import {NButton, NTag, useMessage} from "naive-ui"
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";

const {t} = useI18n();
const userInfoStore = useUserInfoStore();
const themeStore = useThemeStore();
const message = useMessage();

let dialogActive = ref<boolean>(false)

// 表格的字段名
// 主题 工单级别 工单状态 创建时间 最后回复 操作（查看工单/关闭工单）
interface TicketItem {
  id: number  // 工单Id
  title: string // 工单主题
  urgency: number // 紧急程度 1低 2中 3高
  status: boolean // true工单未关闭 false工单关闭
  created_at: string  // 工单创建时间
  last_response: string // 最后一次回复
}

// 假数据
let ticketList = ref<TicketItem[]>([
  {
    id: 1,
    title: '系统登录问题',
    urgency: 2,
    status: true,
    created_at: '2024-09-15 14:22',
    last_response: '2024-09-16 10:00'
  },
  {
    id: 2,
    title: '功能反馈',
    urgency: 1,
    status: false,
    created_at: '2024-09-12 08:30',
    last_response: '2024-09-12 09:45'
  },
  {
    id: 3,
    title: '支付问题',
    urgency: 3,
    status: true,
    created_at: '2024-09-18 16:10',
    last_response: '2024-09-19 11:30'
  },
  {
    id: 3,
    title: '支付问题',
    urgency: 3,
    status: true,
    created_at: '2024-09-18 16:10',
    last_response: '2024-09-19 11:30'
  },
  {
    id: 3,
    title: '支付问题',
    urgency: 3,
    status: true,
    created_at: '2024-09-18 16:10',
    last_response: '2024-09-19 11:30'
  },
  {
    id: 3,
    title: '支付问题',
    urgency: 3,
    status: true,
    created_at: '2024-09-18 16:10',
    last_response: '2024-09-19 11:30'
  },
])

// 点击打开工单后 单独打开一个聊天窗口
let openTicket = (ticket: TicketItem) => {
  message.info(`查看工单：${ticket.title}`);
}

// 关闭工单
let closeTicket = (ticket: TicketItem) => {
  message.success(`工单已关闭：${ticket.title}`);
  ticket.status = false;
}

onMounted(() => {
  themeStore.menuSelected = ''
  themeStore.userPath = '/dashboard/tickets'
})

// 定义表格的列
const columns = [
  {title: '#', key: 'id'},
  {title: '主题', key: 'title'},
  {
    title: '工单级别', key: 'urgency', render(row: TicketItem) {
      let level = '';
      switch (row.urgency) {
        case 1:
          level = '低';
          break;
        case 2:
          level = '中';
          break;
        case 3:
          level = '高';
          break;
      }
      return h(NTag, {
        type: row.urgency === 3 ? 'error' : row.urgency === 2 ? 'warning' : 'success',
        bordered: false,
      }, {default: () => level});
    }
  },
  {
    title: '工单状态', key: 'status', render(row: TicketItem) {
      return h(NTag, {
        type: row.status ? 'info' : 'default',
        bordered: false,
      }, {default: () => row.status ? '未关闭' : '已关闭'});
    }
  },
  {title: '创建时间', key: 'created_at'},
  {title: '最后回复', key: 'last_response'},
  {
    title: '操作', key: 'actions', fixed: 'right', render(row: TicketItem) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          bordered: false,
          style: {marginLeft: '10px'},
          onClick: () => openTicket(row)
        }, {default: () => '查看工单'}),
        h(NButton, {
          size: 'small',
          type: 'error',
          disabled: !row.status,
          style: {marginLeft: '10px'},
          onClick: () => closeTicket(row)
        }, {default: () => '关闭工单'})
      ]);
    }
  }
]

</script>

<template>
  <div class="root">
    <n-card class="tickets-history" :embedded="true" hoverable content-style="padding: 0;" :bordered="false">
      <div class="header">
        <p class="title">{{ t('tickets.userTickets') }}</p>
        <n-button class="add-btn" type="primary" @click="dialogActive = true">{{ t('tickets.addTicket') }}</n-button>
      </div>
    </n-card>

    <n-card :embedded="true" hoverable content-style="padding: 0;" :bordered="false">
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
</template>

<style scoped lang="less">
.root {
  padding: 20px;

  .tickets-history {
    .header {
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      padding: 15px 25px 15px 25px;
      align-items: center;

      .title {
        font-size: 1.1rem;
        font-weight: bold;
        opacity: 0.8;
      }
    }

    margin-bottom: 20px;
  }
}


</style>
