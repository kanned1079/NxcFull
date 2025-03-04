<script setup lang="ts">
import useThemeStore from "@/stores/useThemeStore";
import useSettingStore from "@/stores/useSettingStore"
import instance from "@/axios";
import {onMounted, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useMessage} from "naive-ui";

const {t} = useI18n()
const message = useMessage()
const themeStore = useThemeStore();
const settingStore = useSettingStore();

// let handleSaveOptions = () => {
//   settingStore.saveSetting()
// }

interface Plan {
  id: number;
  group_id: number;
  name: string;
  is_sale: boolean;
  is_renew: boolean;
  capacity_limit: number;
  month_price: number;
  quarter_price: number;
  half_year_price: number;
  year_price: number;
}

type SiteModel =
    | "app_name"
    | "app_sub_name"
    | "app_description"
    | "app_url"
    // | "force_https"
    | "logo_url"
    // | "subscribe_url"
    | "tos_url"
    | "stop_register"
    | "invite_require"
    // | "trial_subscribe"
    // | "trial_time"
    | "currency"
    | "currency_symbol";

interface AppSetting {
  title: string;
  shallow: string;
  model: SiteModel;
  type: string;
  placeholder?: string;
}

// appSettings 用于v-for渲染的数据
const appSettings: AppSetting[] = [
  {
    title: 'adminViews.systemConfig.site.appName.title',
    shallow: 'adminViews.systemConfig.site.appName.shallow',
    model: 'app_name',
    type: 'input',
    placeholder: 'adminViews.systemConfig.site.appName.placeholder'
  },
  {
    title: 'adminViews.systemConfig.site.appSubName.title',
    shallow: 'adminViews.systemConfig.site.appSubName.shallow',
    model: 'app_sub_name',
    type: 'input',
    placeholder: 'adminViews.systemConfig.site.appSubName.placeholder'
  },
  {
    title: 'adminViews.systemConfig.site.appDescription.title',
    shallow: 'adminViews.systemConfig.site.appDescription.shallow',
    model: 'app_description',
    type: 'input-area',
    placeholder: 'adminViews.systemConfig.site.appDescription.placeholder'
  },
  {
    title: 'adminViews.systemConfig.site.appUrl.title',
    shallow: 'adminViews.systemConfig.site.appUrl.shallow',
    model: 'app_url',
    type: 'input',
    placeholder: 'adminViews.systemConfig.site.appUrl.placeholder'
  },
  // {
  //   title: 'adminViews.systemConfig.site.forceHttps.title',
  //   shallow: 'adminViews.systemConfig.site.forceHttps.shallow',
  //   model: 'force_https',
  //   type: 'switch'
  // },
  {
    title: 'adminViews.systemConfig.site.logoUrl.title',
    shallow: 'adminViews.systemConfig.site.logoUrl.shallow',
    model: 'logo_url',
    type: 'input',
    placeholder: 'adminViews.systemConfig.site.logoUrl.placeholder'
  },
  // {
  //   title: 'adminViews.systemConfig.site.subscribeUrl.title',
  //   shallow: 'adminViews.systemConfig.site.subscribeUrl.shallow',
  //   model: 'subscribe_url',
  //   type: 'input',
  //   placeholder: 'adminViews.systemConfig.site.subscribeUrl.placeholder'
  // },
  {
    title: 'adminViews.systemConfig.site.tosUrl.title',
    shallow: 'adminViews.systemConfig.site.tosUrl.shallow',
    model: 'tos_url',
    type: 'input',
    placeholder: 'adminViews.systemConfig.site.tosUrl.placeholder'
  },
  {
    title: 'adminViews.systemConfig.site.stopRegister.title',
    shallow: 'adminViews.systemConfig.site.stopRegister.shallow',
    model: 'stop_register',
    type: 'switch'
  },
  // {
  //   title: 'adminViews.systemConfig.site.inviteRequire.title',
  //   shallow: 'adminViews.systemConfig.site.inviteRequire.shallow',
  //   model: 'invite_require',
  //   type: 'switch'
  // },
  // {
  //   title: 'adminViews.systemConfig.site.trialSubscribe.title',
  //   shallow: 'adminViews.systemConfig.site.trialSubscribe.shallow',
  //   model: 'trial_subscribe',
  //   type: 'select'
  // },
  // {
  //   title: 'adminViews.systemConfig.site.trialTime.title',
  //   shallow: 'adminViews.systemConfig.site.trialTime.shallow',
  //   model: 'trial_time',
  //   type: 'input-number'
  // },
  {
    title: 'adminViews.systemConfig.site.currency.title',
    shallow: 'adminViews.systemConfig.site.currency.shallow',
    model: 'currency',
    type: 'input',
    placeholder: 'adminViews.systemConfig.site.currency.placeholder'
  },
  {
    title: 'adminViews.systemConfig.site.currencySymbol.title',
    shallow: 'adminViews.systemConfig.site.currencySymbol.shallow',
    model: 'currency_symbol',
    type: 'input',
    placeholder: 'adminViews.systemConfig.site.currencySymbol.placeholder'
  }
];


// let plans = [
//   {
//     id: 1, //  对应subscribe_list中的value
//     group_id: 5,
//     name: 'S1订阅', //  对应subscribe_list中的label
//     is_sale: true,  // 取反后对应subscribe_list中的disabled
//     is_renew: false,
//     capacity_limit: 120,
//     month_price: 12,
//     quarter_price: 33,
//     half_year_price: 66,
//     year_price: 129,
//   },
//   {
//     id: 2, //  对应subscribe_list中的value
//     group_id: 6,
//     name: 'S2订阅', //  对应subscribe_list中的label
//     is_sale: true,  // 取反后对应subscribe_list中的disabled
//     is_renew: true,
//     capacity_limit: 120,
//     month_price: 12,
//     quarter_price: 33,
//     half_year_price: 66,
//     year_price: 129,
//   }
// ]


// let subscribe_list = ref([
//   {
//     label: '关闭',
//     value: 0,
//     disabled: false,
//   },
// ])
//
// let getSubscribeList = async () => {
//   try {
//     let {data} = await instance.get('http://localhost:8081/api/admin/v1/plan/kv')
//     if (data.code === 200) {
//       console.log(data)
//       data.plans.forEach((plan: Plan) => {
//         subscribe_list.value.push({
//           label: plan.name,
//           value: plan.id,
//           disabled: plan.is_sale
//         })
//       })
//     }
//   } catch (error) {
//     console.log(error)
//   }
// }

const saveFiled = async (k: string, v: any) => {
  let updateResponse = await settingStore.saveOption('site', k, v)
  if (updateResponse.code === 200)
    message.success(t('adminViews.systemConfig.common.success'))
  else
    message.error(t('adminViews.systemConfig.common.err') + updateResponse.msg || '')
}

onMounted(() => {
  // getSubscribeList()

})

</script>

<script lang="ts">
export default {
  name: 'Site'
}
</script>

<template>
  <div class="root">
    <n-card
        :embedded="true"
        class="security-panel"
        :title="t('adminViews.systemConfig.site.common.title')"
        :bordered="false"
    >
      <n-grid cols="1" responsive="screen" y-gap="16">
        <n-grid-item v-for="setting in appSettings" :key="setting.title" class="grid-item">
          <!-- 每个配置项内部使用 n-grid 实现响应式布局 -->
          <n-grid cols="1 s:1 m:2 l:2" responsive="screen" align-items="center">
            <!-- 左侧：标题 + 描述 -->
            <n-grid-item>
              <div class="describe">
                <p class="title">{{ t(setting.title) }}</p>
                <p class="shallow">{{ t(setting.shallow) }}</p>
              </div>
            </n-grid-item>

            <!-- 右侧：输入框 / 选择框 / 开关 -->
            <n-grid-item>
              <n-input
                  v-if="setting.type === 'input'"
                  v-model:value="settingStore.settings.site[setting.model]"
                  type="text"
                  :placeholder="t(setting.placeholder || '')"
                  size="large"
                  @blur="saveFiled(setting.model, settingStore.settings.site[setting.model])"
              />
              <n-input
                  v-if="setting.type === 'input-area'"
                  v-model:value="settingStore.settings.site[setting.model]"
                  type="textarea"
                  :rows="2"
                  :placeholder="t(setting.placeholder || '')"
                  size="large"
                  @blur="saveFiled(setting.model, settingStore.settings.site[setting.model])"
              />
              <div
                  v-else-if="setting.type === 'switch'"
                  style="text-align: right"
              >
                <n-switch
                    size="medium"
                    v-model:value="settingStore.settings.site[setting.model]"
                    @update:value="saveFiled(setting.model, settingStore.settings.site[setting.model])"
                />
              </div>
<!--              <n-select-->
<!--                  v-else-if="setting.type === 'select'"-->
<!--                  v-model:value="settingStore.settings.site[setting.model]"-->
<!--                  size="large"-->
<!--                  :options="subscribe_list"-->
<!--                  @update:value="saveFiled(setting.model, settingStore.settings.site[setting.model])"-->
<!--              />-->
              <n-input-number
                  v-else-if="setting.type === 'input-number'"
                  v-model:value.number="settingStore.settings.site[setting.model]"
                  type="text"
                  :placeholder="t('adminViews.systemConfig.site.inputNumberPlaceholder')"
                  size="large"
                  @blur="saveFiled(setting.model, settingStore.settings.site[setting.model])"
              />
            </n-grid-item>
          </n-grid>
        </n-grid-item>
      </n-grid>
    </n-card>
  </div>
</template>

<style lang="less" scoped>
.root {
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
</style>