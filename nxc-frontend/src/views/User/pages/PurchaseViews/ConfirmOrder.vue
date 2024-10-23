<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, onBeforeMount, onMounted, ref, onBeforeUnmount} from "vue"
import {NIcon} from 'naive-ui'
import instance from "@/axios/index"
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useApiAddrStore from "@/stores/useApiAddrStore"
import usePaymentStore from "@/stores/usePaymentStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import {useRouter} from "vue-router";
import 'md-editor-v3/lib/preview.css';
let rightSideCardBgColor = computed(() => themeStore.enableDarkMode ? 'rgba(54,55,55,0.8)' : 'rgba(54,56,61,0.8)')
const {t} = useI18n();
const userInfoStore = useUserInfoStore()
const appInfosStore = useAppInfosStore();
const themeStore = useThemeStore();
const apiAddrStore = useApiAddrStore();
const paymentStore = usePaymentStore();
const router = useRouter();

interface ConfirmOrder {
  code: number | null,
  order_id: number | null,
  plan_name: string | null,
  original_price: number | null,
  discount_amount: number | null,
  pay_price: number | null,
  period: string | null,
  created_at: string | null
}

let confirmOrder = ref<ConfirmOrder>()


let goodInfoData = ref([
  {
    title: computed(() => '商品信息'),
    content: computed(() => confirmOrder.value?.plan_name)
  },
  {
    title: computed(() => '周期/类型'),
    content: computed(() => {
      switch (confirmOrder.value?.period) {
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
    content: computed(() => confirmOrder.value?.order_id)
  },
  {
    title: computed(() => '创建日期'),
    content: computed(() => confirmOrder.value?.created_at)
  },
  {
    title: computed(() => '订单号'),
    content: computed(() => confirmOrder.value?.plan_name)
  },

])

let queryOrderInfoAndStatus = async () => {
  try {
    let { data } = await instance.get('/api/user/v1/order/status', {
      params: {
        user_id: userInfoStore.thisUser.id,
        order_id: paymentStore.submittedOrderId.trim()
      }
    });

    if (data.code === 200) {
      console.log('获取成功', data);
      // 如果订单状态已完成或达到某个条件，停止轮询
      if (data.status === 'completed') {
        clearInterval(intervalId); // 清除轮询
        console.log('订单已完成，停止轮询');
      }
    }
  } catch (error: any) {
    console.log(error);
  }
}

let intervalId: any;

let queryOrderInterval = () => {
  intervalId = setInterval(() => {
    queryOrderInfoAndStatus();
  }, 1000); // 每隔5秒查询一次
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
  queryOrderInterval();
})

onBeforeMount(() => {

})

onBeforeUnmount(() => {
  clearInterval(intervalId)
})

</script>

<script lang="ts">
export default {
  name: 'ConfirmOrder',
}
</script>

<template>
  <div class="root" v-if="paymentStore.plan_id_selected != -1">
    <div class="root-layout">
      <div class="l-part">
        <n-card
            class="good-info"
            :embedded="true"
            hoverable
            content-style="padding: 0"
            :bordered="false"
        >
          <n-p class="title">商品信息</n-p>
          <div class="good-info-item" v-for="item in goodInfoData" :key="item.title">
            <p class="item-head">{{ item.title }}</p>
            <p class="item-content">{{ item.content }}</p>
          </div>
        </n-card>

        <n-card
            class="good-info"
            :embedded="true"
            hoverable
            content-style="padding: 0"
            :bordered="false"
        >
          <n-p class="title">订单信息</n-p>
          <div class="good-info-item" v-for="item in orderData" :key="item.title">
            <p class="item-head">{{ item.title }}</p>
            <p class="item-content">{{ item.content }}</p>
          </div>

          <n-hr style="margin: 20px"></n-hr>

          <n-button
              class="change-plan-btn"
              type="primary"
              tertiary
          >
            切换订阅
          </n-button>
          <n-button
              class="cancel-order-btn"
              type="error"
              text
          >
            取消订单
          </n-button>
        </n-card>
      </div>

      <div class="r-part">
        <n-card style="color: white" class="r-part-btn" :embedded="true" hoverable>
          <p class="r-title">您的价格</p>
          1
          <div class="price-small">
            <p class="summary">{{ 111 }} x {{ 222 }}</p>
            <p class="price-small">{{ 222 }} </p>
          </div>
          <n-hr></n-hr>
          <div class="price-discount" v-if="true">
            <p class="coupon-title">1</p>
            <div class="coupon-detail">
              <p class="coupon-name">{{ '优惠券名' }}</p>
              <p class="discount">- {{ 11.11 }} </p>
            </div>
            <n-hr></n-hr>
            22222222222222
          </div>
          <p class="total-title">统计</p>
          <p class="price-large">{{ 11 }} </p>
          <n-button
              @click=""
              class="settleBtn"
              type="primary"
              :bordered="false"
              size="large"
          >
            <n-icon style="margin-right: 5px">
            </n-icon>
            提交
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

      .change-plan-btn {
        margin-left: 20px;
        width: 120px;
      }

      .cancel-order-btn {
        margin-left: 20px;
        width: 80px;

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