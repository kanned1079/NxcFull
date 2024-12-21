import instance from "@/axios/index"

// cancelOrder 取消订单
export const cancelOrder = async (userId: number, orderId: string) => {
    try {
        let {data} = await instance.delete('/api/user/v1/order', {
            params: {
                user_id: userId,
                order_id: orderId
            }
        })
        return data
    } catch (error: any) {
        console.log(error)
    }
}

// getAllMyOrders 获取用户所有的订单历史记录
export const getAllMyOrders = async (userId: number, page: number, size: number) => {
    try {
        let {data} = await instance.get('api/user/v1/orders', {
            params: {
                user_id: userId,
                page: page,
                size: size,
            }
        })
        return data
    } catch (error: any) {
        console.log(error)
    }

}