<script setup lang="ts">
import useSettingStore from "@/stores/useSettingStore";
import {useI18n} from "vue-i18n"
import {useMessage} from "naive-ui"

const {t} = useI18n()
const message = useMessage()
const settingStore = useSettingStore();
type SecurityModel =
    | "email_verify"
    | "email_gmail_limit_enable"
    // | "safe_mode_enable"
    | "secure_path"
    // | "email_whitelist_enable"
    | "recaptcha_enable"
    | "recaptcha_site_key"
    | "ip_register_limit_enable"
    | "ip_register_limit_times"
    | "ip_register_lock_time";

interface SecuritySetting {
  title: string;
  // shallow: string;
  description: string
  modelValue: SecurityModel;
  type: string;
  placeholder?: string;
}

let securitySettingsData: SecuritySetting[] = [
  {
    title: 'adminViews.systemConfig.security.emailVerify.title',
    description: 'adminViews.systemConfig.security.emailVerify.description',
    type: 'switch',
    modelValue: 'email_verify',
  },
  {
    title: 'adminViews.systemConfig.security.gmailAlias.title',
    description: 'adminViews.systemConfig.security.gmailAlias.description',
    type: 'switch',
    modelValue: 'email_gmail_limit_enable',
  },
  // {
  //   title: 'adminViews.systemConfig.security.safeMode.title',
  //   description: 'adminViews.systemConfig.security.safeMode.description',
  //   type: 'switch',
  //   modelValue: 'safe_mode_enable',
  // },
  {
    title: 'adminViews.systemConfig.security.adminPath.title',
    description: 'adminViews.systemConfig.security.adminPath.description',
    type: 'input',
    placeholder: 'adminViews.systemConfig.security.adminPath.placeholder',
    modelValue: 'secure_path',
  },
  // {
  //   title: 'adminViews.systemConfig.security.emailWhitelist.title',
  //   description: 'adminViews.systemConfig.security.emailWhitelist.description',
  //   type: 'switch',
  //   modelValue: 'email_whitelist_enable',
  // },
  {
    title: 'adminViews.systemConfig.security.recaptcha.title',
    description: 'adminViews.systemConfig.security.recaptcha.description',
    type: 'switch',
    modelValue: 'recaptcha_enable',
  },
  {
    title: 'adminViews.systemConfig.security.hCaptchaSiteKey.title',
    description: 'adminViews.systemConfig.security.hCaptchaSiteKey.description',
    type: 'input',
    placeholder: 'adminViews.systemConfig.security.hCaptchaSiteKey.placeholder',
    modelValue: 'recaptcha_site_key',
  },
  {
    title: 'adminViews.systemConfig.security.ipRegisterLimit.title',
    description: 'adminViews.systemConfig.security.ipRegisterLimit.description',
    type: 'switch',
    modelValue: 'ip_register_limit_enable',
  },
  {
    title: 'adminViews.systemConfig.security.registerTimes.title',
    description: 'adminViews.systemConfig.security.registerTimes.description',
    type: 'input-number',
    placeholder: 'adminViews.systemConfig.security.registerTimes.placeholder',
    modelValue: 'ip_register_limit_times',
  },
  {
    title: 'adminViews.systemConfig.security.lockTime.title',
    description: 'adminViews.systemConfig.security.lockTime.description',
    type: 'input-number',
    placeholder: 'adminViews.systemConfig.security.lockTime.placeholder',
    modelValue: 'ip_register_lock_time',
  },
];

const saveFiled = async (k: string, v: any) => {
  let updateResponse = await settingStore.saveOption('security', k, v)
  if (updateResponse.code === 200)
    message.success(t('adminViews.systemConfig.common.success'))
  else
    message.error(t('adminViews.systemConfig.common.err') + updateResponse.msg || '')
}

</script>

<script lang="ts">
export default {
  name: 'Security'
}
</script>

<template>
  <div class="root">
    <n-card
        :embedded="true"
        class="security-panel"
        :title="t('adminViews.systemConfig.security.common.title')"
        :bordered="false"
    >
      <n-grid cols="1" responsive="screen" y-gap="16">
        <n-grid-item v-for="(item, index) in securitySettingsData" :key="index" class="grid-item">
          <n-grid cols="1 s:1 m:2 l:2" responsive="screen" align-items="center">
            <!-- 左侧：标题 + 描述 -->
            <n-grid-item>
              <div class="describe">
                <p class="title">{{ t(item.title) }}</p>
                <p class="shallow">{{ t(item.description) }}</p>
              </div>
            </n-grid-item>

            <!-- 右侧：输入框 / 选择框 / 开关 -->
            <n-grid-item>
              <n-input
                  v-if="item.type === 'input'"
                  size="large"
                  :placeholder="t(item.placeholder || '')"
                  v-model:value="settingStore.settings.security[item.modelValue]"
                  @blur="saveFiled(item.modelValue, settingStore.settings.security[item.modelValue])"
              />
              <n-input-number
                  v-else-if="item.type === 'input-number'"
                  size="large"
                  :placeholder="t(item.placeholder || '')"
                  v-model:value.number="settingStore.settings.security[item.modelValue]"
                  @blur="saveFiled(item.modelValue, settingStore.settings.security[item.modelValue])"
              />
              <div v-else-if="item.type === 'switch'" class="to-right">
                <n-switch
                    size="medium"
                    v-model:value="settingStore.settings.security[item.modelValue]"
                    @update:value="saveFiled(item.modelValue, settingStore.settings.security[item.modelValue])"
                />
              </div>
            </n-grid-item>
          </n-grid>
        </n-grid-item>
      </n-grid>
    </n-card>
  </div>
</template>

<style lang="less" scoped>
.root {
  margin: 0 auto;

  .security-panel {
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
  }
}

.to-right {
  text-align: right;
}
</style>