<script setup lang="ts">
import {computed, h, onMounted, ref} from "vue"
import {useI18n} from "vue-i18n";
import {RefreshOutline as refreshIcon, Search as searchIcon} from '@vicons/ionicons5'
import {NButton, NTag, useMessage} from "naive-ui";
import instance from "@/axios";
import {formatDate} from "@/utils/timeFormat";

interface Order {
  id: number
  user_id: number
  order_id: string
  email: number
  plan_id: number
  group_id: number
  plan_name: string
  period: string
  coupon_id: number
  coupon_number: string
  status: number // 0新购 1续费 2修改订阅
  is_success: boolean
  is_finished: boolean
  price: number
  amount: number
  paid_at: string
  payment_method: string
  created_at: string
  failure_reason: string
}

const message = useMessage()

let animated = ref<boolean>(false)

let showSearchModal = ref<boolean>(false)

let orderList = ref<Order[]>([])

let searchForm = ref<{
  email: string
  sort: 'ASC' | 'DESC'
}>({
  email: '',
  sort: 'DESC',   // 默认按照created_at降序排序
})

const {t} = useI18n()

let pageCount = ref(10)

let dataSize = ref<{ pageSize: number, page: number }>({
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

// 获取状态颜色
function getStatusColor(is_finished: boolean, is_success: boolean): string {
  if (!is_success && !is_finished) {
    return '#e2d849'; // 未支付为黄色
  } else if (is_finished && !is_success) {
    return '#ee3f4d'; // 交易失败为红色
  } else {
    return '#86c166'; // 成功为绿色
  }
}

// 获取状态文本
function getStatusText(is_finished: boolean, is_success: boolean): string {
  if (!is_success && !is_finished) {
    return '未支付';
  } else if (is_finished && !is_success) {
    return '交易失败';
  } else {
    return '成功';
  }
}

const columns = [
  {
    title: '#',
    key: 'id'
  },
  {
    title: '用户邮箱',
    key: 'email'
  },
  {
    title: '订单号',
    key: 'order_id'
  },
  {
    title: '类型',
    key: 'status',
    render(row: Order) {
      return h('p', {}, {
        default: () => {
          switch (row.status) {
            case 0:
              return '新购'
            case 1:
              return '续费'
            case 2:
              return '修改'
          }
        }
      })
    }
  },
  {
    title: '计划名称',
    key: 'plan_name',
    render(row: Order) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: 'default',
      }, {default: () => row.plan_name})
    }
  },
  {
    title: '周期',
    key: 'period',
    render(row: Order) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: 'default',
      }, {
        default: () => {
          switch (row.period) {
            case 'month':
              return '月付'
            case 'quarter':
              return '季付'
            case 'half-year':
              return '半月付'
            case 'year':
              return '年付'
          }
        }
      })
    }
  },
  {
    title: '权限组',
    key: 'group_id',
    render(row: Order) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: 'default',
      }, {
        default: () => {
          return row.group_id.toString() + ' pass'
        }
      });
    }
  },
  {
    title: '实付金额',
    key: 'amount',
    render(row: Order) {
      return h('p', {}, {
        default: () => row.amount.toFixed(2)
      })
    }
  },
  {
    title: '原始价格',
    key: 'price',
    render(row: Order) {
      return h('p', {}, {
        default: () => row.price.toFixed(2)
      })
    }
  },
  // {
  //   title: '订单状态', key: 'is_success', render(row: OrderList) {
  //     return h(NTag, {
  //       // type: row.is_success ? 'success' : 'error',
  //       type: orderStatusTagColor(row.is_finished, row.is_success) as string,
  //       bordered: false,
  //     }, {default: () => orderStatusText(row.is_finished, row.is_success)});
  //   }
  // },
  {
    title: '订单状态',
    key: 'is_success',
    render(row: Order) {
      return h('div', {
        style: {
          display: 'flex',
          alignItems: 'center'
        }
      }, [
        h('div', {
          style: {
            width: '8px',
            height: '8px',
            borderRadius: '50%',
            marginRight: '8px',
            backgroundColor: getStatusColor(row.is_finished, row.is_success)
          }
        }),
        h('span', {}, getStatusText(row.is_finished, row.is_success))
      ]);
    }
  },
  {
    title: '创建时间',
    width: 160,
    key: 'created_at',
    render(row: Order) {
      return h('p', {}, {
        default: () => formatDate(row.created_at)
      })
    }
  },


  //
  // {
  //   title: computed(() => t('adminViews.userMgr.accountStatus')).value,
  //   key: 'status',
  //   render(row: User) {
  //     return h(NTag, {
  //       size: 'small',
  //       bordered: false,
  //       type: row.status ? 'success' : 'warning',
  //     }, {default: () => row.status ? computed(() => t('adminViews.userMgr.normal')).value : computed(() => t('adminViews.userMgr.banned')).value});
  //   }
  // },
  // {
  //   title: computed(() => t('adminViews.userMgr.inviteCode')).value,
  //   key: 'invite_code',
  //   render(row: User) {
  //     return h(NTag, {
  //       size: 'small',
  //       bordered: false,
  //       type: row.invite_code ? 'primary' : 'default',
  //     }, {default: () => row.invite_code ? row.invite_code : computed(() => t('adminViews.userMgr.nullContent')).value});
  //   }
  // },
  // {
  //   title: computed(() => t('adminViews.userMgr.balance')).value,
  //   key: 'balance',
  //   render(row: User) {
  //     return h(
  //         'p',
  //         {
  //           style: {
  //             color: row.balance < 0.00 ? '#d16666' : null,
  //             textDecoration: row.balance < 0.00 ? 'underline' : null  // 如果余额小于0，添加下划线
  //           }
  //         },
  //         {
  //           default: () => appInfoStore.appCommonConfig.currency_symbol + ' ' + row.balance?.toFixed(2).toString()
  //         }
  //     );
  //   }
  // },
  // {
  //   title: computed(() => t('adminViews.userMgr.orderCount')).value,
  //   key: 'order_count',
  //   render(row: User) {
  //     return h(NTag, {size: 'small', bordered: false, type: 'default'}, {default: () => row.order_count});
  //   }
  // },
  // {
  //   title: computed(() => t('adminViews.userMgr.planCount')).value,
  //   key: 'plan_count',
  //   render(row: User) {
  //     return h(NTag, {size: 'small', bordered: false, type: 'default'}, {default: () => row.plan_count});
  //   }
  // },
  {
    title: computed(() => t('adminViews.userMgr.actions')).value,
    fixed: 'right',
    key: 'actions', render(row: Order) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
              size: 'small',
              type: 'primary',
              secondary: true,
              bordered: false,
              onClick: () => {
                // 编辑用户
                // Object.assign(editUser.value, row)
                // showEditUserDrawer.value = true
                message.info('查看订单详情', row.id as number)
              },
            },
            {default: () => '查看详情'}),
        // h(NButton, {
        //   size: 'small',
        //   type: row.status ? 'error' : 'warning',
        //   secondary: true,
        //   disabled: false,
        //   style: {marginLeft: '10px'},
        //   onClick: () => handleBlock2UnblockUserById(row)
        // }, {default: () => row.status ? computed(() => t('adminViews.userMgr.banUser')).value : computed(() => t('adminViews.userMgr.unbanUser')).value})
      ]);
    },
  }
];

let getAllOrders = async () => {
  try {
    let {data} = await instance.get('/api/admin/v1/order', {
      params: {
        page: dataSize.value.page,
        size: dataSize.value.pageSize,
        search_email: searchForm.value.email.trim(),
        sort: searchForm.value.sort
      }
    })
    if (data.code === 200) {
      console.log(data)
      orderList.value = []
      data.orders.forEach((order: Order) => orderList.value.push(order))
      pageCount.value = data.page_count
      animated.value = true
    } else {
      message.error('未知错误' + data.msg || '')
    }
  } catch (err: any) {
    message.error(err + '')
  }
}


onMounted(async () => {

  await getAllOrders()
  animated.value = true
})


</script>

<script lang="ts">
export default {
  name: 'OrderMgr',
}
</script>

<template>
  <div class="root">
    <n-card hoverable :bordered="false" :embedded="true" class="card1" :title="'订单管理'">
      <n-button class="btn" secondary type="primary" size="medium" @click="showSearchModal=true">
        <n-icon size="14" style="padding-right: 8px">
          <searchIcon/>
        </n-icon>
        {{ t('adminViews.userMgr.query') }}
      </n-button>
      <n-button class="btn" tertiary type="primary" size="medium"
                @click="searchForm.email=''; animated=false; getAllOrders()">
        <n-icon size="14" style="padding-right: 8px">
          <refreshIcon/>
        </n-icon>
        {{ '重置搜索' }}
      </n-button>
    </n-card>
  </div>

  <transition name="slide-fade">
    <div style="padding: 20px" v-if="animated">
      <n-card :embedded="true" hoverable :bordered="false" content-style="padding: 0px">
        <n-data-table
            striped
            class="table"
            :columns="columns"
            :data="orderList"
            :pagination="false"
            :bordered="true"
            style=""
            :scroll-x="1000"
        />
      </n-card>
      <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">
        <n-pagination
            size="medium"
            v-model:page.number="dataSize.page"
            :page-count="pageCount"
            @update:page="animated=false; getAllOrders()"
        />
        <n-select
            style="width: 160px; margin-left: 20px"
            v-model:value.number="dataSize.pageSize"
            size="small"
            :options="dataCountOptions"
            :remote="true"
            @update:value="animated=false; dataSize.page = 1; getAllOrders()"
        />
      </div>
    </div>
  </transition>


  <n-modal
      :title="'搜索'"
      v-model:show="showSearchModal"
      preset="dialog"
      :positive-text="t('adminViews.userMgr.query')"
      :negative-text="t('adminViews.userMgr.cancel')"
      style="width: 480px"
      @positive-click="animated=false; getAllOrders()"
      @negative-click="showSearchModal=false;"
      @before-leave="searchForm.email=''"
      :show-icon="false"
  >
    <div style="margin-top: 30px"></div>
    <n-form-item path="email" :label="t('adminViews.userMgr.userEmail')">
      <n-input
          autofocus
          @keyup.enter="showSearchModal=false;animated=false; getAllOrders()"
          :placeholder="t('adminViews.userMgr.inputUserEmail')"
          v-model:value="searchForm.email"
      ></n-input>
    </n-form-item>


    <n-form-item path="sort" :label="'排序算法'">
      <n-select
          :options="[{label: '升序', value: 'ASC'},{label: '降序', value: 'DESC'}]"
          :default-value="'ASC'"
      >
      </n-select>
    </n-form-item>


  </n-modal>

</template>

<style scoped lang="less">
.root {
  padding: 20px 20px 0 20px;

  .card1 {
    .btn {
      margin-right: 10px;
    }
  }
}

</style>