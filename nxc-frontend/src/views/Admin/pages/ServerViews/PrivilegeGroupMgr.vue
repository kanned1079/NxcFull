<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onMounted, ref} from 'vue'
import {type FormInst, NButton, NIcon, type NotificationType, useMessage, useNotification} from "naive-ui";

import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import instance from "@/axios";
import {BarChartOutlined, UserOutlined} from '@vicons/antd' // 引入所需圖標

const {t} = useI18n()
const message = useMessage()

let modifyFunc = ref<string>('add')
let modifyId = ref<number | null>(null)
let pagination = ref({pageSize: 4})
// let page = ref<number>(1)
let pageCount = ref(10)

let showLoading = ref<boolean>(false)

let animated = ref<boolean>(false)

let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})


interface GroupItem {
  id: number
  name: string
  user_count?: number
  plan_count?: number
  updated_at?: string
  created_at?: string
  deleted_at?: string
}

// 假數據
const groupList = ref<GroupItem[]>([])

// 定義表格列
const columns = ref(
    [
      {title: '權限組ID', key: 'id'},
      {title: '組名稱', key: 'name'},
      {
        title: '用戶數量',
        key: 'user_count',
        render(row: GroupItem) {
          return h('div', {style: {display: 'flex', flexDirection: 'row',}}, [
            h(NIcon, {size: 18, style: {marginRight: 5}}, {default: () => h(UserOutlined)}), // 用戶圖標
            `${row.hasOwnProperty('user_count') ? row.user_count : 0}`,
          ])
        }
      },
      {
        title: '計畫數量',
        key: 'plan_count',
        render(row: GroupItem) {
          return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
            h(NIcon, {size: 18, style: {marginRight: 5}}, {default: () => h(BarChartOutlined)}), // 計畫圖標
            `${row.hasOwnProperty('plan_count') ? row.plan_count : 0}`,
          ])
        }
      },
      {
        title: '操作',
        key: 'actions',
        fixed: 'right',
        render(row: GroupItem) {
          return h('div', {style: 'display: flex; gap: 8px'}, [
            h(NButton, {
              size: 'small',
              type: 'primary',
              secondary: true,
              onClick: () => editGroup(row)
            }, {default: () => '編輯'}),
            h(NButton, {
              size: 'small',
              type: 'error',
              secondary: true,
              onClick: () => deleteGroup(row.id)
            }, {default: () => '刪除'})
          ])
        }
      }
    ]
)


// 編輯組操作
const editGroup = (row: GroupItem) => {
  // message.info(`編輯權限組：${id}`)
  modifyFunc.value = 'edit'
  newGroupName.value = row.name
  showModal.value = true
  modifyId.value = row.id
}

// 刪除組操作
const deleteGroup = async (id: number) => {
  animated.value = false
  message.warning(`刪除權限組：${id}`)
  let {data} = await instance.delete('http://localhost:8081/api/admin/v1/groups', {
    data: {
      id,
    }
  })
  if (data.code === 200) {
    console.log('删除成功')
    await getAllGroups()
  }
}

const formRef = ref<FormInst | null>(null)

const notification = useNotification();
const themeStore = useThemeStore()
const apiAddrStore = useApiAddrStore()

interface PrivilegeGroup {
  id: number
  name: string
  plan_count: number
  user_count: number
}

let showModal = ref(false)

let formValue = ref({
  privilege_group: {
    name: '',
  }
})

let rules = {
  privilege_group: {
    required: true,
    message: '权限组名称是必填的',
    trigger: 'blur'
  }
}


let newGroupName = ref<string>('')

let notify = (type: NotificationType, title: string, meta?: string) => {
  notification[type]({
    content: title,
    meta: meta,
    duration: 2500,
    keepAliveOnHover: true
  })
}

let handleAddGroup = async () => {
  modifyFunc.value = 'add'
  newGroupName.value = ''
  showModal.value = true
}

let submitNewGroup = async () => {
  console.log('添加权限组')
  try {
    let {data} = await instance.post('http://localhost:8081/api/admin/v1/groups', {
      group_name: newGroupName.value
    })
    if (data.code === 200) {
      notify('success', '添加成功')
    } else {
      notify('error', '添加失败', data.error)
    }
  } catch (error) {
    console.log(error)
  }
}

let closeModal = () => {
  showModal.value = false
}

// let groupList = ref<PrivilegeGroup>([])

let getAllGroups = async () => {
  showLoading.value = true
  try {
    let {data} = await instance.get('http://localhost:8081/api/admin/v1/groups', {
      params: {
        page: dataSize.value.page,
        size: dataSize.value.pageSize,
      }
    })
    if (data.code === 200) {
      groupList.value = []
      showLoading.value = false
      console.log('获取到的权限组列表', data)
      data.group_list.forEach((group: PrivilegeGroup) => groupList.value.push(group))
      pageCount.value = data.page_count
      // groupList.value = data.group_list
      animated.value = true
    }
  } catch (error) {
    console.log(error)
  }
}

let submit = async () => {
  if (modifyFunc.value === 'add') {
    await submitNewGroup()
  } else {
    await submitUpdate()
  }
  await getAllGroups()
}

let submitUpdate = async () => {
  let {data} = await instance.put('/api/admin/v1/groups', {
    id: modifyId.value,
    name: newGroupName.value,
  })
  if (data.code === 200) {
    console.log('修改成功')
    notify('success', '修改权限组成功')
  } else {
    notify('error', '修改权限组失败')
  }
}

let dataCountOptions = [
  {
    label: computed(() => t('pagination.perPage10')).value,
    value: 10,
  },
  {
    label: computed(() => t('pagination.perPage20')).value,
    value: 20,
  },
  {
    label: computed(() => t('pagination.perPage50')).value,
    value: 50,
  },
  {
    label: computed(() => t('pagination.perPage100')).value,
    value: 100,
  },
]

onMounted(async () => {
  console.log('挂载权限组管理')
  themeStore.contentPath = '/admin/dashboard/group'
  themeStore.menuSelected = 'privilege-group-mgr'
  await getAllGroups()
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'PrivilegeGroupMgr',
}
</script>

<template>
  <div style="padding: 20px 20px 0 20px">
    <n-card hoverable :embedded="true" title="权限组管理" :bordered="false">
      <n-button type="primary" :bordered="false" class="add-btn" @click="handleAddGroup">添加权限组</n-button>
    </n-card>
  </div>

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card style="margin-top: 20px" :embedded="true" hoverable :bordered="false" content-style="padding: 0;">
        <!--      在這裏放置表格-->
        <!--      表格有五列 列名分別是 權限組ID， 組名稱， 用戶數量（前面使用n-icon放置一個用戶圖標）， 計畫數量（前面使用n-icon放置一個圖標， 操作（內部為兩個按鈕 水平排列分別為編輯和刪除）-->
        <n-spin :show="showLoading">
          <n-data-table
              striped
              :columns="columns"
              :data="groupList"
              :scroll-x="900"
              :remote="true"
          />
        </n-spin>
      </n-card>
      <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">
        <n-pagination
            size="medium"
            v-model:page.number="dataSize.page"
            :page-count="pageCount"
            @update:page="animated=false; getAllGroups() "
        />
        <n-select
            style="width: 160px; margin-left: 20px"
            v-model:value.number="dataSize.pageSize"
            size="small"
            :options="dataCountOptions"
            :remote="true"
            @update:value="animated=false; dataSize.page = 1; getAllGroups()"
        />
      </div>
    </div>
  </transition>

  <n-modal
      :title="modifyFunc==='add'?'新建权限组':'修改名称'"
      v-model:show="showModal"
      preset="dialog"
      :positive-text="modifyFunc==='add'?'确认添加':'确认修改'"
      negative-text="算了"
      style="width: 480px; border: 0"
      @positive-click="submit"
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
      <n-form-item label="权限组名称" path="privilege_group.name">
        <n-input v-model:value="newGroupName" placeholder="输入权限组名称"/>
      </n-form-item>
    </n-form>

    <!--    <template #footer>-->
    <!--      尾部-->
    <!--    </template>-->


  </n-modal>
</template>

<style lang="less" scoped>
.root {
  //padding: 20px;
  padding: 0 20px 20px 20px;

  .add-btn {
    margin-top: 10px;
  }

  .btn-pagination {
    margin-top: 20px;
    display: flex;
    flex-direction: row;
    justify-content: right;
  }
}

</style>