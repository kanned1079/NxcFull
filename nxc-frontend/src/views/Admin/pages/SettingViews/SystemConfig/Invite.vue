<script setup lang="ts">
import useSettingStore from "@/stores/useSettingStore";
import {useI18n} from "vue-i18n"
import {useMessage} from "naive-ui"

const {t} = useI18n()
const message = useMessage()
const settingStore = useSettingStore()

// 定义 InviteModel 类型
type InviteModel =
    | "invite_rebate_enable"
    | "invite_rebate_rate"
    | "discount_info"
    | "invite_info";

interface PaymentSetting {
  title: string; // 配置项标题
  shallow: string; // 配置项描述
  model: InviteModel; // 配置项对应的字段
  type: string; // 表单组件类型
  placeholder?: string; // 输入框占位符，可选
}

const paymentSettings: PaymentSetting[] = [
  {
    title: 'adminViews.systemConfig.inviteAndRebate.inviteRebateEnable.title',
    shallow: 'adminViews.systemConfig.inviteAndRebate.inviteRebateEnable.description',
    model: 'invite_rebate_enable',
    type: 'switch',
  },
  {
    title: 'adminViews.systemConfig.inviteAndRebate.inviteRebateRate.title',
    shallow: 'adminViews.systemConfig.inviteAndRebate.inviteRebateRate.description',
    model: 'invite_rebate_rate',
    type: 'input-number',
    placeholder: 'adminViews.systemConfig.inviteAndRebate.inviteRebateRate.placeholder',
  },
  {
    title: 'adminViews.systemConfig.inviteAndRebate.discountInfo.title',
    shallow: 'adminViews.systemConfig.inviteAndRebate.discountInfo.description',
    model: 'discount_info',
    type: 'input',
    placeholder: 'adminViews.systemConfig.inviteAndRebate.discountInfo.placeholder',
  },
  {
    title: 'adminViews.systemConfig.inviteAndRebate.inviteInfo.title',
    shallow: 'adminViews.systemConfig.inviteAndRebate.inviteInfo.description',
    model: 'invite_info',
    type: 'input',
    placeholder: 'adminViews.systemConfig.inviteAndRebate.inviteInfo.placeholder',
  },
];

const saveFiled = async (k: string, v: any) => {
  let updateResponse = await settingStore.saveOption('invite', k, v)
  if (updateResponse.code === 200)
    message.success(t('adminViews.systemConfig.common.success'))
  else
    message.error(t('adminViews.systemConfig.common.err') + updateResponse.msg || '')
}

</script>

<script lang="ts">
export default {
  name: 'Subscribe'
}
</script>

<template>

  <n-card
      class="root"
      :embedded="true"
      :title="t('adminViews.systemConfig.inviteAndRebate.common.title')"
      :bordered="false"
  >
    <!--    <n-alert title="Warning 警告" type="warning" style="margin-bottom: 30px" :bordered="false">-->
    <!--      如果你更改了本页配置，需要对队列服务进行重启。&nbsp;另外本页配置优先级高于.env中邮件配置。-->
    <!--    </n-alert>-->

    <div
        v-for="setting in paymentSettings"
        :key="setting.model"
        class="item"
    >
  <span class="l-content">
    <div class="describe">
      <p class="title">{{ t(setting.title) }}</p>
      <p class="shallow">{{ t(setting.shallow) }}</p>
    </div>
  </span>

      <span class="r-content to-right" v-if="setting.type === 'switch'">
    <n-switch
        size="medium"
        v-model:value="settingStore.settings.invite[setting.model]"
        @update:value="saveFiled(setting.model, settingStore.settings.invite[setting.model])"
    />
  </span>

      <span class="r-content" v-if="setting.type === 'input-number'">
    <n-input-number
        size="large"
        :placeholder="t(setting.placeholder || '')"
        v-model:value="settingStore.settings.invite[setting.model]"
        @blur="saveFiled(setting.model, settingStore.settings.invite[setting.model])"
    />
  </span>

      <span class="r-content" v-if="setting.type === 'input'">
    <n-input
        size="large"
        :placeholder="t(setting.placeholder || '')"
        v-model:value="settingStore.settings.invite[setting.model]"
        @blur="saveFiled(setting.model, settingStore.settings.invite[setting.model])"
    />
  </span>
    </div>
  </n-card>
</template>

<style lang="less" scoped>
.root {
  .warm {
    background-color: #ffefd1;
    height: 50px;
    border: 0;
    line-height: 50px;
    text-align: left;

    p {
      color: #956d34;
      margin-left: 20px;
      font-size: 1rem;
    }
  }
}

.item {
  height: 50px;
  display: flex;
  margin-bottom: 30px;

  .l-content {
    flex: 1;

    .describe {
      .title {
        font-weight: bold;
      }

      .shallow {
        margin-top: 5px;
        opacity: 0.5;
      }
    }
  }

  .r-content {
    margin-left: 30px;
    flex: 0.8;
    justify-content: center;
    line-height: 50px;
  }
}

.to-right {
  text-align: right;
}

//.n-card {
//  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
//  border: 0;
//}
</style>