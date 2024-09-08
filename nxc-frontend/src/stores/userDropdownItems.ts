import {defineStore} from "pinia";
import renderIcon from "@/utils/iconFormator";
import {
    Pencil as EditIcon,
    LogOutOutline as LogoutIcon,
    PersonCircleOutline as UserIcon
} from '@vicons/ionicons5'
// defineStore中的参数<'文件名'>, {<配置对象>}
const useUserDropDown = defineStore('userdropdown', {
    state: () => {  // state 需要是一个函数 有返回值
        return {
            options: [
                {
                    label: '用户资料',
                    key: 'profile',
                    icon: renderIcon(UserIcon)
                },
                {
                    label: '编辑用户资料',
                    key: 'editProfile',
                    icon: renderIcon(EditIcon)
                },
                {
                    label: '退出登录',
                    key: 'logout',
                    icon: renderIcon(LogoutIcon)
                }
            ]
        }
    },
})

export default useUserDropDown;