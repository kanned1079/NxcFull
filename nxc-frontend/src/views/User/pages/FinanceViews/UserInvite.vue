<script lang="ts" setup>
import {ref, computed, onMounted, h} from "vue";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {
  CheckmarkDoneOutline as passOrderIcon, ChevronDownOutline as downIcon,
  LinkOutline as linkIcon,
  PauseOutline as closeOrderIcon
} from "@vicons/ionicons5"
import useAppInfosStore from "@/stores/useAppInfosStore";
import {useI18n} from "vue-i18n";
import useThemeStore from "@/stores/useThemeStore";
import instance from "@/axios";
import {NButton, NDropdown, NIcon, NTag, useMessage} from "naive-ui";
import renderIcon from "@/utils/iconFormator";
import {formatDate} from "@/utils/timeFormat";

const {t} = useI18n();
const userInfoStore = useUserInfoStore();
const appInfoStore = useAppInfosStore();
const themeStore = useThemeStore();
const message = useMessage();

let animated = ref<boolean>(false);
let showCreateInviteCodeMention = ref<boolean>(false);

let myFaCode = ref<string>('')

let pageCount = ref(10)

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

const columns = [
  {
    title: '#',
    key: 'id'
  },
  {
    title: '邮箱地址',
    key: 'email'
  },
  {
    title: '注册时间',
    key: 'created_at'
  },
];


interface MyInvitedUser {
  id: number
  email: string
  created_at: string
}

let myInvitedUserList = ref<MyInvitedUser[]>([]);

let handleCreateMyInviteCode = async () => {
  try {
    animated.value = false
    let {data} = await instance.post('/api/user/v1/invite/code', {
      user_id: userInfoStore.thisUser.id,
    })
    if (data.code === 200) {
      console.log('created ok.');
      await handleGetMyInviteCode()
    } else {
      message.error('创建失败' + data.msg || '');
    }
  } catch (err: any) {
    console.log(err
    )
  }
}

let handleGetMyInviteCode = async () => {
  try {
    let {data} = await instance.get('/api/user/v1/invite/code', {
      params: {
        user_id: userInfoStore.thisUser.id,
      }
    })
    if (data.code === 200) {
      console.log(data)
      myFaCode.value = data.my_invite_code
      animated.value = true
    } else if (data.code === 404) {
      console.log('无')
    } else {
      console.log(data.msg)
    }
  } catch (err: any) {
    console.log(err)
  }
}

let handleGetMyInvitedUserList = async () => {
  try {
    let {data} = await instance.get('/api/user/v1/invite/users', {
      params: {
        user_id: userInfoStore.thisUser.id,
        page: dataSize.value.page,
        size: dataSize.value.pageSize,
      }
    })
    if (data.code === 200) {
      console.log(data)
      pageCount.value = data.page_count
      data.user_list.forEach((item: MyInvitedUser) => myInvitedUserList.value.push(item))
    } else if (data.code === 404) {
      console.log('无')
    } else {
      console.log(data.msg)
    }
  } catch (err: any) {
    console.log(err)
  }
}


onMounted(async () => {
  themeStore.userPath = '/dashboard/invite'

  await handleGetMyInviteCode()

  if (myFaCode.value !== '') {
    await handleGetMyInvitedUserList()
  }

  animated.value = true;
})

</script>

<script lang="ts">
export default {
  name: "UserInvite",
}
</script>

<template>
  <div style="padding: 20px">
    <n-card
        hoverable
        :embedded="true"
        :bordered="false"
        :title="'我的邀请'"
    ></n-card>

    <n-card
        hoverable
        :bordered="false"
        :embedded="true"
        content-style="padding: 0"
        style="margin-top: 10px"
    >
      <n-alert type="success" :bordered="false">
        当您邀请的用户充值成功时，您可以获得其充值金额的10%。
      </n-alert>
    </n-card>


  </div>


  <transition name="slide-fade">
    <div style="padding: 0 20px 20px 20px" v-if="animated">

      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
      >

        <div class="code-mgr-main">
          <p class="code-mgr-main-title">邀请码管理</p>
          <n-button
              type="primary"
              secondary
              size="small"
              @click="showCreateInviteCodeMention=true"
          >
            {{ '生成邀请码' }}
          </n-button>
        </div>

        <n-table
            :bordered="false"
            :bottom-bordered="false"
            class="code-info"
        >
          <n-tr style=" margin-left: 200px">
            <n-td style="background-color: rgba(0,0,0,0.0);">{{ '邀请码' }}</n-td>
            <n-td style="background-color: rgba(0,0,0,0.0);">
              <n-tag>
                {{ myFaCode !== '' ? myFaCode : '您还没有邀请码' }}
              </n-tag>
            </n-td>
          </n-tr>
          <n-tr style=" margin-left: 200px">
            <n-td style="background-color: rgba(0,0,0,0.0);">{{ '邀请链接' }}</n-td>
            <n-td style="background-color: rgba(0,0,0,0.0);">
              <n-tag>
                {{
                  myFaCode !== '' ? ` ${appInfoStore.appCommonConfig.app_url}/register?code=${myFaCode}` : '请先生成邀请码'
                }}
              </n-tag>
            </n-td>
          </n-tr>
        </n-table>


      </n-card>


      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
          style="margin-top: 20px"
          content-style="padding: 0"
      >
        <div style="padding: 20px">
          <p style="font-weight: bold; font-size: 1.1rem">我邀请的用户</p>

        </div>

        <n-data-table
            striped
            class="table"
            :columns="columns"
            :data="myInvitedUserList"
            :pagination="false"
            :bordered="true"
        />
      </n-card>

      <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">
        <n-pagination
            size="medium"
            v-model:page.number="dataSize.page"
            :page-count="pageCount"
            @update:page="animated=false; handleGetMyInvitedUserList()"
        />
        <n-select
            style="width: 160px; margin-left: 20px"
            v-model:value.number="dataSize.pageSize"
            size="small"
            :options="dataCountOptions"
            :remote="true"
            @update:value="animated=false; dataSize.page = 1; handleGetMyInvitedUserList()"
        />
      </div>

    </div>
  </transition>

  <n-modal
      v-model:show="showCreateInviteCodeMention"
      preset="dialog"
      title="确认"
      content="请注意，邀请码创建后不可关闭。"
      positive-text="确认"
      negative-text="算了"
      @positive-click="handleCreateMyInviteCode"
      @negative-click="showCreateInviteCodeMention=false"
  />

</template>

<style lang="less" scoped>

.code-mgr-main {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;

  .code-mgr-main-title {
    font-size: 1.1rem;
    font-weight: bold;
  }
}

.code-info {
  margin-top: 20px;
  background-color: rgba(0, 0, 0, 0.0);
}

.code-div {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: baseline;
  font-size: 1rem;

  .code-title {
    margin-right: 6px;
  }
}
</style>