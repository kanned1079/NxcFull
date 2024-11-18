<script setup lang="ts">
import GoodInfo from "@/views/User/pages/PurchaseViews/Parts/GoodInfo.vue";
import OrderInfo from "@/views/User/pages/PurchaseViews/Parts/OrderInfo.vue";
import {BookOutline as docIcon, ChevronForwardOutline as nextIcon, BagHandleOutline as toPayIcon} from "@vicons/ionicons5"
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

let animated = ref<boolean>(false)

let goodInfoData = ref([
  {
    title: computed(() => t('orderPartUniversal.orderDataHex.goodInfo')),
    content: computed(() => paymentStore.orderDetail.plan_name).value
  },
  {
    title: computed(() => t('orderPartUniversal.orderDataHex.cycleOrType')),
    content: computed(() => { // pass
      switch (paymentStore.orderDetail.period) {
        case 'month': {
          return computed(() => t('orderPartUniversal.period.monthPay')).value
        }
        case 'quarter': {
          return computed(() => t('orderPartUniversal.period.quarterPay')).value
        }
        case 'half-year': {
          return computed(() => t('orderPartUniversal.period.halfYearPay')).value
        }
        case 'year': {
          return computed(() => t('orderPartUniversal.period.yearPay')).value
        }
      }
    })
  }
])

let orderData = ref([
  {
    title: computed(() => t('orderPartUniversal.orderDataHex.orderNumber')),
    content: computed(() => paymentStore.orderDetail.order_id).value
  },
  {
    title: computed(() => t('orderPartUniversal.orderDataHex.createdAt')),
    content: computed(() => paymentStore.orderDetail.created_at ? formatDate(paymentStore.orderDetail.created_at) : '').value
  },
  {
    title: computed(() => t('orderPartUniversal.orderDataHex.amount')),
    content: computed(() => paymentStore.orderDetail.amount.toFixed(2) + ' ' + appInfoStore.appCommonConfig.currency).value
  },
  {
    title: computed(() => t('orderPartUniversal.orderDataHex.paidAt')),
    content: computed(() => paymentStore.orderDetail.paid_at ? formatDate(paymentStore.orderDetail.paid_at) : '未完成支付').value
  },
  {
    title: '失败原因',
    content: paymentStore.orderDetail.failure_reason
  }
])

let toConfirmOrder = () => {
  paymentStore.submittedOrderId = paymentStore.orderDetail.order_id
  router.push({
    path: '/dashboard/purchase/confirm'
  })
}


onMounted(() => {
  themeStore.userPath = '/dashboard/orders/details'
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'PaymentResult'
}
</script>

<template>
  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card class="result-card" :embedded="true" :bordered="false" hoverable>

        <n-result
            v-if="paymentStore.orderDetail.is_success && paymentStore.orderDetail.is_finished"
            status="success"
            :title="t('orderDetail.finished')"
            :description="t('orderDetail.finishedAndSuccessDescription')">
          <template #footer>
            <div style="display: flex; flex-direction: row; align-items: center; justify-content: flex-end;">
              <n-button type="primary" quaternary @click="router.push({path: '/dashboard/document'})">
                <n-icon style="margin-right: 5px" size="16">
                  <docIcon/>
                </n-icon>
                <!--              查看使用教程-->
                {{ t('orderDetail.useManual') }}
                <n-icon style="margin-left: 5px" size="16">
                  <nextIcon/>
                </n-icon>
              </n-button>
            </div>
          </template>
        </n-result>

        <n-result
            v-if="!paymentStore.orderDetail.is_success && !paymentStore.orderDetail.is_finished"
            status="info"
            :title="t('orderDetail.payPending')"
            :description="t('orderDetail.pendingDescription')"
        >
          <template #footer>
            <div style="display: flex; flex-direction: row; align-items: center; justify-content: flex-end;">
              <n-button type="primary" quaternary @click="toConfirmOrder">
                <n-icon style="margin-right: 5px" size="16">
                  <toPayIcon/>
                </n-icon>
                <!--              去支付-->
                   {{ t('orderDetail.toPay') }} 
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
            :title="t('orderDetail.outDate')"
            :description="t('orderDetail.outDateDescription')"
        >
          <template #footer>
            <div style="display: flex; flex-direction: row; align-items: center; justify-content: flex-end;">
              <n-button type="primary" quaternary @click="router.push({path: '/dashboard/purchase'})">
                <!--              选择新的订阅计划-->
                {{ t('orderDetail.chooseNewPlan') }}
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
  </transition>

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