import {reactive, ref} from 'vue';
import {defineStore} from 'pinia';

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
}

export interface AppCommonConfig {
    app_name: string;
    app_sub_name: string;
    app_description: string;
    app_url: string;
    logo: string;
    user_bg: string;
    admin_bg: string;
    currency: string;
    currency_symbol: string;
    stop_register: boolean;
}

const useAppInfosStore = defineStore('appInfosStore', () => {
    const registerPageConfig = reactive<RegisterPageConfig>({
        app_name: 'Nxc Cloud International',
        app_sub_name: '全球站点',
        app_description: '全球站点',
        app_url: 'http://localhost:5173',
        email_whitelist_suffix: false,
        is_email_verify: true,
        email_gmail_limit_enable: false,
        is_invite_force: true,
        is_recaptcha: false,
        logo: 'logo.svg',
        recaptcha_site_key: 'password',
        tos_url: 'https://ikanned.com:24444/',
    });

    const appCommonConfig = ref<AppCommonConfig>({
        app_name: 'Nxc Cloud',
        app_sub_name: '全球站点',
        app_description: '穿过县境上长长的隧道，便是雪国。夜空下，大地一片莹白，火车在信号所前停下来。',
        app_url: 'http://localhost:5173',
        logo: 'https://ikanned.com:2444/d/Upload/NXC/links3.svg',
        user_bg: '',
        admin_bg: '',
        currency: 'USDT',
        currency_symbol: '$',
        stop_register: false,   // false允许注册 true关闭注册
    })

    return {
        registerPageConfig,
        appCommonConfig,
        // subscribeList,
    };
}, {
    persist: true,
});

export default useAppInfosStore;
