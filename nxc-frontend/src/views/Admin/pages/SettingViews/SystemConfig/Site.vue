<script setup lang="ts">
import useThemeStore from "@/stores/useThemeStore";
import useSettingStore from "@/stores/useSettingStore"
import instance from "@/axios";
import {onMounted, ref} from "vue";

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
    | "force_https"
    | "logo_url"
    | "subscribe_url"
    | "tos_url"
    | "stop_register"
    | "invite_require"
    | "trial_subscribe"
    | "trial_time"
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
let appSettings: AppSetting[] = [
  {
    title: "站点名称",
    shallow: "用于显示需要站点名称的地方。",
    model: "app_name",
    type: "input",
    placeholder: "站点名称"
  },
  {
    title: "站点副标题",
    shallow: "一般显示在主要标题的下面。",
    model: "app_sub_name",
    type: "input",
    placeholder: "副标题"
  },
  {
    title: "站点描述",
    shallow: "用于显示需要站点描述的地方。",
    model: "app_description",
    type: "input",
    placeholder: "站点描述"
  },
  {
    title: "站点网址",
    shallow: "当前网站最新网址，将会在邮件等需要用于网址处体现。",
    model: "app_url",
    type: "input",
    placeholder: "站点网址"
  },
  {
    title: "强制 HTTPS",
    shallow: "当站点没有使用 HTTPS，CDN 或反代开启强制 HTTPS 时需要开启。",
    model: "force_https",
    type: "switch"
  },
  {
    title: "LOGO",
    shallow: "用于显示需要 LOGO 的地方。",
    model: "logo_url",
    type: "input",
    placeholder: "logo 的 url 地址"
  },
  {
    title: "订阅 URL",
    shallow: "用于订阅所使用，留空则为站点 URL。如需多个订阅 URL 随机获取请使用逗号进行分割。",
    model: "subscribe_url",
    type: "input",
    placeholder: "订阅 url"
  },
  {
    title: "用户条款(TOS)URL",
    shallow: "用于跳转到用户条款(TOS)",
    model: "tos_url",
    type: "input",
    placeholder: "用户条款地址"
  },
  {
    title: "停止新用户注册",
    shallow: "开启后任何人都将无法进行注册。",
    model: "stop_register",
    type: "switch"
  },
  {
    title: "强制邀请",
    shallow: "开启后当新用户注册时必需填写邀请码。",
    model: "invite_require",
    type: "switch"
  },
  {
    title: "注册试用",
    shallow: "选择需要试用的订阅，如果没有选项请先前往订阅管理添加。",
    model: "trial_subscribe",
    type: "select"
  },
  {
    title: "试用时间(小时)",
    shallow: "新用户注册时订阅试用事件。",
    model: "trial_time",
    type: "input-number"
  },
  {
    title: "货币单位",
    shallow: "仅用于展示使用，更改后系统中所有的货币单位都将发生变更。",
    model: "currency",
    type: "input",
    placeholder: "CNY"
  },
  {
    title: "货币符号",
    shallow: "仅用于展示使用，更改后系统中所有的货币单位都将发生变更。",
    model: "currency_symbol",
    type: "input",
    placeholder: "¥"
  }
]

let plans = [
  {
    id: 1, //  对应subscribe_list中的value
    group_id: 5,
    name: 'S1订阅', //  对应subscribe_list中的label
    is_sale: true,  // 取反后对应subscribe_list中的disabled
    is_renew: false,
    capacity_limit: 120,
    month_price: 12,
    quarter_price: 33,
    half_year_price: 66,
    year_price: 129,
  },
  {
    id: 2, //  对应subscribe_list中的value
    group_id: 6,
    name: 'S2订阅', //  对应subscribe_list中的label
    is_sale: true,  // 取反后对应subscribe_list中的disabled
    is_renew: true,
    capacity_limit: 120,
    month_price: 12,
    quarter_price: 33,
    half_year_price: 66,
    year_price: 129,
  }
]


let subscribe_list = ref([
  {
    label: '关闭',
    value: 0,
    disabled: false,
  },
])

let getSubscribeList = async () => {
  try {
    let {data} = await instance.get('http://localhost:8081/api/admin/v1/plan/kv')
    if (data.code === 200) {
      console.log(data)
      data.plans.forEach((plan: Plan) => {
        subscribe_list.value.push({
          label: plan.name,
          value: plan.id,
          disabled: plan.is_sale
        })
      })
    }
  } catch (error) {
    console.log(error)
  }
}

onMounted(() => {
  getSubscribeList()

})

</script>

<script lang="ts">
export default {
  name: 'Site'
}
</script>

<template>
  <div class="root">
    <n-card :embedded="true" class="security-panel" title="站点" :bordered="false">
      <div v-for="setting in appSettings" :key="setting.title" class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">{{ setting.title }}</p>
            <p class="shallow">{{ setting.shallow }}</p>
          </div>
        </span>
        <span class="r-content" v-if="setting.type === 'input'">
          <n-input
              v-model:value="settingStore.settings.site[setting.model]"
              type="text"
              :placeholder="`${setting.placeholder}`"
              size="large"
              @blur="settingStore.saveOption('site', setting.model, settingStore.settings.site[setting.model])"
          />
        </span>
        <span class="r-content" v-if="setting.type === 'switch'" style="text-align: right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.site[setting.model]"
              @update:value="settingStore.saveOption('site', setting.model, settingStore.settings.site[setting.model])"
          />
        </span>
        <span class="r-content" v-if="setting.type === 'select'">
          <n-select
              v-model:value="settingStore.settings.site[setting.model]"
              size="large"
              :options="subscribe_list"
              @update:value="settingStore.saveOption('site', setting.model, settingStore.settings.site[setting.model])"
          />
        </span>
        <span class="r-content" v-if="setting.type === 'input-number'">
          <n-input-number
              v-model:value.number="settingStore.settings.site[setting.model]"
              type="text"
              placeholder="请输入"
              size="large"
              @blur="settingStore.saveOption('site', setting.model, settingStore.settings.site[setting.model])"
          />
        </span>
      </div>


      <!--      <div class="item">-->
      <!--        <span class="l-content">-->
      <!--          <div class="describe">-->
      <!--            <p class="title">站点名称</p>-->
      <!--            <p class="shallow">用于显示需要站点名称的地方。</p>-->
      <!--          </div>-->
      <!--        </span>-->
      <!--        <span class="r-content">-->
      <!--          <n-input-->
      <!--              v-model:value="settingStore.settings.site.app_name"-->
      <!--              type="text"-->
      <!--              placeholder="请输入站点名称"-->
      <!--              size="large"-->
      <!--              @blur="settingStore.saveOption('site','app_name', settingStore.settings.site.app_name)"-->
      <!--          />-->
      <!--        </span>-->
      <!--      </div>-->

      <!--      <div class="item">-->
      <!--        <span class="l-content">-->
      <!--          <div class="describe">-->
      <!--            <p class="title">站点描述</p>-->
      <!--            <p class="shallow">用于显示需要站点描述的地方。</p>-->
      <!--          </div>-->
      <!--        </span>-->
      <!--        <span class="r-content">-->
      <!--          <n-input-->
      <!--              v-model:value="settingStore.settings.site.app_description"-->
      <!--              type="text"-->
      <!--              placeholder="请输入站点描述"-->
      <!--              size="large"-->
      <!--              @blur="settingStore.saveOption('site','app_description', settingStore.settings.site.app_description)"-->
      <!--          />-->
      <!--        </span>-->
      <!--      </div>-->

      <!--      <div class="item">-->
      <!--        <span class="l-content">-->
      <!--          <div class="describe">-->
      <!--            <p class="title">站点网址</p>-->
      <!--            <p class="shallow">当前网站最新网址，将会在邮件等需要用于网址处体现。</p>-->
      <!--          </div>-->
      <!--        </span>-->
      <!--        <span class="r-content">-->
      <!--          <n-input-->
      <!--              v-model:value="settingStore.settings.site.app_url"-->
      <!--              type="text"-->
      <!--              placeholder="请输入站点网址"-->
      <!--              size="large"-->
      <!--              @blur="settingStore.saveOption('site','app_url', settingStore.settings.site.app_url)"-->
      <!--          />-->
      <!--        </span>-->
      <!--      </div>-->

      <!--      <div class="item">-->
      <!--        <span class="l-content">-->
      <!--          <div class="describe">-->
      <!--            <p class="title">强制HTTPS</p>-->
      <!--            <p class="shallow">当站点没有使用HTTPS，CDN或反代开启强制HTTPS时需要开启。</p>-->
      <!--          </div>-->
      <!--        </span>-->
      <!--        <span class="r-content" style="text-align: right">-->
      <!--          <n-switch-->
      <!--              size="medium"-->
      <!--              v-model:value="settingStore.settings.site.force_https"-->
      <!--              @update:value="settingStore.saveOption('site','force_https', settingStore.settings.site.force_https)"-->
      <!--          />-->
      <!--        </span>-->
      <!--      </div>-->

      <!--      <div class="item">-->
      <!--        <span class="l-content">-->
      <!--          <div class="describe">-->
      <!--            <p class="title">LOGO</p>-->
      <!--            <p class="shallow">用于显示需要LOGO的地方。</p>-->
      <!--          </div>-->
      <!--        </span>-->
      <!--        <span class="r-content">-->
      <!--          <n-input-->
      <!--              v-model:value="settingStore.settings.site.logo_url"-->
      <!--              type="text"-->
      <!--              placeholder="请输入logo的url地址"-->
      <!--              size="large"-->
      <!--              @blur="settingStore.saveOption('site','logo_url', settingStore.settings.site.logo_url)"-->
      <!--          />-->
      <!--        </span>-->
      <!--      </div>-->

      <!--      <div class="item">-->
      <!--        <span class="l-content">-->
      <!--          <div class="describe">-->
      <!--            <p class="title">订阅URL</p>-->
      <!--            <p class="shallow">用于订阅所使用，留空则为站点URL。如需多个订阅URL随机获取请使用逗号进行分割。</p>-->
      <!--          </div>-->
      <!--        </span>-->
      <!--        <span class="r-content">-->
      <!--          <n-input-->
      <!--              v-model:value="settingStore.settings.site.subscribe_url"-->
      <!--              type="text"-->
      <!--              placeholder="请输入订阅url"-->
      <!--              size="large"-->
      <!--              @blur="settingStore.saveOption('site','subscribe_url', settingStore.settings.site.subscribe_url)"-->
      <!--          />-->
      <!--        </span>-->
      <!--      </div>-->

      <!--      <div class="item">-->
      <!--        <span class="l-content">-->
      <!--          <div class="describe">-->
      <!--            <p class="title">用户条款(TOS)URL</p>-->
      <!--            <p class="shallow">用于跳转到用户条款(TOS)</p>-->
      <!--          </div>-->
      <!--        </span>-->
      <!--        <span class="r-content">-->
      <!--          <n-input-->
      <!--              v-model:value="settingStore.settings.site.tos_url"-->
      <!--              type="text"-->
      <!--              placeholder="请输入用户条款地址"-->
      <!--              size="large"-->
      <!--              @blur="settingStore.saveOption('site','tos_url', settingStore.settings.site.tos_url)"-->
      <!--          />-->
      <!--        </span>-->
      <!--      </div>-->

      <!--      <div class="item">-->
      <!--        <span class="l-content">-->
      <!--          <div class="describe">-->
      <!--            <p class="title">停止新用户注册</p>-->
      <!--            <p class="shallow">开启后任何人都将无法进行注册。</p>-->
      <!--          </div>-->
      <!--        </span>-->
      <!--        <span class="r-content" style="text-align: right">-->
      <!--          <n-switch-->
      <!--              size="medium"-->
      <!--              v-model:value="settingStore.settings.site.stop_register"-->
      <!--              @update:value="settingStore.saveOption('site','stop_register', settingStore.settings.site.stop_register)"-->
      <!--          />-->
      <!--        </span>-->
      <!--      </div>-->

      <!--      <div class="item">-->
      <!--        <span class="l-content">-->
      <!--          <div class="describe">-->
      <!--            <p class="title">注册试用</p>-->
      <!--            <p class="shallow">选择需要试用的订阅，如果没有选项请先前往订阅管理添加。</p>-->
      <!--          </div>-->
      <!--        </span>-->
      <!--        <span class="r-content">-->

      <!--          <n-select-->
      <!--              v-model:value="settingStore.settings.site.trial_subscribe"-->
      <!--              size="large"-->
      <!--              :options="subscribe_list"-->
      <!--              @update:value="-->
      <!--                console.log('选择的订阅id: ', settingStore.settings.site.trial_subscribe);-->
      <!--                settingStore.saveOption('site','trial_subscribe', settingStore.settings.site.trial_subscribe)"-->
      <!--          />-->
      <!--        </span>-->
      <!--      </div>-->

      <!--      <div class="item">-->
      <!--        <span class="l-content">-->
      <!--          <div class="describe">-->
      <!--            <p class="title">试用时间(小时)</p>-->
      <!--            <p class="shallow">新用户注册时订阅试用事件。</p>-->
      <!--          </div>-->
      <!--        </span>-->
      <!--        <span class="r-content">-->
      <!--          <n-input-number-->
      <!--              v-model:value.number="settingStore.settings.site.trial_time"-->
      <!--              type="text"-->
      <!--              placeholder="请输入"-->
      <!--              size="large"-->
      <!--              @update:value="settingStore.saveOption('site','trial_time', settingStore.settings.site.trial_time)"-->
      <!--          />-->
      <!--        </span>-->
      <!--      </div>-->

      <!--      <div class="item">-->
      <!--        <span class="l-content">-->
      <!--          <div class="describe">-->
      <!--            <p class="title">货币单位</p>-->
      <!--            <p class="shallow">仅用于展示使用，更改后系统中所有的货币单位都将发生变更。</p>-->
      <!--          </div>-->
      <!--        </span>-->
      <!--        <span class="r-content">-->
      <!--          <n-input-->
      <!--              v-model:value="settingStore.settings.site.currency"-->
      <!--              type="text"-->
      <!--              placeholder="CNY"-->
      <!--              size="large"-->
      <!--              @update:value="settingStore.saveOption('site','currency', settingStore.settings.site.currency)"-->
      <!--          />-->
      <!--        </span>-->
      <!--      </div>-->

      <!--      <div class="item" style="margin-bottom: 0">-->
      <!--        <span class="l-content">-->
      <!--          <div class="describe">-->
      <!--            <p class="title">货币符号</p>-->
      <!--            <p class="shallow">仅用于展示使用，更改后系统中所有的货币单位都将发生变更。</p>-->
      <!--          </div>-->
      <!--        </span>-->
      <!--        <span class="r-content">-->
      <!--          <n-input-->
      <!--              v-model:value="settingStore.settings.site.currency_symbol"-->
      <!--              type="text"-->
      <!--              placeholder="¥"-->
      <!--              size="large"-->
      <!--              @update:value="settingStore.saveOption('site','currency_symbol', settingStore.settings.site.currency_symbol)"-->
      <!--          />-->
      <!--        </span>-->
      <!--      </div>-->

    </n-card>

  </div>
</template>

<style lang="less" scoped>
.root {
  min-width: 900px;

  .security-panel {
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
  }
}

//.n-card {
//  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
//  border: 0;
//}
</style>