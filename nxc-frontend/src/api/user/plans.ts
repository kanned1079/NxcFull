import instance from "@/axios/index"

export const handleGetAllPlans = async (isUser: boolean, page: number, size: number) => {
    // let apiAddrStore = useApiAddrStore()
    try {
        let {data} = await instance.get('/api/user/v1/plan', {
            params: {
                is_user: isUser,
                // is_user: false
                page: page,
                size: size,
            }
        })
        return data
    } catch (error) {
        console.log(error)
        return null
    }
}