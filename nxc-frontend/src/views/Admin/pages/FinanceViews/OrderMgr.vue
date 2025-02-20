<script setup lang="ts">
import {computed, h, onBeforeMount, onMounted, ref, type ComputedRef} from "vue"
import {useI18n} from "vue-i18n";
import useThemeStore from "@/stores/useThemeStore";
import {
  CheckmarkDoneOutline as passOrderIcon,
  ChevronDownOutline as downIcon,
  PauseOutline as closeOrderIcon,
  RefreshOutline as refreshIcon,
  Search as searchIcon,
} from "@vicons/ionicons5"

import {type DataTableColumns, NButton, NDropdown, NIcon, NTag, useMessage} from "naive-ui";
import instance from "@/axios";
import {formatDate} from "@/utils/timeFormat";
import renderIcon from "@/utils/iconFormator";
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue";

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

interface PrivilegeGroup {
  id: number
  name: string
}

const {t} = useI18n()
const i18nPrefix = 'adminViews.orderMgr'
const themeStore = useThemeStore()
const message = useMessage()

let animated = ref<boolean>(false)

let showSearchModal = ref<boolean>(false)

let orderList = ref<Order[]>([])
let privilegeGroupList = ref<PrivilegeGroup[]>([])

let searchForm = ref<{
  email: string
  sort: 'ASC' | 'DESC'
}>({
  email: '',
  sort: 'DESC',   // 默认按照created_at降序排序
})


let pageCount = ref(10)

let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

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
function getStatusText(is_finished: boolean, is_success: boolean): ComputedRef<string> {
  if (!is_success && !is_finished) {
    return computed(() => t(`${i18nPrefix}.tradeWaiting`));
  } else if (is_finished && !is_success) {
    return  computed(() => t(`${i18nPrefix}.tradeFailure`));
  } else {
    return  computed(() => t(`${i18nPrefix}.tradeSuccess`));
  }
}

const i18nTablePrefix = 'adminViews.orderMgr.table'
const columns = computed<DataTableColumns<Order>>(() => [
  {
    title: t(`${i18nTablePrefix}.id`),
    key: 'id'
  },
  {
    title: t(`${i18nTablePrefix}.orderId`),
    key: 'order_id'
  },
  {
    title: t(`${i18nTablePrefix}.email`),
    key: 'email'
  },
  {
    title: t(`${i18nTablePrefix}.status.title`),
    key: 'status',
    render(row: Order) {
      return h('p', {}, {
        default: () => {
          switch (row.status) {
            case 0:
              return t(`${i18nTablePrefix}.status.t1`)
            case 1:
              return t(`${i18nTablePrefix}.status.t2`)
            case 2:
              return t(`${i18nTablePrefix}.status.t3`)
          }
        }
      })
    }
  },
  {
    title: t(`${i18nTablePrefix}.planName`),
    key: 'plan_name',
    render(row: Order) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: 'primary',
      }, {default: () => row.plan_name})
    }
  },
  {
    title: t(`${i18nTablePrefix}.period`),
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
              return t('period.month')
            case 'quarter':
              return t('period.quarter')
            case 'half-year':
              return t('period.halfYear')
            case 'year':
              return t('period.year')
          }
        }
      })
    }
  },
  {
    title: t(`${i18nTablePrefix}.group`),
    key: 'group_id',
    render(row: Order) {
      // 根据 group_id 找到对应的 name
      const groupName = privilegeGroupList.value.find((group: PrivilegeGroup) => group.id === row.group_id)?.name || '-';

      return h(NTag, {
        size: 'small',
        bordered: false,
        type: 'default',
      }, {
        default: () => groupName, // 显示对应的权限组名称
      });
    }
  },
  {
    title: t(`${i18nTablePrefix}.amount`),
    key: 'amount',
    render(row: Order) {
      return h('p', {}, {
        default: () => row.amount.toFixed(2)
      })
    }
  },
  {
    title: t(`${i18nTablePrefix}.price`),
    key: 'price',
    render(row: Order) {
      return h('p', {}, {
        default: () => row.price.toFixed(2)
      })
    }
  },
  {
    title: t(`${i18nTablePrefix}.isSuccess.title`),
    key: 'is_success',
    width: 100,
    render(row: Order) {
      return h('div', {
        style: {
          display: 'flex',
          alignItems: 'center',
          flexDirection: 'row'
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
        // 始终显示的状态文本
        h('p', {
          style: {
            margin: 0,
            fontSize: '14px',
            color: '#666'
          }
        }, getStatusText(row.is_finished, row.is_success).value),

        // 判断订单状态，决定是否显示操作按钮
        !row.is_finished && !row.is_success
            ? h(NDropdown, {
              options: [
                {label: t(`${i18nTablePrefix}.isSuccess.cancelOrder`), key: 'cancel', icon: renderIcon(closeOrderIcon)},
                {label: t(`${i18nTablePrefix}.isSuccess.passOrder`), key: 'approve', icon: renderIcon(passOrderIcon)},
              ],
              onSelect: (key) => {
                if (key === 'cancel') {
                  // console.log('取消订单', row.user_id, row.order_id);
                  handleManualCancelOrder(row.order_id, row.user_id);
                } else if (key === 'approve') {
                  // console.log('通过订单', row.user_id, row.order_id);
                  handleManualPassOrder(row.order_id, row.user_id);
                }
              },
              type: 'primary',
              style: {
                width: '140px',
              },
            }, {
              default: () => h(NButton, {
                size: 'small',
                type: 'default',
                text: true,
                style: {
                  display: 'flex',
                  alignItems: 'center',
                }
              }, [
                h(NIcon, {
                  size: 14,
                  style: {marginLeft: '4px'}
                }, {
                  default: () => h(downIcon) // 替换为实际的下拉图标组件
                })
              ])
            })
            : null
      ]);
    }
  },
  {
    title: t(`${i18nTablePrefix}.createdAt`),
    width: 160,
    key: 'created_at',
    render(row: Order) {
      return h('p', {}, {
        default: () => formatDate(row.created_at)
      })
    }
  },
  {
    title: t(`${i18nTablePrefix}.action.title`),
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
                message.info('查看订单详情' + row.id)
              },
            },
            {default: () => t(`${i18nTablePrefix}.action.showDetail`)}),
      ]);
    },
  }
])

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

let handleManualCancelOrder = async (orderId: string, userId: number) => {
  try {
    animated.value = false
    let {data} = await instance.delete('/api/admin/v1/order', {
      params: {
        user_id: userId,
        order_id: orderId
      }
    })
    if (data.code === 200) {
      if (data.cancelled) {
        message.success('order cancelled successfully')
        await getAllOrders()
        animated.value = true
      } else {
        message.warning('order cancelled failure.')
      }
    } else {
      console.log('request failure', data.msg)
    }
  } catch (err: any) {
    console.log(err + '')
  }
}

let handleManualPassOrder = async (orderId: string, userId: number) => {
  try {
    animated.value = false
    let {data} = await instance.put('/api/admin/v1/order', {
      user_id: userId,
      order_id: orderId
    })
    if (data.code === 200) {
      if (data.passed) {
        message.success('order passed successfully')
        await getAllOrders()
        animated.value = true
      } else {
        message.warning('order passed failure.')
      }
    } else {
      console.log('request failure', data.msg)
    }
  } catch (err: any) {
    console.log(err + '')
  }
}

let getAllGroups = async () => {
  try {
    let {data} = await instance.get('/api/admin/v1/groups/kv',)
    if (data.code === 200) {
      privilegeGroupList.value = []
      // data.group_list.forEach((group: PrivilegeGroup) => groupList.push(group))
      data.group_list.forEach((group: PrivilegeGroup) => privilegeGroupList.value.push(group))
    } else {
      message.error(data.msg || 'Error fetch.')
    }
  } catch (error) {
    console.log(error)
  }
}

onBeforeMount(() => {
  themeStore.menuSelected = 'order-manager'
  themeStore.breadcrumb = t('adminViews.orderMgr.title')
})

onMounted(async () => {
  await getAllGroups()
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

  <PageHead
      :title="t('adminViews.orderMgr.title')"
      :description="t('adminViews.orderMgr.description')"
  >
    <n-button
        tertiary
        type="primary"
        size="medium"
        class="btn-right"
        @click="showSearchModal=true">
      <template #icon>
        <n-icon>
          <searchIcon/>
        </n-icon>
      </template>

      {{ t(`${i18nPrefix}.search`) }}
    </n-button>
    <n-button
        tertiary
        type="primary"
        size="medium"
        class="btn-right"
        @click="searchForm.email=''; animated=false; getAllOrders()">
      <template #icon>
        <n-icon>
          <refreshIcon/>
        </n-icon>
      </template>

      {{ t(`${i18nPrefix}.resetSearch`) }}
    </n-button>
  </PageHead>

  <!--  <div class="root">-->
  <!--    <n-card hoverable :bordered="false" :embedded="true" class="card1" :title="'订单管理'">-->
  <!--      <n-button class="btn" secondary type="primary" size="medium" @click="showSearchModal=true">-->
  <!--        <n-icon size="14" style="padding-right: 8px">-->
  <!--          <searchIcon/>-->
  <!--        </n-icon>-->
  <!--        {{ t('adminViews.userMgr.query') }}-->
  <!--      </n-button>-->
  <!--      <n-button class="btn" tertiary type="primary" size="medium"-->
  <!--                @click="searchForm.email=''; animated=false; getAllOrders()">-->
  <!--        <n-icon size="14" style="padding-right: 8px">-->
  <!--          <refreshIcon/>-->
  <!--        </n-icon>-->
  <!--        {{ '重置搜索' }}-->
  <!--      </n-button>-->
  <!--    </n-card>-->
  <!--  </div>-->

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
            :scroll-x="1600"
        />
      </n-card>

      <DataTableSuffix
          v-model:data-size="dataSize"
          v-model:page-count="pageCount"
          v-model:animated="animated"
          :update-data="getAllOrders"
      />
    </div>
  </transition>


  <n-modal
      :title="t(`${i18nPrefix}.search`)"
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


    <n-form-item path="sort" :label="t(`${i18nPrefix}.searchModal.sort.title`)">
      <n-select
          :options="[{label: t(`${i18nPrefix}.searchModal.sort.ASC`), value: 'ASC'},{label: t(`${i18nPrefix}.searchModal.sort.DESC`), value: 'DESC'}]"
          :default-value="'DESC'"
          v-model:value="searchForm.sort"
      >
      </n-select>
    </n-form-item>

  </n-modal>

</template>

<style scoped lang="less">
.root {
  padding: 0 20px 0 20px;

  .card1 {
    .btn {
      margin-right: 10px;
    }
  }
}

.btn-right {
  margin-right: 10px;
}

</style>