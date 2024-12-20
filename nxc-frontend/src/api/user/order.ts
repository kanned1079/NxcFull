import instance from "@/axios/index"

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
        // message.error("unknown err: " + error)
    }
}


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