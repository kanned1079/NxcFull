<script setup lang="ts">
import {onMounted, ref} from 'vue'
import {type FormInst, type NotificationType, useNotification} from "naive-ui";

import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import instance from "@/axios";

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

let groupList = ref<PrivilegeGroup>([])

let getAllGroups = async () => {
  try {
    groupList.value = []
    let {data} = await instance.get(apiAddrStore.apiAddr.admin.getAllPrivilegeGroups)
    if (data.code === 200) {
      console.log('获取到的权限组列表', data)
      data.group_list.forEach((group: PrivilegeGroup) => groupList.push(group))
    }
  }catch (error) {
    console.log(error)
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
    <n-card hoverable :embedded="true" title="权限组管理">
      <n-button class="add-btn" @click="handleAddGroup">添加权限组</n-button>
    </n-card>
    <n-modal
        title="新建权限组"
        v-model:show="showModal"
        preset="dialog"
        positive-text="确认添加"
        negative-text="算了"
        style="width: 480px"
        @positive-click="submitNewGroup"
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
  </div>
</template>

<style lang="less" scoped>
.root {
  padding: 20px;
  min-width: 900px;

  .add-btn {
    margin-top: 10px;
  }
}

.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  padding: 0;
  border: 0;
}
</style>