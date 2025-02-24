import instance from "@/axios";

interface Plan {
    id: number | null
    group_id?: number | null
    is_renew?: boolean
    is_sale?: boolean
    name: string
    sort: number | null
    capacity_limit: number | null
    residue: number | null
    describe?: string
    month_price?: number | null
    quarter_price?: number | null
    half_year_price?: number | null
    year_price?: number | null
    created_at?: string
    updated_at?: number | null
    deleted_at?: string
}

export const handleFetchPlanKv = async () => {
    try {
        let {data} = await instance.get('/api/admin/v1/plan/kv')
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleSubmitNewPlan = async (plan: Plan) => {
    try {
        let {data} = await instance.post('/api/admin/v1/plan', {
            ...plan,
        })
        return data;
    } catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleUpdatePlan = async (plan: Plan) => {

}