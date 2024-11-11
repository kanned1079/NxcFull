<script setup lang="ts">
import FormSuffix from "@/views/utils/FormSuffix.vue";
import {useI18n} from "vue-i18n";
import {onMounted, ref, h} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import type {FormInst} from 'naive-ui'
import {NButton, NButtonGroup, NSwitch, NTag, useMessage} from 'naive-ui'
import useNoticesStore from "@/stores/useNoticesStore";
import instance from "@/axios";
import {formatTimestamp} from "@/utils/timeFormat"

const {t} = useI18n();
const apiAddrStore = useApiAddrStore();

const themeStore = useThemeStore()
const noticesStore = useNoticesStore()
const message = useMessage()

let animated = ref<boolean>(false)

const formRef = ref<FormInst | null>(null)

let pageCount = ref(10)
let dataSize = ref<{pageSize: number, page: number}>({
  pageSize: 10,
  page: 1,
})

let dataCountOptions = [
  {
    label: '10条数据/页',
    value: 10,
  },
  {
    label: '20条数据/页',
    value: 20,
  },
  {
    label: '50条数据/页',
    value: 50,
  },
  {
    label: '100条数据/页',
    value: 100,
  },
]

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

let rules = {
  name: {
    required: true,
    message: '优惠券名是必须的',
    trigger: 'blur'
  }
}

// noticesStore.getAllNotices()

let showLoading = ref<boolean>(false)

// 显示表单
let showModal = ref(false)
let range = ref<[number, number]>([1183135260000, Date.now()])
let plans = ref<{label: string, value: number}[]>([])

// 订阅计划的键值
let getPlanKV = async () => {
  plans.value = []
  try {
    let {data} = await instance.get('http://localhost:8081/api/admin/v1/plan/kv')
    if (data.code === 200) {
      console.log(data)
      data.plans.forEach((item: {id: number, name: string}) =>  plans.value.push({
        label: item.name,
        value: item.id,
      }))
    }
  }catch (error) {
    console.log(error)
  }

}


// couponList 从后端获取的所有优惠券列表
// let couponList = ref<Coupon[]>([])
let couponList = ref<Coupon[]>([])

// 在此模版中用于渲染表格
// 表格的字段名为 优惠券ID，是否启用（此处设计为一个switch可用于单独关闭或启用该优惠券），优惠券码（使用n-tag），剩余使用次数，启用时间，过期时间，操作（包含编辑和删除功能两个按钮）

let getAllCoupons = async () => {
  showLoading.value = true
  // 这里是使用get方法 获取所有的优惠券列表
  // 请求地址 /admin/v1/coupon/fetch
  try {
    let {data} = await instance.get('http://localhost:8081/api/admin/v1/coupon', {
      params: {
        page: dataSize.value.page,
        size: dataSize.value.pageSize,
      }
    });
    if (data.code === 200) {
      couponList.value = []

      // couponList.value = data.coupons;
      data.coupon_list.forEach((item: Coupon, index: number) => {
        couponList.value.push(item)
      })
      pageCount.value = data.page_count
    } else {
      message.error('获取优惠券列表失败');
    }
    showLoading.value = false
  } catch (error) {
    console.error(error);
    message.error('请求优惠券列表失败');
  }
}

let activateACoupon = async (id: number, value: boolean) => {
  console.log(id, value)
  // 这里是使用post方法 当switch改变时调用来启用或关闭优惠券
  // 请求地址 /admin/v1/coupon/status/update
  // 请求体 {优惠券id， 状态}
  try {
    let {data} = await instance.put('http://localhost:8081/api/admin/v1/coupon/status', {
      id: id,
      status: value
    });
    if (data.code === 200) {
      message.success('优惠券状态更新成功');
    } else {
      message.error('更新优惠券状态失败');
    }
    await getAllCoupons();
  } catch (error) {
    console.error(error);
    message.error('请求更新优惠券状态失败');
  }
}

let updateACoupon = async (row: Coupon) => {
  // 这里是点击编辑一个优惠券后 使用post方法保存到服务器
  // 请求地址 /admin/v1/coupon/update
  // 请求体 {优惠券id， 该优惠券结构}
  try {
    let {data} = await instance.put('http://localhost:8081/api/admin/v1/coupon', row);
    if (data.code === 200) {
      message.success('优惠券更新成功');
      await getAllCoupons(); // 更新成功后刷新列表
    } else {
      message.error('更新优惠券失败');
    }
  } catch (error) {
    console.error(error);
    message.error('请求更新优惠券失败');
  }
}

let deleteACoupon = async (id: number) => {
  // 点击删除后 使用post方法 用于删除一个优惠券
  // 请求地址 /admin/v1/coupon/delete
  // 请求体 {优惠券id}
  console.log(id)
  try {
    let {data} = await instance.delete('http://localhost:8081/api/admin/v1/coupon', {
      params: {
        coupon_id: id,
      }
    });
    if (data.code === 200) {
      message.success('优惠券删除成功');
      await getAllCoupons(); // 删除成功后刷新列表
    } else {
      message.error('删除优惠券失败');
    }
  } catch (error) {
    console.error(error);
    message.error('请求删除优惠券失败');
  }
}

const columns = [
  {
    title: '优惠券ID',
    key: 'id'
  },
  {
    title: '是否启用',
    key: 'enabled',
    render(row: Coupon) {
      return h(NSwitch, {
        value: row.enabled,
        onUpdateValue: (value) => activateACoupon(row.id, value),
      });
    }
  },
  {
    title: '名称',
    key: 'name'
  },
  {
    title: '优惠券码',
    key: 'code',
    render(row: Coupon) {
      return h(NTag, {}, {default: () => row.code});
    }
  },
  {
    title: '剩余使用次数',
    key: 'residue',
  },
  {
    title: '启用时间',
    key: 'start_time',
    render(row: Coupon) {
      return h('p', {}, {default: () => formatTimestamp(row.start_time)});
    }
  },
  {
    title: '过期时间',
    key: 'end_time',
    render(row: Coupon) {
      return h('p', {}, {default: () => formatTimestamp(row.end_time)});
    }
  },
  {
    title: '操作',
    key: 'actions',
    render(row: Coupon) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          secondary: true,
          bordered: false,
          style: {marginLeft: '10px'},
          onClick: () => updateACoupon(row)
        }, {default: () => '编辑'}),
        h(NButton, {
          size: 'small',
          type: 'error',
          secondary: true,
          disabled: false,
          style: {marginLeft: '10px'},
          onClick: () => deleteACoupon(row.id)
        }, {default: () => '删除'})
      ]);
    },
    fixed: 'right',
  }
];


let handleAddCoupon = () => {
  getPlanKV()
  showModal.value = true
}

let submitCoupon = async () => {
  formValue.value.coupon.start_time = range.value[0]
  formValue.value.coupon.end_time = range.value[1]
  console.log(formValue.value)
  try {
    let {data} = await instance.post('http://localhost:8081/api/admin/v1/coupon', {
      ...formValue.value.coupon
    })
    if (data.code === 200) {
      message.success('添加优惠券成功')
    } else {
      console.log(data.msg)
    }
  } catch (error) {
    console.log(error)
  }
}

let closeModal = () => {
  showModal.value = false
}

onMounted(async () => {
  console.log('couponMgr挂载')
  await getAllCoupons()

  themeStore.contentPath = '/admin/dashboard/coupon';
  themeStore.menuSelected = 'coupon-mgr'

  animated.value = true

})

</script>

<script lang="ts">
export default {
  name: 'CouponMgr',
}
</script>

<template>
  <div style="padding: 20px 20px 0 20px">
    <n-card title="优惠券管理" hoverable :embedded="true" class="card-root" :bordered="false">
      <n-button type="primary" :bordered="false" class="add-btn" @click="handleAddCoupon">添加优惠券</n-button>
    </n-card>
  </div>

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

      <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">
        <n-pagination
            size="medium"
            v-model:page.number="dataSize.page"
            :page-count="pageCount"
            @update:page="getAllCoupons"
        />
        <n-select
            style="width: 160px; margin-left: 20px"
            v-model:value.number="dataSize.pageSize"
            size="small"
            :options="dataCountOptions"
            :remote="true"
            @update:value="dataSize.page = 1; getAllCoupons()"
        />
      </div>


    </div>
  </transition>



  <n-modal
      title="新建优惠券"
      v-model:show="showModal"
      preset="dialog"
      positive-text="确认添加"
      negative-text="算了"
      style="width: 480px;"
      @positive-click="submitCoupon"
      @negative-click="closeModal"
      :show-icon="false"
  >
    <!--          pass-->
    <div style="margin-top: 30px"></div>

    <n-form
        ref="formRef"
        :model="formValue"
        :rules="rules"
    >
      <n-form-item label="优惠券名称" path="coupon.name">
        <n-input v-model:value="formValue.coupon.name" placeholder="输入优惠券名称"/>
      </n-form-item>
      <n-form-item label="优惠券码" path="coupon.code">
        <n-input v-model:value="formValue.coupon.code" placeholder="自定义优惠券码（留空随机生成）"/>
      </n-form-item>
      <n-form-item label="优惠信息（百分比）" path="coupon.percent_off">
        <n-input-number
            :style="{width: '100%'}"
            v-model:value.number="formValue.coupon.percent_off"
            placeholder="输入优惠百分比"
        />
      </n-form-item>
      <n-form-item label="优惠券有效期" path="coupon.name">
        <n-date-picker v-model:value="range" type="daterange" style="width: 100%" clearable/>
      </n-form-item>
      <n-form-item label="最大使用次数" path="coupon.capacity">
        <n-input-number
            :style="{width: '100%'}"
            v-model:value.number="formValue.coupon.capacity"
            placeholder="限制最大使用限制，用完则无法使用（为空则不限制）"
        />
      </n-form-item>
      <n-form-item label="每个用户可使用次数" path="coupon.per_user_limit">
        <n-input-number
            :style="{width: '100%'}"
            v-model:value.number="formValue.coupon.per_user_limit"
            placeholder="限制每个用户可使用次数（为空则不限制）"
        />
      </n-form-item>
      <n-form-item label="指定订阅" path="coupon.per_user_limit">
        <n-select
            :options="plans"
            v-model:value.number="formValue.coupon.plan_limit"
            placeholder="限制指定订阅可使用优惠（为空则不限制）"
        />
      </n-form-item>
    </n-form>
    <!--          pass-->
    <template #footer>
      尾部
    </template>


  </n-modal>
</template>

<style lang="less" scoped>
.root {
  padding: 0 20px 20px 20px;

  .card-root {
    .add-btn {
      margin-top: 10px;
    }
  }
}
</style>