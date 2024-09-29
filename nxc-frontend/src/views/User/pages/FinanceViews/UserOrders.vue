<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onBeforeMount, onMounted, ref} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import {NButton, NTag, useMessage} from "naive-ui"



const {t} = useI18n();
const themeStore = useThemeStore();



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


onBeforeMount(() => {

})

onMounted(() => {
  themeStore.userPath = '/dashboard/orders'
  themeStore.menuSelected = 'user-orders'
})

</script>

<script lang="ts">
export default {
  name: 'UserOrders',
}
</script>

<template>
  <div class="root">
    <n-card :bordered="false" :embedded="true" hoverable :title="t('userOrders.myOrders')"></n-card>

    <n-card class="order-table" :bordered="false" :embedded="true" hoverable content-style="padding: 0;">
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

  .order-table {
    margin-top: 20px;
  }

}
</style>