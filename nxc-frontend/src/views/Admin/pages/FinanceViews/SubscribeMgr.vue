<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from "vue";
import useThemeStore from "@/stores/useThemeStore";
// import useApiAddrStore from "@/stores/useApiAddrStore";
import usePaymentStore from "@/stores/usePaymentStore";
// import useUserInfoStore from "@/stores/useUserInfoStore";
import {
  type DrawerPlacement,
  type FormInst,
  type DataTableColumns,
  NIcon,
  type FormRules,
} from 'naive-ui'
import {NButton, NSwitch, NTag, useMessage, useDialog} from 'naive-ui'
import {AddOutline as AddIcon} from "@vicons/ionicons5"

import instance from "@/axios";
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue";
import useTablePagination from "@/hooks/useTablePagination";
// import useConfirmDialog from "@/hooks/useConfirmDialog"

import {
  handleDeletePlanById,
  handleSubmitNewPlan,
  handleUpdatePlan,
  handleUpdatePlanRowIsSaleOrRenew
} from "@/api/admin/plan";
import {handleGetAllGroupsKv} from "@/api/admin/groups";
// import {MdEditor} from "md-editor-v3";

const {t} = useI18n();
const i18nPrefix = 'adminViews.planMgr'
const dialog = useDialog()
let animated = ref<boolean>(false)

const paymentStore = usePaymentStore();
const message = useMessage()

let editType = ref<'add' | 'update'>('add') // 这里可以选择的是add和update

const [dataSize, pageCount] = useTablePagination()

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

const formRef = ref<FormInst | null>(null)
let formValue = ref<{ plan: Plan }>({
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
  }
})

// console.log(formValue.value.plan.name)

let rules = {
  plan: {
    name: {
      required: true,
      message: '',
      trigger: 'blur'
    },
    describe: {
      required: true,
      message: '',
      trigger: 'blur'
    },
    is_sale: {
      required: false,
    },
    is_renew: {
      required: false,
    },
    group_id: {
      required: true,
      message: '',
      type: 'number',
      trigger: ['blur', 'change']
    },
    capacity_limit: {
      required: true,
      type: 'number',
      message: '',
      trigger: ['blur', 'change'],
    },
    residue: {
      required: true,
      type: 'number',
      message: '',
      trigger: ['blur', 'change', 'change'],
      validator() {
        if (formValue.value.plan.residue && formValue.value.plan.capacity_limit)
          if (formValue.value.plan.residue > formValue.value.plan.capacity_limit) {
            return new Error('capacity cannot less than residue')
          } else return null
      }
    },
    month_price: {
      required: false,
      type: 'number',
      message: '',
      validator: () => formValue.value.plan.month_price?formValue.value.plan.month_price <= 0 ? new Error('price cannot less than zero') : null:  new Error('field is null')
    },
    quarter_price: {
      required: false,
      type: 'number',
      message: '',
      // validator: () => formValue.value.plan.quarter_price <= 0 ? new Error('price cannot less than zero') : null,
      validator: () => formValue.value.plan.quarter_price?formValue.value.plan.quarter_price <= 0 ? new Error('price cannot less than zero') : null:  new Error('field is null')

    },
    half_year_price: {
      required: false,
      type: 'number',
      message: '',
      // validator: () => formValue.value.plan.half_year_price <= 0 ? new Error('price cannot less than zero') : null,
      validator: () => formValue.value.plan.half_year_price?formValue.value.plan.half_year_price <= 0 ? new Error('price cannot less than zero') : null:  new Error('field is null')
    },
    year_price: {
      required: false,
      type: 'number',
      message: '',
      // validator: () => formValue.value.plan.year_price <= 0 ? new Error('price cannot less than zero') : null,
      validator: () => formValue.value.plan.year_price?formValue.value.plan.year_price <= 0 ? new Error('price cannot less than zero') : null:  new Error('field is null')
    },
    sort: {
      required: false,
    }
  } as FormRules
}

let groupOptions = ref<{ label: string, value: number }[]>([])

const themeStore = useThemeStore();

const active = ref(false)
const placement = ref<DrawerPlacement>('right')
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

let cancelAdd = () => {
  active.value = false
}

let submitNewPlan = async () => {
  animated.value = false
  if (!formValue.value.plan.name.trim()) {
    message.error(t('adminViews.common.formatNotAllowed'))
    return
  }
  let data = await handleSubmitNewPlan(formValue.value.plan)
  if (data && data.code === 200) {
    active.value = false
    message.success(t('adminViews.common.addSuccess'))
    await reloadPlanList()
  } else {
    message.success(t('adminViews.common.addFailure'))
  }
}

let updatePlan = async () => {
  let data = await handleUpdatePlan(formValue.value.plan)
  if (data && data.code === 200) {
    message.success(t('adminViews.common.updateSuccess'))
    await reloadPlanList()

  } else {
    message.success(t('adminViews.common.updateFailure') + data.msg || '')
  }
}

let validPlanForm = async (type: string) => {
  if (formRef.value) {
    await formRef.value.validate((err: any) => {
      if (!err) {
        animated.value = false
        type === 'add' ? submitNewPlan() : updatePlan()
      } else {
        message.error(t('adminViews.common.checkForm'))
      }
    })
    active.value = false
  }
}

let submitPlan = async () => {
  console.log(formValue.value)
  await validPlanForm(editType.value)
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
  let data = await handleGetAllGroupsKv()
  if (data && data.code === 200) {
    groupOptions.value = []
    if (data.group_list)
      data.group_list.forEach((group: GroupItemFromServer) => groupOptions.value.push({
        label: group.name,
        value: group.id,
      }))
  } else {
    message.error(t('adminViews.common.fetchDataFailure') + data.msg || '')
  }
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
      return h(NTag, {size: 'small', type: 'default', bordered: false}, {default: () => row.hasOwnProperty('sort') ? row.sort?.toString() : 'null'});
    }
  },
  {
    // title: '权限组',
    title: t('adminViews.planMgr.table.group'),
    key: 'group',
    render(row: Plan) {
      const group = groupOptions.value.find(option => option.value === row.group_id);
      return h(NTag, {size: 'small', type: 'primary', bordered: false}, {default: () => group ? group.label : '未知'});
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
          onClick: () => confirmDeletePlanWarn(row.id as number)
        }, {default: () => t('adminViews.planMgr.table.operateBtn.delete')})
      ]);
    }
  },
])

const confirmDeletePlanWarn = (id: number) => dialog.error({
  title: t('adminViews.common.dialog.delete'),
  content: t('adminViews.planMgr.mention.common.delMention'),
  positiveText: t('adminViews.common.confirm'),
  negativeText: t('adminViews.common.cancel'),
  showIcon: false,
  actionStyle: {
    marginTop: '20px',
  },
  contentStyle: {
    marginTop: '20px',
  },
  onPositiveClick: () => handleDelete(id as number),
})


let handleUpdate = (row: Plan) => {
  editType.value = 'update'
  // formValue.value.plan = row
  Object.assign(formValue.value.plan, row)
  active.value = true
}

let handleDelete = async (planId: number) => {
  animated.value = false
  let data = await handleDeletePlanById(planId)
  if (data.code === 200) {
    // notify('success', '成功', '删除成功')
    message.success(t('adminViews.common.deleteSuccess'))
    await reloadPlanList()
  } else {
    message.success(t('adminViews.common.deleteFailure'))
  }
}

let updateRowIsSale = async (id: number, val: boolean) => {
  let data = await handleUpdatePlanRowIsSaleOrRenew('sale', id, val)
  if (data && data.code === 200) {
    message.success(t('adminViews.common.success'))
    await reloadPlanList()
  } else {
    message.success(t('adminViews.common.failure'))
  }
}

let updateRowRenew = async (id: number, val: boolean) => {
  let data = await handleUpdatePlanRowIsSaleOrRenew('renew', id, val)
  if (data && data.code === 200) {
    message.success(t('adminViews.common.success'))
    await reloadPlanList()
  } else {
    message.success(t('adminViews.common.failure'))
  }
}

let reloadPlanList = async () => {
  // let data = await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize)
  // data?animated.value = true:message.error('Fetch failure.');
  let {page_count, success} = await paymentStore.getAllPlans(dataSize.value.page, dataSize.value.pageSize)
  pageCount.value = page_count as number
  // success ? null : setTimeout(() => animated.value = true, 2000)
  if (!success) message.error('adminViews.common.fetchDataFailure')
  animated.value = true
}

onBeforeMount(async () => {
  themeStore.breadcrumb = 'adminViews.planMgr.title'
  themeStore.menuSelected = 'subscription-manager'
})

onMounted(async () => {
  themeStore.contentPath = '/admin/dashboard/subscribemanager'
  await getAllGroups()

  await reloadPlanList()

  animated.value = true

  // themeStore.
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
            :scroll-x="1600"
        />
      </n-card>

      <DataTableSuffix
          v-model:data-size="dataSize"
          v-model:page-count="pageCount"
          v-model:animated="animated"
          :update-data="reloadPlanList"
      />

      <n-drawer v-model:show="active" width="40%" :placement="placement">
        <n-drawer-content :title="t('adminViews.planMgr.addNewPlan')">
          <n-form
              ref="formRef"
              :model="formValue"
              :rules="rules"
              :scroll-x="800"
          >
            <n-form-item :label="t('adminViews.planMgr.form.items.name.title')" path="plan.name">
              <n-input v-model:value="formValue.plan.name"
                       clearable
                       :placeholder="t('adminViews.planMgr.form.items.name.placeholder')"/>
            </n-form-item>
            <n-form-item :label="t('adminViews.planMgr.form.items.describe.title')" path="plan.describe">
              <n-input type="textarea" :rows="8" v-model:value="formValue.plan.describe"
                       :placeholder="t('adminViews.planMgr.form.items.describe.placeholder')"/>
            </n-form-item>
            <n-form-item v-if="editType==='add'" :label="t('adminViews.planMgr.form.items.isSale.title')"
                         path="plan.is_sale">
              <n-switch v-model:value="formValue.plan.is_sale"/>
            </n-form-item>
            <n-form-item v-if="editType==='add'" :label="t('adminViews.planMgr.form.items.isRenew.title')"
                         path="plan.is_renew">
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
            <n-form-item :label="t('adminViews.planMgr.form.items.planResidue.title')" path="plan.residue">
              <n-input-number :style="{width: '100%'}" v-model:value.number="formValue.plan.residue"
                              :placeholder="t('adminViews.planMgr.form.items.planResidue.placeholder')"/>
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
            <n-button style="margin-right: 15px; width: 80px" @click="cancelAdd" secondary type="primary">
              {{ t('operate.cancel') }}
            </n-button>
            <n-button style="width: 80px" type="primary" @click="submitPlan">{{
                editType === 'add' ? t('operate.add') : t('operate.update')
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