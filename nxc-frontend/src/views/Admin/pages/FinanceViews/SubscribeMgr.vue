<script setup lang="ts" name="SubscribeMgr">
import {onMounted, ref} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import type { NotificationType, DrawerPlacement, FormInst } from 'naive-ui'
import { useNotification } from 'naive-ui'
import instance from "@/axios";
import {useMessage} from 'naive-ui'
// import {MdEditor} from "md-editor-v3";

const notification = useNotification()
const apiAddrStore = useApiAddrStore();
const formRef = ref<FormInst | null>(null)
const message = useMessage()
let formValue = ref({
  plan: {
    name: '',
    describe: '',
    is_sale: false,
    is_renew: false,
    group_id: null,
    capacity_limit: null,
    month_price: null,
    quarter_price: null,
    half_year_price: null,
    year_price: null,
    sort: null,
  }
})

let rules = {
  plan: {
    name: {
      required: true,
      message: '文档标题是必填的',
      trigger: 'blur'
    },
    describe: {
      required: true,
      message: '描述区域是必填的',
      trigger: 'blur'
    },
    is_sale: {
      required: false,
    },
    is_renew: {
      required: false,
    },
    group_id: {
      required: false,
    },
    capacity_limit: {
      required: true,
      message: '文档标题是必填的',
      trigger: 'changed'
    },
    month_price: {
      required: false,
    },
    quarter_price: {
      required: false
    },
    half_year_price: {
      required: false,
    },
    year_price: {
      required: false,
    },
    sort: {
      required: false,
    }
  }
}



let groupOptions = ref<{label: string, value: number}[]>([])

const themeStore = useThemeStore();

const active = ref(false)
const placement = ref<DrawerPlacement>('right')
const activate = (place: DrawerPlacement) => {
  active.value = true
  placement.value = place
}

let handleAddSubscribe = () => {
  console.log('处理添加一个订阅')
  getAllGroups()  // 拉取群组信息
  active.value = true // 打开模态框
}

let getAllSubscribeItems = async () => {
  let {data} = await instance.get('/api/admin/get-all-subscribes')
  console.log('获取所有的订阅', data)
}

let cancelAdd = () => {
  active.value = false
}

interface RespData {
  code: number
  msg?: string
  error?: string
  [key: string]: any; // 允许其他字段，键是字符串，值可以是任意类型
}

let submitNewPlan = async () => {
  console.log(formValue.value)
  try {
    let {data} = await instance.post(apiAddrStore.apiAddr.admin.addNewPlan, {
      ...formValue.value.plan,
    })
    if (data.code === 200) {
      notify('success', '成功', '成功添加新的订阅')
    } else {
      notify('error', '添加出错', data.error)
    }
  }catch (error) {
    console.log(error)
  }
}

interface GroupItemFromServer {
  id: number
  label: string
  name: string
  created_at?: string
  updated_at?: string
  deleted_at?: string
}

let getAllGroups = async () => {
  try {
    // groupList.value = []
    let {data} = await instance.get(apiAddrStore.apiAddr.admin.getAllPrivilegeGroups)
    if (data.code === 200) {
      console.log('获取到的权限组列表', data)
      // data.group_list.forEach((group: PrivilegeGroup) => groupList.push(group))
      data.group_list.forEach((group: GroupItemFromServer) => groupOptions.value.push({
        label: group.name,
        value: group.id,
      }))
    }
  }catch (error) {
    console.log(error)
  }
}

let notify = (type: NotificationType, title: string, meta: string) => {
  notification[type]({
    content: title,
    meta: meta,
    duration: 2500,
    keepAliveOnHover: true
  })
}

onMounted(() => {
  console.log('SubscribeMgr挂载')
  // getAllSubscribeItems()

  themeStore.menuSelected = 'subscribe-manager'
  themeStore.contentPath = '/admin/dashboard/subscribemanager'


  // themeStore.
})

</script>

<template>
  <div class="root">
    <n-card class="card" hoverable :embedded="true" title="订阅管理" :bordered="false">
      <n-button type="primary" :bordered="false" class="add-btn" @click="handleAddSubscribe">添加订阅</n-button>
    </n-card>

    <n-drawer v-model:show="active" width="60%" :placement="placement">
      <n-drawer-content title="添加订阅">
        <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
        >
          <n-form-item label="订阅名称" path="plan.name">
            <n-input v-model:value="formValue.plan.name" placeholder="输入文档标题"/>
          </n-form-item>
          <n-form-item label="订阅描述" path="plan.describe">
            <n-input type="textarea" :rows="8" v-model:value="formValue.plan.describe" placeholder="输入订阅描述"/>
          </n-form-item>
          <n-form-item label="启用销售" path="plan.is_sale" >
            <n-switch v-model:value="formValue.plan.is_sale" />
          </n-form-item>
          <n-form-item label="允许续费" path="plan.is_renew" >
            <n-switch v-model:value="formValue.plan.is_renew"/>
          </n-form-item>
          <n-form-item label="群组ID" path="plan.group_id">
            <n-select :options="groupOptions" v-model:value.number="formValue.plan.group_id" placeholder="选择所属群组"/>
          </n-form-item>
          <n-form-item label="最大允许用户数目" path="plan.capacity_limit">
            <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.capacity_limit" placeholder="输入最大承载的用户数"/>
          </n-form-item>
          <n-form-item label="排序" path="plan.sort">
            <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.sort" placeholder="用于前端排序"/>
          </n-form-item>
          <n-form-item label="月付价格" path="plan.month_price">
            <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.month_price" placeholder="输入月付价格"/>
          </n-form-item>
          <n-form-item label="季付价格" path="plan.quarter_price">
            <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.quarter_price" placeholder="输入季付价格"/>
          </n-form-item>
          <n-form-item label="半年付价格" path="plan.half_year_price">
            <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.half_year_price" placeholder="输入半年付价格"/>
          </n-form-item>
          <n-form-item label="年付价格" path="plan.year_price">
            <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.year_price" placeholder="输入年付价格"/>
          </n-form-item>
        </n-form>
        <div style="display: flex; flex-direction: row; justify-content: right; margin-top: 20px">
          <n-button style="margin-right: 15px; width: 80px" @click="cancelAdd" secondary type="primary">取消</n-button>
          <n-button style="width: 80px" type="primary" @click="submitNewPlan">提交</n-button>
        </div>
      </n-drawer-content>
    </n-drawer>

  </div>
</template>

<style scoped lang="less">
.root {
  padding: 20px;

  .card {
    .add-btn {
      margin-top: 10px;
    }
  }
}
</style>