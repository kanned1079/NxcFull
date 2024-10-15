import {ref} from 'vue'
import {defineStore} from "pinia";
// import {useRouter} from "vue-router";
// // import {useMessage} from "naive-ui"


const useThemeStore = defineStore('themeStore', () => {
    let menuSelected = ref<string>('etcd')
    let path = ref<string>('/')

    return {
        menuSelected,
        path,
        // handleSelected,
    }

}, {
    persist: true,
})

export default useThemeStore;