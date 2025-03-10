import {reactive} from 'vue'
import {defineStore} from "pinia";
import instance from "@/axios";
import useApiAddrStore from "@/stores/useApiAddrStore";

interface SiteSettings {
    app_name: string;
    app_sub_name: string;
    app_description: string;
    app_url: string;
    'force_https': boolean;
    logo_url: string;
    subscribe_url: string;
    tos_url: string;
    stop_register: boolean;
    invite_require: boolean;
    trial_time: number;
    trial_subscribe: number;
    currency: string;
    currency_symbol: string;
    // [key: any]: any;
}

interface SecuritySettings {
    email_verify: boolean;
    email_gmail_limit_enable: boolean;
    safe_mode_enable: boolean;
    secure_path: string;
    email_whitelist_enable: boolean;
    recaptcha_enable: boolean;
    recaptcha_site_key: string;
    ip_register_limit_enable: boolean;
    ip_register_limit_times: number;
    ip_register_lock_time: number;
}

interface FrontendSettings {
    frontend_theme_sidebar: boolean;
    frontend_theme_header: boolean;
    frontend_theme: string;
    frontend_background_url: string;
}

// interface SubscribeSettings {
//     user_modify_enable: boolean;
//     show_info_in_sub: boolean;
// }

interface InviteRebateSettings {
    invite_rebate_enable: boolean
    invite_rebate_rate: number
    discount_info: string
    invite_info: string
}

interface ServerSettings {
    server_token: string;
    server_pull_interval: number;
    server_push_interval: number;
}

interface WelcomeSettings {
    app_sub_description: string
    why_choose_us_hint: string
    bilibili_official_link: string
    youtube_official_link: string
    instagram_link: string
    wechat_official_link: string
    filing_number: string
    page_suffix: string
}

interface SendmailSettings {
    email_host: string;
    email_port: number;
    email_encryption: string;
    email_username: string;
    email_password: string;
    email_from_address: string;
    email_template: string;
}

interface NoticeSettings {
    notice_name: string;
    bark_host: string;
    bark_group: string;
}

interface MyAppSettings {
    download_enabled: boolean;
    win_download: string;
    osx_download: string;
    linux_download: string;
    android_download: string;
    ios_download: string;
}

interface Settings {
    site: SiteSettings;
    security: SecuritySettings;
    frontend: FrontendSettings;
    // subscribe: SubscribeSettings;
    invite: InviteRebateSettings;
    server: ServerSettings;
    welcome: WelcomeSettings;
    sendmail: SendmailSettings;
    notice: NoticeSettings;
    my_app: MyAppSettings;
}

const useSettingStore = defineStore('SettingStore', () => {
    // 管理员系统设置界面所有设置

    const settings = reactive({
        site: {
            app_name: '',           // 站点名称
            app_sub_name: '',   // 站点副标题
            app_description: '',    // 站点描述
            app_url: '',            // 站点网址
            force_https: false,     // 强制HTTPS
            logo_url: '',           // LOGO
            subscribe_url: '',      // 订阅URL
            tos_url: '',            // 用户条款TOS
            stop_register: false,   // 停止新用户注册
            invite_require: false,  // 邀请码必填
            trial_time: 6,          // 试用时间
            trial_subscribe: 1,    // 注册试用
            currency: '',           // 货币单位
            currency_symbol: '',    // 货币符号
        } as SiteSettings,
        security: {
            email_verify: true,                 // 邮箱验证
            email_gmail_limit_enable: false,    // 禁止使用Gmail多别名
            safe_mode_enable: false,            // 安全模式
            secure_path: '/admin',              // 后台路径
            email_whitelist_enable: false,      // 邮箱后缀白名单
            recaptcha_enable: false,            // 启用Google reCAPTCHA
            recaptcha_site_key: '',
            ip_register_limit_enable: false,    // IP注册限制
            ip_register_limit_times: 5,         // 显示次数
            ip_register_lock_time: 120,         // 惩罚时间
        },
        frontend: {
            frontend_theme_sidebar: false,  // 边栏风格
            frontend_theme_header: false,   // 头部风格
            frontend_theme: 'bambooGreen',  // 主题色
            frontend_background_url: '',    // 背景
        },
        // subscribe: {
        //     user_modify_enable: true,   // 允许用户修改订阅
        //     show_info_in_sub: false,    // 在订阅中展示订阅信息
        // },
        invite: {
            invite_rebate_enable: false,
            invite_rebate_rate: 0.00,
            discount_info: '',
        },
        server: {
            server_token: '@zeBw2cSe6V^kCrz3uJQSd=FJU', // 通讯密钥
            server_pull_interval: 0,                    // 节点拉取动作轮询间隔
            server_push_interval: 0,                    // 节点推送动作轮询间隔
        },
        welcome: {
            app_sub_description: '',
            why_choose_us_hint: '',
            bilibili_official_link: '',
            youtube_official_link: '',
            instagram_link: '',
            wechat_official_link: '',
            filing_number: '',
            page_suffix: '',
        },
        sendmail: {
            email_host: '',         // SMTP服务器地址
            email_port: 465,        // SMTP服务端口
            email_encryption: '',   // SMTP加密方式
            email_username: '',     // SMTP账号
            email_password: '',     // SMTP密码
            email_from_address: '', // 发件地址
            email_template: ''      // 邮件模板
        },
        notice: {
            notice_name: '',    // 显示名称
            bark_host: '',      // bark接入点
            bark_group: '',     // Bark群组
        },
        my_app: {
            download_enabled: false,
            win_download: '',       // windows端软件下载地址
            osx_download: '',       // osx端软件下载地址
            linux_download: '',
            android_download: '',   // 安卓端软件下载地址
            ios_download: '',
        }
    } as Settings)

    // 从服务器拉取数据
    // let getSetting = async () => {
    //     let {data} = await instance.get('')
    //     if (data.code == 200) {
    //         console.log(data)
    //     }
    // }

    // let saveSetting = async () => {
    //     let apiAddrStore = useApiAddrStore()
    //     console.log('保存数据到数据库')
    //     let { data } = await instance.post(apiAddrStore.apiAddr.admin.getAllSystemConfig, settings)
    //     console.log('设置保存到服务器的返回结果', data)
    //     // console.log(JSON.stringify(settings))
    // }

    const saveOption = async (category: string, key: string, value: any) => {
        // let apiAddrStore = useApiAddrStore()
        // console.log('保存单个键值到数据库', key, value)
        let {data} = await instance.put('/api/admin/v1/setting', {
            category,
            key,
            value
        }, {
            headers: {
                'Content-Type': 'application/json',
            }
        })
        console.log(data)
        // if (data.code === 200) return true; return data.msg || ''
        return data || false
        // console.log('单个键值的返回结果', data)
    }

    let loadSetting = async () => {
        // let apiAddrStore = useApiAddrStore()
        try {
            console.log('从数据库读取配置')
            let {data} = await instance.get('/api/admin/v1/setting')
            console.log(data)
            if (data.code == 200) {
                Object.assign(settings, data.settings)

            }
        } catch (error: any) {
            console.error(error.toString())
        }
    }

    return {
        settings,
        // saveSetting,
        loadSetting,
        saveOption
    }

})

export default useSettingStore;