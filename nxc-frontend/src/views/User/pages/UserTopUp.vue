<script setup lang="ts">
import {ref, onMounted, computed} from "vue"
import {useRouter} from "vue-router";
import {type MenuOption, useMessage} from "naive-ui"
import {useI18n} from "vue-i18n";
import useAppInfosStore from "@/stores/useAppInfosStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";
import renderIcon from "@/utils/iconFormator";
import {
  CashOutline as cashIcon,
  LogoWechat,
  LogoAlipay,
  ChevronForwardOutline as toRight,
  LogoApple,
  Pencil as penIcon,
  Add as addIcon,
} from "@vicons/ionicons5"

const {t} = useI18n();
const message = useMessage();
const router = useRouter()
const appInfoStore = useAppInfosStore();
const userInfoStore = useUserInfoStore();
const themeStore = useThemeStore();
let animated = ref<boolean>(false)

let topUpAmount = ref<number | null>(0)
let rightSideCardBgColor = computed(() => themeStore.enableDarkMode ? 'rgba(54,55,55,0.8)' : 'rgba(54,56,61,0.8)')


interface QuickBalance {
  key: number
  value: number
}

let paymentMethodsList = ref<MenuOption[]>([
  {
    label: '微信支付',
    key: 'wechat',
    icon: renderIcon(LogoWechat),
  },
  {
    label: '支付宝支付',
    key: 'alipay',
    icon: renderIcon(LogoAlipay),
  },
  {
    label: 'Apple Pay',
    key: 'apple',
    icon: renderIcon(LogoApple),
  },
])

let quickSelectBalanceList = ref<MenuOption[]>([
  {
    label: '100' + ' ' + appInfoStore.appCommonConfig.currency,
    key: 100,
    icon: renderIcon(cashIcon)
  },
  {
    label: '200' + ' ' + appInfoStore.appCommonConfig.currency,
    key: 200,
    icon: renderIcon(cashIcon)
  },
  {
    label: '500' + ' ' + appInfoStore.appCommonConfig.currency,
    key: 500,
    icon: renderIcon(cashIcon)
  },
  {
    label: '1000' + ' ' + appInfoStore.appCommonConfig.currency,
    key: 1000,
    icon: renderIcon(cashIcon)
  },
  {
    label: '2000' + ' ' + appInfoStore.appCommonConfig.currency,
    key: 2000,
    icon: renderIcon(cashIcon)
  },
  {
    label: '其他金额',
    key: -1,
    icon: renderIcon(addIcon)
  },
])

let showCustomTopUpInput = ref<boolean>(false)

let handleSelectTopUpAmount = (amount: number) => {
  if (amount !== -1) {
    topUpAmount.value = amount
    showCustomTopUpInput.value = false
  } else {
    topUpAmount.value = null
    showCustomTopUpInput.value = true
  }
}



onMounted(() => {
  themeStore.userPath = '/dashboard/topup'
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: "UserTopUp",
}
</script>

<template>
  <div style="padding: 20px">
    <n-card
        hoverable
        :embedded="true"
        :bordered="false"
        title="加值"
    >
    </n-card>
  </div>

  <transition name="slide-fade">
    <div class="main" v-if="animated">

      <n-grid
          cols="3"
          responsive="screen"
          item-responsive
          :x-gap="15"
          :y-gap="15"
      >
        <!-- 第一个盒子 -->
        <n-grid-item
            span="3 s:3 m:2"
        >
          <div class="box1">
            <n-card
                hoverable
                :embedded="true"
                :bordered="false"
                style="transition: ease 200ms"
            >
              <n-h3 style="font-weight: bold">选择充值金额</n-h3>
              <n-h4 style="font-weight: bold">快速选择</n-h4>

              <n-menu
                  :indent="20"
                  mode="vertical"
                  :options="quickSelectBalanceList"
                  @update:value="handleSelectTopUpAmount"
              />

              <n-collapse-transition :show="showCustomTopUpInput">
                <div v-if="showCustomTopUpInput" style="margin-top: 30px">
                  <n-h4 style="font-weight: bold">自定义金额</n-h4>
                  <div>
                    <n-input-number
                        autofocus
                        class="amount-input"
                        size="large"
                        v-model:value.number="topUpAmount"
                        :placeholder="'输入要充值的金额'"
                        :max="99999999"
                    ></n-input-number>
                  </div>
                </div>
              </n-collapse-transition>

              <div style="width: 100%; display: flex; flex-direction: row; justify-content: flex-end; margin-top: 40px">
                <div>
                  充值遇到问题？
                  <n-button
                      type="primary"
                      text
                      @click="router.push('/dashboard/tickets')"
                  >
                    联系客服
                    <n-icon style="margin-left: 4px">
                      <toRight/>
                    </n-icon>
                  </n-button>
                </div>

              </div>
            </n-card>
          </div>
        </n-grid-item>
        <!-- 第二个盒子 -->
        <n-grid-item
            span="3 s:3 m:1">
          <div class="payment-methods-box">
            <n-card
                hoverable
                :embedded="true"
                :bordered="false"
            >
              <n-h3 style="font-weight: bold">支付方式</n-h3>
              <n-menu
                  :indent="20"
                  mode="vertical"
                  :options="paymentMethodsList"
                  :default-value="'wechat'"
                  @update:expanded-keys=""
              />
            </n-card>
          </div>
          <div class="box2">
            <n-card
                hoverable
                :embedded="true"
                :bordered="false"
                class="r-part-card"
            >
              <n-h3 class="r-part-title">
                您的金额
              </n-h3>

              <div class="new-amount-main">
                <p class="new-amount-main-hex">充值</p>
                <div class="new-amount-main-price-part">
                  <p class="new-amount-main-price"> + {{topUpAmount?topUpAmount.toFixed(2):0.00.toFixed(2) }}</p>
                  <p class="new-amount-main-price-currency">{{ appInfoStore.appCommonConfig.currency }}</p>
                </div>
              </div>


              <div class="new-amount-main">
                <p class="new-amount-main-hex">账户余额</p>
                <div class="new-amount-main-price-part" style="opacity: 0.8">
                  <p class="new-amount-main-price" style="font-size: 1rem !important;">{{ userInfoStore.thisUser.balance.toFixed(2) }}</p>
                  <p class="new-amount-main-price-currency">{{ appInfoStore.appCommonConfig.currency }}</p>
                </div>
              </div>

              <hr style="opacity: 0.3; margin: 10px 0 10px 0;"></hr>

              <div class="new-amount-main">
                <p class="new-amount-main-hex">余额合计</p>
                <div class="new-amount-main-price-part">
                  <p class="new-amount-main-price" style="font-size: 1rem !important;">{{ (userInfoStore.thisUser.balance + (topUpAmount?topUpAmount:0.00)).toFixed(2) }}</p>
                  <p class="new-amount-main-price-currency">{{ appInfoStore.appCommonConfig.currency }}</p>
                </div>
              </div>


              <n-button
                  type="primary"
                  :bordered="false"
                  class="submit-top-up-btn"
              >
                提交
              </n-button>
            </n-card>
          </div>
        </n-grid-item>
      </n-grid>


    </div>
  </transition>

</template>

<style scoped lang="less">
.main {
  padding: 0 20px 20px 20px;
}

.left-root {

}

.money-add-item {
  width: 100%;
  display: flex;
  flex-direction: row;
}

.amount-input {
  width: 70%;
}

@media (max-width: 1024px) {
  .amount-input {
    width: 100%;
  }
}

.r-part-card {
  background-color: v-bind(rightSideCardBgColor);

  .r-part-title {
    font-weight: bold;
    color: #fff;
  }

  .new-amount-main {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: baseline; /* 让子元素在容器内底部对齐 */

    .new-amount-main-hex {
      color: #fff;
      opacity: 0.8;
      font-size: 1rem;
    }

    .new-amount-main-price-part {
      display: flex;
      align-items: baseline; /* 让内部的文字也底部对齐 */

      .new-amount-main-price {
        color: white;
        font-size: 2rem;
      }

      .new-amount-main-price-currency {
        color: white;
        font-size: 1rem;
        margin-left: 0.2rem; /* 添加一点间距，避免价格和货币符号贴得太近 */
      }
    }
  }

  .submit-top-up-btn {
    width: 100%;
    margin-top: 30px;
  }


}

.payment-methods-box {
  margin-bottom: 15px;

}

</style>