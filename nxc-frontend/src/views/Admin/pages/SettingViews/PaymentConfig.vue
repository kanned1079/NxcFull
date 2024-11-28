<script setup lang="ts">
import {computed, onMounted, ref} from 'vue'
import useThemeStore from "@/stores/useThemeStore";
import {useMessage} from "naive-ui";
import {LogoAlipay, LogoApple, LogoWechat,CheckmarkCircle} from "@vicons/ionicons5"
import instance from "@/axios";

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
      enabled: paymentListKv.value.find((payment) => payment.system === method.system)?.enabled ?? false,
    }))
);

let editSystemName = ref<string>('');
let showDrawer = ref<boolean>(false);

let addPaymentMethodClick = () => {
  message.info('Developing...')
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
    let { data } = await instance.get('/api/admin/v1/payment');
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
      message.error(data.msg + '');
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
      message.success('Saved successfully')
    } else {
      message.error(data.msg + '');
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
      message.success('Alter successfully')

    } else {
      message.error(data.msg + '');
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
      message.success('Details get successfully')
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
  <div style="padding: 20px 20px 10px 20px">
    <n-card
        hoverable
        :embedded="true"
        :bordered="false"

    >
      <div style="display: flex; justify-content: space-between; align-items: center">
        <p style="font-weight: bold; font-size: 1.1rem">支付设置</p>
        <n-button
            type="primary"
            :bordered="false"
            class="add-btn"
            @click="addPaymentMethodClick"
        >
          添加支付方式
        </n-button>
      </div>
    </n-card>

    <!--    <n-alert-->
    <!--        type="error"-->
    <!--        :bordered="false"-->
    <!--        style="margin: 20px 0 10px 0;"-->
    <!--    >-->
    <!--      请注意，务必先配置支付方式的信息再进行启用。-->
    <!--    </n-alert>-->

    <n-alert
        type="warning"
        :bordered="true"
        style="margin: 15px 0 0 0"
        :title="'注意事项'"
    >

      <n-ul>
        <n-li style="font-weight: bold">
          务必先配置支付方式的信息再进行启用，这真的很重要。
        </n-li>
        <n-li>
          修改付款方式配置时，如果显示为"---"则代表该选项已经被设置且非空，如果需要重新修改直接输入新值保存即可。
        </n-li>
<!--        <n-li>-->
<!--          修改付款方式配置后，该付款方式会默认被禁用，测试无误后您需要重新启用。-->
<!--        </n-li>-->
        <n-li>
          由于配置信息有三级缓存，如遇到失败时，可以在重启Redis*和order微服务*后再试。
        </n-li>
        <n-li>
          目前仅支持支付宝支付，但是您也可以先配置微信和ApplePay，等待项目进一步完善后可以在更新日志中查看是否支持。
        </n-li>
<!--        <n-li>-->
<!--          如出现收款未到账的问题，首先排除开发者问题，项目中不存在任何开发者个人的收款信息。-->
<!--        </n-li>-->

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
            <n-tag
              type="success"
              size="small"
              :bordered="false"
              style="margin-left: 10px"
              v-if="item.enabled"
            >
              Using
              <template #icon>
                <n-icon :component="CheckmarkCircle" />
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
              {{ item.enabled ? '禁用' : '启用' }}
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
          <n-input-number v-model:value.number="alipayConfig.discount" size="medium"
                          placeholder="有时或许会有一点优惠"></n-input-number>
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
            @click="handleSavePaymentMethodBySystemName"
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
          <n-input-number v-model:value.number="wechatConfig.discount" size="medium"
                          placeholder="输入优惠金额"></n-input-number>
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
            @click="handleSavePaymentMethodBySystemName"
        >
          保存
        </n-button>
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
            @click="handleSavePaymentMethodBySystemName"
        >
          保存
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


</style>