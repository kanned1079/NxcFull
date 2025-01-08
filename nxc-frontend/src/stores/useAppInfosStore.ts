import {reactive, ref} from 'vue';
import {defineStore} from 'pinia';
import instance from "@/axios/authInstance"
import useThemeStore from './useThemeStore';

export interface RegisterPageConfig {
    // app_name: string;
    // app_sub_name: string;
    // app_description: string;
    // app_url: string;
    email_whitelist_suffix: boolean;
    email_gmail_limit_enable: boolean;
    email_verify: boolean;
    stop_register: boolean;
    invite_require: boolean;
    recaptcha_enable: boolean;
    recaptcha_site_key: string;
    // logo: string;
    tos_url: string;

}

export interface AppCommonConfig {
    app_name: string;
    app_sub_name: string;
    app_description: string;
    app_url: string;
    logo_url: string;
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
        // app_name: 'Nxc Cloud International1',
        // app_sub_name: '全球站点',
        // app_description: '全球站点',
        // app_url: 'http://localhost:5173',
        email_whitelist_suffix: false,   // 邮箱白名单
        email_gmail_limit_enable: true, // 是否允许Gmail多别名
        email_verify: true,  // 是否开启邮件验证
        invite_require: true,  // 邀请码是否必填
        recaptcha_enable: false,
        recaptcha_site_key: '',
        // logo: 'logo.svg',
        tos_url: 'https://ikanned.com:24444/',
        stop_register: false,
    });

    const appCommonConfig = ref<AppCommonConfig>({
        app_name: '',
        app_sub_name: '',
        app_description: '',
        app_url: '',
        logo_url: '',
        user_bg: '',
        frontend_background_url: '',
        frontend_theme: '',
        currency: '',
        currency_symbol: '',
        stop_register: false,   // false允许注册 true关闭注册
        secure_path: '',
        safe_mode_enable: false,
    })

    // const getCommonConfig = async () => {
    //     try {
    //         let {data} = await instance.get('/api/app/v1/env/root', {
    //             params: {
    //                 lang: 'en',
    //             }
    //         })
    //         if (data.code == 200) {
    //             console.log(data)
    //             Object.assign(appCommonConfig.value, data.config)
    //             // appCommonConfig.value.app_name = data.config.app_name
    //             // appCommonConfig.value.app_sub_name = data.config.app_sub_name
    //             // appCommonConfig.value.app_description = data.config.app_description
    //             // appCommonConfig.value.app_url = data.config.app_url
    //             // appCommonConfig.value.logo = data.config.logo_url
    //             //
    //             // appCommonConfig.value.frontend_theme = data.config.frontend_theme
    //             // appCommonConfig.value.frontend_background_url = data.config.frontend_background_url
    //             // appCommonConfig.value.currency = data.config.currency
    //             // appCommonConfig.value.currency_symbol = data.config.currency_symbol
    //             // appCommonConfig.value.stop_register = data.config.stop_register
    //
    //         }
    //     } catch (err: any) {
    //
    //     }
    // }

    const getRegisterPageConfig = async () => {
        console.log('獲取註冊頁面配置')
    }

    return {
        registerPageConfig,
        appCommonConfig,
        // subscribeList,
        // getCommonConfig,
        getRegisterPageConfig
    };
}, {
    persist: false,
});

export default useAppInfosStore;
