<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {h, ref, onMounted} from "vue"
import useApiAddrStore from "@/stores/useApiAddrStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";
import instance from "@/axios";

const {t} = useI18n();
const apiAddrStore = useApiAddrStore();
const userInfoStore = useUserInfoStore();
const themeStore = useThemeStore();

interface Key {
  key_id: number  // 密鑰id
  order_id: number  // 訂單id
  key: string // 序列號
  plan_name: string // 訂閱計畫名稱
  is_used: boolean  // 是否被激活過
  is_active: boolean  // 是否可用
  created_at: string  // 密鑰創建日期
  expired_at: string  // 密鑰到期日期
}

let myKeys = ref<Key>([])

let testKeys = ref<Key[]>([
  {
    key_id: 101,
    order_id: 5001,
    key: "SN1234ABC",
    plan_name: "Basic Plan",
    is_used: true,
    is_active: false,
    created_at: "2024-01-10T08:00:00Z",
    expired_at: "2024-07-10T08:00:00Z",
  },
  {
    key_id: 102,
    order_id: 5002,
    key: "SN5678DEF",
    plan_name: "Pro Plan",
    is_used: false,
    is_active: true,
    created_at: "2024-03-15T09:00:00Z",
    expired_at: "2025-03-15T09:00:00Z",
  },
  {
    key_id: 103,
    order_id: 5003,
    key: "SN9101GHI",
    plan_name: "Premium Plan",
    is_used: true,
    is_active: true,
    created_at: "2024-04-20T11:30:00Z",
    expired_at: "2025-04-20T11:30:00Z",
  },
  {
    key_id: 104,
    order_id: 5004,
    key: "SN1112JKL",
    plan_name: "Basic Plan",
    is_used: false,
    is_active: false,
    created_at: "2024-06-01T12:00:00Z",
    expired_at: "2024-12-01T12:00:00Z",
  },
  {
    key_id: 105,
    order_id: 5005,
    key: "SN1314MNO",
    plan_name: "Pro Plan",
    is_used: false,
    is_active: true,
    created_at: "2024-08-05T13:45:00Z",
    expired_at: "2025-08-05T13:45:00Z",
  }
])


// getAllMyKeys 获取所有用户的密钥
let getAllMyKeys = async () => {
  try {
    let {data} = await instance.get(apiAddrStore.apiAddr.user.getAllMyKeys, {
      params: {
        user_id: userInfoStore.thisUser.id,
      }
    })
    if (data.code === 200) {
      console.log(data)
      //
    }
  } catch (error) {
    console.log(error)
  }
}

onMounted(() => {
  themeStore.userPath = '/dashboard/keys'
  themeStore.menuSelected = 'user-keys'
  getAllMyKeys()
})

</script>

<script lang="ts">
export default {
  name: 'UserKeys',
}
</script>

<template>
  <div class="root">
      <n-card class="title" :embedded="true" hoverable :bordered="false" :title="t('userKeys.myKeys')">

      </n-card>

      <n-card class="key-card" :embedded="true" hoverable :bordered="false" content-style="padding: 0;" v-for="item in testKeys" :key="item.key_id">
        <div class="key-item" >
          <div class="l-part">
            <p class="key-id">{{ item.key }}</p>
            <div style="display: flex; flex-direction: row">
              <n-tag :bordered="false" size="small" style="margin-right: 10px" :type="item.is_active?'success':'error'">{{ item.is_active?'活躍':'不可用' }}</n-tag>
              <n-tag :bordered="false" size="small" v-if="item.is_used" type="warning">已過期</n-tag>

            </div>
          </div>
          <div class="r-part">
<!--            <p class="key">{{ item.key }}</p>-->
<!--            <p class="plan-name">{{ item.plan_name }}</p>-->
          </div>
        </div>
      </n-card>
  </div>
</template>

<style scoped lang="less">
.root {
  padding: 20px;
  .title {

  }
  .key-card {
    margin-top: 20px;
    .key-item {
      padding: 20px;
      display: flex;
      flex-direction: row;
      flex-wrap: wrap;
      .l-part {
        flex: 3;
       .key-id {
         font-size: 1.2rem;
         opacity: 0.8;
         padding-bottom: 5px;
       }
        .order-id {
          font-size: 1.1rem;
          opacity: 0.8;
        }
      }
      .r-part {
        flex: 2;
      }
    }
  }
}
</style>