<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from "vue";
import useThemeStore from "@/stores/useThemeStore";
// import useApiAddrStore from "@/stores/useApiAddrStore";
import usePaymentStore from "@/stores/usePaymentStore";
// import useUserInfoStore from "@/stores/useUserInfoStore";
import {type DrawerPlacement, type FormInst, type NotificationType, type DataTableColumns, NIcon} from 'naive-ui'
import {NButton, NSwitch, NTag, useMessage, useNotification} from 'naive-ui'
import {AddOutline as AddIcon} from "@vicons/ionicons5"

import instance from "@/axios";
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue";
// import {MdEditor} from "md-editor-v3";

const {t} = useI18n();
let animated = ref<boolean>(false)

const notification = useNotification()
// const apiAddrStore = useApiAddrStore();
const paymentStore = usePaymentStore();
// const userInfoStore = useUserInfoStore();
const formRef = ref<FormInst | null>(null)
const message = useMessage()

let editType = ref<'add' | 'update'>('add') // 这里可以选择的是add和update

let pageCount = ref(10)
let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

interface Plan {
  id: number | null
  group_id?: number | null
  is_renew?: boolean
  is_sale?: boolean
  name: string
  sort: number | null
  capacity_limit: number | null
  residue: number | null
  describe?: string
  month_price?: number | null
  quarter_price?: number | null
  half_year_price?: number | null
  year_price?: number | null
  created_at?: string
  updated_at?: number | null
  deleted_at?: string
}

let formValue = ref({
  plan: {
    id: null,
    group_id: null,
    name: '',
    describe: '',
    is_sale: false,
    is_renew: false,
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
// const activate = (place: DrawerPlacement) => {
//   active.value = true
//   placement.value = place
// }

let handleAddSubscribe = () => {
  editType.value = 'add'  // 添加操作
  Object.assign(formValue.value.plan, {
    id: null,
    group_id: null,
    name: '',
    describe: '',
    is_sale: false,
    is_renew: false,
    capacity_limit: null,
    residue: null,
    month_price: null,
    quarter_price: null,
    half_year_price: null,
    year_price: null,
    sort: null,
  })
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

// interface RespData {
//   code: number
//   msg?: string
//   error?: string
//
//   [key: string]: any; // 允许其他字段，键是字符串，值可以是任意类型
// }

let submitNewPlan = async () => {
  console.log(formValue.value)
  try {
    let {data} = await instance.post('/api/admin/v1/plan', {
      ...formValue.value.plan,
    })
    if (data.code === 200) {
      notify('success', '成功', '成功添加新的订阅')
      let {page_count, success} = await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize)
      if (success) {
        pageCount.value = page_count as number
      } else {
        console.log('pull failure')
      }
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
    let {data} = await instance.get('/api/admin/v1/groups/kv',)
    if (data.code === 200) {
      groupOptions.value = []
      // data.group_list.forEach((group: PrivilegeGroup) => groupList.push(group))
      data.group_list.forEach((group: GroupItemFromServer) => groupOptions.value.push({
        label: group.name,
        value: group.id,
      }))
    } else {
      message.error(data.msg || 'Error fetch.')
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


const columns = computed<DataTableColumns<Plan>>(() => [
  {
    title: t('adminViews.planMgr.table.id'),
    key: 'id'
  },
  {
    title: t('adminViews.planMgr.table.isSale'),
    key: 'is_sale',
    render(row: Plan) {
      return h(NSwitch, {
        value: row.is_sale,
        onUpdateValue: (value) => updateRowIsSale(row.id as number, value)
      });
    }
  },
  {
    title: t('adminViews.planMgr.table.isRenew'),
    key: 'is_renew',
    render(row: Plan) {
      return h(NSwitch, {
        value: row.is_renew,
        onUpdateValue: (value) => updateRowRenew(row.id as number, value)
      });
    }
  },
  {
    title: t('adminViews.planMgr.table.sort'),
    key: 'sort',
    render(row: Plan) {
      return h(NTag, {}, {default: () => row.hasOwnProperty('sort') ? row.sort?.toString() : 'null'});
    }
  },
  {
    // title: '权限组',
    title: t('adminViews.planMgr.table.group'),
    key: 'group',
    render(row: Plan) {
      const group = groupOptions.value.find(option => option.value === row.group_id);
      return h(NTag, {}, {default: () => group ? group.label : '未知'});
    }
  },
  {
    // title: '名称',
    title: t('adminViews.planMgr.table.name'),
    key: 'name'
  },
  {
    // title: '数量',
    title: t('adminViews.planMgr.table.nums'),
    key: 'capacity_limit'
  },
  {
    // title: '余量',
    title: t('adminViews.planMgr.table.residue'),
    key: 'residue'
  },
  {
    // title: '月付价格',

    title: t('period.month'),
    key: 'month_price',
    render(row: Plan): any {
      return row.hasOwnProperty('month_price') ? row.month_price?.toFixed(2) : 'null'
    }
  },
  {
    // title: '季付价格',
    title: t('period.quarter'),
    key: 'quarter_price',
    render(row: Plan): any {
      return row.hasOwnProperty('quarter_price') ? row.quarter_price?.toFixed(2) : 'null'
    }
  },
  {
    // title: '半年付价格',
    title: t('period.halfYear'),
    key: 'halfYear_price',
    render(row: Plan): any {
      return row.hasOwnProperty('half_year_price') ? row.half_year_price?.toFixed(2) : 'null'
    }
  },
  {
    // title: '年付价格',
    title: t('period.year'),
    key: 'year_price',
    render(row: Plan): any {
      return row.hasOwnProperty('year_price') ? row.year_price?.toFixed(2) : 'null'
    }
  },
  {
    // title: '操作',
    title: t('adminViews.planMgr.table.operate'),
    key: 'actions',
    fixed: 'right',
    render(row: Plan) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          secondary: true,
          bordered: false,
          style: {marginLeft: '10px'},
          onClick: () => handleUpdate(row)
        }, {default: () => t('adminViews.planMgr.table.operateBtn.update')}),
        h(NButton, {
          size: 'small',
          type: 'error',
          secondary: true,
          style: {marginLeft: '10px'},
          onClick: () => handleDelete(row.id as number)
        }, {default: () => t('adminViews.planMgr.table.operateBtn.delete')})
      ]);
    }
  },
])

let updatePlan = async () => {
  console.log('updatePlan')
  // console.log(...row)
  try {
    let {data} = await instance.put('/api/admin/v1/plan', {
      ...formValue.value.plan,
      // ...row
    })
    if (data.code === 200) {
      notify('success', '成功', '修改订阅成功')
      // await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize)
      let {page_count, success} = await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize)
      // pageCount.value = page_count as number
      if (success) {
        pageCount.value = page_count as number
        animated.value = false
        active.value = false
      } else {
        message.error('更新失败')
      }
      animated.value = true

    } else {
      // notify('error', '修改出错', data.error)
      message.error('修改出错' + data.msg || '')
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
    animated.value = false
    let {data} = await instance.delete('/api/admin/v1/plan', {
      // ...formValue.value.plan,
      // ...row
      params: {
        plan_id: planId,
      }
    })
    if (data.code === 200) {
      notify('success', '成功', '删除成功')
      // await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize) ? animated.value = true : setTimeout(() => animated.value = true, 3000)
      let {page_count, success} = await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize)
      // pageCount.value = page_count as number
      if (success) {
        pageCount.value = page_count as number
        animated.value = false
      }

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
      // await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize)
      let {page_count} = await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize)
      pageCount.value = page_count as number
    } else {
      notify('error', '更新失败')

    }
  } catch (error: any) {
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
      // await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize)
      let {page_count} = await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize)
      pageCount.value = page_count as number
    } else {
      notify('error', '更新失败')

    }
  } catch (error: any) {
    console.log(error)
  }
}

let reloadPlanList = async () => {
  // let data = await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize)
  // data?animated.value = true:message.error('Fetch failure.');
  let {page_count, success} = await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize)
  pageCount.value = page_count as number
  success ? animated.value = true : setTimeout(() => animated.value = true, 2000)
}


onMounted(async () => {
  themeStore.menuSelected = 'subscription-manager'
  themeStore.contentPath = '/admin/dashboard/subscribemanager'
  await getAllGroups()
  // await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize)
  let {page_count} = await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize)
  pageCount.value = page_count as number

  animated.value = true

  // themeStore.
})

onBeforeMount(async () => {


})

</script>

<script lang="ts">
export default {
  name: 'SubscribeMgr',
}
</script>

<template>

  <div>
    <PageHead
        :title="t('adminViews.planMgr.title')"
        :description="t('adminViews.planMgr.description')"
    >
      <n-button
          tertiary
          type="primary"
          size="medium"
          class="btn-right"
          @click="handleAddSubscribe">
        <template #icon>
          <n-icon>
            <AddIcon/>
          </n-icon>
        </template>

        {{ t('adminViews.planMgr.addNewPlan') }}
      </n-button>
    </PageHead>
  </div>

  <transition name="slide-fade">
    <div class="root" v-if="animated">
<!--      <n-card class="card" hoverable :embedded="true" :title="t('adminViews.planMgr.title')" :bordered="false">-->
<!--        <n-button type="primary" :bordered="false" class="add-btn" @click="handleAddSubscribe">添加订阅</n-button>-->
<!--      </n-card>-->

      <n-card :embedded="true" :bordered="false" hoverable content-style="padding: 0;" style="margin-top: 20px">
        <!--      在此处放置表格-->
        <!--      表格的列名为 #ID， 启用销售（数据为bool，显示为一个n-switch），允许续费（数据为bool，显示为一个n-switch）,排序（数据为数值，显示为一个n-tag），权限组（数值为string，显示为一个n-tag），名称，数量，余量，月付价格，季付价格，半年付价格，年付价格，操作（显示为两个按钮，使用h函数，将按钮放置为一个div中，使用flex布局横向排列）-->
        <n-data-table
            striped
            size="medium"
            :columns="columns"
            :data="paymentStore.plan_list"
            scroll-x="1000"
        />
      </n-card>

      <DataTableSuffix
          v-model:data-size="dataSize"
          v-model:page-count="pageCount"
          v-model:animated="animated"
          :update-data="reloadPlanList"
      />

      <n-drawer v-model:show="active" width="60%" :placement="placement">
        <n-drawer-content :title="t('adminViews.planMgr.addNewPlan')">
          <n-form
              ref="formRef"
              :model="formValue"
              :rules="rules"
              :scroll-x="800"
          >
            <n-form-item :label="t('adminViews.planMgr.form.items.name.title')" path="plan.name">
              <n-input v-model:value="formValue.plan.name" :placeholder="t('adminViews.planMgr.form.items.name.placeholder')"/>
            </n-form-item>
            <n-form-item :label="t('adminViews.planMgr.form.items.describe.title')" path="plan.describe">
              <n-input type="textarea" :rows="8" v-model:value="formValue.plan.describe" :placeholder="t('adminViews.planMgr.form.items.describe.placeholder')"/>
            </n-form-item>
            <n-form-item v-if="editType==='add'" :label="t('adminViews.planMgr.form.items.isSale.title')" path="plan.is_sale">
              <n-switch v-model:value="formValue.plan.is_sale"/>
            </n-form-item>
            <n-form-item v-if="editType==='add'" :label="t('adminViews.planMgr.form.items.isRenew.title')" path="plan.is_renew">
              <n-switch v-model:value="formValue.plan.is_renew"/>
            </n-form-item>
            <n-form-item :label="t('adminViews.planMgr.form.items.group.title')" path="plan.group_id">
              <n-select :options="groupOptions" v-model:value.number="formValue.plan.group_id"
                        :placeholder="t('adminViews.planMgr.form.items.group.placeholder')"/>
            </n-form-item>
            <n-form-item :label="t('adminViews.planMgr.form.items.capacityLimit.title')" path="plan.capacity_limit">
              <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.capacity_limit"
                              :placeholder="t('adminViews.planMgr.form.items.capacityLimit.placeholder')"/>
            </n-form-item>
            <n-form-item :label="t('adminViews.planMgr.form.items.sort.title')" path="plan.sort">
              <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.sort"
                              :placeholder="t('adminViews.planMgr.form.items.sort.placeholder')"/>
            </n-form-item>
            <n-form-item :label="t('period.month')" path="plan.month_price">
              <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.month_price"
                              :placeholder="t('adminViews.planMgr.form.items.periodPlaceholder.month')"/>
            </n-form-item>
            <n-form-item :label="t('period.quarter')" path="plan.quarter_price">
              <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.quarter_price"
                              :placeholder="t('adminViews.planMgr.form.items.periodPlaceholder.quarter')"/>
            </n-form-item>
            <n-form-item :label="t('period.halfYear')" path="plan.half_year_price">
              <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.half_year_price"
                              :placeholder="t('adminViews.planMgr.form.items.periodPlaceholder.halfYear')"/>
            </n-form-item>
            <n-form-item :label="t('period.year')" path="plan.year_price">
              <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.year_price"
                              :placeholder="t('adminViews.planMgr.form.items.periodPlaceholder.year')"/>
            </n-form-item>
          </n-form>
          <div style="display: flex; flex-direction: row; justify-content: right; margin-top: 20px">
            <n-button style="margin-right: 15px; width: 80px" @click="cancelAdd" secondary type="primary">{{ t('operate.cancel') }}
            </n-button>
            <n-button style="width: 80px" type="primary" @click="submitPlan">{{
                editType === 'add' ? t('operate.update') : t('operate.add')
              }}
            </n-button>
          </div>
        </n-drawer-content>
      </n-drawer>

    </div>
  </transition>

</template>

<style scoped lang="less">
.root {
  margin: 20px 20px 0 20px;

  .card {
    .add-btn {
      margin-top: 10px;
    }
  }
}
</style>