import instance from "@/axios/index"

export const handleTestServerLatency = async () => {
    try {
        let {data} = await instance.get('/api/admin/v1/server/latency/test')
        return data.code === 200
    } catch (err: any) {
        console.log(err)
        return false
    }
}

export const handleFetchServerStatus = async (code: number, page: number, size: number) => {
    try {
        let {data} = await instance.get('/api/admin/v1/server/status/fetch', {
            params: {
                code: code,
                page: page,
                size: size
            }
        })
        return data
    } catch (err: any) {
        console.log(err)
    }
}

export const handleDeletePreviousLog = async () => {
    try {
        let {data} = await instance.delete('/api/admin/v1/server/log/delete')
        return data
    } catch (err: any) {

    }
}