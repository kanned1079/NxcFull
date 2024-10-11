import {reactive} from 'vue'
import {defineStore} from "pinia";
import instance from "@/axios";
// import {useMessage} from "naive-ui"


interface Notice {
    id: number,
    title: string,
    content: string,
    show: boolean,
    img_url?: string,
    tags?: string,
    created_at: string,
    updated_at?: string,
    deleted_at?: string,
}

const useNoticesStore = defineStore('NoticesStore', () => {



    let noticesArr = reactive<Notice[]>([])

    return {
        noticesArr,
        // getAllNotices,
        // deleteNotice,
        // aa
    }

}, {
    // persist: true,
})

export default useNoticesStore;