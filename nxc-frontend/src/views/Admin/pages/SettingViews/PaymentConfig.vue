<script setup lang="ts">
import {ref} from 'vue'
import type { DrawerPlacement } from 'naive-ui'
import useThemeStore from "@/stores/useThemeStore";

const themeStore = useThemeStore();
let active = ref(false);
const placement = ref<DrawerPlacement>('right')

let handleAddPaymentMethods = () => {
  console.log('handleAddPaymentMethods');
  active.value = !active.value
}

let paymentMethods = [
  {
    id: '0001',
    enabled: true,
    name: 'AlipayF2F',
    api_addr: 'https://alipay.com',
    notice_addr: 'https://api.ikanned.com'
  },
  {
    id: '0002',
    enabled: false,
    name: 'WeChat',
    api_addr: 'https://alipay.com',
    notice_addr: 'https://api.ikanned.com'
  },
  {
    id: '0002',
    enabled: false,
    name: 'WeChat',
    api_addr: 'https://alipay.com',
    notice_addr: 'https://api.ikanned.com'
  },
  {
    id: '0002',
    enabled: false,
    name: 'WeChat',
    api_addr: 'https://alipay.com',
    notice_addr: 'https://api.ikanned.com'
  },
  {
    id: '0002',
    enabled: false,
    name: 'WeChat',
    api_addr: 'https://alipay.com',
    notice_addr: 'https://api.ikanned.com'
  },
  {
    id: '0002',
    enabled: false,
    name: 'WeChat',
    api_addr: 'https://alipay.com',
    notice_addr: 'https://api.ikanned.com'
  },
]

</script>

<script lang="ts">
export default {
  name: 'PaymentConfig'
}
</script>


<template>
  <div class="root-root">
    <n-infinite-scroll style="height: 100vh" :distance="10">
      <n-card hoverable class="root" :embedded="true" title="支付设置" :bordered="false">
        <n-button type="primary" :bordered="false" class="add-btn" @click="handleAddPaymentMethods">添加支付方式</n-button>
        <div class="table">
          <n-table :single-line="false">
            <thead>
            <tr>
              <th>ID</th>
              <th>启用</th>
              <th>显示名称</th>
              <th>支付接口</th>
              <th>通知地址</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="item in paymentMethods" :key="item.id">
              <td>{{ item.id }}</td>
              <td>{{ item.enabled?'是':'否' }}</td>
              <td>{{ item.name }}</td>
              <td>{{ item.api_addr }}</td>
              <td>{{ item.notice_addr }}</td>
            </tr>

            </tbody>
          </n-table>
        </div>
      </n-card>


    </n-infinite-scroll>
  </div>

  <n-drawer v-model:show="active" :width="502" :placement="placement">
    <n-drawer-content title="添加支付方式">
      <div class="add-panel">
        <div class="item">
          <p class="title">显示名称</p>
          <n-input size="medium" placeholder="用于前端显示使用"></n-input>
        </div>

        <div class="item">
          <p class="title">图标URL(选填)</p>
          <n-input size="medium" placeholder="用于前端显示使用(https://x.com/icon.svg)"></n-input>
        </div>

        <div class="item">
          <p class="title">显示名称</p>
          <n-input size="medium" placeholder="网关的通知将会发送到该域名(https://x.com)"></n-input>
        </div>

        <div class="item">
          <p class="title">显示名称</p>
          <n-input size="medium" placeholder="用于前端显示使用"></n-input>
        </div>
      </div>
    </n-drawer-content>
  </n-drawer>


</template>

<style lang="less" scoped>
.root-root {
  padding: 20px;
}
.root {
  min-width: 900px;

  .add-btn {
    margin-top: 10px;
  }

  .table {
    margin-top: 30px;
  }
}
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

</style>