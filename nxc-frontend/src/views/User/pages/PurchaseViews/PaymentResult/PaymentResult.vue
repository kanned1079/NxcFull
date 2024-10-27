<script setup lang="ts">
import GoodInfo from "@/views/User/pages/PurchaseViews/Parts/GoodInfo.vue";
import OrderInfo from "@/views/User/pages/PurchaseViews/Parts/OrderInfo.vue";
import {BookOutline as docIcon, ChevronForwardOutline as nextIcon} from "@vicons/ionicons5"
import {computed, onMounted, ref} from "vue"
import {useRouter} from "vue-router";
import useAppInfosStore from "@/stores/useAppInfosStore";
import useThemeStore from "@/stores/useThemeStore";
import {useI18n} from "vue-i18n";
import usePaymentStore from "@/stores/usePaymentStore";
import {formatDate} from "@/utils/timeFormat";

const {t} = useI18n()
const router = useRouter()
const appInfoStore = useAppInfosStore()
const themeStore = useThemeStore()
const paymentStore = usePaymentStore()


let goodInfoData = ref([
  {
    title: computed(() => '商品信息'),
    content: computed(() => paymentStore.orderDetail.plan_name)
  },
  {
    title: computed(() => '周期/类型'),
    content: computed((): string => { // pass
      switch (paymentStore.orderDetail.period) {
        case 'month': {
          return computed(() => '月付')
        }
        case 'quarter': {
          return computed(() => '季付')
        }
        case 'half-year': {
          return computed(() => '半年付')
        }
        case 'year': {
          return computed(() => '年付')
        }
      }
    })
  }
])

let orderData = ref([
  {
    title: computed(() => '订单号'),
    content: computed(() => paymentStore.orderDetail.order_id)
  },
  {
    title: computed(() => '创建日期'),
    content: computed(() => paymentStore.orderDetail.created_at ? formatDate(paymentStore.orderDetail.created_at) : '')
  },
  {
    title: computed(() => '支付金额'),
    content: computed(() => paymentStore.orderDetail.amount.toFixed(2) + ' ' + appInfoStore.appCommonConfig.currency)
  },
  {
    title: computed(() => '支付时间'),
    content: computed(() => paymentStore.orderDetail.paid_at ? formatDate(paymentStore.orderDetail.paid_at) : '未完成支付')
  },
])


onMounted(() => {
  themeStore.userPath = '/dashboard/orders/details'
})

</script>

<script lang="ts">
export default {
  name: 'PaymentResult'
}
</script>

<template>
  <div class="root">
    <n-card class="result-card" :embedded="true" :bordered="false" hoverable>

      <n-result v-if="paymentStore.orderDetail.is_success && paymentStore.orderDetail.is_finished" status="success"
                title="已完成" description="订单已成功支付并开通">
        <template #footer>
          <div style="text-align: right">
            <n-button type="primary" quaternary @click="router.push({path: '/dashboard/document'})">
              <n-icon style="margin-right: 5px" size="16">
                <docIcon/>
              </n-icon>
              查看使用教程
              <n-icon style="margin-left: 5px" size="16">
                <nextIcon/>
              </n-icon>
            </n-button>
          </div>
        </template>
      </n-result>

      <n-result
          v-if="!paymentStore.orderDetail.is_success && !paymentStore.orderDetail.is_finished" s
          tatus="info"
          title="尚未支付"
          description="订单暂时保留可以继续执行支付"
      >
        <template #footer>
          <div style="text-align: right">
            <n-button type="primary" quaternary @click="router.push({path: '/dashboard/document'})">
              <n-icon style="margin-right: 5px" size="16">
                <docIcon/>
              </n-icon>
              去支付
              <n-icon style="margin-left: 5px" size="16">
                <nextIcon/>
              </n-icon>
            </n-button>
          </div>
        </template>
      </n-result>

      <n-result
          v-if="!paymentStore.orderDetail.is_success && paymentStore.orderDetail.is_finished"
          status="404"
          title="已失效"
          description="由于您取消了订单或未在指定时间内完成支付，因此该订单已取消，您可以重新选取订阅。"
      >
        <template #footer>
          <div style="text-align: right">
            <n-button type="primary" quaternary @click="router.push({path: '/dashboard/purchase'})">
              选择新的订阅计划
              <n-icon style="margin-left: 5px" size="16">
                <nextIcon/>
              </n-icon>
            </n-button>
          </div>
        </template>
      </n-result>


    </n-card>

    <GoodInfo :goodInfoData="goodInfoData"></GoodInfo>

    <OrderInfo :orderData="orderData"></OrderInfo>


  </div>
</template>

<style scoped lang="less">
.root {
  margin: 20px;

  .result-card {
    margin-bottom: 20px;
    padding: 40px 0 0 0;
  }
}

</style>