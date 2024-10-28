<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {useRouter} from "vue-router";
import {h, onBeforeMount, onMounted, ref, computed} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import usePaymentStore from "@/stores/usePaymentStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {NButton, NTag, useMessage} from "naive-ui"
import instance from "@/axios/index"
import {formatDate} from "@/utils/timeFormat"

const message = useMessage()
const router = useRouter()
const paymentStore = usePaymentStore()

let pageCount = ref(10)

let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

let dataCountOptions = [
  {
    label: '10条数据/页',
    value: 10,
  },
  {
    label: '20条数据/页',
    value: 20,
  },
  {
    label: '50条数据/页',
    value: 50,
  },
  {
    label: '100条数据/页',
    value: 100,
  },
]


const {t} = useI18n();
const themeStore = useThemeStore();
const userInfoStore = useUserInfoStore()

interface OrderList {
  id: number
  order_id: string
  user_id: number
  email: string
  plan_id: number
  plan_name: string
  payment_method: string
  period: string
  coupon_id: number
  status: number
  is_success: boolean
  is_finished: boolean
  failure_reason: string
  discount_amount: number
  amount: number
  paid_at: any
  created_at: string
  updated_at: string
  deleted_at: any
}

let isOrderCanBeCancelled = ref<boolean>(false)

let orderStatusTagColor = (is_finished: boolean, is_success: boolean): string => {
  if (!is_success && !is_finished) {
    return 'warning'
  } else if (is_finished && !is_success) {
    return 'error'
  } else {
    return 'success'
  }
}

let orderStatusText = (is_finished: boolean, is_success: boolean): string => {
  if (!is_success && !is_finished) {
    return '未支付'
  } else if (is_finished && !is_success) {
    return '交易失败'
  } else {
    return '成功'
  }
}


let orderList = ref<OrderList[]>([])

// 定义表格的列
const columns = [
  {title: '#', key: 'order_id'},
  {title: '订阅名', key: 'plan_name'},
  {
    title: '周期', key: 'period', render(row) {
      const periodLabel = computed(() => {
        switch (row.period) {
          case 'month':
            return t('userOrders.period.monthPay');
          case 'quarter':
            return t('userOrders.period.quarterPay');
          case 'half-year':
            return t('userOrders.period.halfYearPay');
          case 'year':
            return t('userOrders.period.yearPay');
          default:
            return ''; // 若 period 无匹配值
        }
      });

      return h(NTag, { bordered: false }, { default: () => periodLabel.value });
    }
  },
  {
    title: '订单金额',
    key: 'amount',
    render(row: OrderList) {
      return h('span', {}, {default: () => row.amount.toFixed(2)});
    }
  },
  {
    title: '订单状态', key: 'is_success', render(row: OrderList) {
      return h(NTag, {

        // type: row.is_success ? 'success' : 'error',
        type: orderStatusTagColor(row.is_finished, row.is_success),
        bordered: false,
      }, {default: () => orderStatusText(row.is_finished, row.is_success)});
    }
  },
  {
    title: '创建时间',
    key: 'created_at',
    render(row: OrderList) {
      return h('span', {}, {default: () => formatDate(row.created_at)});
    }
  },
  {
    title: '操作', key: 'actions', render(row: OrderList) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          bordered: false,
          onClick: () => showOrderDetails(row)
        }, {default: () => '订单详情'}),
        h(NButton, {
          size: 'small',
          type: 'error',
          disabled: !(!row.is_success && !row.is_finished),
          style: {marginLeft: '10px'},
          onClick: () => cancelOrder(row)
        }, {default: () => '取消订单'})
      ]);
    },
    width: 200,
    fixed: 'right'
  }
];

let showOrderDetails = (row: OrderList) => {
  console.log('显示订单详情', row.id)
  paymentStore.orderDetail.order_id = row.order_id
  // paymentStore.orderDetail.period = () => {
  //   switch (row.period) {
  //     case 'month': {
  //       return t('userOrders.period.monthPay')
  //     }
  //     case 'quarter': {
  //       return t('userOrders.period.quarterPay')
  //     }
  //     case 'half-year': {
  //       return t('userOrders.period.halfYearPay')
  //     }
  //     case 'year': {
  //       return t('userOrders.period.yearPay')
  //     }
  //   }
  // }
  paymentStore.orderDetail.plan_name = row.plan_name
  paymentStore.orderDetail.period = row.period
  paymentStore.orderDetail.amount = row.amount
  paymentStore.orderDetail.created_at = row.created_at
  paymentStore.orderDetail.paid_at = row.paid_at
  paymentStore.orderDetail.failure_reason = row.failure_reason
  paymentStore.orderDetail.is_finished = row.is_finished
  paymentStore.orderDetail.is_success = row.is_success
  router.push({
    path: "/dashboard/orders/details"
  })
}

let cancelOrder = async (row: OrderList) => {
  try {
    let {data} = await instance.delete('/api/user/v1/order', {
      params: {
        user_id: userInfoStore.thisUser.id,
        order_id: row.order_id
      }
    })
    if (data.code === 200) {
      console.log(data)
      message.info('订单已取消')
      // router.back()
    } else {
      message.error("未知错误" + data.msg || '' as string)
    }
    await getAllMyOrders()
  } catch (error: any) {
    console.log(error)
    message.error("未知错误" + error)
  }
}


let getAllMyOrders = async () => {
  orderList.value = []
  let {data} = await instance.get('api/user/v1/orders', {
    params: {
      user_id: userInfoStore.thisUser.id,
      page: dataSize.value.page,
      size: dataSize.value.pageSize,
    }
  })
  console.log('我的所有订单', data)
  pageCount.value = data.page_count
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

    <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">
      <n-pagination
          size="medium"
          v-model:page.number="dataSize.page"
          :page-count="pageCount"
          @update:page="getAllMyOrders"
      />
      <n-select
          style="width: 160px; margin-left: 20px"
          v-model:value.number="dataSize.pageSize"
          size="small"
          :options="dataCountOptions"
          :remote="true"
          @update:value="dataSize.page = 1; getAllMyOrders()"
      />
    </div>
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