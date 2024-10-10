import {ref} from 'vue'
import {defineStore} from "pinia";
import instance from "@/axios";
// import {useMessage} from "naive-ui"
import useApiAddrStore from "@/stores/useApiAddrStore";

export interface Plan {
    id: number
    group_id?: number
    is_renew?: boolean
    is_sale?: boolean
    name: string
    capacity: number
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

const usePaymentStore = defineStore('paymentStore', () => {
    let plan_list = ref<Plan[]>([])

    let getAllPlans = async () => {
        let apiAddrStore = useApiAddrStore()
        try {
            plan_list.value = []
            let {data} = await instance.get('http://localhost:8081/api/user/v1/plan')
            if (data.code === 200) {
                data.plans.forEach((item: Plan) => plan_list.value.push(item))
                // plan_list.value = data.plans
                console.log(plan_list.value)
            } else {
                console.log('获取订阅数据失败')
            }
        } catch (error) {
            console.log(error)
        }
    }

    let plan_id_selected = ref<number>(-1)

    let confirmOrder = ref<ConfirmOrder | undefined>()

    return {
        plan_list,
        getAllPlans,
        plan_id_selected,
        confirmOrder
    }

}, {
    persist: true,
})

export default usePaymentStore;