import {defineStore} from "pinia";
import {reactive} from "vue";
// defineStore中的参数<'文件名'>, {<配置对象>}
const useSiteInfo = defineStore('siteInfo', {
    state: () => {  // state 需要是一个函数 有返回值
        return {
            siteName: 'Nxc Networks',
            siteInfo: {
                app_name: '',
                app_description:'',
                app_url: '',
                force_https: false,
                logo: null,
                subscribe_url: '',
                tos_url: '',
                stop_register: false,
                trial_time: 0,
                trial_subscribe: '',
                subscribe_list: [{
                    label: 'Everybody\'s Got Something to Hide Except Me and My Monkey',
                    value: 'song0',
                    disabled: true
                },
                    {
                        label: 'Drive My Car',
                        value: 'song1'
                    },
                    {
                        label: 'Norwegian Wood',
                        value: 'song2'
                    },],
                currency: 'CNY',
                currency_symbol: 'null',

            }
        }
    }
})

export default useSiteInfo;