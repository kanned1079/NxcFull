import instance from "@/axios/index"

export const handleCreateMyInviteCode = async (userId: number) => {
    try {
        // animated.value = false
        let {data} = await instance.post('/api/user/v1/invite/code', {
            user_id: userId,
        })
       return data
    } catch (err: any) {
        console.log(err)
    }
}

export const handleGetInviteMsg = async () => {
    try {
        let {data} = await instance.get('/api/user/v1/invite/banner')
        return data
    } catch (err: any) {
        console.log(err)
    }
}

export const handleGetMyInviteCode = async (userId: number) => {
    try {
        let {data} = await instance.get('/api/user/v1/invite/code', {
            params: {
                user_id: userId
            }
        })
        return data
    } catch (err: any) {
        console.log(err)
    }
}

export const handleGetMyInvitedUserList = async (userId: number, page: number, size: number) => {
    try {
        let {data} = await instance.get('/api/user/v1/invite/users', {
            params: {
                user_id: userId,
                page: page,
                size: size,
            }
        })
        return data
    } catch (err: any) {
        console.log(err)
    }
}