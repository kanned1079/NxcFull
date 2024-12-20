import instance from "@/axios";

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

export const placeOrder = async (userId: number, orderId: string) => {
    // 点击提交按钮立刻停止定时器
    // intervalId.value ? clearInterval(intervalId.value) : null
    try {
        let {data} = await instance.put('/api/user/v1/order', {
            user_id: userId,
            order_id: orderId.trim(),
        })
        return data
    } catch (error: any) {
        console.log(error)
        // router.back()
        return null
    }
}

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
