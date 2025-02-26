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
    const appVersion = ref<string>('')

    const registerPageConfig = ref<RegisterPageConfig>({
        email_whitelist_suffix: false,   // 邮箱白名单
        email_gmail_limit_enable: true, // 是否允许Gmail多别名
        email_verify: true,  // 是否开启邮件验证
        invite_require: true,  // 邀请码是否必填
        recaptcha_enable: false,
        recaptcha_site_key: '',
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

    // const dataSize = ref<{ pageSize: number, page: number }>({
    //     pageSize: 10,
    //     page: 1,
    // })

    const defaultTablePagination = ref<{page: number, size: number}>({
        page: 1,
        size: 10,
    })

    const updatePaginationSize = (newSize: number) => {
        defaultTablePagination.value.size = newSize
        localStorage.setItem('pagination', JSON.stringify({page: 1, size: newSize}))
    }

    return {
        registerPageConfig,
        appCommonConfig,
        appVersion,
        defaultTablePagination,
        updatePaginationSize,
    };
}, {
    persist: false,
});

export default useAppInfosStore;
