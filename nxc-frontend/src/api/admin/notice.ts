import instance from "@/axios";

type Notice = {
    id?: number | undefined;
    title: string;
    content: string;
    tags: string; 
    img_url?: string | undefined;
}

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

export const handleUpdateNoticeEnabled = async (id: number, isShow: boolean) => {
    try {
        let {data} = await instance.put('/api/admin/v1/notice/status', {
            id,
            is_show: isShow,
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleDeleteNoticeById = async (noticeId: number) => {
    try {
        let {data} = await instance.delete('/api/admin/v1/notice',
            {
                params: {notice_id: noticeId,},
            },
        )
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleAddNewNotice = async (notice: Notice) => {
    try {
        let {data} = await instance.post('/api/admin/v1/notice', {
            ...notice
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleEditNoticeByIdReq = async (notice: Notice) => {
    try {
        let {data} = await instance.put('/api/admin/v1/notice', {
            ...notice
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}