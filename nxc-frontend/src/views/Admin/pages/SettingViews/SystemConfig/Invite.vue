<script setup lang="ts">
import useThemeStore from '@/stores/useThemeStore'
import useSettingStore from "@/stores/useSettingStore";
const themeStore = useThemeStore()
const settingStore = useSettingStore()
let handok = () => {

}

let sendTestMail = () => {

}

</script>

<script lang="ts">
export default {
  name: 'Subscribe'
}
</script>

<template>

  <n-card class="root" :embedded="true" title="支付和返利" :bordered="false">
<!--    <n-alert title="Warning 警告" type="warning" style="margin-bottom: 30px" :bordered="false">-->
<!--      如果你更改了本页配置，需要对队列服务进行重启。&nbsp;另外本页配置优先级高于.env中邮件配置。-->
<!--    </n-alert>-->

    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">启用用户返利</p>
            <p class="shallow">如果启用则在被邀请的用户充值时，按照下面设置的充值比例给予用户返利。</p>
          </div>
        </span>
      <span class="r-content" style="text-align: right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.invite.invite_rebate_enable"
              @update:value="settingStore.saveOption('invite','invite_rebate_enable', settingStore.settings.invite.invite_rebate_enable)"
          />
        </span>
    </div>

    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">返利比例</p>
            <p class="shallow">设置返利的金额比例。</p>
          </div>
        </span>
      <span class="r-content" style="">
          <n-input-number
              size="large"
              v-model:value="settingStore.settings.invite.invite_rebate_rate"
              @blur="settingStore.saveOption('invite','invite_rebate_rate', settingStore.settings.invite.invite_rebate_rate)"
          />
        </span>
    </div>

    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">优惠信息</p>
            <p class="shallow">设置优惠信息，它将会展示在充值页面顶部。</p>
          </div>
        </span>
      <span class="r-content" style="">
          <n-input
              size="large"
              v-model:value="settingStore.settings.invite.discount_info"
              :placeholder="'设置优惠信息'"
              @blur="settingStore.saveOption('invite','discount_info', settingStore.settings.invite.discount_info)"
          />
        </span>
    </div>

    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">邀请信息</p>
            <p class="shallow">设置邀请信息，它将会展示在用户邀请页面，用于显示返利比例。</p>
          </div>
        </span>
      <span class="r-content" style="">
          <n-input
              size="large"
              v-model:value="settingStore.settings.invite.invite_info"
              :placeholder="'设置返利信息'"
              @blur="settingStore.saveOption('invite','invite_info', settingStore.settings.invite.invite_info)"
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