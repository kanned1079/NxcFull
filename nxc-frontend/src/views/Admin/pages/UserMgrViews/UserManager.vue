<script setup lang="ts">
import {NButton, NTag, useMessage} from 'naive-ui'
import {h, onMounted, ref} from 'vue'
// PersonAddOutline
import {PersonAddOutline as addUserIcon, RefreshOutline as refreshIcon, Search as searchIcon,} from '@vicons/ionicons5'
import useThemeStore from "@/stores/useThemeStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import instance from "@/axios";
import {formatDate} from "@/utils/timeFormat";

const appInfoStore = useAppInfosStore()
const themeStore = useThemeStore();
const message = useMessage()

interface User {
  id?: number
  invite_user_id?: number
  email?: string
  is_admin?: boolean
  is_staff?: boolean
  status?: boolean
  privilege_group?: string
  order_count?: number
  plan_count?: number
  balance?: number
  created_at?: string
}

let showSearchNModal = ref<boolean>(false)
let showAddUserModal = ref<boolean>(false)

let users = ref<User[]>([])

let editUser = ref<User>({})

// 定义表格的列
const columns = [
  {
    title: '#',
    key: 'id'
  },
  // {
  //   title: '被邀请码',
  //   key: 'invite_user_id',
  //   render(row: User) {
  //     return h('p', null, {default: () => row.invite_user_id ? row.invite_user_id : '-'})
  //   }
  // },
  {
    title: '邮箱',
    key: 'email'
  },
  {
    title: '注册日期',
    key: 'created_at',
    render(row: User) {
      return h('p', {}, {
        default: () => formatDate(row.created_at)
      })
    }
  },
  {
    title: '管理员',
    key: 'is_admin',
    render(row: User) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.is_admin ? 'success' : 'default',
      }, {default: () => row.is_admin ? 'YES' : 'NO'});
    }
  },
  {
    title: '员工',
    key: 'is_staff',
    render(row: User) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.is_staff ? 'success' : 'default',
      }, {default: () => row.is_staff ? 'YES' : 'NO'});
    }
  },
  {
    title: '帐户状态',
    key: 'status',
    render(row: User) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.status ? 'success' : 'warning',
      }, {default: () => row.status ? '正常' : '封禁'});
    }
  },
  // {
  //   title: '权限组',
  //   key: 'privilege_group',
  //   render(row: User) {
  //     return h(NTag, {size: 'small', bordered: false, type: 'default'}, {default: () => row.privilege_group});
  //   }
  // },
  {
    title: '余额',
    key: 'balance',
    render(row: User) {
      return h(
          'p',
          {
            style: {
              color: row.balance < 0.00 ? '#d16666' : null,
              textDecoration: row.balance < 0.00 ? 'underline' : null  // 如果余额小于0，添加下划线
            }
          },
          {
            default: () => appInfoStore.appCommonConfig.currency_symbol + ' ' + row.balance?.toFixed(2).toString()
          }
      );
    }
  },
  {
    title: '订单数量',
    key: 'order_count',
    render(row: User) {
      return h(NTag, {size: 'small', bordered: false, type: 'default'}, {default: () => row.order_count});
    }
  },
  {
    title: '计划数',
    key: 'plan_count',
    render(row: User) {
      return h(NTag, {size: 'small', bordered: false, type: 'default'}, {default: () => row.plan_count});
    }
  },
  {
    title: '操作', key: 'actions', render(row: User) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          secondary: true,
          bordered: false,
          onClick: () => function () {
          }
        }, {default: () => '编辑资料'}),
        h(NButton, {
          size: 'small',
          type: row.status ? 'error' : 'warning',
          secondary: true,
          disabled: false,
          style: {marginLeft: '10px'},
          onClick: () => handleBlock2UnblockUserById(row)
        }, {default: () => row.status ? '封禁用户' : '解封用户'})
      ]);
    },
    width: 200,
    fixed: 'right'
  }
];

let pageCount = ref(10)

let animated = ref<boolean>(false)

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

let searchEmail = ref<string>('')

let getAllUsers = async () => {
  try {
    let {data} = await instance.get('/api/admin/v1/users', {
      params: {
        page: dataSize.value.page,
        size: dataSize.value.pageSize,
        email: searchEmail.value.trim(),
      }
    })
    if (data.code === 200) {
      users.value = []
      pageCount.value = data.page_count
      console.log(data)
      if (data.users)
        data.users.forEach((user: User) => users.value.push(user))
      else
        message.warning('无记录')
      animated.value = true
    }
  } catch (err: any) {
    message.error(err + '')
  }
}

let handleBlock2UnblockUserById = async (row: User) => {
  try {
    // animated.value = false
    let {data} = await instance.patch('/api/admin/v1/users', {
      user_id: row.id
    })
    if (data.code === 200) {
      message.success('操作成功')
      await getAllUsers()
    } else {
      message.error('未知错误' + data.msg || '')
    }
  } catch (err: any) {
    message.error(err + '')
  }
}

let handleAddNewUser = async () => {
  try {
    let {data} = await instance.post('/api/admin/v1/user', {})
  } catch (err: any) {
    message.error(err + '')
  }
}

let editUsrInfo = async () => {
  try {
    let {data} = await instance.put()
  } catch (err: any) {
    message.error(err + '')
  }
}

onMounted(async () => {
  themeStore.menuSelected = 'user-manager'
  themeStore.contentPath = '/admin/dashboard/usermanager'
  await getAllUsers()
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'UserManager'
}
</script>

<template>
  <div class="root">
    <n-card hoverable :bordered="false" :embedded="true" class="card1" title="用户管理">
      <n-button class="btn" tertiary type="primary" size="medium" @click="showSearchNModal=true">
        <n-icon size="14" style="padding-right: 8px">
          <searchIcon/>
        </n-icon>
        查询
      </n-button>
      <n-button class="btn" tertiary type="primary" size="medium"
                @click="searchEmail=''; animated=false; getAllUsers()">
        <n-icon size="14" style="padding-right: 8px">
          <refreshIcon/>
        </n-icon>
        重置
      </n-button>
      <n-button class="btn" tertiary type="primary" size="medium" @click="showAddUserModal=true">
        <n-icon size="14" style="padding-right: 8px">
          <addUserIcon/>
        </n-icon>
        新建用户
      </n-button>
    </n-card>
  </div>

  <transition name="slide-fade">
    <div style="padding: 20px" v-if="animated">
      <n-card class="form-root-card" hoverable :embedded="true" :bordered="false" content-style="padding: 0px">
        <n-data-table
            striped
            size="large"
            :bordered="false"
            :columns="columns"
            :data="users"
        />
      </n-card>

      <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">
        <n-pagination
            size="medium"
            v-model:page.number="dataSize.page"
            :page-count="pageCount"
            @update:page="animated=false; getAllUsers()"
        />
        <n-select
            style="width: 160px; margin-left: 20px"
            v-model:value.number="dataSize.pageSize"
            size="small"
            :options="dataCountOptions"
            :remote="true"
            @update:value="animated=false; dataSize.page = 1; getAllUsers()"
        />
      </div>
    </div>

  </transition>

  <n-modal
      title="搜索用户"
      v-model:show="showSearchNModal"
      preset="dialog"
      :positive-text="'查询'"
      negative-text="算了"
      style="width: 480px"
      @positive-click="animated=false; getAllUsers()"
      @negative-click="showSearchNModal=false;"
      @before-leave="searchEmail=''"
      :show-icon="false"
  >
    <div style="margin-top: 30px"></div>
    <n-form-item path="email" label="用户邮箱" autofocus>
      <n-input
          autofocus
          @keyup.enter="showSearchNModal=false;animated=false; getAllUsers()"
          placeholder="输入用户邮箱（模糊搜索）"
          v-model:value="searchEmail"
      />
    </n-form-item>
  </n-modal>

  <n-modal
      title="新建用户"
      v-model:show="showAddUserModal"
      preset="dialog"
      :positive-text="'查询'"
      negative-text="算了"
      style="width: 480px"
      @positive-click="animated=false; getAllUsers()"
      @negative-click="showSearchNModal=false;"
      @before-leave="searchEmail=''"
      :show-icon="false"
  >
    <div style="margin-top: 30px"></div>
    <n-form-item path="email" label="邮箱" autofocus>
      <n-input
          autofocus
          @keyup.enter="showSearchNModal=false;animated=false; getAllUsers()"
          placeholder="输入用户邮箱（模糊搜索）"
          v-model:value="searchEmail"
      />
    </n-form-item>
    <n-form-item path="email" label="密码" autofocus>
      <n-input
          type="password"
          autofocus
          @keyup.enter="showSearchNModal=false;animated=false; getAllUsers()"
          placeholder="输入用户邮箱（模糊搜索）"
          v-model:value="searchEmail"
      />
    </n-form-item>
  </n-modal>

</template>

<style lang="less" scoped>
.root {
  padding: 20px 20px 0 20px;

  .card1 {
    .btn {
      margin-right: 10px;
    }
  }
}


</style>