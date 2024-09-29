<script setup lang="ts">
import {h, onMounted, ref} from 'vue'
import {type FormInst, NButton, NIcon, type NotificationType, useMessage, useNotification} from "naive-ui";

import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import instance from "@/axios";
import {BarChartOutlined, UserOutlined} from '@vicons/antd' // 引入所需圖標
const message = useMessage()

let modifyFunc = ref<string>('add')
let modifyId = ref<number>(null)
let pagination =  ref({ pageSize: 4 })
let page = ref<number>(1)
let pageCount = ref(18)

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
          return h('div', {style: {display: 'flex', flexDirection: 'row', }}, [
            h(NIcon, {size: 18, style: {marginRight: 5}}, {default: () => h(UserOutlined)}), // 用戶圖標
            ` ${row.user_count}`
          ])
        }
      },
      {
        title: '計畫數量',
        key: 'plan_count',
        render(row: GroupItem) {
          return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
            h(NIcon, {size: 18, style: {marginRight: 5}}, {default: () => h(BarChartOutlined)}), // 計畫圖標
            ` ${row.plan_count}`
          ])
        }
      },
      {
        title: '操作',
        key: 'actions',
        width: 200,
        fixed: 'right',
        render(row: GroupItem) {
          return h('div', {style: 'display: flex; gap: 8px'}, [
            h(NButton, {size: 'small', type: 'primary', onClick: () => editGroup(row)}, {default: () => '編輯'}),
            h(NButton, {size: 'small', type: 'error', onClick: () => deleteGroup(row.id)}, {default: () => '刪除'})
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
  message.warning(`刪除權限組：${id}`)
  let {data} = await instance.delete('/api/admin/v1/groups', {
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
  created_at?: string
  updated_at?: string
  deleted_at?: string
}

let showModal = ref(false)

let formValue = ref({
  privilege_group: {
    name: '',
  } as PrivilegeGroup
})

let rules = {
  privilege_group: {
    required: true,
    message: '权限组名称是必填的',
    trigger: 'blur'
  }
}


let newGroupName = ref<string>('')

let notify = (type: NotificationType, title: string, meta: string) => {
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
    let {data} = await instance.post(apiAddrStore.apiAddr.admin.addNewPrivilegeGroup, {
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
  try {
    // groupList.value = []
    let {data} = await instance.get(apiAddrStore.apiAddr.admin.getAllPrivilegeGroups)
    if (data.code === 200) {
      console.log('获取到的权限组列表', data)
      // data.group_list.forEach((group: PrivilegeGroup) => groupList.value.push(group))
      groupList.value = data.group_list
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
  console.log(12345)
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

onMounted(() => {
  console.log('挂载权限组管理')
  themeStore.contentPath = '/admin/dashboard/group'
  themeStore.menuSelected = 'privilege-group-mgr'
  getAllGroups()
})

</script>

<script lang="ts">
export default {
  name: 'PrivilegeGroupMgr',
}
</script>

<template>
  <div class="root">
    <n-card hoverable :embedded="true" title="权限组管理" :bordered="false">
      <n-button type="primary" :bordered="false" class="add-btn" @click="handleAddGroup">添加权限组</n-button>
    </n-card>

    <n-card style="margin-top: 20px" :embedded="true" hoverable :bordered="false" content-style="padding: 0;">
      <!--      在這裏放置表格-->
      <!--      表格有五列 列名分別是 權限組ID， 組名稱， 用戶數量（前面使用n-icon放置一個用戶圖標）， 計畫數量（前面使用n-icon放置一個圖標， 操作（內部為兩個按鈕 水平排列分別為編輯和刪除）-->
      <n-data-table
          :columns="columns"
          :data="groupList"
          :scroll-x="900"
      />
    </n-card>

    <div class="btn-pagination">
      <n-pagination  v-model:page="page" :page-count="pageCount" />
    </div>

  </div>

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

    <template #footer>
      尾部
    </template>


  </n-modal>
</template>

<style lang="less" scoped>
.root {
  padding: 20px;

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