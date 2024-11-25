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
import instance from '@/axios/index'
import {formatDate, formatTimestamp} from "@/utils/timeFormat";


const {t} = useI18n();
const message = useMessage();
const router = useRouter()
const appInfoStore = useAppInfosStore();
const userInfoStore = useUserInfoStore();
const themeStore = useThemeStore();
let animated = ref<boolean>(false)

let topUpAmount = ref<number | null>(0)
let rightSideCardBgColor = computed(() => themeStore.enableDarkMode ? 'rgba(54,55,55,0.8)' : 'rgba(54,56,61,0.8)')

interface PaymentMethodsKv {
  id: number
  name: string
}

let paymentMethodsKv = ref<PaymentMethodsKv[]>([
  {
    id: 1,
    name: "wechat",
  },
  {
    id: 2,
    name: "alipay",
  },
  {
    id: 3,
    name: "apple",
  }
])

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

// let pay = async () => {
//   try {
//     let {data} = await instance.post('https://openapi.alipay.com/gateway.do', {
//       app_auth_token: '',
//       biz_content: JSON.stringify({
//         out_trade_no: '20241125010101001',
//         total_amount: '0.02',
//         subject: 'Iphone6 16G',
//         product_code: 'QR_CODE_OFFLINE',  // 订单码支付
//         seller_id: '2088722846526358',
//         body: 'Iphone6 16G',
//         goods_detail: [
//           {
//             goods_name: 'ipad',
//             quantity: 1,
//             price: '2000',
//             goods_id: 'apple-01',
//             goods_category: '34543238',
//             categories_tree: '124868003|126232002|126252004',
//             show_url: 'http://www.alipay.com/xxx.jpg',
//           },
//         ],
//         extend_params: {
//           sys_service_provider_id: '2088511833207846',
//         },
//         business_params: {
//           mc_create_trade_ip: '127.0.0.1',
//         },
//         discountable_amount: '0.01',
//         store_id: 'NJ_001',
//         operator_id: 'yx_001',
//         terminal_id: 'NJ_T_001',
//         merchant_order_no: '20161008001',
//       }),
//     }, {
//       headers: {
//         'Content-Type': 'application/x-www-form-urlencoded',
//       },
//       params: {
//         charset: 'UTF-8',
//         method: 'alipay.trade.precreate',
//         format: 'json',
//         sign: 'sign', // 使用签名工具生成的签名
//         app_id: '2021004199650824',
//         version: '1.0',
//         sign_type: 'RSA2',
//         timestamp: formatTimestamp(Date.now()), // 当前时间戳，格式：yyyy-MM-dd HH:mm:ss
//       }
//     })
//   } catch (err: any) {
//     message.error(err + '')
//   }
// }

const Params = {
  charset: 'UTF-8',
  method: 'alipay.trade.precreate',
  format: 'json',
  app_id: '2021004199650824',
  version: '1.0',
  sign_type: 'RSA2',
  timestamp: '2024-11-25 22:24:58',
}

const bizContent = {
  out_trade_no: '20241125010101001',
  total_amount: '0.02',
  subject: 'Iphone6 16G',
  product_code: 'QR_CODE_OFFLINE',  // 订单码支付
  seller_id: '2088722846526358',
  body: 'Iphone6 16G',
  goods_detail: [
    {
      goods_name: 'ipad',
      quantity: 1,
      price: '2000',
      goods_id: 'apple-01',
      goods_category: '34543238',
      categories_tree: '124868003|126232002|126252004',
      show_url: 'http://www.alipay.com/xxx.jpg',
    },
  ],
  extend_params: {
    sys_service_provider_id: '2088511833207846',
  },
  business_params: {
    mc_create_trade_ip: '127.0.0.1',
  },
  discountable_amount: '0.01',
  store_id: 'NJ_001',
  operator_id: 'yx_001',
  terminal_id: 'NJ_T_001',
  merchant_order_no: '20161008001',
}

const queryString = new URLSearchParams(Params).toString();
let result = queryString+ `biz_content=${JSON.stringify(bizContent)}`;
result = result.replace(/\+/g, ' ');

console.log(result);

let pay = async () => {
  try {
    let {data} = await instance.post('https://openapi.alipay.com/gateway.do', {
      biz_content: bizContent,
    }, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      params: {
        ...Params,
        sign: 'RjFYJcgAl+Lt/VewpH/v1G/1lDv1kHTWTkQ9wmny6cfV+FEl/x5XjWxi1ddhVGHJjyX7dr9E5+Rky3oTwTkwcz+PrxFfysjaNjdeXNFODKIWU5Z68upqkx9tCik7Zwu+YgLj9Bw2WnXtNMWbHK08U+j/VBuYvgwixmJvqGxcb6UwCl6KH+DEXiww/z+1n2ZIhQnIbi6P+m3vGdp5kjbCufCYqflGKNI9ZCNwMGQ5s9qt7DE6Uy9VXXzPACIqEVJehmop78K5dSP8NMDTHoqN0nuqJR8lnHgZ55xfhdF/jgKBPfLEj7kHDVw2tx3/W8kVhR721axSzl8ioBK+LZey1A=='
      },

    })
  } catch (err: any) {
    message.error(err + '')
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

                  <div style="display: flex; flex-direction: row; align-items: baseline">
                    <n-h4 style="font-weight: bold">自定义金额</n-h4>
                    <p style="opacity: 0.5; margin-left: 10px; font-size: 0.7rem">最大金额：10000000 {{ appInfoStore.appCommonConfig.currency }}</p>
                  </div>
                  <div>
                    <n-input-number
                        autofocus
                        class="amount-input"
                        size="large"
                        v-model:value.number="topUpAmount"
                        :placeholder="'输入要充值的金额'"
                        :max="10000000"
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
                  @click="pay"
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