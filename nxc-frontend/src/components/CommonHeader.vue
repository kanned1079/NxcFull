<script setup lang="ts" name="CommonHeader">
import {ref} from 'vue'
import {Sunny as sunIcon, MoonSharp as moonIcon, PersonCircle as userIcon} from '@vicons/ionicons5'
import useThemeStore from "@/stores/useThemeStore";
// import type {GlobalTheme} from 'naive-ui'
import { useRouter } from 'vue-router';

import useUserDropDown from "@/stores/userDropdownItems";
import useUserInfoStore from "@/stores/useUserInfoStore";
const userDropdownStore = useUserDropDown()
const themeStore = useThemeStore();
const userInfoStore = useUserInfoStore();
let thisUser = userInfoStore.thisUser;

const theme = themeStore.getTheme;

let options = userDropdownStore.options

const router = useRouter();


let handleSelect = (key: string | number) => {
  switch (key) {
    case 'profile': {
      console.log('个人中心')
      break
    }
    case 'editProfile': {
      console.log('修改个人资料')
      break
    }
    case 'logout': {
      console.log('退出登录')
      // sessionStorage.removeItem('isAuthed')
      // userInfoStore.setAndSaveAuthStatus(false)
      // sessionStorage.removeItem('token')
      // userInfoStore.isAuthed = false as boolean
      // router.push({ path: '/admin/login' })
      userInfoStore.logout()
      break
    }
  }
}

let handleChangeTheme = () => {
  console.log('修改主题颜色')
  themeStore.enableDarkMode = !themeStore.enableDarkMode;
  themeStore.saveEnableDarkMode();
}
</script>

<template>
  <div class="root">
    <div class="l-content">
      <p class="txt">
        仪表板
      </p>
    </div>
    <div class="r-content">
    <span class="all">
      <div class="color-switch">
      <n-button quaternary class="btn" size="medium" @click="handleChangeTheme">
        <template #icon>
          <n-icon v-if="!themeStore.enableDarkMode" size="20"><sunIcon/></n-icon>
          <n-icon v-else size="18"><moonIcon/></n-icon>
        </template>
      </n-button>
    </div>
      <div class="info">
        <n-dropdown
            @select="handleSelect"
            :options="options"
            placement="bottom"
            size="large"
            style="width: 180px;"
            content-style="{backgroundColor='#e3e5e7'}"
            class="dd">

          <n-button quaternary style="font-size: 1rem; color: white; align-items: center">
            <n-icon size="20" style="padding-right: 10px"><userIcon/></n-icon> {{ thisUser.email }}
          </n-button>
        </n-dropdown>
      </div>
    </span>

    </div>
  </div>
</template>

<style lang="less" scoped>
.root {
  display: flex;
  justify-content: space-between;
  background-color: v-bind("themeStore.getTheme.topHeaderBgColor");

  .l-content {
    width: 100px;
    line-height: 52px;
    padding-left: 30px;

    .txt {
      color: v-bind('theme.topHeaderTextColor')
    }
  }

  .r-content {
    width: 280px;
    min-width: 240px;
    display: flex;
    align-items: center;
    margin-right: 20px;

    .all {
      display: flex;
      align-items: center; /* 使子元素上下居中 */

      .color-switch {
        margin-right: 5px;
        .btn {
          color: white;
        }
      }

      .info {
        width: 240px;
        .dd {
          display: flex;
          align-items: center;
          .btn {
            font-size: 3rem;
            color: white;
          }
        }
      }
    }
  }
}
</style>