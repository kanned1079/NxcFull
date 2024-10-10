<script setup lang="ts" name="Setting">
import {useRouter} from 'vue-router';
import useThemeStore from "@/stores/useThemeStore";
import {
  BagHandleSharp as SubIcon,
  PeopleSharp as UserIcon,
  ReorderFourSharp as OrderIcon,
  SettingsSharp as settingIcon,
} from '@vicons/ionicons5'

const router = useRouter();
const themeStore = useThemeStore();

// 方法
let enterSetting = () => {
  console.log('settings')
}
</script>

<template>

  <n-card hoverable @click="router.push({
    path: '/admin/dashboard/systemconfig',
  })" :embedded="true" :bordered="false">
    <div>
      <n-icon size="25">
        <settingIcon/>
      </n-icon>
      <p style="margin-top: 8px; font-size: 15px">系统设置</p>
    </div>
  </n-card>

  <n-card hoverable :embedded="true" :bordered="false">
    <div>
      <n-icon size="25">
        <SubIcon/>
      </n-icon>
      <p style="margin-top: 8px; font-size: 15px">订单管理</p>
    </div>
  </n-card>

  <n-card hoverable :embedded="true" :bordered="false">
    <div>
      <n-icon size="25">
        <OrderIcon/>
      </n-icon>
      <p style="margin-top: 8px; font-size: 15px">订阅管理</p>
    </div>
  </n-card>

  <n-card hoverable style="margin-right: 20px" :embedded="true" :bordered="false">
    <div>
      <n-icon size="25">
        <UserIcon/>
      </n-icon>
      <p style="margin-top: 8px; font-size: 15px">用户管理</p>
    </div>
  </n-card>

</template>

<style lang="less" scoped>
.n-card {
  max-width: 25%;
  margin: 20px 0 20px 20px;
  //background-color: #4cae4c;
  height: 100px;
}
//.n-card {
//  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
//  border: 0;
//}
</style>