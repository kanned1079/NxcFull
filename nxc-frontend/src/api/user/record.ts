import instance from "@/axios/index"

export const handleGetAllMyActivateLog = async (userId: number, page: number, size: number) => {
    try {
        let {data} = await instance.get('/api/user/v1/activation', {
            params: {
                user_id: userId,
                page: page,
                size: size,
            }
        })
        return data


        // if (data.code === 200) {
            // activateRecordList.value = []
            // data.log.forEach((log: ActivateRecord) => activateRecordList.value.push(log))
            // pageCount.value = data.page_count
            // animated.value = true
        // }

    } catch (error: any) {
        console.log(error)
        // message.error("unknown err: " + error)
    }
}

export const handleUnbindById = async (userId: number, recordId: number) => {
    // animated.value = false
    try {
        let {data} = await instance.delete('/api/user/v1/activation', {
            params: {
                user_id: userId,
                activation_id: recordId,
            }
        })
            return data
    } catch (error: any) {
        console.log(error)
        // message.error("unknown err: " + error)
    }
}
