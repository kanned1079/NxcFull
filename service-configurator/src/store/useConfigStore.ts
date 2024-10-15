import {ref} from 'vue'
import {defineStore} from "pinia";
import {internalSelectionDark} from "naive-ui";
// import {useRouter} from "vue-router";
// // import {useMessage} from "naive-ui"



interface EtcdConfig {
    endpoints: string
    port: number
    username?: string
    password?: string
}

interface MySqlConfig {
    protocol: string
    host: string
    port: number
    username: string
    password: string
    database: string
}

interface RedisConfig {
    host: string
    port: number
    username: string
    password: string
}

interface MqConfig {
    host: string
    port: number
    username: string
    password: string
}

interface ApiServerConfig {
    listen_ip: string
    port : number
}



interface Config {
    EtcdConfig,
    MySqlConfig,
    RedisConfig,
    ApiServerConfig,
}

const useConfigStore = defineStore('configStore', () => {
    let etcdConfig = ref<EtcdConfig>({})
    let mysqlConfig = ref<MySqlConfig>({
        protocol: 'tcp',
    })
    let redisConfig = ref<RedisConfig>({})
    let apiServerConfig = ref<ApiServerConfig>({})
    let mqConfig = ref<MqConfig>({})

    // saveEtcdConfig2LocalStorage 将etcd数据存储到本地存储
    let saveEtcdConfig2LocalStorage = () => {
        localStorage.setItem('etcdConfig', JSON.stringify(etcdConfig.value))
    }

    // 读取etcd的配置文件
    let loadEtcdConfigFromLocalStorage = () => {
        let storedEtcdConfig = localStorage.getItem('etcdConfig')
        if (storedEtcdConfig) {
            try {
                etcdConfig.value = JSON.parse(storedEtcdConfig)
            } catch (error) {
                console.error('Error parsing etcd config:', error)
            }
        }
    }

    // 提交所有的配置信息
    let submitConfig2Server = async () => {
        let postData = {
            etcd_config: {
                ...etcdConfig
            },
            config: {
                mysql_config: mysqlConfig.value,
                redis_config: redisConfig.value,
                api_server_config: apiServerConfig.value,
            }
        }

        console.log(postData)
    }




    return {
        etcdConfig,
        saveEtcdConfig2LocalStorage,
        loadEtcdConfigFromLocalStorage,
        mysqlConfig,
        redisConfig,
        apiServerConfig,
        mqConfig,
        submitConfig2Server
    }

}, {
    persist: false,
})

export default useConfigStore;