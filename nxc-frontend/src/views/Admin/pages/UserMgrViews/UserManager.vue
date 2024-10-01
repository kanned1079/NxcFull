<script setup lang="ts">
import type {DataTableColumns} from 'naive-ui'
import {computed, onMounted} from 'vue'
import { NButton, NTag, useMessage, NDropdown } from 'naive-ui'
import {h} from 'vue'
// PersonAddOutline
import {PersonAddOutline as addUserIcon,} from '@vicons/ionicons5'
import useThemeStore from "@/stores/useThemeStore";
import instance from "@/axios";
import useApiAddrStore from "@/stores/useApiAddrStore";
const apiAddrStore = useApiAddrStore();

const themeStore = useThemeStore();

const rsp = {
  "code": 500,
  "msg": "ok",
  "users": [
    {
      "id": 1,
      "invite_user_id": 0,
      "name": "kanna",
      "email": "admin@ikanned.com",
      "isAdmin": true,
      "is_staff": false,
      "balance": 29.8,
      "last_login": "2019-10-16T15:45:00+08:00",
      "last_login_ip": "",
      "createdAt": "2024-09-03T00:04:25+08:00",
      "updatedAt": "2024-09-03T00:04:25+08:00",
      "deletedAt": null
    }
  ]
}

interface RowData {
  key: number
  id: number
  email: string
  status: string
  privilege_group: string
  tags: string[]
  end_time: string
  balance: number
  created_at: string
}

// let getStatus = computed((status: boolean) => status?'正常':'封禁')

let createColumns = ({sendMail}: { sendMail: (rowData: RowData) => void }): DataTableColumns<RowData> => {
  return [
    {
      title: 'ID',
      key: 'id'
    },
    {
      title: '用户名',
      key: 'name'
    },
    {
      title: '邮箱',
      key: 'email'
    },
    {
      title: '是否管理员',
      key: 'status'
    },
    {
      title: '是否员工',
      key: 'status'
    },
    {
      title: 'Tags',
      key: 'tags',
      render(row) {
        const tags = row.tags.map((tagKey) => {
          return h(
              NTag,
              {
                style: {
                  marginRight: '6px'
                },
                type: 'info',
                bordered: false
              },
              {
                default: () => tagKey
              }
          )
        })
        return tags
      }
    },
    {
      title: '权限组',
      key: 'privilege_group',
    },
    {
      title: '到期时间',
      key: 'end_time',
    },
    {
      title: '余额',
      key: 'balance',
    },
    {
      title: '加入时间',
      key: 'created_at',
    },

    {
      title: 'Action',
      key: 'actions',
      render(row) {
        return h(
            NButton,
            // NDropdown,
            {
              size: 'small',
              onClick: () => sendMail(row)
            },
            {
              // NButton,
              default: () => 'Send Email'
             }
        )
      }
    }
  ]
}

let createData = (): RowData[] => {
  return [
    {
      key: 0,
      id: 11,
      email: 'xxx@xxx.com1',
      status: 'ok',
      privilege_group: 'A1',
      end_time: '2019-03-21',
      balance: 123.32,
      created_at: '2020-24-23',
      tags: ['正常'],
    },
    {
      key: 1,
      id: 11,
      email: 'xxx@xxx.com1',
      status: 'ok',
      privilege_group: 'A1',
      end_time: '2019-03-21',
      balance: 123.32,
      created_at: '2020-24-23',
      tags: ['正常'],
    },
    {
      key: 2,
      id: 11,
      email: 'xxx@xxx.com1',
      status: 'ok',
      privilege_group: 'A1',
      end_time: '2019-03-21',
      balance: 123.32,
      created_at: '2020-24-23',
      tags: ['正常'],
    },
    {
      key: 3,
      id: 11,
      email: 'xxx@xxx.com1',
      status: 'ok',
      privilege_group: 'A1',
      end_time: '2019-03-21',
      balance: 123.32,
      created_at: '2020-24-23',
      tags: ['正常'],
    },
    {
      key: 4,
      id: 11,
      email: 'xxx@xxx.com1',
      status: 'ok',
      privilege_group: 'A1',
      end_time: '2019-03-21',
      balance: 123.32,
      created_at: '2020-24-23',
      tags: ['正常'],
    },

  ]
}

let data = createData()
let columns = createColumns({
  sendMail: (rowData: RowData) => {
    console.log(`send mail to ${rowData.id}`)
  }
})
let pagination =  {
  pageSize: 10,
}

let getAllUsers = async () => {
  let {data} = await instance.get('/api/admin/v1/users')
  console.log(data)
}

onMounted(() => {
  themeStore.menuSelected = 'user-manager'
  themeStore.contentPath = '/admin/dashboard/usermanager'
  getAllUsers()

})

</script>

<script lang="ts">
export default {
  name: 'UserManager'
}
</script>

<template>
  <div class="root">
    <n-card :embedded="true" class="card1" title="用户管理">
      <n-button class="btn" tertiary type="primary" size="medium">过滤器</n-button>
      <n-button class="btn" tertiary type="primary" size="medium">操作</n-button>
      <n-button class="btn" tertiary type="primary" size="medium">
        <n-icon size="14px" style="padding-right: 8px">
          <addUserIcon/>
        </n-icon>
        新建用户
      </n-button>
      <n-data-table
          size="large"
          :bordered="false"
          :columns="columns"
          :data="data"
          :pagination="pagination"
          style="margin-top: 20px"
      />
    </n-card>

  </div>

</template>

<style lang="less" scoped>
.root {
  padding: 20px;

  .card1 {
    .btn {
      margin-right: 10px;
    }
  }
}

.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}
</style>