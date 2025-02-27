<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {type FormInst, NButton, NIcon, NTag, useMessage, type FormRules, type DataTableColumns, useDialog} from 'naive-ui'
import {computed, h, onMounted, onBeforeMount, ref} from 'vue'
import {
  PeopleOutline as usersIcon,
  PersonAddOutline as addUserIcon,
  RefreshOutline as refreshIcon,
  Search as searchIcon
} from '@vicons/ionicons5'
import useThemeStore from "@/stores/useThemeStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
// import instance from "@/axios";
import useTablePagination from "@/hooks/useTablePagination";
import {formatDate} from "@/utils/timeFormat";
import {hashPassword} from "@/utils/encryptor";
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue";
import {
  handleAddNewUserReq,
  handleBlock2UnblockUserById,
  handleEditUserInfoReq,
  handleGetAllUsers
} from "@/api/admin/user";
import {handleDeleteGroup} from "@/api/admin/groups";

const {t} = useI18n()
const appInfoStore = useAppInfosStore()
const themeStore = useThemeStore();
const message = useMessage()
const dialog = useDialog()

interface User {
  id: number
  invite_user_id?: number
  invite_code?: string
  name?: string
  email?: string
  password?: string
  is_admin: boolean
  is_staff?: boolean
  status?: boolean
  privilege_group?: string
  order_count?: number
  plan_count?: number
  balance?: number
  created_at?: string
  updated_at?: string
  deleted_at?: string
}

let showSearchNModal = ref<boolean>(false)
let showAddUserModal = ref<boolean>(false)
let showEditUserDrawer = ref<boolean>(false)

const addUserForm = ref<FormInst | null>(null)
const editUserForm = ref<FormInst | null>(null);
const addRules = computed<FormRules>(() => {
  return {
    email: [
      {required: true, message: t('adminViews.userMgr.enterEmail'), trigger: 'blur'},
      {type: 'email', message: t('adminViews.userMgr.enterValidEmail'), trigger: 'blur'},
    ],
    password: [
      {required: true, message: t('adminViews.userMgr.enterPassword'), trigger: 'blur'},
      {min: 6, message: t('adminViews.userMgr.passwordMinLength'), trigger: 'blur'},
      {max: 20, message: t('adminViews.userMgr.passwordMaxLength'), trigger: 'blur'},
      {
        pattern: /^(?=.*[a-zA-Z])(?=.*\d)(?=.*[!@#$%^&*(),.?":{}|<>])/, // 至少包含一个字母、数字和特殊字符
        message: t('adminViews.userMgr.passwordStrength'),
        trigger: 'blur'
      }
    ]
  };
})

const editRules = {
  email: [
    {required: true, message: computed(() => t('adminViews.userMgr.enterEmail')).value, trigger: 'blur'},
    {type: 'email', message: computed(() => t('adminViews.userMgr.enterValidEmail')).value, trigger: 'blur'},
  ],
};

const submitForm = async () => {
  try {
    // 等待表单验证结果
    const valid = await editUserForm.value?.validate();
    if (valid) {
      console.log('提交表单:', editUser);
      showEditUserDrawer.value = false;
      message.success(t('adminViews.userMgr.validationSuccess'));
      await handleEditUserInfo()
    }
  } catch (error) {
    message.error(t('adminViews.userMgr.validationFailed'));
  }
};

const submitAddForm = async () => {
  try {
    const valid = await addUserForm.value?.validate();
    if (valid) {
      await handleAddNewUser()
    }
  } catch (error) {
    showAddUserModal.value = true;
    message.error(t('adminViews.userMgr.validationFailed'));
  }
};

let cancelEdit = () => {
  showEditUserDrawer.value = false;
}

let users = ref<User[]>([])

let newUser = ref<{
  email: string,
  password: string
}>({
  email: '',
  password: ''
})

let editUser = ref<User>({
  id: 0,
  invite_user_id: 0,
  invite_code: '',
  email: '',
  password: '',
  is_admin: false,
  is_staff: false,
  status: false,
  privilege_group: '',
  order_count: 0,
  plan_count: 0,
  balance: 0,
  created_at: '',
})

// 处理 is_admin 的计算属性
const isAdminString = computed({
  get: () => editUser.value.is_admin ? 'true' : 'false',
  set: (value) => {
    editUser.value.is_admin = value === 'true';
  },
});

// 处理 is_staff 的计算属性
const isStaffString = computed({
  get: () => editUser.value.is_staff ? 'true' : 'false',
  set: (value) => {
    editUser.value.is_staff = value === 'true';
  },
});

// 定义表格的列

const columns = computed<DataTableColumns<User>>(() => [
  {
    title: t('adminViews.userMgr.email'),
    key: 'email',
    render(row: User) {
      return h('p', {
        style: row.deleted_at ? {
          opacity: 1,
          color: '#787d7b',
        } : null
      }, row.email);
    }
  },
  {
    title: t('adminViews.userMgr.registerDate'),
    key: 'created_at',
    render(row: User) {
      return h('p', {}, {
        default: () => formatDate(row.created_at as string),
      });
    }
  },
  {
    title: t('adminViews.userMgr.isAdmin'),
    key: 'is_admin',
    render(row: User) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.is_admin ? 'success' : 'default',
      }, {default: () => row.is_admin ? t('adminViews.userMgr.yes') : t('adminViews.userMgr.no')});
    }
  },
  {
    title: t('adminViews.userMgr.isStaff'),
    key: 'is_staff',
    render(row: User) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.is_staff ? 'success' : 'default',
      }, {default: () => row.is_staff ? t('adminViews.userMgr.yes') : t('adminViews.userMgr.no')});
    }
  },
  {
    title: t('adminViews.userMgr.accountStatus'),
    key: 'status',
    render(row: User) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.deleted_at ? 'error' : (row.status ? 'success' : 'warning'),
      }, {
        default: () => row.deleted_at ? t('adminViews.userMgr.deleted') : (row.status ? t('adminViews.userMgr.normal') : t('adminViews.userMgr.banned'))
      });
    }
  },
  {
    title: t('adminViews.userMgr.inviteCode'),
    key: 'invite_code',
    render(row: User) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.invite_code ? 'primary' : 'default',
      }, {default: () => row.invite_code || t('adminViews.userMgr.nullContent')});
    }
  },
  {
    title: t('adminViews.userMgr.balance'),
    key: 'balance',
    render(row: User) {
      return h(
          'p',
          {
            style: {
              color: row.balance !== undefined && row.balance < 0.00 ? '#d16666' : null,
              textDecoration: row.balance !== undefined && row.balance < 0.00 ? 'underline' : null
            },
          },
          {default: () => appInfoStore.appCommonConfig.currency_symbol + ' ' + row.balance?.toFixed(2).toString()}
      );
    }
  },
  {
    title: t('adminViews.userMgr.orderCount'),
    key: 'order_count',
    render(row: User) {
      return h(NTag, {size: 'small', bordered: false, type: 'default'}, {default: () => row.order_count});
    }
  },
  {
    title: t('adminViews.userMgr.planCount'),
    key: 'plan_count',
    render(row: User) {
      return h(NTag, {size: 'small', bordered: false, type: 'default'}, {default: () => row.plan_count});
    }
  },
  {
    title: t('adminViews.userMgr.actions'),
    fixed: 'right',
    key: 'actions',
    render(row: User) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          secondary: true,
          disabled: !!row.deleted_at,
          bordered: false,
          onClick: () => {
            // 编辑用户
            Object.assign(editUser.value, row);
            showEditUserDrawer.value = true;
          },
        }, {default: () => t('adminViews.userMgr.editProfile')}),
        h(NButton, {
          size: 'small',
          type: row.status ? 'error' : 'warning',
          secondary: true,
          disabled: Boolean(row.deleted_at) || Boolean(row.is_admin),
          style: {marginLeft: '10px'},
          onClick: () => block2UnblockUserById(row)
        }, {default: () => row.status ? t('adminViews.userMgr.banUser') : t('adminViews.userMgr.unbanUser')})
      ]);
    },
  }
]);

let animated = ref<boolean>(false)
let [dataSize, pageCount] = useTablePagination()

let searchEmail = ref<string>('')

let getAllUsers = async () => {
  let data = await handleGetAllUsers(dataSize.value.page, dataSize.value.pageSize, searchEmail.value.trim())
  if (data && data.code === 200) {
    users.value = []
    if (data.users) {
      data.users.forEach((user: User) => users.value.push(user))
      pageCount.value = data.page_count || 0
    } else {
      message.error(t('adminViews.common.fetchDataFailure') + data.msg || '')
    }
    animated.value = true
  }
}

let block2UnblockUserById = async (row: User) => {
  const runDelTask = async () => {
    let data = await handleBlock2UnblockUserById(row.id)
    if (data && data.code === 200) {
      t('adminViews.common.success')
      await getAllUsers()
    } else {
      message.error(t('adminViews.common.failure') + data.msg || '')
    }
  }
  dialog.error({
    title: t('adminViews.userMgr.mention.title'),
    content: () => {
      return h('div', {}, [
        h('p', {style: {fontWeight: 'bold', fontSize: '1rem', opacity: '0.8'}}, row.name?row.name:row.email),
        h('p', {style: {marginTop: '4px'}}, t('adminViews.userMgr.mention.content', {appName: appInfoStore.appCommonConfig.app_name}))
      ])
    },
    positiveText: t('adminViews.common.confirm'),
    negativeText: t('adminViews.common.cancel'),
    showIcon: false,
    actionStyle: {
      marginTop: '20px',
    },
    contentStyle: {
      marginTop: '20px',
    },
    onPositiveClick: async () => {
      animated.value = false
      await runDelTask()
    }
  })



}

let handleAddNewUser = async () => {
  animated.value = false
  let hashedPwd = await hashPassword(newUser.value.password)
  let data = await handleAddNewUserReq(newUser.value.email.trim(), hashedPwd)
  if (data && data.code === 200) {
    message.success(t('adminViews.common.success'))
    await getAllUsers()
  } else {
    message.error(t('adminViews.common.failure') + data.msg || '')
  }
}

let handleEditUserInfo = async () => {
  animated.value = false
  editUser.value.password = editUser.value.password?.trim() ? await hashPassword(editUser.value.password) : '';
  let data = await handleEditUserInfoReq(
      editUser.value.id,
      editUser.value.password,
      editUser.value.invite_code || '',
      editUser.value.is_admin,
      editUser.value.is_staff || false,
      editUser.value.balance || 0.00,
  )
  if (data.code === 200) {
    message.success(t('adminViews.common.success'))
    await getAllUsers()
  } else {
    message.error(t('adminViews.common.failure') + data.msg || '')
  }
}

onBeforeMount(() => {
  themeStore.breadcrumb = 'adminViews.userMgr.userManager'
  themeStore.menuSelected = 'user-manager'
})

onMounted(async () => {
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

  <PageHead
      :title="t('adminViews.userMgr.userManager')"
      :description="t('adminViews.userMgr.description')"
  >
    <template #default>
      <n-button
          tertiary
          type="primary"
          size="medium"
          class="btn-right"
          @click="showSearchNModal=true"
      >
        {{ t('adminViews.userMgr.query') }}
        <template #icon>
          <n-icon>
            <searchIcon/>
          </n-icon>
        </template>
      </n-button>

      <n-button
          tertiary
          type="primary"
          size="medium"
          class="btn-right"
          @click="searchEmail=''; animated=false; getAllUsers()">
        <template #icon>
          <n-icon>
            <refreshIcon/>
          </n-icon>
        </template>

        {{ t('adminViews.userMgr.reset') }}
      </n-button>

      <n-button
          tertiary
          type="primary"
          size="medium"
          class="btn-right"
          @click="showAddUserModal=true">

        <template #icon>
          <n-icon>
            <addUserIcon/>
          </n-icon>
        </template>
        {{ t('adminViews.userMgr.addNewUser') }}
      </n-button>
    </template>
  </PageHead>
  <transition name="slide-fade">
    <div style="padding: 20px" v-if="animated">
      <n-card class="form-root-card" hoverable :embedded="true" :bordered="false" content-style="padding: 0px">
        <n-data-table
            striped
            size="large"
            :bordered="true"
            :columns="columns"
            :data="users"
            :scroll-x="1200"
        />
      </n-card>


      <DataTableSuffix
          v-model:data-size="dataSize"
          v-model:page-count="pageCount"
          v-model:animated="animated"
          :update-data="getAllUsers"
      />
    </div>
  </transition>

  <n-modal
      :title="t('adminViews.userMgr.searchUser')"
      v-model:show="showSearchNModal"
      preset="dialog"
      :positive-text="t('adminViews.userMgr.query')"
      :negative-text="t('adminViews.userMgr.cancel')"
      style="width: 480px"
      @positive-click="animated=false; getAllUsers()"
      @negative-click="showSearchNModal=false;"
      @before-leave="searchEmail=''"
      :show-icon="false"
  >
    <div style="margin-top: 30px"></div>
    <n-form-item path="email" :label="t('adminViews.userMgr.userEmail')" autofocus>
      <n-input
          autofocus
          @keyup.enter="showSearchNModal=false;animated=false; getAllUsers()"
          :placeholder="t('adminViews.userMgr.inputUserEmail')"
          v-model:value="searchEmail"
      />
    </n-form-item>
  </n-modal>

  <n-modal
      :title="t('adminViews.userMgr.addNewUser')"
      v-model:show="showAddUserModal"
      preset="dialog"
      :positive-text="t('adminViews.userMgr.submit')"
      :negative-text="t('adminViews.userMgr.cancel')"
      style="width: 480px"
      @positive-click="submitAddForm"
      @negative-click="showAddUserModal=false; newUser.email=''; newUser.password=''"
      :show-icon="false"
  >
    <div style="margin-top: 30px"></div>

    <n-form
        ref="addUserForm"
        :model="newUser"
        :rules="addRules"
        label-placement="top"
        label-width="120px"
    >
      <n-form-item path="email" :label="t('adminViews.userMgr.email')">
        <n-input
            autofocus
            @keyup.enter="submitAddForm"
            :placeholder="t('adminViews.userMgr.inputEmail')"
            v-model:value="newUser.email"
        />
      </n-form-item>
      <n-form-item path="password" :label="t('adminViews.userMgr.password')">
        <n-input
            type="password"
            autofocus
            @keyup.enter="submitAddForm"
            :placeholder="t('adminViews.userMgr.inputPassword')"
            v-model:value="newUser.password"
        />
      </n-form-item>
    </n-form>
  </n-modal>

  <n-drawer
      width="40%"
      placement="right"
      v-model:show="showEditUserDrawer"
  >
    <n-drawer-content :title="t('adminViews.userMgr.editUser')">
      <n-form
          ref="editUserForm"
          :model="editUser"
          :rules="editRules"
          label-placement="top"
          label-width="120px"
          scroll-x="1000"
      >
        <!-- 邮箱 必填 -->
        <n-form-item :label="t('adminViews.userMgr.email')" path="email">
          <n-input :disabled="true" v-model:value="editUser.email"></n-input>
        </n-form-item>

        <!-- 密码 选填 -->
        <n-form-item :label="t('adminViews.userMgr.password')" path="password">
          <n-input
              :autofocus="false"
              showPasswordOn="click"
              type="password"
              v-model:value="editUser.password"/>
        </n-form-item>

        <!-- 用户邀请码 选填 -->
        <n-form-item :label="t('adminViews.userMgr.inviteCode')" path="invite_code">
          <n-input clearable v-model:value="editUser.invite_code"></n-input>
        </n-form-item>

        <!-- 是否将用户设置为管理员 必填 -->
        <n-form-item :label="t('adminViews.userMgr.isAdmin')" path="is_admin">
          <n-select
              :options="[{ label: t('adminViews.userMgr.no'), value: 'false' }, { label: t('adminViews.userMgr.yes'), value: 'true' }]"
              v-model:value="isAdminString"
          ></n-select>
        </n-form-item>

        <!-- 是否将用户设置为员工 必填 -->
        <n-form-item :label="t('adminViews.userMgr.isStaff')" path="is_staff">
          <n-select
              :options="[{ label: t('adminViews.userMgr.no'), value: 'false' }, { label: t('adminViews.userMgr.yes'), value: 'true' }]"
              v-model:value="isStaffString"
          ></n-select>
        </n-form-item>

        <!-- 用户余额 必填 -->
        <n-form-item :label="t('adminViews.userMgr.balance')" path="balance">
          <n-input-number style="width: 100%" v-model:value.number="editUser.balance"></n-input-number>
        </n-form-item>

        <div style="display: flex; justify-content: flex-end; margin-top: 20px">
          <n-button style="width: auto; margin-right: 15px" @click="cancelEdit" type="primary" secondary>
            {{ t('adminViews.userMgr.cancel') }}
          </n-button>
          <n-button style="width: auto" type="primary" @click="submitForm">
            {{ t('adminViews.userMgr.submit') }}
          </n-button>
        </div>
      </n-form>
    </n-drawer-content>
  </n-drawer>

</template>

<style lang="less" scoped>
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