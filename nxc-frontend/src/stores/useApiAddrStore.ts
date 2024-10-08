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
            getAllNoticesList: '/api/user/v1/notices/get',  // 获取所有的通知
            getAllDocumentList: '/api/user/v1/document/get', // 获取所有的文档列表
            getAllPlanList: '/api/user/v1/plan/get', // 获取所有的订阅计划列表
            verifyCoupon: '/api/user/v1/coupon/verify',
            saveOrder: '/api/user/v1/subscribe/save',  // 提交订单
            verifyOldPassword: '/api/user/v1/auth/passcode/verify',
            applyNewPassword: '/api/user/v1/auth/passcode/update',
            getMyPlanList: '/api/user/v1/plan/info/fetch',
            getAllMyKeys: '/api/user/v1/keys/fetch',

        },
        admin: {
            adminLogin: '/api/admin/v1/login',  // 管理员登录
            getSystemStatus: '/api/admin/v1/os/status/get',     // 获取系统信息
            saveSingleConfig: '/api/admin/v1/settings/set', // 保存单个设置
            getAllSystemConfig: '/api/admin/v1/settings/get',     // 获取所有系统设置
            getAllUsersList: '/api/admin/v1/users/get',     // 获取所有用户列表
            getAllNoticesList: '/api/admin/v1/Document/get',    // 获取所有通知列表

            addANotice: '/api/admin/v1/notice/add',     // 添加一条通知
            deleteANotice: '/api/admin/v1/notice/delete', // 删除一条通知
            // 邮件测试部分
            sendTestMail: '/api/admin/v1/mail/test',
            // 添加新文档
            addNewDocument: '/api/admin/v1/document/add',
            addNewCoupon: '/api/admin/v1/coupon/add',   // 添加新的优惠券
            addNewPrivilegeGroup: '/api/admin/v1/groups',   // 添加新的权限组
            getAllPrivilegeGroups: '/api/admin/v1/groups', // 获取所有的权限组列表
            updatePrivilegeGroup: '/api/admin/v1/groups',
            deletePrivilegeGroup: '/api/admin/v1/groups/delete',
            addNewPlan: '/api/admin/v1/plans/add',   // 添加新的订阅信息
            getAllCoupons: '/api/admin/v1/coupon/fetch',    // 获取所有的优惠券列表
            updateCoupon: '/api/admin/v1/coupon/update', // 更新优惠券信息
            updateCouponStatus: '/api/admin/v1/coupon/status/update',    // 优惠券是否启用
            deleteCoupon: '/api/admin/v1/coupon/delete',    // 删除优惠券
            getPlanKVList: 'api/admin/v1/plans/kv/fetch',
        }
    })


    return {
        apiAddr,
    }

}, {
    persist: true,
})

export default useApiAddrStore;