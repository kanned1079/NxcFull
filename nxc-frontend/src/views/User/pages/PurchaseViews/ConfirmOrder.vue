<script setup lang="ts">
import ConfirmOrderBefore from "@/views/User/pages/PurchaseViews/ConfirmOrderBefore.vue";
import GoodInfo from "@/views/User/pages/PurchaseViews/Parts/GoodInfo.vue";
import OrderInfo from "@/views/User/pages/PurchaseViews/Parts/OrderInfo.vue";
import {useI18n} from "vue-i18n";
import {computed, onBeforeMount, onBeforeUnmount, onMounted, ref} from "vue"
import {NIcon, useMessage, useNotification} from 'naive-ui'
import {cancelOrder, placeOrder, queryOrderInfoAndStatus,} from "@/api/user/confirm";
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import usePaymentStore from "@/stores/usePaymentStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import {useRouter} from "vue-router";
import 'md-editor-v3/lib/preview.css';
import {formatDate} from "@/utils/timeFormat";
import {CheckmarkCircleOutline as confirmIcon} from "@vicons/ionicons5"

let rightSideCardBgColor = computed(() => themeStore.enableDarkMode ? 'rgba(54,55,55,0.8)' : 'rgba(54,56,61,0.8)')
const {t} = useI18n();
const message = useMessage()
const notify = useNotification()
const userInfoStore = useUserInfoStore()
const appInfosStore = useAppInfosStore();
const themeStore = useThemeStore();
const paymentStore = usePaymentStore();
const router = useRouter();

// let showAnimation = ref<boolean>(false)
let animated = ref<boolean>(false)
let showLoading = ref<boolean>(true)

interface ConfirmOrder {
  code: number | null,
  order_id: string,
  plan_name: string | null,
  price: number, // 原始价格
  discount_amount?: number, // 抵折金额
  amount: number, // 最终价格
  pay_price: number | null,
  period: string | null,
  coupon_name?: string,
  created_at: string | null
}

let confirmOrder = ref<ConfirmOrder>({
  code: 0,
  order_id: '',
  plan_name: '',
  price: 0.00,
  discount_amount: 0.00,
  amount: 0.00,
  pay_price: 0.00,
  period: '',
  created_at: '',
  coupon_name: '',
})

let periodCode2Text = computed(() => { // pass
  switch (confirmOrder.value?.period) {
    case 'month': {
      return computed(() => t('userConfirmOrder.monthPay')).value
    }
    case 'quarter': {
      return computed(() => t('userConfirmOrder.quarterPay')).value
    }
    case 'half-year': {
      return computed(() => t('userConfirmOrder.halfYearPay')).value
    }
    case 'year': {
      return computed(() => t('userConfirmOrder.yearPay')).value
    }
  }
})


let goodInfoData = ref([
  {
    title: computed(() => t('userConfirmOrder.goodInfo')),
    content: computed(() => confirmOrder.value?.plan_name)
  },
  {
    title: computed(() => t('userConfirmOrder.cycleOrType')),
    content: computed(() => { // pass
      switch (confirmOrder.value?.period) {
        case 'month': {
          return computed(() => t('userConfirmOrder.monthPay')).value
        }
        case 'quarter': {
          return computed(() => t('userConfirmOrder.quarterPay')).value
        }
        case 'half-year': {
          return computed(() => t('userConfirmOrder.halfYearPay')).value
        }
        case 'year': {
          return computed(() => t('userConfirmOrder.yearPay')).value
        }
      }
    })
  }
])

let orderData = ref([
  {
    title: computed(() => t('userConfirmOrder.orderNumber')),
    content: computed(() => confirmOrder.value?.order_id)
  },
  {
    title: computed(() => t('userConfirmOrder.createdAt')),
    content: computed(() => confirmOrder.value?.created_at ? formatDate(confirmOrder.value?.created_at) : '')
  },
  // {
  //   title: computed(() => '订单号'),
  //   content: computed(() => confirmOrder.value?.plan_name)
  // },

])

let intervalId = ref<number | null>();

let callQueryOrderInfoAndStatus = async () => {
  // try {
  // let {data} = await instance.get('/api/user/v1/order/status', {
  //   params: {
  //     user_id: userInfoStore.thisUser.id,
  //     order_id: paymentStore.submittedOrderId.trim()
  //   }
  // });
  let data = await queryOrderInfoAndStatus(userInfoStore.thisUser.id, paymentStore.submittedOrderId)
  if (data.code === 200) {
    console.log('获取成功', data);
    Object.assign(confirmOrder.value, data.order_info)
    console.log('订单号：', data.order_info.order_id)
    setTimeout(async () => {
      showLoading.value = false
      animated.value = true
    }, 2000)
    // showLoading.value = false
    // 如果订单状态已完成或达到某个条件，停止轮询
    if (data.order_info.is_finished && data.order_info.is_success) {  // 订单完成且成功
      intervalId.value ? clearInterval(intervalId.value) : null
      message.info('订单成功')
      message.info('去往下一个页面')
      // toOrderResult(data.code)
      toOrderResult(data.code)
    } else if (data.order_info.is_finished && !data.order_info.is_success) {
      message.info('订单完成但未成功')
      // clearInterval(intervalId)
      intervalId.value ? clearInterval(intervalId.value) : null
      toOrderResult(data.code)  // 去往下一个页面
    }
    console.log('订单未支付，进行下一次轮询查询')
  } else {
    console.log('订单已超时，停止轮询');
    message.error(t('userConfirmOrder.orderExpired'))
    // clearInterval(intervalId); // 清除轮询
    intervalId.value ? clearInterval(intervalId.value) : null
    router.back()

  }
// pass
//   } catch (error: any) {
//     console.log(error);
//   }
}


let queryOrderInterval = () => {
  intervalId.value = setInterval(() => {
    // queryOrderInfoAndStatus();
    callQueryOrderInfoAndStatus()
  }, 3000); // 每隔3秒查询一次
}

let toOrderResult = (code: number) => {
  console.log('转到结果页', code)
  // router.push({
  //   path: '',
  // })
  router.push({
    path: '/dashboard/orders'
  })
}

let callPlaceOrder = async () => {
  // 点击提交按钮立刻停止定时器
  intervalId.value ? clearInterval(intervalId.value) : null
  // try {
  //   let {data} = await instance.put('/api/user/v1/order', {
  //     user_id: userInfoStore.thisUser.id,
  //     order_id: confirmOrder.value.order_id,
  //   })
  let data = await placeOrder(userInfoStore.thisUser.id, confirmOrder.value.order_id)
  if (!data) return router.back()
  if (data.code === 200) {
    console.log(data)
    if (data.placed && data.key_generated) {
      console.log("订单成功")
      // clearInterval(intervalId) // 停止定时器
      // intervalId.value?clearInterval(intervalId.value):null
      message.success(t('userConfirmOrder.paySuccessfully'))
      setTimeout(async () => {
        await router.push({
          path: '/dashboard/keys'
        })
      }, 1000)
    } else {
      message.error('unknow err')
    }
  } else if (data.code == 402) {
    // 用户余额不足
    message.warning(t('userConfirmOrder.balanceNotEnough'))
    // await queryOrderInfoAndStatus()
    await callQueryOrderInfoAndStatus()
  } else {
    message.error(t('userConfirmOrder.orderErrorOccur') + data.msg)
    // clearInterval(intervalId)
    intervalId.value ? clearInterval(intervalId.value) : null
    router.back()
  }
  // } catch (error: any) {
  //   console.log(error)
  //   router.back()
  // }
}

let callCancelOrder = async () => {
  // try {
  //   let {data} = await instance.delete('/api/user/v1/order', {
  //     params: {
  //       user_id: userInfoStore.thisUser.id,
  //       order_id: confirmOrder.value.order_id,
  //     }
  //   })
  let data = await cancelOrder(userInfoStore.thisUser.id, confirmOrder.value.order_id)
  if (data.code === 200) {
    console.log(data)
    message.info(t('userConfirmOrder.orderCancelled'))
    // clearInterval(intervalId)
    intervalId.value ? clearInterval(intervalId.value) : null
    router.back()
  }
  // } catch (error: any) {
  //   console.log(error)
  //
  // }
}

let reChoosePlan = async () => {
  // await cancelOrder()
  await callCancelOrder()
  router.back() // 此处只需要返回一次即可 因为cancelOrder中也有一次back
}

onMounted(() => {
  console.log('settlement挂载', paymentStore.plan_id_selected)
  // appendCycleOptions()
  themeStore.userPath = '/dashboard/purchase/confirm'

  console.log(paymentStore.confirmOrder)
  // if (paymentStore.confirmOrder)
  //   confirmOrder.value = paymentStore.confirmOrder

  // queryOrderInfoAndStatus()
  // 开始轮询
  // setTimeout(() => showAnimation.value = true, 0)
  // showAnimation.value = true
  // animated.value = true
})

onBeforeMount(async () => {
  // await queryOrderInfoAndStatus()
  await callQueryOrderInfoAndStatus()
  queryOrderInterval();
})

onBeforeUnmount(() => {
  // clearInterval(intervalId)
  intervalId.value ? clearInterval(intervalId.value) : null
})

</script>

<script lang="ts">
export default {
  name: 'ConfirmOrder',
}
</script>

<template>
  <!--  <div class="root" v-if="paymentStore.plan_id_selected != -1">-->

  <ConfirmOrderBefore v-if="showLoading"></ConfirmOrderBefore>

  <transition name="slide-fade">
    <div v-if="animated" class="root">
      <div class="root-layout">
        <div class="l-part">
          <GoodInfo :goodInfoData="goodInfoData"></GoodInfo>
          <OrderInfo :orderData="orderData"></OrderInfo>
          <n-card hoverable :bordered="false" :embedded="true" content-style="padding: 10px">
            <n-button
                @click="reChoosePlan"
                class="change-plan-btn"
                type="primary"
                tertiary
            >
              {{ t('userConfirmOrder.switchPlan') }}
            </n-button>
            <n-button
                @click="callCancelOrder"
                class="cancel-order-btn"
                type="error"
                text
            >
              <!--            取消订单-->
              {{ t('userConfirmOrder.cancelOrder') }}

            </n-button>
          </n-card>
        </div>

        <div class="r-part">
          <n-card style="color: white" class="r-part-btn" :bordered="false" :embedded="true" hoverable>
            <p class="r-title">{{ t('userConfirmOrder.yourPrice') }}</p>

            <div class="price-small">
              <p class="summary">{{ confirmOrder.plan_name }} x {{ periodCode2Text }}</p>
              <p class="price-small">{{ confirmOrder.price }} {{ appInfosStore.appCommonConfig.currency }}</p>
            </div>
            <n-hr></n-hr>
            <div class="price-discount" v-if="confirmOrder.discount_amount != 0">
              <p class="coupon-title">{{ t('userConfirmOrder.couponOffset') }}</p>
              <div class="coupon-detail">
                <p class="coupon-name">{{ confirmOrder.coupon_name }}</p>
                <p class="discount">- {{ confirmOrder.discount_amount?.toFixed(2) }}
                  {{ appInfosStore.appCommonConfig.currency }}</p>
              </div>
              <n-hr v-if="confirmOrder.discount_amount != 0"></n-hr>
            </div>
            <p class="total-title">{{ t('userConfirmOrder.price') }}</p>
            <p class="price-large">{{ confirmOrder.amount.toFixed(2) }} {{ appInfosStore.appCommonConfig.currency }}</p>
            <n-button
                @click="callPlaceOrder"
                class="settleBtn"
                type="primary"
                :bordered="false"
                size="large"
            >
              <n-icon style="margin-right: 5px">
                <confirmIcon/>
              </n-icon>
              <!--            提交-->
              {{ t('userConfirmOrder.submit') }}
            </n-button>
          </n-card>
        </div>

      </div>
    </div>
  </transition>

</template>

<style scoped lang="less">

.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}

.fade-slide-enter-active, .fade-slide-leave-active {
  transition: opacity 0.3s, transform 0.3s;
}

.fade-slide-enter, .fade-slide-leave-to {
  opacity: 0;
  transform: translateY(10px);
}

.fade-slide-enter-to {
  opacity: 1;
  transform: translateY(0);
}

.root {
  width: auto;
  margin: 20px;

  .root-layout {
    display: flex;
    flex-direction: row;

    .l-part {
      flex: 2;

      .good-info {
        padding-bottom: 15px;
        margin-bottom: 20px;

        .title {
          font-size: 1.1rem;
          font-weight: 400;
          background-color: rgba(216, 216, 216, 0.1);
          padding: 10px 20px 15px 20px;
          border-radius: 3px 3px 0 0;
        }

        .good-info-item {
          display: flex;
          flex-direction: row;
          margin: 0 0 10px 20px;
          //background-color: #66afe9;

          .item-head {
            font-size: 0.9rem;
            opacity: 0.8;
            font-weight: 400;
            min-width: 100px;
            //background-color: #66afe9;
          }

          .item-content {
            font-size: 0.9rem;
            margin-left: 120px;
            opacity: 1;
            font-weight: 400;
          }

        }
      }


    }

    .change-plan-btn {
      width: 120px;
    }

    .cancel-order-btn {
      margin-left: 20px;
      width: 80px;

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
        //margin-top: 20px;

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
        //margin-top: 20px;

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


</style>