<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, onMounted, onBeforeMount, ref} from 'vue'
import useThemeStore from "@/stores/useThemeStore";
import {NButton, NIcon, useMessage} from "naive-ui";
import {CheckmarkCircle, LogoAlipay, LogoApple, LogoWechat} from "@vicons/ionicons5"
import {AddOutline as AddIcon} from "@vicons/ionicons5"

import instance from "@/axios";
import PageHead from "@/views/utils/PageHead.vue";

const {t} = useI18n()
let animated = ref<boolean>(false)

const message = useMessage()
const themeStore = useThemeStore();
let showModal = ref(false);


interface Alipay {
  app_id: string
  app_cert_public: string
  app_private_key: string
  alipay_root_cert: string
  alipay_cert_public_key: string
  discount: number
  enabled: boolean
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
  enabled: false,

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
interface PaymentListKv {
  system: string;
  discount: number;
  enabled: boolean;
}

let paymentListKv = ref<PaymentListKv[]>([]);

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
    title: 'adminViews.payConfig.alipay.title',
  },
  {
    id: 1,
    system: 'wechat',
    title: 'adminViews.payConfig.wechat.title',
  },
  {
    id: 2,
    system: 'apple',
    title: 'adminViews.payConfig.apple.title',
  }
] as PaymentConfig[]

// 动态计算 enabled 属性
let computedPaymentMethods = computed(() =>
    paymentMethodList.map((method) => ({
      ...method,
      enabled: paymentListKv.value.find((payment) => payment.system === method.system)?.enabled ?? false,
    }))
);

let editSystemName = ref<string>('');
let showDrawer = ref<boolean>(false);

let addPaymentMethodClick = () => {
  message.info('Developing...')
  // TODO 添加支付方式功能
}

let conf = computed(() => {
  switch (editSystemName.value) {
    case 'alipay':
      return alipayConfig.value;
    case 'wechat':
      return wechatConfig.value;
    case 'apple':
      return appleConfig.value
  }
})

let handleGetAllPaymentMethodsBasic = async () => {
  try {
    let {data} = await instance.get('/api/admin/v1/payment');
    if (data.code === 200) {
      paymentListKv.value = []

      // 遍历服务器返回的支付方式配置并更新到本地
      data.conf.forEach((item: PaymentListKv) => {
        switch (item.system) {
          case 'alipay':
            alipayConfig.value.discount = item.discount;
            alipayConfig.value.enabled = item.enabled;
            break;
          case 'wechat':
            wechatConfig.value.discount = item.discount;
            wechatConfig.value.enabled = item.enabled;
            break;
          case 'apple':
            appleConfig.value.discount = item.discount;
            appleConfig.value.enabled = item.enabled;
            break;
        }
        paymentListKv.value.push(item);
        animated.value = true;
      });

    } else {
      message.error(t('adminViews.payConfig.common.alterSuccess') + data.msg || '');
    }
  } catch (err: any) {
    console.log(err);
  }
};

let handleSavePaymentMethodBySystemName = async () => {
  try {
    animated.value = false;
    let {data} = await instance.post('/api/admin/v1/payment', {
      ...conf.value,
    }, {
      params: {
        system: editSystemName.value,
      }
    })
    if (data.code === 200) {
      animated.value = true;
      await handleGetAllPaymentMethodsBasic()
      message.success(t('adminViews.payConfig.common.saveSuccess'))
    } else {
      message.error(t('adminViews.payConfig.common.alterSuccess') + data.msg || '');
    }
  } catch (err: any) {
    console.log(err)
  }
}

let handleChangeEnabledClick = async (system: string) => {
  try {
    animated.value = false;
    let {data} = await instance.patch('/api/admin/v1/payment', {
      system: system,
    })
    if (data.code === 200) {
      // animated.value = true;
      await handleGetAllPaymentMethodsBasic()
      message.success(t('adminViews.payConfig.common.alterSuccess'))

    } else {
      message.error(t('adminViews.payConfig.common.alterSuccess') + data.msg ||'');
    }
  } catch (err: any) {
    console.log(err)
  }
}

let handleGetPaymentMethodDetailsBySystemName = async (system: string) => {
  try {
    let {data} = await instance.get('/api/admin/v1/payment/details', {
      params: {
        system: system,
      }
    })
    if (data.code === 200) {
      switch (system) {
        case 'alipay':
          Object.assign(alipayConfig.value, data.details);
          break
        case 'wechat':
          Object.assign(wechatConfig.value, data.details);
          break
        case 'apple':
          Object.assign(appleConfig.value, data.details);
          break
      }

      animated.value = true;

    } else {
      message.error(data.msg + '');
    }
  } catch (err: any) {
    console.log(err)
  }
}

onBeforeMount(() => {
  themeStore.menuSelected = 'payment-config'
})

onMounted(async () => {

  await handleGetAllPaymentMethodsBasic()
  animated.value = true;
})

</script>

<script lang="ts">
export default {
  name: 'PaymentConfig'
}
</script>


<template>

  <div style="padding: 0 20px 10px 20px">
    <PageHead
        :title="t('adminViews.payConfig.title')"
        :description="t('adminViews.payConfig.description')"
    >
      <n-button
          tertiary
          type="primary"
          size="medium"
          class="btn-right"
          @click="addPaymentMethodClick">
        <template #icon>
          <n-icon>
            <AddIcon/>
          </n-icon>
        </template>

        {{ t('adminViews.payConfig.addPaymentMethod') }}
      </n-button>
    </PageHead>
<!--    <n-card-->
<!--        hoverable-->
<!--        :embedded="true"-->
<!--        :bordered="false"-->
<!--    >-->
<!--      <div style="display: flex; justify-content: space-between; align-items: center">-->
<!--        <p style="font-weight: bold; font-size: 1.1rem">{{ t('adminViews.payConfig.title') }}</p>-->
<!--        <n-button-->
<!--            type="primary"-->
<!--            :bordered="false"-->
<!--            class="add-btn"-->
<!--            @click="addPaymentMethodClick"-->
<!--        >-->
<!--          {{ t('adminViews.payConfig.addPaymentMethod') }}-->
<!--        </n-button>-->
<!--      </div>-->
<!--    </n-card>-->



    <n-alert
        type="warning"
        :bordered="true"
        style="margin: 15px 0 0 0"
        :title="t('adminViews.payConfig.attention.title')"
    >
      <n-ul>
        <n-li :style="i == 1?{fontWeight: 'bold'}:null" v-for="i in 2" :key="i">
          {{ t(`adminViews.payConfig.attention.point${i}`) }}
        </n-li>
      </n-ul>
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
          :style="{ opacity: item.enabled? 1 : 0.8 }"
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

            <p class="payment-item-card-inner-l-title">{{ t(item.title) }}</p>
            <n-tag
                type="success"
                size="small"
                :bordered="false"
                style="margin-left: 10px"
                v-if="item.enabled"
            >
              Using
              <template #icon>
                <n-icon :component="CheckmarkCircle"/>
              </template>
            </n-tag>
          </div>
          <div class="payment-item-card-inner-r">
            <n-button
                class="payment-item-card-inner-r-btn"
                :type="item.enabled? 'warning': 'success'"
                secondary
                style="margin-right: 10px"
                @click="handleChangeEnabledClick(item.system)"
            >
              {{ item.enabled? t('adminViews.payConfig.disableBtnHint') : t('adminViews.payConfig.enableBtnHint') }}
            </n-button>
            <n-button
                class="payment-item-card-inner-r-btn"
                type="primary"
                secondary
                @click="
                handleGetPaymentMethodDetailsBySystemName(item.system);
                editSystemName=item.system as string;
                showDrawer=true"
            >
              {{ t('adminViews.payConfig.setupPaymentMethod') }}
            </n-button>
          </div>
        </div>
      </n-card>
    </div>
  </transition>

  <n-drawer v-model:show="showDrawer" :width="'40%'" :placement="'right'">
    <n-drawer-content :title="t('adminViews.payConfig.common.detail', {method: t('adminViews.payConfig.alipay.title')})" v-if="editSystemName==='alipay'">
      <n-alert
          type="info"
          :show-icon="true"
          style="margin: 0 0 20px 0;"
      >
        {{ t('adminViews.payConfig.common.fillAttention') }}
      </n-alert>
      <div class="add-panel">
        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.alipay.config.appId') }}</p>
          <n-input v-model:value.trim="alipayConfig.app_id" size="medium" placeholder="app_id"></n-input>
        </div>

        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.alipay.config.appPublicKeyCertContent') }}</p>
          <n-input
              v-model:value.trim="alipayConfig.app_cert_public"
              size="medium"
              placeholder="app_cert_public"
              type="textarea"
              :rows="3"
          ></n-input>
        </div>

        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.alipay.config.appPrivateKey') }}</p>
          <n-input v-model:value.trim="alipayConfig.app_private_key" size="medium"
                   placeholder="app_private_key"
                   type="textarea"
                   :rows="3"></n-input>
        </div>

        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.alipay.config.alipayRootCert') }}</p>
          <n-input v-model:value.trim="alipayConfig.alipay_root_cert" size="medium"
                   placeholder="alipay_root_cert"
                   type="textarea"
                   :rows="3"></n-input>
        </div>

        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.alipay.config.alipayPublicKeyCertContent') }}</p>
          <n-input v-model:value.trim="alipayConfig.alipay_cert_public_key" size="medium"
                   placeholder="alipay_cert_public_key"
                   type="textarea"
                   :rows="3"></n-input>
        </div>

        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.common.discountAmount') }}</p>
          <n-input-number v-model:value.number="alipayConfig.discount" size="medium"
                          :placeholder="t('adminViews.payConfig.common.discountPlaceholder')"></n-input-number>
        </div>

      </div>
      <div style="display: flex; flex-direction: row; justify-content: flex-end; margin-top: 50px">
        <n-button
            type="default"
            secondary
            style="width: 90px; margin-right: 15px"
            @click="showDrawer=false; Object.assign(alipayConfig, { app_id: '', app_cert_public: '',app_private_key: '',alipay_root_cert: '',alipay_cert_public_key: '',discount: 0.00,})"
        >
          {{ t('adminViews.payConfig.common.cancelBtnHint') }}
        </n-button>
        <n-button
            type="primary"
            secondary
            style="width: 90px"
            @click="handleSavePaymentMethodBySystemName"
        >
          {{ t('adminViews.payConfig.common.saveConfigBtnHint') }}
        </n-button>
      </div>
    </n-drawer-content>

    <n-drawer-content :title="t('adminViews.payConfig.common.detail', {method: t('adminViews.payConfig.wechat.title')})" v-if="editSystemName === 'wechat'">
      <n-alert
          type="info"
          :show-icon="true"
          style="margin: 0 0 20px 0;"
      >
        {{ t('adminViews.payConfig.common.fillAttention') }}
      </n-alert>
      <div class="add-panel">
        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.wechat.config.mchId') }}</p>
          <n-input v-model:value.trim="wechatConfig.mch_id" size="medium" placeholder="mch_id"></n-input>
        </div>

        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.wechat.config.mchCertSerial') }}</p>
          <n-input v-model:value.trim="wechatConfig.serial_no" size="medium" placeholder="serial_no"></n-input>
        </div>

        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.wechat.config.apiV3Key') }}</p>
          <n-input v-model:value.trim="wechatConfig.api_v3_key" size="medium" placeholder="api_v3_key"></n-input>
        </div>

        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.wechat.config.privateKey') }}</p>
          <n-input v-model:value.trim="wechatConfig.private_key" size="medium" placeholder="private_key"></n-input>
        </div>

        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.common.discountAmount') }}</p>
          <n-input-number v-model:value.number="wechatConfig.discount" size="medium"
                          :placeholder="t('adminViews.payConfig.common.discountPlaceholder')"></n-input-number>
        </div>
      </div>
      <div style="display: flex; flex-direction: row; justify-content: flex-end; margin-top: 50px">
        <n-button
            type="default"
            secondary
            style="width: 90px; margin-right: 15px"
            @click="showDrawer=false; Object.assign(alipayConfig, { app_id: '', app_cert_public: '',app_private_key: '',alipay_root_cert: '',alipay_cert_public_key: '',discount: 0.00,})"
        >
          {{ t('adminViews.payConfig.common.cancelBtnHint') }}
        </n-button>
        <n-button
            type="primary"
            secondary
            style="width: 90px"
            @click="handleSavePaymentMethodBySystemName"
        >
          {{ t('adminViews.payConfig.common.saveConfigBtnHint') }}
        </n-button>
      </div>
    </n-drawer-content>

    <n-drawer-content :title="t('adminViews.payConfig.common.detail', {method: t('adminViews.payConfig.apple.title')})" v-if="editSystemName === 'apple'">
      <n-alert
          type="info"
          :show-icon="true"
          style="margin: 0 0 20px 0;"
      >
        {{ t('adminViews.payConfig.common.fillAttention') }}
      </n-alert>
      <div class="add-panel">
        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.apple.config.issuerId') }}</p>
          <n-input v-model:value.trim="appleConfig.iss" size="medium" placeholder="issuer_id"></n-input>
        </div>

        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.apple.config.bundleId') }}</p>
          <n-input v-model:value.trim="appleConfig.bid" size="medium" placeholder="bundle_id"></n-input>
        </div>

        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.apple.config.privateKeyId') }}</p>
          <n-input v-model:value.trim="appleConfig.kid" size="medium" placeholder="key_id"></n-input>
        </div>

        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.apple.config.privateKeyContent') }}</p>
          <n-input v-model:value.trim="appleConfig.private_key" size="medium" placeholder="private_key"></n-input>
        </div>

        <div class="item">
          <p class="title">{{ t('adminViews.payConfig.common.discountAmount') }}</p>
          <n-input-number v-model:value="appleConfig.discount" size="medium"
                          :placeholder="t('adminViews.payConfig.common.discountPlaceholder')"></n-input-number>
        </div>
      </div>
      <div style="display: flex; flex-direction: row; justify-content: flex-end; margin-top: 50px">
        <n-button
            type="default"
            secondary
            style="width: 90px; margin-right: 15px"
            @click="showDrawer=false; Object.assign(alipayConfig, { app_id: '', app_cert_public: '',app_private_key: '',alipay_root_cert: '',alipay_cert_public_key: '',discount: 0.00,})"
        >
          {{ t('adminViews.payConfig.common.cancelBtnHint') }}
        </n-button>
        <n-button
            type="primary"
            secondary
            style="width: 90px"
            @click="handleSavePaymentMethodBySystemName"
        >
          {{ t('adminViews.payConfig.common.saveConfigBtnHint') }}
        </n-button>
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

.item {
  .title {
    margin-top: 40px;
  }
}


</style>