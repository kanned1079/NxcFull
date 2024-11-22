<script lang="ts" setup>
import {ref, computed, onMounted, h} from "vue";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {useI18n} from "vue-i18n";
import useThemeStore from "@/stores/useThemeStore";
import instance from "@/axios";

const {t} = useI18n();
const userInfoStore = useUserInfoStore();
const themeStore = useThemeStore();

let animated = ref<boolean>(false);

let myFaCode = ref<string>('')

let handleCreateMyInviteCode = async () => {
  try {

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
      myFaCode.value = data.my_fa_code
    }
  } catch (err: any) {
    console.log(err)
  }
}


onMounted(() => {
  themeStore.userPath = '/dashboard/invite'
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
  </div>

  <transition name="slide-fade">
    <div style="padding: 0 20px 20px 20px">
      <n-card
        hoverable
        :embedded="true"
        :bordered="false"
      >
        <n-h3 style="font-weight: bold">您的邀请码</n-h3>

        <n-h4>{{ userInfoStore.thisUser.inviteCode!==''?userInfoStore.thisUser.inviteCode:'您还没有邀请码' }}</n-h4>

        <n-button
          type="primary"
          secondary
        >
          {{ '生成邀请码' }}
        </n-button>

      </n-card>

    </div>
  </transition>

</template>

<style lang="less" scoped>

</style>