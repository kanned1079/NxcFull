import instance from "@/axios/index";

export const verifyTicket = async (couponCode: string, planId: number, userId: number) => {
    try {
        let {data} = await instance.post('http://localhost:8081/api/user/v1/coupon/verify', {
            coupon_code: couponCode,  // 优惠券码
            plan_id: planId,               // 订阅计划id
            user_id: userId,  // 用户id 判断是否使用过
        });
        return data
    } catch (err: any) {
        // 处理网络错误或后端500错误
        console.log(err)
        // couponInfo.value.verified = false
        // await notify('error', t('newSettlement.notify.couponInvalid'), error.toString());

    }
}

export const saveOrder = async (planId: number, userId: number, couponId: number, period: string) => {
    console.log('提交订单准备支付')
    try {
        let {data} = await instance.post('/api/user/v1/order', {
            plan_id: planId,
            user_id: userId,
            coupon_id: couponId,
            period: period || 'month',
        })
       return data
    } catch (error: any) {
        console.log(error)
    }
}