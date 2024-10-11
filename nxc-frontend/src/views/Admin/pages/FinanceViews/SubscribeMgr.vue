<script setup lang="ts">
import {onBeforeMount, onMounted, ref} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import usePaymentStore from "@/stores/usePaymentStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import type {DrawerPlacement, FormInst, NotificationType} from 'naive-ui'
import {NButton, NSwitch, NTag, useMessage, useNotification} from 'naive-ui'
import instance from "@/axios";
// import {MdEditor} from "md-editor-v3";

const notification = useNotification()
const apiAddrStore = useApiAddrStore();
const paymentStore = usePaymentStore();
const userInfoStore = useUserInfoStore();
const formRef = ref<FormInst | null>(null)
const message = useMessage()

let editType = ref<'add' | 'update'>('add') // 这里可以选择的是add和update

interface Plan {
  id: number
  group_id?: number
  is_renew?: boolean
  is_sale?: boolean
  name: string
  capacity_limit: number
  residue: number
  describe?: string
  month_price?: number
  quarter_price?: number
  half_year_price?: number
  year_price?: number
  created_at: string
  updated_at?: number
  deleted_at?: string
}

let formValue = ref({
  plan: {
    name: '',
    describe: '',
    is_sale: false,
    is_renew: false,
    group_id: null,
    capacity_limit: null,
    residue: null,
    month_price: null,
    quarter_price: null,
    half_year_price: null,
    year_price: null,
    sort: null,
  } as Plan
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


let groupOptions = ref<{ label: string, value: number }[]>([])

const themeStore = useThemeStore();

const active = ref(false)
const placement = ref<DrawerPlacement>('right')
const activate = (place: DrawerPlacement) => {
  active.value = true
  placement.value = place
}

let handleAddSubscribe = () => {
  console.log('处理添加一个订阅')
  editType.value = 'add'  // 添加操作
  formValue.value.plan = {}
  getAllGroups()  // 拉取群组信息
  active.value = true // 打开模态框
}

// let getAllSubscribeItems = async () => {
//   let {data} = await instance.get('/api/admin/get-all-subscribes')
//   console.log('获取所有的订阅', data)
// }

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
    let {data} = await instance.post('/api/admin/v1/plan', {
      ...formValue.value.plan,
    })
    if (data.code === 200) {
      notify('success', '成功', '成功添加新的订阅')
    } else {
      notify('error', '添加出错', data.error)
    }
  } catch (error) {
    console.log(error)
  }
}

let submitPlan = async () => {
  console.log(formValue.value)
  switch (editType.value) {
    case 'add': {
      await submitNewPlan()
      break
    }
    case 'update': {
      await updatePlan()
      break
    }
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
    let {data} = await instance.get('/api/admin/v1/groups/kv', )
    if (data.code === 200) {
      console.log('获取到的权限组列表', data)
      // data.group_list.forEach((group: PrivilegeGroup) => groupList.push(group))
      data.group_list.forEach((group: GroupItemFromServer) => groupOptions.value.push({
        label: group.name,
        value: group.id,
      }))
    }
  } catch (error) {
    console.log(error)
  }
}

let notify = (type: NotificationType, title: string, meta?: string) => {
  notification[type]({
    content: title,
    meta: meta,
    duration: 2500,
    keepAliveOnHover: true
  })
}


// let planlist = ref<Plan[]>([])

const columns = [
  {
    title: '#',
    key: 'id'
  },
  {
    title: '启用销售',
    key: 'is_sale',
    render(row: Plan) {
      return h(NSwitch, {
        value: row.is_sale,
        onUpdateValue: (value) => updateRowIsSale(row.id, value)
      });
    }
  },
  {
    title: '允许续费',
    key: 'is_renew',
    render(row: Plan) {
      return h(NSwitch, {
        value: row.is_renew,
        onUpdateValue: (value) => updateRowRenew(row.id, value)
      });
    }
  },
  {
    title: '排序',
    key: 'sort',
    render(row: Plan) {
      return h(NTag, {}, {default: () => row.hasOwnProperty('sort') ? row.sort.toString() : 'null'});
    }
  },
  {
    title: '权限组',
    key: 'group',
    render(row: Plan) {
      const group = groupOptions.value.find(option => option.value === row.group_id);
      return h(NTag, {}, {default: () => group ? group.label : '未知'});
    }
  },
  {
    title: '名称',
    key: 'name'
  },
  {
    title: '数量',
    key: 'capacity_limit'
  },
  {
    title: '余量',
    key: 'residue'
  },
  {
    title: '月付价格',
    key: 'month_price',
    render(row: Plan): any {
      return row.hasOwnProperty('month_price') ? row.month_price.toFixed(2) : 'null'
    }
  },
  {
    title: '季付价格',
    key: 'quarter_price',
    render(row: Plan): any {
      return row.hasOwnProperty('quarter_price') ? row.quarter_price.toFixed(2) : 'null'
    }
  },
  {
    title: '半年付价格',
    key: 'halfYear_price',
    render(row: Plan): any {
      return row.hasOwnProperty('half_year_price') ? row.half_year_price.toFixed(2) : 'null'
    }
  },
  {
    title: '年付价格',
    key: 'year_price',
    render(row: Plan): any {
      return row.hasOwnProperty('year_price') ? row.year_price.toFixed(2) : 'null'
    }
  },
  {
    title: '操作',
    key: 'actions',
    fixed: 'right',
    render(row: RowData) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          bordered: false,
          style: {marginLeft: '10px'},
          onClick: () => handleUpdate(row)
        }, {default: () => 'Update'}),
        h(NButton, {
          size: 'small',
          type: 'error',
          style: {marginLeft: '10px'},
          onClick: () => handleDelete(row.id)
        }, {default: () => 'Delete'})
      ]);
    }
  }
];

let updatePlan = async (row: Plan) => {
  console.log('updatePlan')
  // console.log(...row)
  try {
    let {data} = await instance.put('/api/admin/v1/plan', {
      ...formValue.value.plan,
      // ...row
    })
    if (data.code === 200) {
      notify('success', '成功', '修改订阅成功')
    } else {
      notify('error', '修改出错', data.error)
    }
  } catch (error) {
    console.log(error)
  }
}

let handleUpdate = (row: Plan) => {
  console.log(row)
  editType.value = 'update'
  formValue.value.plan = row
  active.value = true
}

let handleDelete = async (planId: number) => {
  console.log(planId)
  try {
    let {data} = await instance.delete('/api/admin/v1/plan', {
      // ...formValue.value.plan,
      // ...row
      params: {
        plan_id: planId,
      }
    })
    if (data.code === 200) {
      notify('success', '成功', '删除成功')
      await paymentStore.getAllPlans()
    } else {
      notify('error', '删除失败', data.error)
    }
  } catch (error) {
    console.log(error)
  }
}

let updateRowIsSale = async (id: number, val: boolean) => {
  console.log('updateRowIsSale: ', val)
  try {
    let {data} = await instance.put('/api/admin/v1/plan/sale', {
      id,
      is_sale: val,
    })
    if (data.code === 200) {
      notify('success', '更新成功')
      await paymentStore.getAllPlans()
    } else {
      notify('error', '更新失败')

    }
  } catch (error: Error) {
    console.log(error)
  }
}

let updateRowRenew = async (id: number, val: boolean) => {
  console.log(val)
  try {
    let {data} = await instance.put('/api/admin/v1/plan/renew', {
      id,
      is_renew: val,
    })
    if (data.code === 200) {
      notify('success', '更新成功')
      await paymentStore.getAllPlans()
    } else {
      notify('error', '更新失败')

    }
  } catch (error: Error) {
    console.log(error)
  }
}


onMounted(() => {
  console.log('SubscribeMgr挂载')
  // getAllSubscribeItems()

  themeStore.menuSelected = 'subscription-manager'
  themeStore.contentPath = '/admin/dashboard/subscribemanager'


  // themeStore.
})

onBeforeMount(async () => {
  await getAllGroups()
  await paymentStore.getAllPlans()

})

</script>

<script lang="ts">
export default {
  name: 'SubscribeMgr',
}
</script>

<template>
  <div class="root">
    <n-card class="card" hoverable :embedded="true" title="订阅管理" :bordered="false">
      <n-button type="primary" :bordered="false" class="add-btn" @click="handleAddSubscribe">添加订阅</n-button>
    </n-card>

    <n-card :embedded="true" :bordered="false" hoverable content-style="padding: 0;" style="margin-top: 20px">
      <!--      在此处放置表格-->
      <!--      表格的列名为 #ID， 启用销售（数据为bool，显示为一个n-switch），允许续费（数据为bool，显示为一个n-switch）,排序（数据为数值，显示为一个n-tag），权限组（数值为string，显示为一个n-tag），名称，数量，余量，月付价格，季付价格，半年付价格，年付价格，操作（显示为两个按钮，使用h函数，将按钮放置为一个div中，使用flex布局横向排列）-->
      <n-data-table
          size="medium"
          :columns="columns"
          :data="paymentStore.plan_list"
          scroll-x="1000"
      />
    </n-card>

    <n-drawer v-model:show="active" width="60%" :placement="placement">
      <n-drawer-content title="添加订阅">
        <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :scroll-x="800"
        >
          <n-form-item label="订阅名称" path="plan.name">
            <n-input v-model:value="formValue.plan.name" placeholder="输入文档标题"/>
          </n-form-item>
          <n-form-item label="订阅描述" path="plan.describe">
            <n-input type="textarea" :rows="8" v-model:value="formValue.plan.describe" placeholder="输入订阅描述"/>
          </n-form-item>
          <n-form-item label="启用销售" path="plan.is_sale">
            <n-switch v-model:value="formValue.plan.is_sale"/>
          </n-form-item>
          <n-form-item label="允许续费" path="plan.is_renew">
            <n-switch v-model:value="formValue.plan.is_renew"/>
          </n-form-item>
          <n-form-item label="群组ID" path="plan.group_id">
            <n-select :options="groupOptions" v-model:value.number="formValue.plan.group_id"
                      placeholder="选择所属群组"/>
          </n-form-item>
          <n-form-item label="最大允许用户数目" path="plan.capacity_limit">
            <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.capacity_limit"
                            placeholder="输入最大承载的用户数"/>
          </n-form-item>
          <n-form-item label="排序" path="plan.sort">
            <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.sort"
                            placeholder="用于前端排序"/>
          </n-form-item>
          <n-form-item label="月付价格" path="plan.month_price">
            <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.month_price"
                            placeholder="输入月付价格"/>
          </n-form-item>
          <n-form-item label="季付价格" path="plan.quarter_price">
            <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.quarter_price"
                            placeholder="输入季付价格"/>
          </n-form-item>
          <n-form-item label="半年付价格" path="plan.half_year_price">
            <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.half_year_price"
                            placeholder="输入半年付价格"/>
          </n-form-item>
          <n-form-item label="年付价格" path="plan.year_price">
            <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.year_price"
                            placeholder="输入年付价格"/>
          </n-form-item>
        </n-form>
        <div style="display: flex; flex-direction: row; justify-content: right; margin-top: 20px">
          <n-button style="margin-right: 15px; width: 80px" @click="cancelAdd" secondary type="primary">取消</n-button>
          <n-button style="width: 80px" type="primary" @click="submitPlan">{{
              editType === 'add' ? '添加' : '修改'
            }}
          </n-button>
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