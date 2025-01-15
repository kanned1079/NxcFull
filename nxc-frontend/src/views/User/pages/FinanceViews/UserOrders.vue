<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {useRouter} from "vue-router";
import {computed, h, onMounted, ref} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import usePaymentStore from "@/stores/usePaymentStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {NButton, NTag, useMessage, type DataTableColumns} from "naive-ui"
// import instance from "@/axios/index"
import {cancelOrder, getAllMyOrders} from "@/api/user/order";
import {formatDate} from "@/utils/timeFormat"
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";

const {t} = useI18n()
const message = useMessage()
const router = useRouter()
const paymentStore = usePaymentStore()

let pageCount = ref(10)

let animated = ref<boolean>(false)

let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

const themeStore = useThemeStore();
const userInfoStore = useUserInfoStore()

interface OrderListItem {
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

// let isOrderCanBeCancelled = ref<boolean>(false)

let orderStatusTagColor = (is_finished: boolean, is_success: boolean): "warning" | "error" | "success" => {
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
    return computed(() => t('userOrders.orderStatusTags.notPay')).value as string // 還沒支付
  } else if (is_finished && !is_success) {
    return computed(() => t('userOrders.orderStatusTags.cancelled')).value as string  // 交易失敗
  } else {
    return computed(() => t('userOrders.orderStatusTags.success')).value as string  // 交易成功
  }
}


let orderList = ref<OrderListItem[]>([])

// 定义表格的列
const columns = computed<DataTableColumns<OrderListItem>>(() => [
  {
    title: t('userOrders.orderId'),
    key: 'order_id'
  },
  {
    title: t('userOrders.planName'),
    key: 'plan_name'
  },
  {
    title: t('userOrders.planCycle'),
    key: 'period',
    render(row: OrderListItem) {
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

      return h(NTag, {bordered: false}, {default: () => periodLabel.value});
    }
  },
  {
    title: t('userOrders.orderPrice'),
    key: 'amount',
    render(row: OrderListItem) {
      return h('span', {}, {default: () => row.amount.toFixed(2)});
    }
  },
  {
    title: t('userOrders.orderStatus'),
    key: 'is_success',
    render(row: OrderListItem) {
      return h(NTag, {
        // type: row.is_success? 'success' : 'error',
        type: orderStatusTagColor(row.is_finished, row.is_success),
        bordered: false,
      }, {default: () => orderStatusText(row.is_finished, row.is_success)});
    }
  },
  {
    title: t('userOrders.createdAt'),
    key: 'created_at',
    render(row: OrderListItem) {
      return h('span', {}, {default: () => formatDate(row.created_at)});
    }
  },
  {
    title: t('userOrders.operate'),
    key: 'actions',
    render(row: OrderListItem) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          secondary: true,
          bordered: false,
          onClick: () => showOrderDetails(row)
        }, {
          default: () => t('userOrders.showDetail')
        }),
        h(NButton, {
          size: 'small',
          type: 'error',
          secondary: true,
          disabled:!(!row.is_success &&!row.is_finished),
          style: {marginLeft: '10px'},
          onClick: () => callCancelOrder(row)
        }, {
          default: () => t('userOrders.cancelOrder')
        })
      ]);
    },
    width: 200,
    fixed: 'right'
  }
]);

let showOrderDetails = (row: OrderListItem) => {

  paymentStore.orderDetail.order_id = row.order_id
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

let callCancelOrder = async (row: OrderListItem) => {
  let data = await cancelOrder(userInfoStore.thisUser.id, row.order_id)
  if (data.code === 200) {
    console.log(data)
    await callGetAllMyOrders()
    message.info(t('userOrders.orderCancelled'))
    // router.back()
  } else {
    message.error("unknown err: " + data.msg || '' as string)
  }
}


let callGetAllMyOrders = async () => {
  let data = await getAllMyOrders(userInfoStore.thisUser.id, dataSize.value.page, dataSize.value.pageSize)
  if (data.code === 200) {
    animated.value = true
    orderList.value = []
    pageCount.value = data.page_count
    data.order_list.forEach((order: OrderListItem) => orderList.value.push(order))
  }
}

onMounted(async () => {
  themeStore.userPath = '/dashboard/orders'
  themeStore.menuSelected = 'user-orders'

  await callGetAllMyOrders()

  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'UserOrders',
}
</script>

<template>
  <div style="padding: 20px 20px 0 20px">
    <n-card :bordered="false" :embedded="true" hoverable :title="t('userOrders.myOrders')"></n-card>
  </div>

  <transition name="slide-fade">
    <div class="root" v-if="animated">

      <n-card class="order-table" :bordered="false" :embedded="true" hoverable content-style="padding: 0;">
        <n-data-table
            striped
            class="table"
            :columns="columns"
            :data="orderList"
            :pagination="false"
            :bordered="true"
            style=""
            :scroll-x="800"
        />
      </n-card>
      <DataTableSuffix
          v-model:data-size="dataSize"
          v-model:page-count="pageCount"
          v-model:animated="animated"
          :update-data="callGetAllMyOrders"
      />

<!--      <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">-->
<!--        <n-pagination-->
<!--            size="medium"-->
<!--            v-model:page.number="dataSize.page"-->
<!--            :page-count="pageCount"-->
<!--            @update:page="animated=false; callGetAllMyOrders()"-->
<!--        />-->
<!--        <n-select-->
<!--            style="width: 160px; margin-left: 20px"-->
<!--            v-model:value.number="dataSize.pageSize"-->
<!--            size="small"-->
<!--            :options="dataCountOptions"-->
<!--            :remote="true"-->
<!--            @update:value="animated=false; dataSize.page = 1; callGetAllMyOrders()"-->
<!--        />-->
<!--      </div>-->
    </div>
  </transition>


</template>

<style scoped lang="less">
.root {
  //padding: 20px;
  padding: 0 20px 20px 20px;

  .order-table {
    margin-top: 15px;
  }

}
</style>