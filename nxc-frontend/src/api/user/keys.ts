import instance from "@/axios";

export const handleGetAllMyKeys = async (userId: number, page: number, size: number) => {
    try {
        let {data} = await instance.get('/api/user/v1/keys', {
            params: {
                user_id: userId,
                page: page,
                size: size,
            }
        })
      return data
    } catch (error: any) {
        console.error(error)
        return false
    }
}