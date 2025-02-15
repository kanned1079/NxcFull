import instance from "@/axios";

export const handleFetchPlanKv = async () => {
    try {
        let {data} = await instance.get('/api/admin/v1/plan/kv')
        return data
    } catch (err: any) {
        console.error(err)
    }
}