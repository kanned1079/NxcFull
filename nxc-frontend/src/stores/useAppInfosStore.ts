import {reactive, ref} from 'vue';
import {defineStore} from 'pinia';
import instance from "@/axios/authInstance"
import useThemeStore from './useThemeStore';

export interface RegisterPageConfig {
    app_name: string;
    app_sub_name: string;
    app_description: string;
    app_url: string;
    email_whitelist_suffix: boolean;
    email_gmail_limit_enable: boolean;
    is_email_verify: boolean;
    is_invite_force: boolean;
    is_recaptcha: boolean;
    logo: string;
    recaptcha_site_key: string;
    tos_url: string;
    stop_register: boolean;
}

export interface AppCommonConfig {
    app_name: string;
    app_sub_name: string;
    app_description: string;
    app_url: string;
    logo: string;
    user_bg: string;
    frontend_background_url: string;
    frontend_theme: string;
    currency: string;
    currency_symbol: string;
    stop_register: boolean;
    secure_path: string;
    safe_mode_enable: boolean;
}

const useAppInfosStore = defineStore('appInfosStore', () => {
    const registerPageConfig = ref<RegisterPageConfig>({
        app_name: 'Nxc Cloud International1',
        app_sub_name: '全球站点',
        app_description: '全球站点',
        app_url: 'http://localhost:5173',
        email_whitelist_suffix: true,
        email_gmail_limit_enable: true,
        is_email_verify: false,
        is_invite_force: true,
        is_recaptcha: false,
        logo: 'logo.svg',
        recaptcha_site_key: 'password',
        tos_url: 'https://ikanned.com:24444/',
        stop_register: false,
    });

    const appCommonConfig = ref<AppCommonConfig>({
        app_name: 'Nxc Cloud',
        app_sub_name: '全球站点',
        app_description: '穿过县境上长长的隧道，便是雪国。夜空下，大地一片莹白，火车在信号所前停下来。',
        app_url: 'http://localhost:5173',
        logo: 'https://ikanned.com:2444/d/Upload/NXC/links3.svg',
        user_bg: '',
        frontend_background_url: '',
        frontend_theme: '',
        currency: 'USDT',
        currency_symbol: '$',
        stop_register: false,   // false允许注册 true关闭注册
        secure_path: '',
        safe_mode_enable: false,
    })

    const getCommonConfig = async () => {
        try {
            let {data} = await instance.get('/api/app/v1/env', {
                params: {
                    lang: 'en',
                }
            })
            if (data.code == 200) {
                console.log(data)
                Object.assign(appCommonConfig.value, data.config)
                // appCommonConfig.value.app_name = data.config.app_name
                // appCommonConfig.value.app_sub_name = data.config.app_sub_name
                // appCommonConfig.value.app_description = data.config.app_description
                // appCommonConfig.value.app_url = data.config.app_url
                // appCommonConfig.value.logo = data.config.logo_url
                //
                // appCommonConfig.value.frontend_theme = data.config.frontend_theme
                // appCommonConfig.value.frontend_background_url = data.config.frontend_background_url
                // appCommonConfig.value.currency = data.config.currency
                // appCommonConfig.value.currency_symbol = data.config.currency_symbol
                // appCommonConfig.value.stop_register = data.config.stop_register

            }
        } catch (err: any) {

        }
    }

    const getRegisterPageConfig = async () => {
        console.log('獲取註冊頁面配置')
    }

    return {
        registerPageConfig,
        appCommonConfig,
        // subscribeList,
        getCommonConfig,
        getRegisterPageConfig
    };
}, {
    persist: false,
});

export default useAppInfosStore;
