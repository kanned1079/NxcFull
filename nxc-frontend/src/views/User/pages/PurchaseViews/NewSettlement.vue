<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, onBeforeMount, onBeforeUnmount, onMounted, ref} from "vue"
import type {MenuOption, NotificationType} from 'naive-ui'
import {NIcon, useMessage, useNotification} from 'naive-ui'
import useThemeStore from "@/stores/useThemeStore";
import usePaymentStore from "@/stores/usePaymentStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import {useRouter} from "vue-router";
import {MdPreview} from 'md-editor-v3';
import 'md-editor-v3/lib/preview.css';
import {
  CheckmarkCircleOutline as settlementIcon,
  ChevronBackOutline as backIcon,
  TicketOutline as ticketIcon
} from "@vicons/ionicons5"
// import instance from "@/axios";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {saveOrder, verifyTicket} from "@/api/user/settlement";

const {t} = useI18n();
const notification = useNotification()
const message = useMessage()
const userInfoStore = useUserInfoStore();
const appInfosStore = useAppInfosStore();
const themeStore = useThemeStore();
// const apiAddrStore = useApiAddrStore();
const paymentStore = usePaymentStore();
const router = useRouter();

let animate = ref<boolean>(false)

interface Plan {
  id: number
  group_id?: number
  is_renew?: boolean
  is_sale?: boolean
  name: string
  capacity_limit: number
  describe?: string
  month_price?: number
  quarter_price?: number
  half_year_price?: number
  year_price?: number
  created_at: string
  updated_at?: number
  deleted_at?: string
}

let rightSideCardBgColor = computed(() => themeStore.enableDarkMode ? 'rgba(54,55,55,0.8)' : 'rgba(54,56,61,0.8)')
const id = 'preview-only';
let plan: Plan = paymentStore.plan_list[paymentStore.plan_id_selected]
let selectedCycle = ref<string>('month')
let resultPrice = ref<number>(0)
// let ticketVerified = ref<boolean>(false)

let couponCode = ref<string>('')
let period = ref<string>()

let showCycleName = computed(() => {
  if (selectedCycle.value === 'month') {
    resultPrice.value = plan.month_price || 0
    return computed(() => t('newSettlement.monthPay'))
  }
  if (selectedCycle.value === 'quarter') {
    resultPrice.value = plan.quarter_price || 0
    return computed(() => t('newSettlement.quarterPay'))
  }
  if (selectedCycle.value === 'half-year') {
    resultPrice.value = plan.half_year_price || 0
    return computed(() => t('newSettlement.halfYearPay'))
  }
  if (selectedCycle.value === 'year') {
    resultPrice.value = plan.year_price || 0
    return computed(() => t('newSettlement.yearPay'))
  }
})

let cycleSelect = (key: string) => {
  selectedCycle.value = key
  period.value = key
}

const menuOptions: MenuOption[] = []

let appendCycleOptions = () => {
  if (plan.month_price) {
    menuOptions.push({
      label: t('newSettlement.monthPay'),
      key: 'month',
      disabled: false,
    })
  }
  if (plan.quarter_price) {
    menuOptions.push({
      label: t('newSettlement.quarterPay'),
      key: 'quarter',
      disabled: false,
    })
  }
  if (plan.half_year_price) {
    menuOptions.push({
      label: t('newSettlement.halfYearPay'),
      key: 'half-year',
      disabled: false,
    })
  }
  if (plan.year_price) {
    menuOptions.push({
      label: t('newSettlement.yearPay'),
      key: 'year',
      disabled: false,
    })
  }
}

// 优惠券信息
interface Coupon {
  id: number,
  name: string,
  verified: boolean,
  percent_off: number,
}

let couponInfo = ref<Coupon>({
  id: 0,
  name: '',
  verified: false,
  percent_off: 0.0,
})

let callVerifyTicket = async () => {
  if (couponCode.value.trim() === '') {
    notify('error', t('newSettlement.err'), t('newSettlement.notify.couponIsNull'))
    return
  }
  let data = await verifyTicket(couponCode.value, plan.id, userInfoStore.thisUser.id)
  if (data.code === 200) {
    if (data.verified) {
      // ticketVerified.value = true;
      couponInfo.value.id = data.id;
      couponInfo.value.verified = true;
      couponInfo.value.percent_off = data.percent_off;
      couponInfo.value.name = data.name;
      notify('success', t('newSettlement.notify.passTitle'), t('newSettlement.notify.couponVerified'));
    } else {
      couponInfo.value.verified = false
      notify('error', t('newSettlement.notify.couponInvalid'), data.msg);
    }
  } else {
    couponInfo.value.verified = false
    // 处理非200状态码的错误
    notify('error', t('newSettlement.notify.couponInvalid'), data.msg);
  }

}

// 抵折金额
let discountPrice = computed(() => couponInfo.value.verified ? ((couponInfo.value.percent_off / 100) * resultPrice.value) : 0)

let callSaveOrder = async () => {
  // try {
  //   let {data} = await instance.post('/api/user/v1/order', {
  //     plan_id: plan.id,
  //     user_id: userInfoStore.thisUser.id,
  //     coupon_id: couponInfo.value.id || null,
  //     period: period.value || 'month',
  //   })
  let data = await saveOrder(plan.id, userInfoStore.thisUser.id, couponInfo.value.id, period.value || 'month')
  if (data.code === 200) {
    console.log('订单创建成功')
    // Object.assign(data, paymentStore.confirmOrder)
    // paymentStore.confirmOrder = data
    paymentStore.submittedOrderId = data.order_id as string   // 仅保存订单号
    console.log('选择支付方式')
    // 执行异步
    setTimeout(async () => {
      await router.push({path: '/dashboard/purchase/confirm'})
    }, 500) // 延迟500ms

  }
  // } catch (error: any) {
  //   console.log(error)
  //   message.error(error)
  // }
}

let notify = (type: NotificationType, title: string, meta: string) => {
  notification[type]({
    content: title,
    meta: meta,
    duration: 2500,
    keepAliveOnHover: true
  })
}

onMounted(() => {

  console.log('settlement挂载', paymentStore.plan_id_selected)
  // appendCycleOptions()
  themeStore.userPath = '/dashboard/purchase/settlement'

  animate.value = true

})

onBeforeMount(() => {
  console.log('settlement挂载前', paymentStore.plan_id_selected)
  if (paymentStore.plan_id_selected < 0) {
    router.back()
  } else {
    appendCycleOptions()

  }
})

onBeforeUnmount(() => {
  animate.value = false
})


</script>

<script lang="ts">
export default {
  name: 'NewSettlement',
}
</script>

<template>
  <div style="margin: 20px 20px 0 20px; font-size: 1rem">
    <n-button type="primary" text @click="router.back()">
      <n-icon style="margin: 0 5px 0 2px;">
        <backIcon/>
      </n-icon>
      重新選取
    </n-button>
  </div>

  <transition name="slide-fade">
    <div class="root" v-if="paymentStore.plan_id_selected != -1 && animate">
      <div class="root-layout">
        <div class="l-part">
          <n-card class="l-info-head" :embedded="true" hoverable content-style="padding: 0;">
            <p class="title">{{ plan.name }}</p>
            <MdPreview
                style="background-color: rgba(0,0,0,0)"
                :theme="themeStore.enableDarkMode?'dark':'light'"
                :editorId="id"
                :modelValue="plan.describe"
                :preview-theme="'github'"
            />
          </n-card>
          <n-card class="l-info-btn" :embedded="true" hoverable content-style="padding: 0;">
            <p class="detail-title">{{ t('newSettlement.payCycle') }}</p>
            <n-menu
                :options="menuOptions"
                @update:value="cycleSelect"
                :value="selectedCycle"
                style="font-size: 0.9rem"
            />
          </n-card>
        </div>

        <div class="r-part">
          <n-card class="r-part-head" :embedded="true" hoverable content-style="padding: 10px">
            <!--          <n-input-group>-->
            <div class="r-border" style="display: flex; flex-direction: row; border-radius: 5px; padding: 5px">
              <n-input
                  size="large"
                  :bordered="true"
                  class="ticket-input"
                  style="margin-right: 10px"
                  v-model:value="couponCode"
                  :placeholder="t('newSettlement.couponPlaceholder')"
              />
              <n-button
                  style="color: white"
                  size="large"
                  :keyboard="true"
                  type="primary"
                  @click="callVerifyTicket"
                  :bordered="false"
              >
                <n-icon style="margin-right: 10px">
                  <ticketIcon/>
                </n-icon>
                {{ t('newSettlement.verifyBtn') }}
              </n-button>
            </div>

          </n-card>
          <n-card style="color: white" class="r-part-btn" :embedded="true" hoverable>
            <p class="r-title">{{ t('newSettlement.orderTotalTitle') }}</p>
            <div class="price-small">
              <p class="summary">{{ plan.name }} x {{ showCycleName }}</p>
              <p class="price-small">{{ resultPrice.toFixed(2) }} {{ appInfosStore.appCommonConfig.currency }}</p>
            </div>
            <n-hr></n-hr>
            <div class="price-discount" v-if="couponInfo.verified">
              <p class="coupon-title">{{ t('newSettlement.coupon') }}</p>
              <div class="coupon-detail">
                <p class="coupon-name">{{ couponInfo.name }}</p>
                <p class="discount">- {{ discountPrice.toFixed(2) }} {{ appInfosStore.appCommonConfig.currency }}</p>
              </div>
              <n-hr></n-hr>
            </div>
            <p class="total-title">{{ t('newSettlement.total') }}</p>
            <p class="price-large">{{ (resultPrice - discountPrice).toFixed(2) }}
              {{ appInfosStore.appCommonConfig.currency }}</p>
            <n-button
                @click="callSaveOrder"
                class="settleBtn"
                type="primary"
                :bordered="false"
                size="large"
            >
              <n-icon style="margin-right: 5px">
                <settlementIcon/>
              </n-icon>
              {{ t('newSettlement.submitOrder') }}
            </n-button>
          </n-card>
        </div>

      </div>
    </div>
  </transition>

</template>

<style scoped lang="less">
.root {
  width: auto;
  margin: 20px;

  .root-layout {
    display: flex;
    flex-direction: row;

    .l-part {
      flex: 2;

      .l-info-head {
        margin-bottom: 20px;

        .title {
          font-size: 1.5rem;
          font-weight: 600;
          margin: 20px 0 0 20px;
        }
      }

      .l-info-btn {
        display: flex;
        flex-direction: row;
        justify-content: space-between;

        .detail-title {
          font-size: 1.2rem;
          margin: 20px 0 20px 20px;
        }
      }
    }

    .r-part {
      flex: 1;

      .r-part-head {
        //background-color: v-bind(rightSideCardBgColor);

        .r-border {
          border: rgba(62, 151, 195, 0.0) 1px solid;

        }

        .r-border:hover {
          //border: rgba(62, 151, 195, 0.5) 1px solid;
        }
      }

      .r-part-btn {
        //background-color: rgba(145, 145, 145, 0.2);
        background-color: v-bind(rightSideCardBgColor);
        color: white;
        margin-top: 20px;

        .r-title {
          font-size: 1.2rem;
          font-weight: 600;
          margin: 0 0 20px 0;
          color: white;
        }

        .price-small {
          display: flex;
          flex-direction: row;
          justify-content: space-between;
          margin-bottom: 5px;

          .summary {
          }

          .price-small {
          }
        }

        .price-discount {
          margin-bottom: 10px;

          .coupon-title {
            font-size: 1rem;
            font-weight: 400;
            opacity: 0.6;
            margin-bottom: 20px;
          }

          .coupon-detail {
            display: flex;
            flex-direction: row;
            justify-content: space-between;
            margin-bottom: 5px;
          }
        }

        .total-title {
          font-size: 1rem;
          font-weight: 400;
          opacity: 0.6;
        }

        .price-large {
          margin-top: 10px;
          font-size: 2rem;
          font-weight: 600;
        }

        .settleBtn {
          width: 100%;
          margin-top: 10px;
          color: white;
        }
      }
    }

    @media (max-width: 1000px) {
      .r-part {
        margin-top: 20px;

      }
    }
    @media (min-width: 1000px) {
      .l-part {
        margin-right: 20px;
      }
    }
  }

  @media (max-width: 1000px) {
    .root-layout {
      flex-direction: column;
    }
  }
}

.n-card {
  //background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
  transition: transform 200ms ease;
}

.n-card:hover {
  transform: translateY(-5px);

}

//.n-input {
//  background-color: v-bind(rightSideCardBgColor);
//  color: white;
//}

.n-input__input-el {
  color: white;
}
</style>