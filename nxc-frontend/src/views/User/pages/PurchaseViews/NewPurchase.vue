<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onBeforeMount, onMounted, ref, computed} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import {useMessage} from "naive-ui";
import {useRouter} from "vue-router";
import {MdPreview} from 'md-editor-v3';
import usePaymentStore from "@/stores/usePaymentStore";
import 'md-editor-v3/lib/preview.css';

const {t} = useI18n();
const id = 'preview-only';
const text = ref('# Hello Editor');
const hoveredIndex = ref<number | null>(null);

let animated = ref<boolean>(false)
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
  console.log('purchase', plan_id)
  paymentStore.plan_id_selected = plan_id // 存储选择的订阅下标
  router.push({
    path: '/dashboard/purchase/settlement'
  })
}

// 计算当前悬浮的卡片索引状态
const cardClasses = (index: number) =>
    computed(() => ({
      'is-hovered': hoveredIndex.value === index,
      'is-not-hovered': hoveredIndex.value !== null && hoveredIndex.value !== index,
    })).value;


onMounted(() => {
  themeStore.menuSelected = 'user-buy-plan'
  themeStore.userPath = '/dashboard/purchase'

  animated.value = true
})

onBeforeMount(() => {
  paymentStore.getAllPlans(0, 0)
})

</script>

<script lang="ts">
export default {
  name: 'NewPurchase',
}
</script>

<!--<template>-->
<!--  <n-p style="font-size: 1.2rem; margin: 20px 0 0 25px; font-weight: 550; opacity: 0.7">-->
<!--    {{ t('newPurchase.headerPlaceholder') }}-->
<!--  </n-p>-->

<!--  <transition name="slide-fade">-->
<!--    <div v-if="animated" style="padding: 20px">-->
<!--      <n-grid cols="1 s:1 m:2 l:3" responsive="screen" x-gap="20px" y-gap="20px">-->
<!--        <n-grid-item-->
<!--            v-for="(item, index) in paymentStore.plan_list as Plan[]"-->
<!--            :key="item.id"-->
<!--            class="plan-item"-->
<!--            content-style="padding: 0;"-->
<!--            :bordered="false"-->
<!--            @click="toPurchase(index)"-->
<!--        >-->

<!--          <n-card-->
<!--              class="plan-detail"-->
<!--              :embedded="true"-->
<!--              :title="item.name"-->
<!--              content-style="padding: 0;"-->
<!--              :style="shadowBg"-->
<!--              :bordered="false"-->
<!--          >-->
<!--            <div class="price-item">-->
<!--              <div class="price-blk">-->
<!--                <p class="price">{{ item.month_price?.toFixed(2) }}</p>-->
<!--                <p class="price-symbol">{{ appInfosStore.appCommonConfig.currency }}</p>-->
<!--              </div>-->
<!--              <p class="pay-method">{{ item.month_price ? t('newPurchase.monthPay') : t('newPurchase.moreMethod') }}</p>-->
<!--            </div>-->

<!--            <MdPreview-->
<!--                style="background-color: rgba(0,0,0,0)"-->
<!--                :theme="themeStore.enableDarkMode?'dark':'light'"-->
<!--                :editorId="`preview-${index}`"-->
<!--                :modelValue="item.describe || ''"-->
<!--            />-->
<!--            <n-button-->
<!--                @click=""-->
<!--                type="primary"-->
<!--                class="purchase"-->
<!--                :bordered="false"-->
<!--            >-->
<!--              {{ t('newPurchase.purchaseBtn') }}-->
<!--            </n-button>-->
<!--          </n-card>-->

<!--        </n-grid-item>-->
<!--      </n-grid>-->
<!--    </div>-->
<!--  </transition>-->


<!--</template>-->

<template>
    <n-p style="font-size: 1.2rem; margin: 20px 0 0 25px; font-weight: 550; opacity: 0.7">
      {{ t('newPurchase.headerPlaceholder') }}
    </n-p>
  <transition name="slide-fade">
    <div v-if="animated" style="padding: 20px">
      <n-grid cols="1 s:1 m:2 l:3" responsive="screen" x-gap="20px" y-gap="20px">
        <n-grid-item
            v-for="(item, index) in paymentStore.plan_list as Plan[]"
            :key="item.id"
            class="plan-item"
            content-style="padding: 0;"
            :bordered="false"
            @click="toPurchase(index)"
            @mouseover="hoveredIndex = index"
            @mouseleave="hoveredIndex = null"
        >
          <n-card
              class="plan-detail"
              :class="cardClasses(index)"
              :embedded="true"
              :title="item.name"
              content-style="padding: 0;"
              :style="shadowBg"
              :bordered="false"
          >
            <div class="price-item">
              <div class="price-blk">
                <p class="price">{{ item.month_price?.toFixed(2) }}</p>
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
                @click=""
                type="primary"
                class="purchase"
                :bordered="false"
            >
              {{ t('newPurchase.purchaseBtn') }}
            </n-button>
          </n-card>
        </n-grid-item>
      </n-grid>
    </div>
  </transition>
</template>


<!--<style scoped lang="less">-->

<!--.plan-item {-->
<!--  background-color: rgba(0, 0, 0, 0);-->

<!--  .plan-detail {-->
<!--    box-shadow: 0 1px 2px -2px rgba(0, 0, 0, .04), 0 3px 6px 0 rgba(0, 0, 0, .03), 0 5px 12px 4px rgba(0, 0, 0, .02);-->

<!--    .price-item {-->
<!--      background-color: rgba(188, 188, 188, 0.2);-->
<!--      padding: 0 0 10px 20px;-->

<!--      .price-blk {-->
<!--        display: flex;-->
<!--        flex-direction: row;-->
<!--        align-items: baseline;-->
<!--        padding: 15px 0 15px 0;-->

<!--        .price {-->
<!--          font-size: 2.5rem;-->
<!--          font-weight: 400;-->
<!--          opacity: 0.9;-->
<!--        }-->

<!--        .price-symbol {-->
<!--          font-size: 1.2rem;-->
<!--          font-weight: 400;-->
<!--          margin-left: 5px;-->
<!--          opacity: 0.8;-->
<!--        }-->
<!--      }-->

<!--      .pay-method {-->
<!--        font-size: 1rem;-->
<!--        font-weight: 500;-->
<!--        opacity: 0.7;-->
<!--      }-->
<!--    }-->

<!--    .plan-describe {-->
<!--      padding: 10px 0 10px 20px;-->
<!--    }-->

<!--    .purchase {-->
<!--      width: auto;-->
<!--      margin: 20px 0 20px 20px;-->
<!--    }-->
<!--  }-->

<!--  .plan-detail:hover {-->
<!--    transform: translateY(-5px);-->
<!--    //transition: ease 200ms ;-->
<!--  }-->

<!--}-->

<!--.n-card {-->
<!--  transition: transform 200ms ease;-->
<!--}-->
<!--</style>-->

<style scoped lang="less">
.plan-item {
  background-color: rgba(0, 0, 0, 0); /* 背景透明 */

  .plan-detail {
    box-shadow: 0 1px 2px -2px rgba(0, 0, 0, 0.04), 0 3px 6px 0 rgba(0, 0, 0, 0.03), 0 5px 12px 4px rgba(0, 0, 0, 0.02);
    transition: transform 200ms ease, opacity 200ms ease;

    /* 鼠标悬浮时的效果 */

    &.is-hovered {
      transform: translateY(-3px); /* 微微上移 */
      opacity: 1; /* 完全不透明 */
    }

    /* 其他卡片变透明 */

    &.is-not-hovered {
      opacity: 0.7; /* 半透明 */
      filter: blur(0.01px);
    }

    .price-item {
      background-color: rgba(188, 188, 188, 0.2);
      padding: 0 0 10px 20px;

      .price-blk {
        display: flex;
        flex-direction: row;
        align-items: baseline;
        padding: 15px 0;

        .price {
          font-size: 2.5rem;
          font-weight: 400;
          opacity: 0.9;
        }

        .price-symbol {
          font-size: 1.2rem;
          font-weight: 400;
          margin-left: 5px;
          opacity: 0.8;
        }
      }

      .pay-method {
        font-size: 1rem;
        font-weight: 500;
        opacity: 0.7;
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
}

.n-card {
  transition: transform 200ms ease, opacity 200ms ease;
}
</style>
