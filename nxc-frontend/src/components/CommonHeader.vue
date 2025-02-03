<script setup lang="ts">
import {useI18n} from 'vue-i18n'
import {computed, onMounted, reactive, ref} from 'vue'
import {
  ChevronDownOutline as downIcon,
  Language as langIcon,
  List as menuIcon,
  LogOutOutline as LogoutIcon,
  MoonSharp as moonIcon,
  Pencil as EditIcon,
  PersonCircle as userIcon,
  PersonCircleOutline as UserIcon,
  Sunny as sunIcon,
} from '@vicons/ionicons5'
import useThemeStore from "@/stores/useThemeStore";
// import type {GlobalTheme} from 'naive-ui'
import {useRouter} from 'vue-router';
import type {DrawerPlacement} from 'naive-ui'
// import useUserDropDown from "@/stores/userDropdownItems";
import useUserInfoStore from "@/stores/useUserInfoStore";
import CommonAside from "@/components/CommonAside.vue";
import renderIcon from "@/utils/iconFormator";

const {t, locale} = useI18n()

const dropDownBtn = ref<HTMLElement | null>(null);
let dropDownWidth = ref(0)

// const userDropdownStore = useUserDropDown()
const themeStore = useThemeStore();
const userInfoStore = useUserInfoStore();
let thisUser = userInfoStore.thisUser;

const theme = themeStore.getTheme;
// const active = ref(false)
const placement = ref<DrawerPlacement>('right')

// let options = userInfoStore.thisUser.isAdmin ? userDropdownStore.admin_options : userDropdownStore.user_options;

const router = useRouter();

let handleSelect = (key: string | number) => {
  switch (key) {
    case 'profile': {
      console.log('个人中心')
      break
    }
    case 'editProfile': {
      console.log('修改个人资料')
      router.push({
        path: '/dashboard/profile'
      })
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

let langOptions = ref([
  {
    label: '简体中文',
    key: 'zh_CN',
    disabled: false
  },
  {
    label: '繁體中文',
    key: 'zh_HK',
    disabled: false
  },
  {
    label: 'English',
    key: 'en_US',
    disabled: false
  },
  {
    label: '日本語',
    key: 'ja_JP',
    disabled: false
  },
])

let user_options = reactive([
  {
    // label: '用户资料',
    label: computed(() => t('commonHeader.userData')),
    key: 'profile',
    icon: renderIcon(UserIcon)
  },
  {
    label: computed(() => t('commonHeader.editUserData')),
    key: 'editProfile',
    icon: renderIcon(EditIcon)
  },
  {
    label: computed(() => t('commonHeader.logout')),
    key: 'logout',
    icon: renderIcon(LogoutIcon)
  }
])

let admin_options = ref([
  {
    label: '退出登录',
    key: 'logout',
    icon: renderIcon(LogoutIcon)
  }
])

// 设置切换语言
let handleSelectLang = (lang: string) => {
  console.log(lang)
  // lang: en_US, zh_CN, zh_HK, ja_JP
  // const { locale } = useI18n()
  // 动态切换语言
  locale.value = lang
  localStorage.setItem('locale', lang)  // 保存语言到 localStorage
}

const activateMenu = (place: DrawerPlacement) => {
  // active.value = true
  themeStore.menuIsFlippActive = true

  placement.value = place
}

let handleChangeTheme = () => {
  console.log('修改主题颜色')
  themeStore.enableDarkMode = !themeStore.enableDarkMode;
  themeStore.saveEnableDarkMode();
}

onMounted(() => {
  // dropDownWidth.value = dropDownBtn.value.getBoundingClientRect().width || dropDownBtn.value.offsetWidth;
  if (dropDownBtn.value) {
    dropDownWidth.value = dropDownBtn.value.getBoundingClientRect().width || dropDownBtn.value.offsetWidth;
  } else {
    // 可以在这里进行一些默认值的设置或者提示操作，比如设置一个默认宽度
    dropDownWidth.value = 100;
  }
})

</script>

<script lang="ts">
export default {
  name: 'CommonHeader',
}
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
        <!--        菜单-->
        {{ $t('commonHeader.menuTxt') }}
      </n-button>

      <p class="txt" v-if="!themeStore.menuCollapsed">
        仪表板
      </p>
    </div>
    <div class="r-content">
    <span class="all">
      <span class="color-switch">
      <n-button quaternary class="btn" size="medium" @click="handleChangeTheme">
        <template #icon>
          <n-icon v-if="!themeStore.enableDarkMode" size="20"><sunIcon/></n-icon>
          <n-icon v-else size="18"><moonIcon/></n-icon>
        </template>
      </n-button>
    </span>
      <span class="lang-switch">
        <n-dropdown trigger="hover" :options="langOptions" @select="handleSelectLang">
          <n-button quaternary class="btn" size="medium">
            <template #icon>
              <n-icon size="20"><langIcon/></n-icon>
            </template>
          </n-button>
        </n-dropdown>
    </span>
      <span ref="dropDownBtn" class="info">
        <n-dropdown
            @select="handleSelect"
            :options="userInfoStore.thisUser.isAdmin?admin_options:user_options"
            placement="bottom"
            size="large"
            :content-style="{backgroundColor:'#e3e5e7', width:'320px'}"
            class="dd"
            :width="themeStore.menuCollapsed?'100%':dropDownWidth"
        >
          <n-button quaternary style="font-size: 1rem; color: white; align-items: center">
            <n-icon size="20" style="padding-right: 10px"><userIcon/></n-icon>
            <n-p
                v-if="!themeStore.menuCollapsed"
                :style="{color:theme.topHeaderTextColor}"
                style="margin-right: 10px"
            >
              {{ thisUser.name || thisUser.email }}
            </n-p>
            <n-icon size="16" style="font-weight: 800"><downIcon/></n-icon>
          </n-button>
        </n-dropdown>
      </span>
    </span>

    </div>
  </div>



  <n-drawer
      v-model:show="themeStore.menuIsFlippActive"
      :width="themeStore.menuCollapsed?'80%':'320px'"
      :placement="placement"
  >
    <CommonAside />
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

      .lang-switch {
        margin-right: 5px;

        .btn {
          color: white;
        }
      }

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