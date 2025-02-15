<script setup lang="ts">
// import FormSuffix from "@/views/utils/FormSuffix.vue";
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from "vue";
import useThemeStore from "@/stores/useThemeStore";
// import useApiAddrStore from "@/stores/useApiAddrStore";
import {AddOutline as AddIcon} from "@vicons/ionicons5"

import {type DataTableColumns, type FormInst, type FormRules, NButton, NIcon, NSwitch, NTag, useMessage} from 'naive-ui'
// import useNoticesStore from "@/stores/useNoticesStore";
import instance from "@/axios";
import {formatTimestamp} from "@/utils/timeFormat"
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue";
import {
  handleActivateCouponById,
  handleDeleteCouponById,
  handleGetAllCoupon,
  handleUpdateCoupon
} from "@/api/admin/coupon";
import {handleFetchPlanKv} from "@/api/admin/plan";

const {t} = useI18n();
const i18nPrefix: string = 'adminViews.couponMgr';
// const apiAddrStore = useApiAddrStore();


const themeStore = useThemeStore()
// const noticesStore = useNoticesStore()
const message = useMessage()

let animated = ref<boolean>(false)

const formRef = ref<FormInst | null>(null)

let pageCount = ref(10)
let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})


// 从服务器拉取的优惠券列表
interface Coupon {
  id: number  // 优惠券id
  name: string // 优惠券名称
  enabled: boolean // 优惠券是否启用
  code: string  // 优惠券码
  percent_off: number   // 抵折百分比
  start_time: number  // 优惠券启用时间
  end_time: number  // 优惠券过期时间
  per_user_limit: number  // 每个用户限制的使用次数
  capacity: number  // 总共优惠券数量
  residue: number // 剩余数量
  plan_limit: number  // 限制可以使用的订阅计划id
  created_at: string  // gorm创建日期
  updated_at: string  // gorm更新日期
  deleted_at?: string  // gorm删除标记
}

let formValue = ref({
  coupon: {
    name: '',
    code: '',
    percent_off: null,
    start_time: null,
    end_time: null,
    per_user_limit: null,
    capacity: null,
    plan_limit: null,
  } as {
    id?: number
    name: string
    code: string
    percent_off: number | null
    start_time: number | null
    end_time: number | null
    per_user_limit: number | null
    capacity: number | null
    plan_limit: number | null
  }
})

const rules = computed<FormRules>(() => {
  return {
    coupon: {
      name: {
        required: true,
        trigger: 'blur',
        message: t(`${i18nPrefix}.modal.emptyNotAllow`),
      },
      code: {
        required: false,
      },
      percent_off: {
        type: 'number',
        required: true,
        trigger: 'blur',
        message: t(`${i18nPrefix}.modal.emptyNotAllow`),
      },
      capacity: {
        type: 'number',
        required: true,
        trigger: 'blur',
        message: t(`${i18nPrefix}.modal.emptyNotAllow`),
      },
      per_user_limit: {
        type: 'number',
        required: true,
        trigger: 'blur',
        message: t(`${i18nPrefix}.modal.emptyNotAllow`),
      }
    }
  }
})


// noticesStore.getAllNotices()

let showLoading = ref<boolean>(false)
// 显示表单
let showModal = ref(false)
let range = ref<[number, number]>([new Date().setMonth(new Date().getMonth() - 1), new Date().getTime()])
let plans = ref<{ label: string, value: number }[]>([])
let editType = ref<'create' | 'edit'>('create')

// 订阅计划的键值
let getPlanKV = async () => {
  plans.value = []
  // try {
  // let {data} = await instance.get('http://localhost:8081/api/admin/v1/plan/kv')
  let data = await handleFetchPlanKv()
  if (data && data.code === 200) {
    console.log(data)
    data.plans.forEach((item: { id: number, name: string }) => plans.value.push({
      label: item.name,
      value: item.id,
    }))
  } else {
    message.error(t(`${i18nPrefix}.fetchPlanKvFailure`))
  }
  // } catch (error) {
  //   console.log(error)
  // }

}


// couponList 从后端获取的所有优惠券列表
// let couponList = ref<Coupon[]>([])
let couponList = ref<Coupon[]>([])

// 在此模版中用于渲染表格
// 表格的字段名为 优惠券ID，是否启用（此处设计为一个switch可用于单独关闭或启用该优惠券），优惠券码（使用n-tag），剩余使用次数，启用时间，过期时间，操作（包含编辑和删除功能两个按钮）

let getAllCoupons = async () => {
  showLoading.value = true
  // 这里是使用get方法 获取所有的优惠券列表
  let data = await handleGetAllCoupon(dataSize.value.page, dataSize.value.pageSize)
  if (data && data.code === 200) {
    couponList.value = []

    // couponList.value = data.coupons;
    data.coupon_list.forEach((item: Coupon) => {
      couponList.value.push(item)
    })
    pageCount.value = data.page_count
    animated.value = true
  } else {
    message.error(t('adminViews.common.fetchDataFailure'));
  }
  showLoading.value = false
}

let activateACoupon = async (id: number, value: boolean) => {
  // console.log(id, value)
  // 这里是使用post方法 当switch改变时调用来启用或关闭优惠券
  // 请求地址 /admin/v1/coupon/status/update
  // 请求体 {优惠券id， 状态}
  // try {
  //   let {data} = await instance.put('http://localhost:8081/api/admin/v1/coupon/status', {
  //     id: id,
  //     status: value
  //   });
  let data = await handleActivateCouponById(id, value)
  if (data && data.code === 200) {
    // message.success('优惠券状态更新成功');
    message.success(t('adminViews.common.updateSuccess'));
    await getAllCoupons()
  } else {
    // message.error('更新优惠券状态失败');
    message.error(t('adminViews.common.updateFailure'));
  }
  await getAllCoupons();
  // } catch (error) {
  //   console.error(error);
  //   message.error(t('adminViews.common.updateFailure'));
  //   // message.error('请求更新优惠券状态失败');
  // }
}

const updateCouponClick = async (row: Coupon) => {
  Object.assign(formValue.value.coupon, row)
  range.value = [row.start_time, row.end_time]
  editType.value = 'edit'
  await getPlanKV()
  showModal.value = true // 修改好後再顯示
}

let updateACoupon = async () => {
  // 这里是点击编辑一个优惠券后 使用post方法保存到服务器
  // 请求地址 /admin/v1/coupon/update
  // 请求体 {优惠券id， 该优惠券结构}
  formValue.value.coupon.start_time = range.value[0]
  formValue.value.coupon.end_time = range.value[1]
  // try {
  animated.value = false
  // let {data} = await instance.put('http://localhost:8081/api/admin/v1/coupon', formValue.value.coupon);
  let data = await handleUpdateCoupon(formValue.value.coupon)
  if (data && data.code === 200) {
    message.success(t('adminViews.common.updateSuccess'));
    await getAllCoupons(); // 更新成功后刷新列表
  } else {
    message.error(t('adminViews.common.updateFailure'));
  }
  // } catch (error) {
  //   console.error(error);
  //   message.error(t('adminViews.common.updateFailure'));
  // }
}

let deleteACoupon = async (id: number) => {
  // 点击删除后 使用post方法 用于删除一个优惠券
  // 请求地址 /admin/v1/coupon/delete
  // 请求体 {优惠券id}
  // console.log(id)
  // try {
  animated.value = false
  // let {data} = await instance.delete('http://localhost:8081/api/admin/v1/coupon', {
  //   params: {
  //     coupon_id: id,
  //   }
  // });
  let data = await handleDeleteCouponById(id)
  if (data && data.code === 200) {
    message.success(t('adminViews.common.deleteSuccess'));
    await getAllCoupons(); // 删除成功后刷新列表
  } else {
    message.error(t('adminViews.common.deleteFailure'));
  }
  // } catch (error) {
  //   console.error(error);
  //   message.error(t('adminViews.common.deleteFailure'));
  // }
}

const i18nTablePrefix = 'adminViews.couponMgr.table'
const columns = computed<DataTableColumns<Coupon>>(() => [
  {
    title: t(`${i18nTablePrefix}.id`),
    key: 'id'
  },
  {
    title: t(`${i18nTablePrefix}.enabled`),
    key: 'enabled',
    render(row: Coupon) {
      return h(NSwitch, {
        value: row.enabled,
        onUpdateValue: (value) => activateACoupon(row.id, value),
      });
    }
  },
  {
    title: t(`${i18nTablePrefix}.name`),
    key: 'name'
  },
  {
    title: t(`${i18nTablePrefix}.code`),
    key: 'code',
    render(row: Coupon) {
      return h(NTag, {}, {default: () => row.code});
    }
  },
  {
    title: t(`${i18nTablePrefix}.residue`),
    key: 'residue',
  },
  {
    title: t(`${i18nTablePrefix}.startTime`),
    key: 'start_time',
    render(row: Coupon) {
      return h('p', {}, {default: () => formatTimestamp(row.start_time)});
    }
  },
  {
    title: t(`${i18nTablePrefix}.endTime`),
    key: 'end_time',
    render(row: Coupon) {
      return h('p', {}, {default: () => formatTimestamp(row.end_time)});
    }
  },
  {
    title: t(`${i18nTablePrefix}.actions`),
    key: 'actions',
    render(row: Coupon) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          secondary: true,
          bordered: false,
          style: {marginLeft: '10px'},
          onClick: () => updateCouponClick(row)
        }, {default: () => t(`${i18nTablePrefix}.edit`)}),
        h(NButton, {
          size: 'small',
          type: 'error',
          secondary: true,
          disabled: false,
          style: {marginLeft: '10px'},
          onClick: () => deleteACoupon(row.id)
        }, {default: () => t(`${i18nTablePrefix}.delete`)})
      ]);
    },
    fixed: 'right',
  }
])


let handleAddCoupon = async () => {
  // await getPlanKV()
  Object.assign(formValue.value.coupon, {
    name: '',
    code: '',
    percent_off: null,
    start_time: null,
    end_time: null,
    per_user_limit: null,
    capacity: null,
    plan_limit: null,
  });
  let nowDate: Date = new Date()
  range.value = [new Date().setMonth(nowDate.getMonth() - 1), nowDate.getTime()];
  // range.value = [new Date().getTime(), new Date().getTime()];
  editType.value = 'create';
  await getPlanKV()
  showModal.value = true
}

let submitCoupon = async () => {
  animated.value = false
  formValue.value.coupon.start_time = range.value[0]
  formValue.value.coupon.end_time = range.value[1]
  console.log(formValue.value)
  try {
    let {data} = await instance.post('http://localhost:8081/api/admin/v1/coupon', {
      ...formValue.value.coupon
    })
    if (data.code === 200) {
      message.success('添加优惠券成功')
      message.success(t('adminViews.common.addSuccess'))
      await getAllCoupons();
    } else {
      message.error(t('adminViews.common.addSuccess') + data.msg || '')
    }
  } catch (error) {
    console.error(error)
  }
}

let closeModal = () => {
  showModal.value = false
}
onBeforeMount(() => {
  themeStore.breadcrumb = t('adminViews.couponMgr.title')
  themeStore.menuSelected = 'coupon-mgr'
})

onMounted(async () => {
  // console.log('couponMgr挂载')
  await getAllCoupons()

  themeStore.contentPath = '/admin/dashboard/coupon';

  animated.value = true

})

</script>

<script lang="ts">
export default {
  name: 'CouponMgr',
}
</script>

<template>
  <PageHead
      :title="t('adminViews.couponMgr.title')"
      :description="t('adminViews.couponMgr.description')"
  >
    <n-button
        tertiary
        type="primary"
        size="medium"
        class="btn-right"
        @click="handleAddCoupon">
      <template #icon>
        <n-icon>
          <AddIcon/>
        </n-icon>
      </template>

      {{ t(`${i18nPrefix}.modal.newCoupon`) }}
    </n-button>
  </PageHead>

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card :embedded="false" hoverable :bordered="false" content-style="padding: 0;" style="margin-top: 20px">
        <!--      这里是表格的位置-->
        <!--      <n-data-table :columns="columns" :data="couponList" />-->
        <n-spin :show="showLoading">
          <n-data-table
              striped
              class="table"
              :columns="columns"
              :data="couponList"
              :pagination="false"
              :bordered="true"
              :scroll-x="1000"
          />
        </n-spin>

      </n-card>

      <DataTableSuffix
          v-model:data-size="dataSize"
          v-model:page-count="pageCount"
          v-model:animated="animated"
          :update-data="getAllCoupons"
      />

    </div>
  </transition>


  <n-modal
      :title="editType === 'create'?t(`${i18nPrefix}.modal.newCoupon`):t(`${i18nPrefix}.modal.editCoupon`) + ` #${formValue.coupon.id}`"
      v-model:show="showModal"
      preset="dialog"
      :positive-text="editType === 'create'?t(`${i18nPrefix}.modal.confirmAdd`):t(`${i18nPrefix}.modal.confirmEdit`)"
      :negative-text="t(`${i18nPrefix}.modal.cancel`)"
      style="width: 480px;"
      @positive-click="editType === 'create'?submitCoupon():updateACoupon()"
      @negative-click="closeModal"
      :show-icon="false"
  >
    <!--          pass-->
    <div style="margin-top: 30px"></div>

    <n-form
        ref="formRef"
        :model="formValue"
        :rules="rules"
        :show-feedback="true"
    >
      <n-form-item :label="t(`${i18nPrefix}.modal.name.title`)" path="coupon.name">
        <n-input v-model:value="formValue.coupon.name" :placeholder="t(`${i18nPrefix}.modal.name.placeholder`)"/>
      </n-form-item>
      <n-form-item :label="t(`${i18nPrefix}.modal.code.title`)" path="coupon.code">
        <n-input v-model:value="formValue.coupon.code" :placeholder="t(`${i18nPrefix}.modal.code.placeholder`)"/>
      </n-form-item>
      <n-form-item :label="t(`${i18nPrefix}.modal.percentOff.title`)" path="coupon.percent_off">
        <n-input-number
            :style="{width: '100%'}"
            v-model:value.number="formValue.coupon.percent_off"
            :placeholder="t(`${i18nPrefix}.modal.percentOff.placeholder`)"
            :max="100"
            :min="0"
        />
      </n-form-item>
      <n-form-item :label="t(`${i18nPrefix}.modal.activeRange.title`)" path="coupon.name">
        <n-date-picker v-model:value="range" type="datetimerange" style="width: 100%" clearable/>
      </n-form-item>
      <n-form-item :label="t(`${i18nPrefix}.modal.capacity.title`)" path="coupon.capacity">
        <n-input-number
            :style="{width: '100%'}"
            v-model:value.number="formValue.coupon.capacity"
            :placeholder="t(`${i18nPrefix}.modal.capacity.placeholder`)"
        />
      </n-form-item>
      <n-form-item :label="t(`${i18nPrefix}.modal.perUserLimit.title`)" path="coupon.per_user_limit">
        <n-input-number
            :style="{width: '100%'}"
            v-model:value.number="formValue.coupon.per_user_limit"
            :placeholder="t(`${i18nPrefix}.modal.perUserLimit.placeholder`)"
        />
      </n-form-item>
      <n-form-item :label="t(`${i18nPrefix}.modal.planLimit.title`)" path="coupon.per_user_limit">
        <n-select
            :options="plans"
            v-model:value.number="formValue.coupon.plan_limit"
            :placeholder="t(`${i18nPrefix}.modal.planLimit.placeholder`)"
        />
      </n-form-item>
    </n-form>
    <!--          pass-->

  </n-modal>
</template>

<style lang="less" scoped>
.root {
  padding: 0 20px 0 20px;

  .card-root {
    .add-btn {
      margin-top: 10px;
    }
  }
}
</style>