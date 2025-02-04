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