<script setup lang="ts">
import {computed, onMounted, ref} from 'vue'
import useThemeStore from "@/stores/useThemeStore";

import {LogoAlipay, LogoApple, LogoWechat,} from "@vicons/ionicons5"

let animated = ref<boolean>(false)

const themeStore = useThemeStore();
let showModal = ref(false);

// let handleAddPaymentMethods = () => {
//   console.log('handleAddPaymentMethods');
//   active.value = !active.value
// }

interface Alipay {
  app_id: string
  app_cert_public: string
  app_private_key: string
  alipay_root_cert: string
  alipay_cert_public_key: string
  discount: number
}

interface Wechat {
  mch_id: string;       // 商户ID 或者服务商模式的 sp_mchid
  serial_no: string;    // 商户证书的证书序列号
  api_v3_key: string;   // apiV3Key 商户平台获取
  private_key: string;  // 私钥内容
  discount: number;     // 优惠金额
  enabled: boolean;     // 是否启用
}

interface Apple {
  iss: string;        // issuer ID
  bid: string;        // bundle ID
  kid: string;        // private key ID
  private_key: string; // 私钥文件读取后的字符串内容
  discount: number;   // 优惠金额
  enabled: boolean;   // 是否启用
}

let alipayConfig = ref<Alipay>({
  app_id: '',
  app_cert_public: '',
  app_private_key: '',
  alipay_root_cert: '',
  alipay_cert_public_key: '',
  discount: 0.00,
})

let wechatConfig = ref<Wechat>({
  mch_id: '',
  serial_no: '',
  api_v3_key: '',
  private_key: '',
  discount: 0.0,
  enabled: false,
});

let appleConfig = ref<Apple>({
  iss: '',            // 默认值为空字符串
  bid: '',            // 默认值为空字符串
  kid: '',            // 默认值为空字符串
  private_key: '',    // 默认值为空字符串
  discount: 0.0,      // 默认值为 0.0
  enabled: false,     // 默认值为 false
});

// 假设这是从服务器获取的支付的方式的状态列表
let paymentListKv = [
  {
    system: 'alipay',
    enabled: true,
  },
  {
    system: 'wechat',
    enabled: false,
  },
  {
    system: 'apple',
    enabled: true,
  },
]

interface PaymentConfig {
  id: number
  system: 'alipay' | 'wechat' | 'apple',
  title: string,
  enabled: boolean
}

// 这是要渲染到页面上的数据
let paymentMethodList = [
  {
    id: 0,
    system: 'alipay',
    title: '支付宝',
  },
  {
    id: 1,
    system: 'wechat',
    title: '微信支付',
  },
  {
    id: 2,
    system: 'apple',
    title: 'Apple Pay',
  }
] as PaymentConfig[]

// 动态计算 enabled 属性
let computedPaymentMethods = computed(() =>
    paymentMethodList.map((method) => ({
      ...method,
      enabled: paymentListKv.find((payment) => payment.system === method.system)?.enabled ?? false,
    }))
);

let editSystemName = ref<string>('');
let showDrawer = ref<boolean>(false);

onMounted(() => {
  animated.value = true;
})

</script>

<script lang="ts">
export default {
  name: 'PaymentConfig'
}
</script>


<template>
  <div style="padding: 20px 20px 10px 20px">
    <n-card
        hoverable
        :embedded="true"
        :bordered="false"

    >
      <div style="display: flex; justify-content: space-between; align-items: center">
        <p style="font-weight: bold; font-size: 1.1rem">支付设置</p>
        <n-button type="primary" :bordered="false" class="add-btn" @click="">添加支付方式</n-button>
      </div>
    </n-card>

    <n-alert
        type="error"
        :bordered="false"
        style="margin: 20px 0 10px 0;"
    >
      请注意，务必先配置支付方式的信息再进行启用。
    </n-alert>

    <n-alert
        type="warning"
        :bordered="false"
        style="margin: 0"
    >
      目前仅支持支付宝支付，但是您也可以先配置微信和ApplePay，等待项目进一步完善后可以在更新日志中查看是否支持。
    </n-alert>

  </div>

  <transition name="slide-fade">
    <div v-if="animated" style="padding: 0 20px 20px 20px">
      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
          class="payment-item-card"
          v-for="item in computedPaymentMethods"
          :key="item.id"
          content-style="padding: 20px;"
          :style="{ opacity: item.enabled ? 1 : 0.8 }"
      >
        <div class="payment-item-card-inner">
          <div class="payment-item-card-inner-l">
            <n-icon v-if="item.system === 'alipay'" color="#0e9dec" size="30">
              <LogoAlipay/>
            </n-icon>
            <n-icon v-if="item.system === 'wechat'" color="#09b83e" size="30">
              <LogoWechat/>
            </n-icon>
            <n-icon v-if="item.system === 'apple'" :color="themeStore.enableDarkMode?'#fff':'#000'" size="31">
              <LogoApple/>
            </n-icon>

            <p class="payment-item-card-inner-l-title">{{ item.title }}</p>
          </div>
          <div class="payment-item-card-inner-r">
            <n-button
                class="payment-item-card-inner-r-btn"
                :type="item.enabled? 'warning': 'success'"
                secondary
                style="margin-right: 10px"
            >
              {{ item.enabled ? '禁用' : '启用' }}
            </n-button>
            <n-button
                class="payment-item-card-inner-r-btn"
                type="primary"
                secondary
                @click="editSystemName=item.system as string; showDrawer=true"
            >
              配置
            </n-button>
          </div>
        </div>
      </n-card>
    </div>
  </transition>


  <n-drawer v-model:show="showDrawer" :width="'40%'" :placement="'right'">
    <n-drawer-content title="支付宝配置" v-if="editSystemName==='alipay'">

      <n-alert
          type="warning"
          :show-icon="false"
          style="margin: 0 0 20px 0;"
      >
        为确保安全，不显示详细信息，重新填入信息以创建或覆盖配置。
      </n-alert>

      <div class="add-panel">
        <div class="item">
          <p class="title">应用Id</p>
          <n-input v-model:value.trim="alipayConfig.app_id" size="medium" placeholder="app_id"></n-input>
        </div>

        <div class="item">
          <p class="title">应用公钥证书内容</p>
          <n-input v-model:value.trim="alipayConfig.app_cert_public" size="medium"
                   placeholder="app_cert_public"></n-input>
        </div>

        <div class="item">
          <p class="title">应用私钥</p>
          <n-input v-model:value.trim="alipayConfig.app_private_key" size="medium"
                   placeholder="app_private_key"></n-input>
        </div>

        <div class="item">
          <p class="title">支付宝根证书</p>
          <n-input v-model:value.trim="alipayConfig.alipay_root_cert" size="medium"
                   placeholder="alipay_root_cert"></n-input>
        </div>

        <div class="item">
          <p class="title">支付宝公钥证书内容</p>
          <n-input v-model:value.trim="alipayConfig.alipay_cert_public_key" size="medium"
                   placeholder="alipay_cert_public_key"></n-input>
        </div>

        <div class="item">
          <p class="title">优惠金额</p>
          <n-input-number size="medium" placeholder="有时或许会有一点优惠"></n-input-number>
        </div>

      </div>
      <div style="display: flex; flex-direction: row; justify-content: flex-end">
        <n-button
            type="default"
            secondary
            style="width: 90px; margin-right: 15px"
            @click="showDrawer=false; Object.assign(alipayConfig, { app_id: '', app_cert_public: '',app_private_key: '',alipay_root_cert: '',alipay_cert_public_key: '',discount: 0.00,})"
        >
          取消
        </n-button>
        <n-button
            type="primary"
            secondary
            style="width: 90px"
        >
          保存
        </n-button>
      </div>

    </n-drawer-content>

    <n-drawer-content title="微信支付配置" v-if="editSystemName === 'wechat'">
      <div class="add-panel">
        <div class="item">
          <p class="title">商户 ID</p>
          <n-input v-model:value.trim="wechatConfig.mch_id" size="medium" placeholder="mch_id"></n-input>
        </div>

        <div class="item">
          <p class="title">商户证书序列号</p>
          <n-input v-model:value.trim="wechatConfig.serial_no" size="medium" placeholder="serial_no"></n-input>
        </div>

        <div class="item">
          <p class="title">API v3 Key</p>
          <n-input v-model:value.trim="wechatConfig.api_v3_key" size="medium" placeholder="api_v3_key"></n-input>
        </div>

        <div class="item">
          <p class="title">私钥</p>
          <n-input v-model:value.trim="wechatConfig.private_key" size="medium" placeholder="private_key"></n-input>
        </div>

        <div class="item">
          <p class="title">优惠金额</p>
          <n-input-number v-model:value="wechatConfig.discount" size="medium"
                          placeholder="输入优惠金额"></n-input-number>
        </div>
      </div>
    </n-drawer-content>

    <n-drawer-content title="Apple Pay 配置" v-if="editSystemName === 'apple'">
      <div class="add-panel">
        <div class="item">
          <p class="title">Issuer ID</p>
          <n-input v-model:value.trim="appleConfig.iss" size="medium" placeholder="issuer_id"></n-input>
        </div>

        <div class="item">
          <p class="title">Bundle ID</p>
          <n-input v-model:value.trim="appleConfig.bid" size="medium" placeholder="bundle_id"></n-input>
        </div>

        <div class="item">
          <p class="title">Private Key ID</p>
          <n-input v-model:value.trim="appleConfig.kid" size="medium" placeholder="key_id"></n-input>
        </div>

        <div class="item">
          <p class="title">私钥内容</p>
          <n-input v-model:value.trim="appleConfig.private_key" size="medium" placeholder="private_key"></n-input>
        </div>

        <div class="item">
          <p class="title">优惠金额</p>
          <n-input-number v-model:value="appleConfig.discount" size="medium"
                          placeholder="输入优惠金额"></n-input-number>
        </div>
      </div>
    </n-drawer-content>
  </n-drawer>


</template>

<style lang="less" scoped>
//.root-root {
//  padding: 20px;
//}
//
//.root {
//  min-width: 900px;
//
//  .add-btn {
//    margin-top: 10px;
//  }
//
//  .table {
//    margin-top: 30px;
//  }
//}

.add-panel {
  .item {
    height: 80px;
    margin-bottom: 5px;
    text-align: left;

    .title {
      font-weight: 500;
      opacity: 0.9;
      font-size: 15px;
      margin-bottom: 5px;
    }
  }
}

.payment-item-card {
  margin-top: 15px;
  transition: ease 200ms;

  .payment-item-card-inner {
    display: flex;
    flex-direction: row;

    .payment-item-card-inner-l {
      flex: 1;
      display: flex;
      flex-direction: row;
      align-items: center;
      margin-left: 10px;

      .payment-item-card-inner-l-title {
        font-size: 1rem;
        margin-left: 15px;
      }
    }

    .payment-item-card-inner-r {
      flex: 1;
      display: flex;
      justify-content: flex-end;
      margin-right: 30px;

      .payment-item-card-inner-r-btn {
        width: 80px;
      }
    }
  }

  @media (max-width: 400px) {
    .payment-item-card-inner {
      flex-direction: column;

      .payment-item-card-inner-r {
        justify-content: flex-start;
        margin-top: 10px;
      }
    }
  }

}

.payment-item-card:hover {
  transform: translateX(0px) translateY(-3px);
}


</style>