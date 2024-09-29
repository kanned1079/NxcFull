<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onBeforeMount, onMounted, ref, computed} from "vue"
import {type MenuOption, NIcon} from 'naive-ui'
import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore"
import usePaymentStore from "@/stores/usePaymentStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import {useRouter} from "vue-router";
import 'md-editor-v3/lib/preview.css';
import instance from "@/axios";

const {t} = useI18n();
const appInfosStore = useAppInfosStore();
const themeStore = useThemeStore();
const apiAddrStore = useApiAddrStore();
const paymentStore = usePaymentStore();
const router = useRouter();

let goodInfoData = ref([
  {
    title: computed(() => '商品信息'),
    content: computed(() => paymentStore.confirmOrder.plan_name)
  },
  {
    title: computed(() => '周期/类型'),
    content: computed(() => {
      switch (paymentStore.confirmOrder.period) {
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
    content: computed(() => paymentStore.confirmOrder.order_id)
  },
  {
    title: computed(() => '创建日期'),
    content: computed(() => paymentStore.confirmOrder.created_at)
  },
  {
    title: computed(() => '订单号'),
    content: computed(() => paymentStore.confirmOrder.plan_name)
  },

])



onMounted(() => {
  console.log('settlement挂载', paymentStore.plan_id_selected)
  // appendCycleOptions()
  themeStore.userPath = '/dashboard/purchase/confirm'

  console.log(paymentStore.confirmOrder)

})

onBeforeMount(() => {

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

          <n-button class=""> 关闭订单</n-button>
        </n-card>
      </div>

      <div class="r-part">

        <n-card style="color: white" class="r-part-btn" :embedded="true" hoverable>
          <p class="r-title">{{ t('newSettlement.orderTotalTitle') }}</p>
          <div class="price-small">
            <p class="summary">{{ 111 }} x {{ 222 }}</p>
            <p class="price-small">{{ 222 }} {{ appInfosStore.appCommonConfig.currency }}</p>
          </div>
          <n-hr></n-hr>
          <div class="price-discount" v-if="true">
            <p class="coupon-title">{{ t('newSettlement.coupon') }}</p>
            <div class="coupon-detail">
              <p class="coupon-name">{{  }}</p>
              <p class="discount">- {{ 11.11 }} {{ appInfosStore.appCommonConfig.currency }}</p>
            </div>
            <n-hr></n-hr>
          </div>
          <p class="total-title">{{ t('newSettlement.total') }}</p>
          <p class="price-large">{{ 11 }} {{ appInfosStore.appCommonConfig.currency }}</p>
          <n-button
              @click=""
              class="settleBtn"
              type="primary"
              :bordered="false"
              size="large"
          >
            <n-icon style="margin-right: 5px">
            </n-icon>
            {{ t('newSettlement.submitOrder') }}
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
          background-color: rgba(216,216,216,0.1);
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

    .r-part {
      flex: 1;

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




</style>