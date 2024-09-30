<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onBeforeMount, onMounted, ref, h, computed} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {NButton, NTag, useMessage} from "naive-ui"
import instance from "@/axios/index"
import {formatDate} from "@/utils/timeFormat"


const {t} = useI18n();
const themeStore = useThemeStore();
const userInfoStore = useUserInfoStore()

interface OrderList {
  id: number
  user_id: number
  email: string
  plan_id: number
  payment_method: string
  period: string
  coupon_id: number
  status: number
  is_success: boolean
  failure_reason: string
  discount_amount: number
  amount: number
  paid_at: any
  created_at: string
  updated_at: string
  deleted_at: any
}

let orderList = ref<OrderList[]>([])

// 定义表格的列
const columns = [
  { title: '# 订单号', key: 'id' },
  {
    title: '周期', key: 'period', render(row: OrderList) {
      return h(NTag, { bordered: false }, { default: () => {
        switch (row.period) {
          case 'month': {return t('userOrders.period.monthPay')}
          case 'quarter': {return t('userOrders.period.quarterPay')}
          case 'half-year': {return t('userOrders.period.halfYearPay')}
          case 'year': {return t('userOrders.period.yearPay')}
        }
        }
      });
    }
  },
  {
    title: '订单金额',
    key: 'amount',
    render(row: OrderList) {
      return h('span', {}, { default: () => row.amount.toFixed(2) });
    }
  },
  {
    title: '订单状态', key: 'is_success', render(row: OrderList) {
      return h(NTag, {
        type: row.is_success ? 'success' : 'error',
        bordered: false,
      }, { default: () => row.is_success ? '已完成' : '已取消' });
    }
  },
  {
    title: '创建时间',
    key: 'created_at',
    render(row: OrderList) {
      return h('span', {}, { default: () => formatDate(row.created_at) });
    }
  },
  {
    title: '操作', key: 'actions', render(row: OrderList) {
      return h('div', { style: { display: 'flex', flexDirection: 'row' } }, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          bordered: false,
          onClick: () => showOrderDetails(row)
        }, { default: () => '订单详情' }),
        h(NButton, {
          size: 'small',
          type: 'error',
          disabled: !row.is_success,
          style: { marginLeft: '10px' },
          onClick: () => cancelOrder(row)
        }, { default: () => '取消订单' })
      ]);
    },
    width: 200,
    fixed: 'right'
  }
];

let showOrderDetails = (row: OrderList) => {
  console.log('显示订单详情', row.id)
}

let cancelOrder = (row: OrderList) => {
console.log('取消订单', row.id)
}


let getAllMyOrders = async () => {
  orderList.value = []
  let {data} = await instance.get('api/user/v1/orders', {
    params: {
      user_id: userInfoStore.thisUser.id
    }
  })
  console.log('我的所有订单',data)
  data.order_list.forEach((order: OrderList) => orderList.value.push(order))
}


onBeforeMount(() => {
  getAllMyOrders()
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
          :data="orderList"
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