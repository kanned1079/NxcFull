import {defineStore} from "pinia";
import {ref, reactive, computed, toRaw} from "vue"
// import {useRouter} from "vue-router";
import router from "@/router"
import instance from "@/axios";

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
        // licenseActive: false,
        // licenseExpiration: '',
        // licenseId: 0,
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
        sessionStorage.removeItem('token')
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
            // licenseActive: false,
            // licenseExpiration: '',
            // licenseId: 0,
            token: '',
        });
        console.log('是否是管理员:', isAdminBak)
        router.push(isAdminBak ? '/admin/login' : '/login').catch(err => {
            console.error('Failed to navigate:', err);
        });
    }


    let updateUserInfo = async ():Promise<boolean> => {
        try {
            let {data} = await instance.get('/api/user/v1/user', {
                params: {
                    user_id: thisUser.id
                }
            })
            if ((data.code as number) === 200) {
                thisUser.balance = data.user_info.balance
                console.log('ok')
                return true
            } else {
                return false
            }
        } catch (error: any) {
            console.log(error)
            return false
        }
    }

    return {
        isAuthed,
        thisUser,
        setAndSaveAuthStatus,
        updateUserInfo,
        logout
    }

},{
    persist: true,
})

export default useUserInfoStore;