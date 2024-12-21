import instance from "@/axios/index"

// handleCreateMyInviteCode 创建邀请码请求
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

// handleGetInviteMsg 获取头部邀请的返利信息
export const handleGetInviteMsg = async () => {
    try {
        let {data} = await instance.get('/api/user/v1/invite/banner')
        return data
    } catch (err: any) {
        console.log(err)
    }
}

// handleGetMyInviteCode 获取用户的邀请码
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

// handleGetMyInvitedUserList 获取该用户所邀请的用户列表
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