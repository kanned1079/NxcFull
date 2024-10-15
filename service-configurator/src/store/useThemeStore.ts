import {ref} from 'vue'
import {defineStore} from "pinia";
// import {useRouter} from "vue-router";
// // import {useMessage} from "naive-ui"


const useThemeStore = defineStore('themeStore', () => {
    let menuSelected = ref<string>('etcd')
    let path = ref<string>('/')
    let darkThemeEnabled = ref<boolean>(false)

    return {
        menuSelected,
        path,
        darkThemeEnabled
        // handleSelected,
    }

}, {
    persist: true,
})

export default useThemeStore;