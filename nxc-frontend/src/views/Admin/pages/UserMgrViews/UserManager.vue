<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {type FormInst, NButton, NTag, useMessage} from 'naive-ui'
import {h, onMounted, ref, computed} from 'vue'
import {PersonAddOutline as addUserIcon, RefreshOutline as refreshIcon, Search as searchIcon} from '@vicons/ionicons5'
import useThemeStore from "@/stores/useThemeStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import instance from "@/axios";
import {formatDate} from "@/utils/timeFormat";
import {hashPassword} from "@/utils/encryptor";

const {t} = useI18n()
const appInfoStore = useAppInfosStore()
const themeStore = useThemeStore();
const message = useMessage()

interface User {
  id?: number
  invite_user_id?: number
  invite_code?: string
  email?: string
  password?: string
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
let showEditUserDrawer = ref<boolean>(false)

const addUserForm = ref<FormInst | null>(null)
const editUserForm = ref<FormInst | null>(null);
const addRules = {
  email: [
    {required: true, message: computed(() => t('adminViews.userMgr.enterEmail')).value, trigger: 'blur'},
    {type: 'email', message: computed(() => t('adminViews.userMgr.enterValidEmail')).value, trigger: 'blur'},
  ],
  password: [
    {required: true, message: computed(() => t('adminViews.userMgr.enterPassword')).value, trigger: 'blur'},
    {min: 6, message: computed(() => t('adminViews.userMgr.passwordMinLength')).value, trigger: 'blur'},
    {max: 20, message: computed(() => t('adminViews.userMgr.passwordMaxLength')).value, trigger: 'blur'},
    {
      pattern: /^(?=.*[a-zA-Z])(?=.*\d)(?=.*[!@#$%^&*(),.?":{}|<>])/, // 至少包含一个字母、数字和特殊字符
      message: computed(() => t('adminViews.userMgr.passwordStrength')).value,
      trigger: 'blur'
    }
  ]
};
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
      message.success(computed(() => t('adminViews.userMgr.validationSuccess')).value);
      await handleEditUserInfo()
    }
  } catch (error) {
    console.log('表单验证失败', error);
    message.error(computed(() => t('adminViews.userMgr.validationFailed')).value);
  }
};

const submitAddForm = async () => {
  try {
    const valid = await addUserForm.value?.validate();
    if (valid) {
      console.log('提交表单2:', editUser);
      await handleAddNewUser()
    }
  } catch (error) {
    console.log('表单验证失败', error);
    showAddUserModal.value = true;
    message.error(computed(() => t('adminViews.userMgr.validationFailed')).value);
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
const columns = [
  {
    title: '#',
    key: 'id'
  },
  {
    title: computed(() => t('adminViews.userMgr.email')).value,
    key: 'email'
  },
  {
    title: computed(() => t('adminViews.userMgr.registerDate')).value,
    key: 'created_at',
    render(row: User) {
      return h('p', {}, {
        default: () => formatDate(row.created_at)
      })
    }
  },
  {
    title: computed(() => t('adminViews.userMgr.isAdmin')).value,
    key: 'is_admin',
    render(row: User) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.is_admin ? 'success' : 'default',
      }, {default: () => row.is_admin ? computed(() => t('adminViews.userMgr.yes')).value : computed(() => t('adminViews.userMgr.no')).value});
    }
  },
  {
    title: computed(() => t('adminViews.userMgr.isStaff')).value,
    key: 'is_staff',
    render(row: User) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.is_staff ? 'success' : 'default',
      }, {default: () => row.is_staff ? computed(() => t('adminViews.userMgr.yes')).value : computed(() => t('adminViews.userMgr.no')).value});
    }
  },
  {
    title: computed(() => t('adminViews.userMgr.accountStatus')).value,
    key: 'status',
    render(row: User) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.status ? 'success' : 'warning',
      }, {default: () => row.status ? computed(() => t('adminViews.userMgr.normal')).value : computed(() => t('adminViews.userMgr.banned')).value});
    }
  },
  {
    title: computed(() => t('adminViews.userMgr.inviteCode')).value,
    key: 'invite_code',
    render(row: User) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.invite_code ? 'primary' : 'default',
      }, {default: () => row.invite_code ? row.invite_code : computed(() => t('adminViews.userMgr.nullContent')).value});
    }
  },
  {
    title: computed(() => t('adminViews.userMgr.balance')).value,
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
    title: computed(() => t('adminViews.userMgr.orderCount')).value,
    key: 'order_count',
    render(row: User) {
      return h(NTag, {size: 'small', bordered: false, type: 'default'}, {default: () => row.order_count});
    }
  },
  {
    title: computed(() => t('adminViews.userMgr.planCount')).value,
    key: 'plan_count',
    render(row: User) {
      return h(NTag, {size: 'small', bordered: false, type: 'default'}, {default: () => row.plan_count});
    }
  },
  {
    title: computed(() => t('adminViews.userMgr.actions')).value, key: 'actions', render(row: User) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          secondary: true,
          bordered: false,
          onClick: () => {
            // 编辑用户
            Object.assign(editUser.value, row)
            showEditUserDrawer.value = true
          },
        }, {default: () => computed(() => t('adminViews.userMgr.editProfile')).value}),
        h(NButton, {
          size: 'small',
          type: row.status ? 'error' : 'warning',
          secondary: true,
          disabled: false,
          style: {marginLeft: '10px'},
          onClick: () => handleBlock2UnblockUserById(row)
        }, {default: () => row.status ? computed(() => t('adminViews.userMgr.banUser')).value : computed(() => t('adminViews.userMgr.unbanUser')).value})
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
    label: computed(() => t('adminViews.userMgr.dataCountOptions10')).value,
    value: 10,
  },
  {
    label: computed(() => t('adminViews.userMgr.dataCountOptions20')).value,
    value: 20,
  },
  {
    label: computed(() => t('adminViews.userMgr.dataCountOptions50')).value,
    value: 50,
  },
  {
    label: computed(() => t('adminViews.userMgr.dataCountOptions100')).value,
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
        message.warning(computed(() => t('adminViews.userMgr.noRecord')).value)
      animated.value = true
    }
  } catch (err: any) {
    message.error(err + '')
  }
}

let handleBlock2UnblockUserById = async (row: User) => {
  try {
    let {data} = await instance.patch('/api/admin/v1/users', {
      user_id: row.id
    })
    if (data.code === 200) {
      message.success(computed(() => t('adminViews.userMgr.operationSuccess')).value)
      await getAllUsers()
    } else {
      message.error(computed(() => t('adminViews.userMgr.unknownError')).value + data.msg || '')
    }
  } catch (err: any) {
    message.error(err + '')
  }
}

let handleAddNewUser = async () => {
  try {
    animated.value = false
    let hashedPwd = hashPassword(newUser.value.password)
    let {data} = await instance.post('/api/admin/v1/users', {
      email: newUser.value.email,
      password: hashedPwd,
    })
    if (data.code === 200) {
      message.success(computed(() => t('adminViews.userMgr.addUserSuccess')).value)
      await getAllUsers()
    } else {
      message.error(computed(() => t('adminViews.userMgr.unknownError')).value + data.msg || '')
    }
  } catch (err: any) {
    message.error(err + '')
  }
}

let handleEditUserInfo = async () => {
  try {
    animated.value = false
    let {data} = await instance.put('/api/admin/v1/users', {
      ...editUser.value
    })
    if (data.code === 200) {
      message.success(computed(() => t('adminViews.userMgr.updateSuccess')).value)
      await getAllUsers()
    } else {
      message.error(computed(() => t('adminViews.userMgr.unknownError')).value + data.msg || '')
    }
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
    <n-card hoverable :bordered="false" :embedded="true" class="card1" :title="t('adminViews.userMgr.userManager')">
      <n-button class="btn" tertiary type="primary" size="medium" @click="showSearchNModal=true">
        <n-icon size="14" style="padding-right: 8px">
          <searchIcon/>
        </n-icon>
        {{ t('adminViews.userMgr.query') }}
      </n-button>
      <n-button class="btn" tertiary type="primary" size="medium"
                @click="searchEmail=''; animated=false; getAllUsers()">
        <n-icon size="14" style="padding-right: 8px">
          <refreshIcon/>
        </n-icon>
        {{ t('adminViews.userMgr.reset') }}
      </n-button>
      <n-button class="btn" tertiary type="primary" size="medium" @click="showAddUserModal=true">
        <n-icon size="14" style="padding-right: 8px">
          <addUserIcon/>
        </n-icon>
        {{ t('adminViews.userMgr.addNewUser') }}
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
      >
        <!-- 邮箱 必填 -->
        <n-form-item :label="t('adminViews.userMgr.email')" path="email">
          <n-input v-model:value="editUser.email"></n-input>
        </n-form-item>

        <!-- 密码 选填 -->
        <n-form-item :label="t('adminViews.userMgr.password')" path="password">
          <n-input type="password" v-model:value="editUser.password"></n-input>
        </n-form-item>

        <!-- 用户邀请码 选填 -->
        <n-form-item :label="t('adminViews.userMgr.inviteCode')" path="invite_code">
          <n-input v-model:value="editUser.invite_code"></n-input>
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
          <n-button style="width: 80px; margin-right: 15px" @click="cancelEdit" type="primary" secondary>
            {{ t('adminViews.userMgr.cancel') }}
          </n-button>
          <n-button style="width: 80px" type="primary" @click="submitForm">
            {{ t('adminViews.userMgr.submit') }}
          </n-button>
        </div>
      </n-form>
    </n-drawer-content>
  </n-drawer>

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