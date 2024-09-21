<script setup lang="ts" name="CouponMgr">
import {onMounted, ref} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import type {FormInst} from 'naive-ui'
import {NButton, useMessage} from 'naive-ui'
import useNoticesStore from "@/stores/useNoticesStore";
import instance from "@/axios";

const apiAddrStore = useApiAddrStore();

const themeStore = useThemeStore()
const noticesStore = useNoticesStore()
const message = useMessage()

const formRef = ref<FormInst | null>(null)

interface Coupon {
  id: number
  name: string
  code?: string
  percent_off: number
  start_time: number
  end_time: number
  per_user_limit?: number
  capacity: number
  plan_limit?: number
  created_at: string
  updated_at: string
  deleted_at?: string
}

// interface formData {
//   title: string
//   content: string
//   tags: string
//   img_url?: string
// }

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
  } as Coupon
})

let rules = {
  name: {
    required: true,
    message: '优惠券名是必须的',
    trigger: 'blur'
  }
}

// noticesStore.getAllNotices()

// 显示表单
let showModal = ref(false)
let range = ref<[number, number]>([1183135260000, Date.now()])
let plans = ref([
  {
    label: 'S1订阅',
    value: 0,
  },
  {
    label: 'S2订阅',
    value: 1,
  },
  {
    label: 'S1订阅',
    value: 2,
  }
])

let handleAddCoupon = () => {
  showModal.value = true
}

let submitCoupon = async () => {
  formValue.value.coupon.start_time = range.value[0]
  formValue.value.coupon.end_time = range.value[1]
  console.log(formValue.value)
  try {
    let {data} = instance.post(apiAddrStore.apiAddr.admin.addNewCoupon, {
      name: formValue.value.coupon.name,
      code: formValue.value.coupon.code,
      percent_off: formValue.value.coupon.percent_off,
      start_time: formValue.value.coupon.start_time,
      end_time: formValue.value.coupon.end_time,
      capacity: formValue.value.coupon.capacity,
      per_user_limit: formValue.value.coupon.per_user_limit,
      plan_limit: formValue.value.coupon.plan_limit,
    })
    if (data.code === 200) {
      message.success('添加优惠券成功')
    } else {
      console.log(data.msg)
    }
  }catch (error) {
    console.log(error)
  }
}

let closeModal = () => {
  showModal.value = false
}

onMounted(() => {
  console.log('couponMgr挂载')

  themeStore.contentPath = '/admin/dashboard/publicNotice';
  themeStore.menuSelected = 'publicNotice-mgr'

})

</script>

<template>
  <div class="root">
    <n-card title="优惠券管理" hoverable :embedded="true" class="card-root">
      <n-button class="add-btn" @click="handleAddCoupon">添加优惠券</n-button>
    </n-card>

    <n-modal
        title="新建优惠券"
        v-model:show="showModal"
        preset="dialog"
        positive-text="确认添加"
        negative-text="算了"
        style="width: 480px"
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
          <n-date-picker v-model:value="range" type="daterange" clearable/>
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
  </div>
</template>

<style lang="less" scoped>
.root {
  padding: 20px;

  .card-root {
    .add-btn {
      margin-top: 10px;
    }
  }
}
</style>