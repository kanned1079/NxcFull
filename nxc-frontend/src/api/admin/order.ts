import instance from "@/axios"

type Order = {
    id: number
    user_id: number
    order_id: string
    email: number
    plan_id: number
    group_id: number
    plan_name: string
    period: string
    coupon_id: number
    coupon_name: string
    status: number // 0新购 1续费 2修改订阅
    is_success: boolean
    is_finished: boolean
    price: number
    amount: number
    discount_amount: number
    paid_at: string
    payment_method: string
    created_at: string
    failure_reason: string
}

export const h1 = (page: number, size: number) => {
    try {

    } catch (error) {
        console.error(error)
        return false
    }
}

export const handleGetAllOrders = async (page: number, size: number, searchEmail: string, sort: string) => {
    try {
        let {data} = await instance.get('/api/admin/v1/order', {
            params: {
                page: page,
                size: size,
                search_email: searchEmail,
                sort: sort,
            }
        })
        return data
    } catch (error) {
        console.error(error)
        return false
    }
}

export const handleManualCancelOrderById = async (userId: number, orderId: string) => {
    try {
        let {data} = await instance.delete('/api/admin/v1/order', {
            params: {
                user_id: userId,
                order_id: orderId
            }
        })
        return data
    } catch (error) {
        console.error(error)
        return false
    }
}

export const handleManualPassOrderById = async (userId: number, orderId: string) => {
    try {
        let {data} = await instance.put('/api/admin/v1/order', {
            user_id: userId,
            order_id: orderId
        })
        return data
    } catch (error) {
        console.error(error)
        return false
    }
}