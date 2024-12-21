import instance from "@/axios";

// queryOrderInfoAndStatus 请求提交的订单的状态
export const queryOrderInfoAndStatus = async (userId: number, orderId: string) => {
    try {
        let {data} = await instance.get('/api/user/v1/order/status', {
            params: {
                user_id: userId,
                order_id: orderId.trim(),
            }
        });
       return data
// pass
    } catch (error: any) {
        console.log(error);
    }
}

// placeOrder 确认下单
export const placeOrder = async (userId: number, orderId: string) => {
    try {
        let {data} = await instance.put('/api/user/v1/order', {
            user_id: userId,
            order_id: orderId.trim(),
        })
        return data
    } catch (error: any) {
        console.log(error)
        return null
    }
}

// cancelOrder 取消当前的订单
export const cancelOrder = async (userId: number, orderId: string) => {
    try {
        let {data} = await instance.delete('/api/user/v1/order', {
            params: {
                user_id: userId,
                order_id: orderId.trim(),
            }
        })
        return data
    } catch (error: any) {
        console.log(error)
        return error
    }
}
