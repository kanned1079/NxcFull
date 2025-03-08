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
  if (data && updateResponse.code === 200)
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
    <n-grid cols="1" responsive="screen" y-gap="16">
      <n-grid-item v-for="setting in paymentSettings" :key="setting.model" class="grid-item">
        <n-grid cols="1 s:1 m:2 l:2" responsive="screen" align-items="center">
          <!-- 左侧：标题 + 描述 -->
          <n-grid-item>
            <div class="describe">
              <p class="title">{{ t(setting.title) }}</p>
              <p class="shallow">{{ t(setting.shallow) }}</p>
            </div>
          </n-grid-item>

          <!-- 右侧：输入框 / 开关 -->
          <n-grid-item>
            <div v-if="setting.type === 'switch'" class="to-right">
              <n-switch
                  size="medium"
                  v-model:value="settingStore.settings.invite[setting.model]"
                  @update:value="saveFiled(setting.model, settingStore.settings.invite[setting.model])"
              />
            </div>
            <n-input-number
                v-else-if="setting.type === 'input-number'"
                size="large"
                :placeholder="t(setting.placeholder || '')"
                v-model:value="settingStore.settings.invite[setting.model]"
                @blur="saveFiled(setting.model, settingStore.settings.invite[setting.model])"
            />
            <n-input
                v-else-if="setting.type === 'input'"
                size="large"
                :placeholder="t(setting.placeholder || '')"
                v-model:value="settingStore.settings.invite[setting.model]"
                @blur="saveFiled(setting.model, settingStore.settings.invite[setting.model])"
            />
          </n-grid-item>
        </n-grid>
      </n-grid-item>
    </n-grid>
  </n-card>
</template>

<style lang="less" scoped>
.root {
  margin: 0 auto;
}

.grid-item {
  margin-bottom: 16px;
}

.describe {
  .title {
    font-weight: bold;
  }

  .shallow {
    margin-top: 5px;
    opacity: 0.5;
  }
}

.to-right {
  text-align: right;
}
</style>