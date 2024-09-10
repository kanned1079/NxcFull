import {reactive} from 'vue'
import {defineStore} from "pinia";
import instance from "@/axios";
// import {useMessage} from "naive-ui"


const useNoticesStore = defineStore('NoticesStore', () => {

    interface Notices {
        id: int,
        title: string,
        content: string,
        show: boolean,
        img_url?: string,
        tags?: string,
        created_at: string,
        updated_at?: string,
        deleted_at?: string,
    }

    let noticesArr = reactive([]<Notices>)

    // // 获取所有的通知
    // let getAllNotices = async () => {
    //     try {
    //         let {data} = await instance.get('/api/admin/get-all-notices')
    //         if (data.code === 200) {
    //             // console.log(data.notices)
    //             noticesArr = data.notices
    //             console.log('axios获取的数据', noticesArr)
    //             // createData()
    //         } else {
    //             message.error('获取失败')
    //         }
    //     } catch (err) {
    //         message.error('未知错误 ' + err)
    //     }
    // }
    //
    // // 删除一条通知
    // let deleteNotice = async (id: number) => {
    //     const message = useMessage()
    //     try {
    //         let {data} = await instance.delete('/api/admin/delete-notice',
    //             {
    //                 params: {notice_id: notice_id,},
    //             },
    //         )
    //         data.code === 200 ? message.success("删除成功") : message.error("出现错误")
    //     } catch (err) {
    //         message.error("奇怪的错误" + err)
    //     }
    //     // console.log(data)
    // }



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