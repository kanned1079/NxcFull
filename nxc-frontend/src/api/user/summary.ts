import instance from "@/axios/index"

export const checkIsUserHaveOpenTickets = async (userId: number) => {
    try {
        let {data} = await instance.get('/api/user/v1/ticket/check', {
            params: {
                user_id: userId,
            }
        })
      return data
    } catch (error) {
        console.log(error)
    }
}

export const getAllNotices = async (isUser: boolean) => {
    try {
        let {data} = await instance.get('/api/user/v1/notice', {
            params: {
                is_user: isUser,
            }
        })
       return data
    } catch (error) {
        console.log(error)
    }
}

export const getActivePlanList = async (userId: number) => {
    try {
        let {data} = await instance.get('/api/user/v1/plan/summary/fetch', {
            params: {
                user_id: userId,
            }
        })
        return data
    } catch (error) {
        // haveActive.value = false
        console.log(error)
    }
}