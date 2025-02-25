import instance from "@/axios/index"

interface CouponInst {
    id?: number
    name: string
    code: string
    percent_off: number | null
    start_time: number | null
    end_time: number | null
    per_user_limit: number | null
    capacity: number | null
    residue: number | null
    plan_limit: number | null
}

export const handleGetAllCoupon = async (page: number, size: number) => {
    try {
        let {data} = await instance.get('/api/admin/v1/coupon', {
            params: {
                page: page,
                size: size,
            }
        })
        return data
    } catch (err: any) {
        console.error(err)
    }
}

export const handleActivateCouponById = async (id: number, val: boolean) => {
    try {
        let {data} = await instance.put('http://localhost:8081/api/admin/v1/coupon/status', {
            id: id,
            status: val
        });
        return data
    } catch (err: any) {
        console.error(err)
    }
}

export const handleUpdateCoupon = async (coupon: CouponInst) => {
    try {
        let {data} = await instance.put('/api/admin/v1/coupon', coupon);
        return data
    } catch (err: any) {
        console.error(err)
    }
}

export const handleDeleteCouponById = async (couponId: number) => {
    try {
        let {data} = await instance.delete('/api/admin/v1/coupon', {
            params: {
                coupon_id: couponId,
            }
        })
        return data
    } catch (err: any) {
        console.error(err)
    }
}