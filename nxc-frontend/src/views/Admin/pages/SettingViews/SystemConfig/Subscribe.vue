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

  <n-card class="root" :embedded="true" title="订阅管理" :bordered="false">
<!--    <n-alert title="Warning 警告" type="warning" style="margin-bottom: 30px" :bordered="false">-->
<!--      如果你更改了本页配置，需要对队列服务进行重启。&nbsp;另外本页配置优先级高于.env中邮件配置。-->
<!--    </n-alert>-->

    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">允许用户修改订阅</p>
            <p class="shallow">开启后用户将会可以对订阅计划进行变更。</p>
          </div>
        </span>
      <span class="r-content" style="text-align: right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.subscribe.user_modify_enable"
              @update:value="settingStore.saveOption('subscribe','user_modify_enable', settingStore.settings.subscribe.user_modify_enable)"
          />
        </span>
    </div>

    <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">在订阅中展示订阅信息</p>
            <p class="shallow">开启后将会在用户订阅节点时输出订阅信息。</p>
          </div>
        </span>
      <span class="r-content" style="text-align: right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.subscribe.show_info_in_sub"
              @update:value="settingStore.saveOption('subscribe','show_info_in_sub', settingStore.settings.subscribe.show_info_in_sub)"
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