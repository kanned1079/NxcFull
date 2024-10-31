<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onBeforeMount, onMounted, ref} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import {useMessage} from "naive-ui";
import {useRouter} from "vue-router";
// import VueMarkdown from 'vue-markdown';
import {MdPreview} from 'md-editor-v3';
// import instance from "@/axios";
import usePaymentStore from "@/stores/usePaymentStore";
import 'md-editor-v3/lib/preview.css';

const {t} = useI18n();
const id = 'preview-only';
const text = ref('# Hello Editor');

let animated = ref<boolean>(false)

// let isMounted = ref<boolean>(false)

// let transparentDeg = ref<number>(0.0)
// const scrollElement = document.documentElement;

const paymentStore = usePaymentStore();
const message = useMessage()
const apiAddrStore = useApiAddrStore();
const themeStore = useThemeStore();
const appInfosStore = useAppInfosStore();
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

let shadowBg = () => {
  if (themeStore.enableDarkMode) {
    return ({
      filter: 'drop-shadow(5px 5px 10px rgba(220, 220, 220, 0.9))'
    })
  } else {
    return ({
      filter: 'drop-shadow(5px 5px 10px rgba(220, 220, 220, 0.3))'
    })
  }
}

let toPurchase = (plan_id: number) => {
  console.log('订购', plan_id)
  paymentStore.plan_id_selected = plan_id // 存储选择的订阅下标
  router.push({
    path: '/dashboard/purchase/settlement'
  })
}


onMounted(() => {
  themeStore.menuSelected = 'user-buy-plan'
  themeStore.userPath = '/dashboard/purchase'
  // paymentStore.getAllPlans()
  // setTimeout(() => isMounted.value = true, 0)
  animated.value = true
})

onBeforeMount(() => {
  paymentStore.getAllPlans()
})

</script>

<script lang="ts">
export default {
  name: 'NewPurchase',
}
</script>

<template>

  <transition>

  </transition>
  <n-p style="font-size: 1.2rem; margin: 20px 0 0 30px; font-weight: 600;">{{
      t('newPurchase.headerPlaceholder')
    }}
  </n-p>

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card
          v-for="(item, index) in paymentStore.plan_list as Plan[]"
          :key="item.id"
          class="plan-item"
          content-style="padding: 0;"
          :bordered="false"
      >
        <n-card
            class="plan-detail"
            :embedded="true" hoverable
            :title="item.name"
            content-style="padding: 0;"
            :style="shadowBg"
            :bordered="false"
        >
          <div class="price-item">
            <div class="price-blk">
              <p class="price">{{ item.month_price }}</p>
              <p class="price-symbol">{{ appInfosStore.appCommonConfig.currency }}</p>
            </div>
            <p class="pay-method">{{ item.month_price ? t('newPurchase.monthPay') : t('newPurchase.moreMethod') }}</p>
          </div>

          <MdPreview
              style="background-color: rgba(0,0,0,0)"
              :theme="themeStore.enableDarkMode?'dark':'light'"
              :editorId="`preview-${index}`"
              :modelValue="item.describe || ''"
          />
          <n-button
              @click="toPurchase(index)"
              type="primary"
              class="purchase"
              :bordered="false"
          >
            {{ t('newPurchase.purchaseBtn') }}
          </n-button>
        </n-card>

      </n-card>
    </div>
  </transition>


</template>

<style scoped lang="less">
.root {
  padding: 20px;
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  //opacity: 0;
  //transform: translateX(0px) translateY(10px);
  //transition: ease 300ms;

  .plan-item {
    background-color: rgba(0, 0, 0, 0);
    padding: 10px;

    .plan-detail {
      //box-shadow: 5px 5px 20px rgba(221, 221, 221, 0.5);
      //filter: drop-shadow(5px 5px 10px rgba(220, 220, 220, 0.3));

      .price-item {
        background-color: rgba(188, 188, 188, 0.2);
        padding: 0 0 10px 20px;

        .price-blk {
          display: flex;
          flex-direction: row;
          align-items: baseline;
          padding: 15px 0 15px 0;

          .price {
            font-size: 2.5rem;
            font-weight: 400;
          }

          .price-symbol {
            font-size: 1.2rem;
            font-weight: 400;
            margin-left: 5px;
          }
        }

        .pay-method {
          font-size: 1rem;
          font-weight: 400;
          opacity: 0.8;
        }
      }

      .plan-describe {
        padding: 10px 0 10px 20px;
      }

      .purchase {
        width: auto;
        margin: 20px 0 20px 20px;
      }
    }

    .plan-detail:hover {
      transform: translateY(-5px);
      //transition: ease 200ms ;
    }

  }

  @media screen and (min-width: 1200px) {
    .plan-item {
      flex: 0 0 33.333333%;
      max-width: 33.333333%;
    }
  }

}

.n-card {
  transition: transform 200ms ease;
}
</style>