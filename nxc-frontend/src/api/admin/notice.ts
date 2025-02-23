import instance from "@/axios";

export const handleFetchAllNotices = async (page: number, size: number, isUser: boolean) => {
    try {
        let {data} = await instance.get('/api/admin/v1/notice', {
            params: {
                page: page,
                size: size,
                is_user: isUser,
            }
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}