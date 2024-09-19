<script setup lang="ts">
import {onMounted, reactive, ref, computed} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import {useMessage} from "naive-ui";
import {useRouter} from "vue-router";
// import VueMarkdown from 'vue-markdown';
import { MdPreview, MdCatalog } from 'md-editor-v3';
import 'md-editor-v3/lib/preview.css';
import instance from "@/axios";
import usePaymentStore from "@/stores/usePaymentStore";

const id = 'preview-only';
const text = ref('# Hello Editor');
// const scrollElement = document.documentElement;

const paymentStore = usePaymentStore();
const message = useMessage()
const apiAddrStore = useApiAddrStore();
const themeStore = useThemeStore();
const router = useRouter();

// let id = ref<number>(1)

// let medText = ref<string>('')

// let plans = ref([])

// let getAllPlans = async () => {
//   try {
//     let {data} = await instance.get(apiAddrStore.apiAddr.user.getAllPlanList)
//     if (data.code === 200) {
//       plans.value = data.plans
//       console.log(plans.value)
//     } else {
//       message.error("获取数据失败")
//     }
//   } catch (error) {
//     message.error(error)
//   }
// }

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

  paymentStore.getAllPlans()
})

</script>

<script lang="ts">
export default {
  name: 'NewPurchase',
}
</script>

<template>
  <div class="root">
    <n-card
        v-for="(item, index) in paymentStore.plan_list"
        :key="item.id"
        class="plan-item"
        content-style="padding: 0;"
    >
      <n-card
          class="plan-detail"
          :embedded="true" hoverable
          :title="item.name"
          content-style="padding: 0;"
          :style="shadowBg"
      >
        <div class="price-item">
          <div class="price-blk">
            <p class="price">{{ item.month_price }}</p>
            <p class="price-symbol">USD</p>
          </div>
          <p class="pay-method">{{ item.month_price?'月付':'更多选择' }}</p>
        </div>

        <MdPreview
            style="background-color: rgba(0,0,0,0)"
            :theme="themeStore.enableDarkMode?'dark':'light'"
            :editorId="id"
            :modelValue="item.describe"
        />
        <n-button
            @click="toPurchase(index)"
            type="primary"
            class="purchase"
        >订购</n-button>
      </n-card>

    </n-card>


  </div>
</template>

<style scoped lang="less">
.root {
  padding: 20px;
  display: flex;
  flex-wrap: wrap;
  align-items: center;

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
        width: 100px;
        margin: 20px 0 20px 20px;
      }
    }
    .plan-detail:hover {
      transform: translateY(-5px);
      transition: ease 200ms ;

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
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}
</style>