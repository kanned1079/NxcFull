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
    try {
        let {data} = await instance.put('/api/admin/v1/plan', {
            ...plan,
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleDeletePlanById = async (planId: number) => {
    try {
        let {data} = await instance.delete('/api/admin/v1/plan', {
            params: {
                plan_id: planId,
            }
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}
//
// export const handleUpdatePlanRowIsSale = async (id: number, val: boolean) => {
//     try {
//
//     } catch (err: any) {
//         console.error(err)
//         return false
//     }
// }

export const handleUpdatePlanRowIsSaleOrRenew = async (field: 'sale' | 'renew', id: number, val: boolean) => {
    try {
        // let {data} = await instance.put(field==='sale'?'api/admin/v1/plan/sale':'/api/admin/v1/plan/renew', {
        let {data} = await instance.put(`/api/admin/v1/plan/${field === 'sale' ? 'sale' : 'renew'}`, {
            id,
            is_sale: val,
            is_renew: val,
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}