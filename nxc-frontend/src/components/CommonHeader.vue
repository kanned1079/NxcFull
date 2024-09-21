<script setup lang="ts" name="CommonHeader">
import {onMounted, ref} from 'vue'
import {
  ChevronDownOutline as downIcon,
  List as menuIcon,
  MoonSharp as moonIcon,
  PersonCircle as userIcon,
  Sunny as sunIcon,
} from '@vicons/ionicons5'
import useThemeStore from "@/stores/useThemeStore";
// import type {GlobalTheme} from 'naive-ui'
import {useRouter} from 'vue-router';
import type {DrawerPlacement} from 'naive-ui'
import useUserDropDown from "@/stores/userDropdownItems";
import useUserInfoStore from "@/stores/useUserInfoStore";
import CommonAside from "@/components/CommonAside.vue";

let dropDownBtn = ref(null)
let dropDownWidth = ref(0)

const userDropdownStore = useUserDropDown()
const themeStore = useThemeStore();
const userInfoStore = useUserInfoStore();
let thisUser = userInfoStore.thisUser;

const theme = themeStore.getTheme;
const active = ref(false)
const placement = ref<DrawerPlacement>('right')

let options = userInfoStore.thisUser.isAdmin ? userDropdownStore.admin_options : userDropdownStore.user_options;

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

const activateMenu = (place: DrawerPlacement) => {
  active.value = true
  placement.value = place
}

let handleChangeTheme = () => {
  console.log('修改主题颜色')
  themeStore.enableDarkMode = !themeStore.enableDarkMode;
  themeStore.saveEnableDarkMode();
}

onMounted(() => {
  dropDownWidth.value = dropDownBtn.value.getBoundingClientRect().width || dropDownBtn.value.offsetWidth;

})

</script>

<template>
  <div class="root">
    <div class="l-content">
      <n-button
          v-if="themeStore.menuCollapsed"
          quaternary
          class="show-menu-btn"
          size="medium"
          @click="activateMenu('left')">
        <n-icon size="20" style="margin-right: 5px">
          <menuIcon/>
        </n-icon>
        菜单
      </n-button>

      <p class="txt" v-if="!themeStore.menuCollapsed">
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
      <div ref="dropDownBtn" class="info">
        <n-dropdown
            @select="handleSelect"
            :options="options"
            placement="bottom"
            size="large"
            :content-style="{backgroundColor:'#e3e5e7', width:'320px'}"
            class="dd"
            :width="themeStore.menuCollapsed?'100%':dropDownWidth"
        >
          <n-button quaternary style="font-size: 1rem; color: white; align-items: center">
            <n-icon size="20" style="padding-right: 10px"><userIcon/></n-icon>

            <n-p v-if="!themeStore.menuCollapsed" :style="{color:theme.topHeaderTextColor}" style="margin-right: 10px">{{
                thisUser.email
              }}</n-p>
            <n-icon size="16" style="font-weight: 800"><downIcon/></n-icon>

          </n-button>
        </n-dropdown>
      </div>
    </span>

    </div>
  </div>

  <n-drawer v-model:show="active" :width="themeStore.menuCollapsed?'80%':'320px'" :placement="placement">
    <CommonAside></CommonAside>
  </n-drawer>
</template>

<style lang="less" scoped>
.root {
  display: flex;
  justify-content: space-between;
  background-color: v-bind("themeStore.getTheme.topHeaderBgColor");

  .l-content {
    display: flex;
    align-items: center;
    line-height: 52px;
    height: 52px;
    padding-left: 10px;

    .show-menu-btn {
      color: white;
    }

    .txt {
      margin-left: 10px;
      color: v-bind('theme.topHeaderTextColor')
    }
  }


  .r-content {
    display: flex;
    align-items: center;
    margin-right: 15px;

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

        .dd {
          display: flex;
          align-items: center;

          .btn {
            font-size: 3rem;
            color: white;
          }
        }

        .n-dropdown-menu {
          margin-top: 20px;
        }
      }
    }
  }
}
</style>