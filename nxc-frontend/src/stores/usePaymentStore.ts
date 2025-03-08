import {ref} from 'vue'
import {defineStore} from "pinia";
// import instance from "@/axios";
import {handleGetAllPlans} from "@/api/user/plans";
// import {useMessage} from "naive-ui"
import useUserInfoStore from "@/stores/useUserInfoStore";

export interface Plan {
    id: number
    group_id?: number
    is_renew?: boolean
    is_sale?: boolean
    name: string
    capacity_limit: number
    sort: number,
    residue: number
    describe?: string
    month_price?: number
    quarter_price?: number
    half_year_price?: number
    year_price?: number
    created_at: string
    updated_at?: number
    deleted_at?: string
}

interface ConfirmOrder {
    code: number | null,
    order_id: number | null,
    plan_name: string | null,
    original_price: number | null,
    discount_amount: number | null,
    pay_price: number | null,
    period: string | null,
    created_at: string | null
}

interface Order {
    order_id: string
    plan_name: string
    period: string

    is_success: boolean
    is_finished: boolean
    failure_reason: string

    amount: number
    paid_at: any
    created_at: string
}

const usePaymentStore = defineStore('paymentStore', () => {
    const userInfoStore = useUserInfoStore()
    let plan_list = ref<Plan[]>([])

    let getAllPlans = async (page: number, size: number): Promise<{ success: boolean, page_count: number }> => {
        let data = await handleGetAllPlans(!userInfoStore.thisUser.isAdmin, page, size)
        if (data && data.code === 200) {
            plan_list.value = []
            if (data.plans)
                data.plans.forEach((item: Plan) => plan_list.value.push(item))
            return {
                success: true,
                page_count: data.page_count as number,
            }
        } else {
            console.log('err:', data.msg || '')
            return {
                success: false,
                page_count: 0,
            }
        }

    }

    let plan_id_selected = ref<number>(-1)

    let submittedOrderId = ref<string>('666666')

    let confirmOrder = ref<ConfirmOrder | undefined>()

    let orderDetail = ref<Order>({
        order_id: '',
        plan_name: '',
        period: '',
        is_success: false,
        is_finished: false,
        failure_reason: '',
        amount: 0.00,
        paid_at: '',
        created_at: '',
    })


    return {
        plan_list,
        getAllPlans,
        plan_id_selected,
        confirmOrder,
        submittedOrderId,
        orderDetail
    }

}, {
    persist: true,
})

export default usePaymentStore;