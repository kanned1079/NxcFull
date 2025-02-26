import instance from "@/axios/index"

export const handleGetAllUsers = async (
    page: number,
    size: number,
    email: string,
) => {
    try {
        let {data} = await instance.get('/api/admin/v1/users', {
            params: {
                page: page,
                size: size,
                email: email,
            }
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleBlock2UnblockUserById = async (userId: number) => {
    try {
        let {data} = await instance.patch('/api/admin/v1/users', {
            user_id: userId,
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleAddNewUserReq = async (
    email: string,
    hashedPassword: string
) => {
    try {
        let {data} = await instance.post('/api/admin/v1/users', {
            email: email,
            password: hashedPassword,
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleEditUserInfoReq = async (
    id: number,
    password: string,
    inviteCode: string,
    isAdmin: boolean,
    isStaff: boolean,
    balance: number
) => {
    try {
        let {data} = await instance.put('/api/admin/v1/users', {
            id: id,
            password: password,
            invite_code: inviteCode,
            is_admin: isAdmin,
            is_staff: isStaff,
            balance: balance,
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}