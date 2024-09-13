import {reactive} from 'vue'
import {defineStore} from "pinia";
import instance from "@/axios";
// import {useMessage} from "naive-ui"


const useApiAddrStore = defineStore('apiAddrStore', () => {
    let apiAddr = reactive({
        user: {
            userLogin: '/api/user/v1/login',    // 用户登录
            sendAuthMail: '/api/user/v1/register/vc/get',    // 发送验证码
            verifyAuthMail: '/api/user/v1/register/vc/verify',    // 验证代码
            newUserRegister: '/api/user/v1/register/register',   // 新用户注册
            getAllNoticesList: '/api/user/v1/notices/get'
        },
        admin: {
            adminLogin: '/api/admin/v1/login',  // 管理员登录
            getSystemStatus: '/api/admin/v1/os/status/get',     // 获取系统信息
            saveSingleConfig: '/api/admin/v1/settings/set', // 保存单个设置
            getAllSystemConfig: '/api/admin/v1/settings/get',     // 获取所有系统设置
            getAllUsersList: '/api/admin/v1/users/get',     // 获取所有用户列表
            getAllNoticesList: '/api/admin/v1/notices/get',    // 获取所有通知列表
            addANotice: '/api/admin/v1/notice/add',     // 添加一条通知
            deleteANotice: '/api/admin/v1/notice/delete', // 删除一条通知
            // 邮件测试部分
            sendTestMail: '/api/admin/v1/mail/test'
        }
    })


    return {
        apiAddr,
    }

}, {
    persist: true,
})

export default useApiAddrStore;