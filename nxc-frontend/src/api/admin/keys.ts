import instance from "@/axios/index"

export const handleGetAllKeysByAdmin = async (page: number, size: number, search_key_id?: number, search_key_content?: string) => {
    try {
        let {data} = await instance.get('/api/admin/v1/keys', {
            params: {
                page: page,
                size: size,
                key_id: search_key_id || 0,
                key_content: search_key_content || '',
            }
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleBlockKeyById = async (userId: number, keyId: number) => {
    try {
        let {data} = await instance.delete('/api/admin/v1/keys', {
            params: {
                user_id: userId,
                key_id: keyId,
            }
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}