import {reactive} from 'vue'
import {defineStore} from "pinia";
import instance from "@/axios";
// import {useMessage} from "naive-ui"


const useAppInfosStore = defineStore('appInfosStore', () => {

    interface RegisterPageConfig{
        app_name: string
        app_description: string
        app_url: string
        email_whitelist_suffix: boolean
        is_email_verify: boolean
        is_invite_force: boolean
        is_recaptcha: boolean
        logo: string
        recaptcha_site_key: string
        tos_url: string
    }

    let registerPageConfig = reactive<RegisterPageConfig>({
        app_name: 'Nxc Cloud International',
        app_description: '全球站点',
        app_url: 'https://ikanned.com/',
        email_whitelist_suffix: false,
        is_email_verify: true,  // 启用邮件验证
        is_invite_force: true,  // 是否需要邀请码
        is_recaptcha: false,    // 启用google验证
        logo: 'logo.svg',   // 头部logo
        recaptcha_site_key: 'password',
        tos_url: 'https://ikanned.com:24444/',  // 用户条款url
    })


    return {
        registerPageConfig,
        // getAllNotices,
        // deleteNotice,
        // aa
    }

}, {
    persist: true,
})

export default useAppInfosStore;