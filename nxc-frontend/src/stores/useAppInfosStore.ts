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

    // 每次用户打开首页必须的配置
    interface AppCommonConfig {
        app_name: string
        app_sub_name: string
        app_description: string
        app_url: string
        logo: string
        user_bg: string
        admin_bg: string
        currency: string
        currency_symbol: string
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

    let appCommonConfig: AppCommonConfig = {
        app_name: 'Nxc Cloud International',
        app_sub_name: '全球站点',
        app_description: '穿过县境上长长的隧道，便是雪国。夜空下，大地一片莹白，火车在信号所前停下来。 ',
        app_url: '',
        logo: 'https://ikanned.com:24444/d/Upload/NXC/links3.svg    ',
        user_bg: '',
        admin_bg: '',
        currency: 'pt',
        currency_symbol: '$',
    }

    let subscribeList = ref([
        {
            label: 'Everybody\'s Got Something to Hide Except Me and My Monkey',
            value: 'song0',
            disabled: true
        }
    ])


    return {
        registerPageConfig,
        appCommonConfig
        // getAllNotices,
        // deleteNotice,
        // aa
    }

}, {
    persist: true,
})

export default useAppInfosStore;