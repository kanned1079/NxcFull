import {defineStore} from "pinia";
import {ref, reactive, computed, toRaw} from "vue"
// import {useRouter} from "vue-router";
import router from "@/router"

// const router = useRouter()

const useUserInfoStore = defineStore('userInfoS',() => {
    // 数据
    let isAuthed = ref<boolean>(false)
    let thisUser = reactive({
        id: 0,
        inviteUserId: 0,
        name: '',
        email: '',
        isAdmin: false,
        isStaff: false,
        balance: 0.00,
        lastLogin: '',
        lastLoginIp: '0.0.0.0',
        licenseActive: false,
        licenseExpiration: '',
        token: 'ewfesrflhweaifuhiesagfesrgfegfesgfvliehsguu',
    })

    // 方法
    // let getAuthed = computed(() => thisUser)
    let setAndSaveAuthStatus = (status: boolean) => {
        isAuthed.value = status
        sessionStorage.setItem('isAuthed', JSON.stringify(status))
    }

    // logout 登出 设置状态false 写入session 转回登陆页面
    let logout = () => {
        setAndSaveAuthStatus(false)
        let isAdminBak = thisUser.isAdmin
        sessionStorage.removeItem('tokeni')
        Object.assign(thisUser, {
            id: 0,
            inviteUserId: 0,
            name: '',
            email: '',
            isAdmin: false,
            isStaff: false,
            balance: 0.00,
            lastLogin: '',
            lastLoginIp: '0.0.0.0',
            licenseActive: false,
            licenseExpiration: '',
            token: '',
        });
        console.log('是否是管理员:', isAdminBak)
        router.push(isAdminBak ? '/admin/login' : '/login').catch(err => {
            console.error('Failed to navigate:', err);
        });
    }

    return {
        isAuthed,
        thisUser,
        setAndSaveAuthStatus,
        logout
    }

},{
    persist: true,
})

export default useUserInfoStore;