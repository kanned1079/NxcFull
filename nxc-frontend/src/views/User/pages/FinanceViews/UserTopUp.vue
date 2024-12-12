<script setup lang="ts">
import ColoredRibbon from "@/views/utils/ColoredRibbon.vue"
import SnowFall from "@/views/utils/SnowFall.vue";
import {computed, onMounted, onBeforeUnmount, ref} from "vue"
import {useRouter} from "vue-router";
import {type MenuOption, useMessage} from "naive-ui"
import {useI18n} from "vue-i18n";
import useAppInfosStore from "@/stores/useAppInfosStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";
import renderIcon from "@/utils/iconFormator";
import {
  Add as addIcon,
  CashOutline as cashIcon,
  CheckmarkOutline as checkIcon,
  ChevronForwardOutline as toRight,
  ChevronForwardOutline as toRightIcon,
  CloseOutline as errIcon,
  GiftOutline as giftIcon,
  LogoAlipay,
  LogoApple,
  LogoWechat,
  CheckmarkCircle as scannedIcon,
} from "@vicons/ionicons5"
import instance from '@/axios/index'


const {t} = useI18n();
const message = useMessage();
const router = useRouter()
const appInfoStore = useAppInfosStore();
const userInfoStore = useUserInfoStore();
const themeStore = useThemeStore();
let animated = ref<boolean>(false)

let discountMsg = ref<string>('')
let topUpAmount = ref<number | null>()
let rightSideCardBgColor = computed(() => themeStore.enableDarkMode ? 'rgba(54,55,55,0.8)' : 'rgba(54,56,61,0.8)')

let paymentMethodSelected = ref<string>('')

let methodsAvailable = ref<MethodsFromServer[]>([])

let paymentMethodsList = ref<MenuOption[]>([
  {
    label: "微信支付",
    key: "wechat",
    icon: renderIcon(LogoWechat),
  },
  {
    label: "支付宝支付",
    key: "alipay",
    icon: renderIcon(LogoAlipay),
  },
  {
    label: "Apple Pay",
    key: "apple",
    icon: renderIcon(LogoApple),
  },
]);

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

interface MethodsFromServer {
  system: string
  discount: number
  enabled: boolean
}

const handleGetAllPaymentMethods = async () => {
  try {
    const {data} = await instance.get('/api/user/v1/payment/methods');
    if (data.code === 200) {
      // 遍历服务器返回的数据并更新 `paymentMethodsList`
      discountMsg.value = data.discount_msg || ''
      data.conf.forEach((item: MethodsFromServer) => methodsAvailable.value.push(item))
    }
  } catch (err) {
    console.error(err);
  }
};

let getDiscount = computed(() => {
  const selectedMethod = methodsAvailable.value.find(
      (method: MethodsFromServer) => method.system === paymentMethodSelected.value
  );
  return selectedMethod && selectedMethod.enabled && selectedMethod.discount !== 0 && topUpAmount.value && topUpAmount.value >= selectedMethod.discount
      ? selectedMethod.discount
      : 0;
});

let checkMethodEnabled = computed(() => {
  return methodsAvailable.value.some(
      (method: MethodsFromServer) => method.system === paymentMethodSelected.value && method.enabled && resultAmount.value !== 0.00)
})

let resultAmount = computed(() => topUpAmount.value ? (topUpAmount.value - getDiscount.value) : 0.00)

interface TopUpOrderResponse {
  code: number
  created: boolean
  msg: string
  order_id: string
  qr_code: string
}

let topUpOrderResponse = ref<TopUpOrderResponse>({
  code: 0,
  created: false,
  msg: '',
  order_id: '20190101682468247',
  qr_code: 'NO_CONTENT_PLEASE_RETRY'
})

let handleCommitNewTopUpOrder = async () => {
  try {
    let {data} = await instance.post('/api/user/v1/payment/top-up', {
      user_id: userInfoStore.thisUser.id, // 用户id
      payment_method: paymentMethodSelected.value,  // 支付方式
      amount: topUpAmount.value,  // 选择的充值金额
      discount: getDiscount.value,  // 优惠金额
      committed_at: Math.floor(Date.now()), // 当前的毫秒时间戳
      confirmed: true,
    });
    if (data.code === 200) {
      Object.assign(topUpOrderResponse.value, data)
      showQrCodeModal.value = true
      setTimeout(() => {

        startQueryTopUpOrderStatusLoop()
      }, 5000)
    }
  } catch (err: any) {
    console.log(err)
  }
}

let showQrCodeModal = ref<boolean>(false)

let clickToPaymentApp = () => {
  window.open(topUpOrderResponse.value.qr_code, '_blank');
}

/*
const (
	// WaitUserScanQrCode 预生成订单成功后的等待用户扫码
	WaitUserScanQrCode = http.StatusAccepted // 202 - 请求已接受，但尚未处理，等待用户操作

	// ScannedQrCodeWaitUserPay 用户扫描二维码后生成订单等待支付
	ScannedQrCodeWaitUserPay = http.StatusProcessing // 102 - 请求已接收，正在处理中，但尚未完成

	// TradeClosed 未付款交易超时关闭，或支付完成后全额退款
	TradeClosed = http.StatusRequestTimeout // 408 - 请求超时，通常表示支付超时未完成

	// TradeSuccess 交易成功
	TradeSuccess = http.StatusOK // 200 - 请求成功并且已处理完毕

	// TradeFinished TRADE_FINISHED，交易完成，表示交易结束
	TradeFinished = http.StatusGone // 410 - 资源不再可用，交易已经结束
)
*/
let resultCode = ref<number>(0)
let qrCodeIsScanned = computed(() => resultCode.value === 102)
// let qrCodeIsScanned = computed(() => true)

let checkPaymentResult = async () => {
  // message.info('查询支付结果')
  try {
    let {data} = await instance.get('/api/user/v1/payment/top-up/check', {
      params: {
        payment_method: paymentMethodSelected.value,
        order_id: topUpOrderResponse.value.order_id,
        user_id: userInfoStore.thisUser.id,
        invite_user_id: userInfoStore.thisUser.inviteUserId,
      }
    })
    resultCode.value = data.code || 0

    console.log(data)
  } catch (err: any) {
    console.log(err)
  }
}

let queryLoopIntervalId = ref<number | undefined>(undefined)
let topUpIsFinishedSuccess = ref<boolean>(false)

let startQueryTopUpOrderStatusLoop = () => {
  // 异步请求函数
  queryLoopIntervalId.value = setInterval(async () => {
    await checkPaymentResult()
    switch (resultCode.value) {
      case 202: {
        // message.info('等待用户付款')
        break
      }
      case 102: {
        // message.info('扫描二维码成功')
        break
      }
      case 408: {
        // message.error('未付款交易超时关闭')
        clearInterval(queryLoopIntervalId.value? queryLoopIntervalId.value: undefined)
        break
      }
      case 200: {
        // showQrCodeModal.value = false
        topUpIsFinishedSuccess.value = true
        message.success('交易成功')
        clearInterval(queryLoopIntervalId.value? queryLoopIntervalId.value: undefined)

        setTimeout(async () => {
         let updated = await userInfoStore.updateUserInfo();
         if (updated) {
           console.log("用户信息更新成功")
           setTimeout(async () => {
             await router.push({
               path: "/dashboard/profile",
             })
           }, 3000)
         } else {
           message.warning('用户信息可能更新有延迟')
         }
        }, 1000)
        break
      }
      case 401: {
        message.info('交易完成，不可退款')
        clearInterval(queryLoopIntervalId.value? queryLoopIntervalId.value: undefined)
        break
      }
      case 500: {
        message.error('更新用户余额失败，请联系客服')
        clearInterval(queryLoopIntervalId.value? queryLoopIntervalId.value: undefined)
        break
      }

    }
  }, 2000)

}

let showRibbon = ref<boolean>(false)

onMounted(async () => {
  themeStore.userPath = '/dashboard/topup'

  await handleGetAllPaymentMethods()

  animated.value = true
  setTimeout(() => {
    showRibbon.value = true
  }, 500)
})

onBeforeUnmount(() => {
  clearInterval(queryLoopIntervalId.value?queryLoopIntervalId.value:undefined)
})

</script>

<script lang="ts">
export default {
  name: "UserTopUp",
}
</script>

<template>
<!--  <SnowFall :show="showRibbon"></SnowFall>-->
<!--  <ColoredRibbon :show="showRibbon"></ColoredRibbon>-->
  <ColoredRibbon :show="topUpIsFinishedSuccess"></ColoredRibbon>
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
            <n-alert
                type="warning"
                :bordered="false"
                style="margin-bottom: 10px"
                v-if="discountMsg !== ''"
            >
              <template #icon>
                <n-icon size="18">
                  <giftIcon/>
                </n-icon>
              </template>
              {{ discountMsg }}
            </n-alert>
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
                    <p style="opacity: 0.5; margin-left: 10px; font-size: 0.7rem">最大金额：10000000
                      {{ appInfoStore.appCommonConfig.currency }}</p>
                  </div>
                  <div>
                    <n-input-number
                        autofocus
                        class="amount-input"
                        size="large"
                        v-model:value.number="topUpAmount"
                        :placeholder="'输入要充值的金额'"
                        :max="10000000"
                        :min="0.01"
                        :precision="2"
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
                  v-model:value.trim="paymentMethodSelected"
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
                  <p class="new-amount-main-price"> + {{ resultAmount.toFixed(2) }}</p>
                  <p class="new-amount-main-price-currency">{{ appInfoStore.appCommonConfig.currency }}</p>
                </div>
              </div>

              <div class="new-amount-main" v-if="getDiscount !== 0.00">
                <p class="new-amount-main-hex">优惠</p>
                <div class="new-amount-main-price-part" style="opacity: 0.8">
                  <p class="new-amount-main-price" style="font-size: 1rem !important;">
                    - {{ getDiscount.toFixed(2) }}</p>
                  <p class="new-amount-main-price-currency">{{ appInfoStore.appCommonConfig.currency }}</p>
                </div>
              </div>

              <div class="new-amount-main">
                <p class="new-amount-main-hex">账户余额</p>
                <div class="new-amount-main-price-part" style="opacity: 0.8">
                  <p class="new-amount-main-price" style="font-size: 1rem !important;">
                    {{ userInfoStore.thisUser.balance.toFixed(2) }}</p>
                  <p class="new-amount-main-price-currency">{{ appInfoStore.appCommonConfig.currency }}</p>
                </div>
              </div>

              <hr style="opacity: 0.3; margin: 10px 0 10px 0;"></hr>

              <div class="new-amount-main">
                <p class="new-amount-main-hex">余额合计</p>
                <div class="new-amount-main-price-part">
                  <p class="new-amount-main-price" style="font-size: 1rem !important;">
                    {{ (userInfoStore.thisUser.balance + (topUpAmount ? topUpAmount : 0.00)).toFixed(2) }}</p>
                  <p class="new-amount-main-price-currency">{{ appInfoStore.appCommonConfig.currency }}</p>
                </div>
              </div>


              <n-button
                  type="primary"
                  :bordered="false"
                  class="submit-top-up-btn"
                  :disabled="!checkMethodEnabled"
                  @click="handleCommitNewTopUpOrder"
              >
                <template #icon v-if="checkMethodEnabled">
                  <n-icon>
                    <checkIcon/>
                  </n-icon>
                </template>
                <template #icon v-else>
                  <n-icon>
                    <errIcon/>
                  </n-icon>
                </template>
                {{ checkMethodEnabled ? '提交' : '支付方式不可用请选择其他' }}
              </n-button>
            </n-card>
          </div>
        </n-grid-item>
      </n-grid>


    </div>
  </transition>

  <n-modal
      v-model:show="showQrCodeModal"
      class="qr-card"
      preset="card"
      :title="'支付'"
      size="medium"
      style="width: 300px;"
      :bordered="false"
  >


    <div class="qr-body">
      <div class="qr-code-body" v-if="!topUpIsFinishedSuccess">
        <n-spin
            :rotate="false"
            size="large"
            :show="qrCodeIsScanned"
        >
<!--          <n-alert title="啦啦啦" type="success">-->
<!--            明天再打开行李箱。宝贝，挂电话啦。-->
<!--          </n-alert>-->
          <n-qr-code
              :size="160"
              :value="topUpOrderResponse.qr_code"
              error-correction-level="H"
              :color="!qrCodeIsScanned?'#252525':'rgba(25,25,25,0.3)'"
              type="svg"
          />
          <template #icon>
            <n-icon color="#66c18c">
              <scannedIcon />
            </n-icon>
          </template>
          <template #description>
            <p style="color: #66c18c; font-size: 1rem; font-weight: bold">
              扫描成功
            </p>
          </template>
        </n-spin>

        <n-button
            text
            type="default"
            @click="clickToPaymentApp"
            icon-placement="right"
            style="margin-top: 5px"
        >
          <template #icon>
            <n-icon>
              <toRightIcon/>
            </n-icon>
          </template>
          或点击跳转到APP
        </n-button>
      </div>
      <div style="text-align: center" v-if="topUpIsFinishedSuccess">
        <p style="font-size: 1.5rem; font-weight: bold">
          充值成功
        </p>
        <p style="font-size: 1rem; opacity: 0.6; margin-bottom: 30px">
          感谢您的支持
        </p>
      </div>

      <div class="qr-hex">
        <p class="amount">{{ `${resultAmount.toFixed(2)} ${appInfoStore.appCommonConfig.currency}` }}</p>
        <p class="order-detail">{{ `#${topUpOrderResponse.order_id}` }}</p>
      </div>

    </div>
<!--    <div style="width: 100%; display: flex; flex-direction: column; margin-top: 20px">-->
<!--      <n-button-->
<!--          type="success"-->
<!--          secondary-->
<!--          :bordered="false"-->
<!--          @click="checkPaymentResult"-->
<!--      >-->
<!--        我已完成支付-->
<!--      </n-button>-->
<!--    </div>-->

    <template #footer v-if="!topUpIsFinishedSuccess">
      <div style="width: 100%; display: flex; flex-direction: row; justify-content: flex-end; margin-top: 10px">
        <div>
          支付值遇到问题？
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
    </template>
    <template #footer v-else>
      <div style="width: 100%; display: flex; flex-direction: row; justify-content: flex-end; margin-top: 10px">
        <div>

          <n-button
              type="primary"
              text
              @click="router.push('/dashboard/purchase')"
          >
            去购买订阅
            <n-icon style="margin-left: 4px">
              <toRight/>
            </n-icon>
          </n-button>
        </div>

      </div>
    </template>
  </n-modal>

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

.qr-card {
  width: 300px;
  height: 500px;
}

.qr-body {
  width: 100%;

  .qr-code-body {
    margin-top: 5px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .qr-hex {
    margin-top: 10px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    .amount {
      font-size: 1.35rem;
      font-weight: bold;
    }

    .order-detail {
      font-size: 1rem;
      opacity: 0.7;
    }
  }
}

</style>