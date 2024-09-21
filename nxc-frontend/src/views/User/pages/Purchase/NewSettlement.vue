<script setup lang="ts">
import {onBeforeMount, onMounted, ref} from "vue"
import type {MenuOption, NotificationType} from 'naive-ui'
import {NIcon, useNotification} from 'naive-ui'
import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore"
import usePaymentStore from "@/stores/usePaymentStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import {useRouter} from "vue-router";
import {MdPreview} from 'md-editor-v3';
import 'md-editor-v3/lib/preview.css';
import {CheckmarkCircleOutline as settlementIcon, TicketOutline as ticketIcon} from "@vicons/ionicons5"
import instance from "@/axios";
import useUserInfoStore from "@/stores/useUserInfoStore";

const notification = useNotification()
const userInfoStore = useUserInfoStore();
const appInfosStore = useAppInfosStore();
const themeStore = useThemeStore();
const apiAddrStore = useApiAddrStore();
const paymentStore = usePaymentStore();
const router = useRouter();

interface Plan {
  id: number
  group_id?: number
  is_renew?: boolean
  is_sale?: boolean
  name: string
  capacity: number
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


let showCycleName = computed(() => {
  if (selectedCycle.value === 'month') {
    resultPrice.value = plan.month_price
    return '月付'
  }
  if (selectedCycle.value === 'quarter') {
    resultPrice.value = plan.quarter_price
    return '季付'
  }
  if (selectedCycle.value === 'half-year') {
    resultPrice.value = plan.half_year_price
    return '半年付'
  }
  if (selectedCycle.value === 'year') {
    resultPrice.value = plan.year_price
    return '年付'
  }
})

let cycleSelect = (key: string) => {
  console.log(key)
  selectedCycle.value = key
}

const menuOptions: MenuOption[] = []

let appendCycleOptions = () => {
  if (plan.month_price) {
    menuOptions.push({
      label: '月付',
      key: 'month',
      disabled: false,
    })
  }
  if (plan.quarter_price) {
    menuOptions.push({
      label: '季付',
      key: 'quarter',
      disabled: false,
    })
  }
  if (plan.half_year_price) {
    menuOptions.push({
      label: '半年付',
      key: 'half-year',
      disabled: false,
    })
  }
  if (plan.year_price) {
    menuOptions.push({
      label: '年付',
      key: 'year',
      disabled: false,
    })
  }
  console.log(menuOptions)
}

let handleUpdateExpandedKeys = (keys: string) => {
  console.log(keys)
}

// 优惠券信息
let couponInfo = ref({
  name: '',
  verified: false,
  percent_off: 0.0,
})

let verifyTicket = async () => {
  if (couponCode.value.trim() === '') {
    notify('error', '错误', '输入的优惠券码不能为空')
    return
  }
  console.log('验证优惠券')
  try {
    let {data} = await instance.post(apiAddrStore.apiAddr.user.verifyCoupon, {
      coupon_code: couponCode.value.trim(),  // 优惠券码
      plan_id: plan.id,               // 订阅计划id
      user_id: userInfoStore.thisUser.id  // 用户id 判断是否使用过
    });
    if (data.code === 200) {
      if (data.verified) {
        console.log('验证码有效');
        // ticketVerified.value = true;
        couponInfo.value.verified = true;
        couponInfo.value.percent_off = data.percent_off;
        couponInfo.value.name = data.name;
        await notify('success', '验证通过', '优惠券有效');
      } else {
        couponInfo.value.verified = false
        // console.log('未知错误', data.msg);
        await notify('error', '优惠券不可用', data.msg);
      }
    } else {
      couponInfo.value.verified = false
      // 处理非200状态码的错误
      await notify('error', '优惠券不可用', data.msg);
    }
  } catch (error) {
    // 处理网络错误或后端500错误
    couponInfo.value.verified = false
    await notify('error', '优惠券不可用', error.toString());

  }
}

// 抵折金额
let discountPrice = computed(()  => couponInfo.value.verified ? ((couponInfo.value.percent_off / 100) * resultPrice.value) : 0)



let saveOrder = async () => {
  console.log('提交订单准备支付')


  // try{
  //   let {data} = await instance.post(apiAddrStore.apiAddr.user.saveOrder, {
  //     plan_id: plan.id,
  //   })
  //   if (data.code === 200) {
  //     console.log('订单创建成功')
  //   }
  // }catch (error) {
  //   console.log(error)
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

})

onBeforeMount(() => {
  console.log('settlement挂载前', paymentStore.plan_id_selected)
  if (paymentStore.plan_id_selected < 0) {
    router.back(-1)
  } else{
    appendCycleOptions()

  }
})

</script>

<script lang="ts">
export default {
  name: 'NewSettlement',
}
</script>

<template>
  <div class="root" v-if="paymentStore.plan_id_selected != -1">
    <div class="root-layout">
      <div class="l-part">
        <n-card class="l-info-head" :embedded="true" hoverable content-style="padding: 0;">
          <p class="title">{{ plan.name }}</p>
          <MdPreview
              style="background-color: rgba(0,0,0,0)"
              :theme="themeStore.enableDarkMode?'dark':'light'"
              :editorId="id"
              :modelValue="plan.describe"
          />
        </n-card>
        <n-card class="l-info-btn" :embedded="true" hoverable content-style="padding: 0;">
          <p class="detail-title">付款周期</p>
          <n-menu
              :options="menuOptions"
              @update:value="cycleSelect"
              :value="selectedCycle"
              style="font-size: 0.9rem"
          />
        </n-card>
      </div>

      <div class="r-part">
        <n-card class="r-part-head" :embedded="true" hoverable>
          <!--          <n-input-group>-->
          <div class="r-border" style="display: flex; flex-direction: row; border-radius: 5px; padding: 5px">
            <n-input
                size="large"
                :bordered="true"
                class="ticket-input"
                style="margin-right: 10px"
                v-model:value="couponCode"
                placeholder="有优惠券？"
            />
            <n-button
                style="color: white"
                size="large"
                :keyboard="true"
                type="primary"
                @click="verifyTicket"
            >
              <n-icon style="margin-right: 10px">
                <ticketIcon/>
              </n-icon>
              验证
            </n-button>
          </div>

          <!--          </n-input-group>-->
        </n-card>
        <n-card style="color: white" class="r-part-btn" :embedded="true" hoverable>
          <p class="r-title">订单总额</p>
          <div class="price-small">
            <p class="summary">{{ plan.name }} x {{ showCycleName }}</p>
            <p class="price-small">{{ resultPrice }} {{ appInfosStore.appCommonConfig.currency }}</p>
          </div>
          <n-hr></n-hr>
          <div class="price-discount" v-if="couponInfo.verified">
            <p class="coupon-title">优惠券</p>
            <div class="coupon-detail">
              <p class="coupon-name">{{ couponInfo.name }}</p>
              <p class="discount">- {{ discountPrice.toFixed(2) }} {{ appInfosStore.appCommonConfig.currency }}</p>
            </div>
            <n-hr></n-hr>
          </div>
          <p class="total-title">总计</p>
          <p class="price-large">{{ (resultPrice - discountPrice).toFixed(2) }} {{ appInfosStore.appCommonConfig.currency }}</p>
          <n-button
              @click="saveOrder"
              class="settleBtn"
              type="primary"
          >
            <n-icon>
              <settlementIcon/>
            </n-icon>
            下单
          </n-button>
        </n-card>
      </div>

    </div>
  </div>
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
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
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