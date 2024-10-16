import {ref, reactive} from 'vue'
import {defineStore} from "pinia";
import {internalSelectionDark} from "naive-ui";
import {instance} from "@/axios";
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
    host?: string
    port?: number
    username?: string
    password?: string
    database?: number
}

interface RedisConfig {
    host: string
    port: number
    username: string
    password: string
    database: number
}

interface MqConfig {
    host: string
    port: number
    username: string
    password: string
}

interface ApiServerConfig {
    listen_addr: string
    listen_port : number
}



// interface Config {
//     EtcdConfig,
//     MySqlConfig,
//     RedisConfig,
//     ApiServerConfig,
// }

const useConfigStore = defineStore('configStore', () => {
    let etcdConfig = ref<EtcdConfig>({
        endpoints: '',
        port: 2379,
        username: '',
        password: '',
    })
    let mysqlConfig = ref<MySqlConfig>({
        protocol: 'tcp',
        host: '',
        port: 3306,
        username: '',
        password: '',
        database: '',
    })
    let redisConfig = ref<RedisConfig>({
        host: '',
        port: 6379,
        username: '',
        password: '',
        database: 0,
    })
    let apiServerConfig = ref<ApiServerConfig>({
        listen_addr: '',
        listen_port : 8080,
    })
    let mqConfig = ref<MqConfig>({
        host: '',
        port: 5672,
        username: '',
        password: '',
    })

    let etcdClientIsInitialized = ref<boolean>(false)


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
    // let submitConfig2Server = async () => {
    //     let postData = {
    //         etcd_config: {
    //             ...etcdConfig
    //         },
    //         config: {
    //             mysql_config: mysqlConfig.value,
    //             redis_config: redisConfig.value,
    //             api_server_config: apiServerConfig.value,
    //         }
    //     }
    //
    //     console.log(postData)
    // }

    let checkEtcd = async () => {
        try {
            let {data} = await instance.get('/api/v1/config/etcd/check')
            console.log('1', data.initialled)
            if (data.code === 200 && data.initialled) {
                console.log("Etcd已初始化")
                etcdClientIsInitialized.value = true
            } else {
                etcdClientIsInitialized.value = false
            }
        } catch (error: any) {
            console.log(error)
        }

    }






    return {
        etcdConfig,
        saveEtcdConfig2LocalStorage,
        loadEtcdConfigFromLocalStorage,
        mysqlConfig,
        redisConfig,
        apiServerConfig,
        mqConfig,
        // submitConfig2Server,
        etcdClientIsInitialized,
        checkEtcd
    }

}, {
    persist: true,
})

export default useConfigStore;