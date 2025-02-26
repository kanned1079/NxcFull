import instance from "@/axios/index"

export const handleGetAllTicket = async (
    pendingPage: number,
    pendingSize: number,
    finishedPage: number,
    finishedSize: number,
) => {
    try {
        let {data} = await instance.get('/api/admin/v1/ticket', {
            params: {
                pending_page: pendingPage,
                pending_size: pendingSize,
                finished_page: finishedPage,
                finished_size: finishedSize,
            }
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleCloseTicketById = async (userId: number, ticketId: number) => {
    try {
        let {data} = await instance.delete('/api/user/v1/ticket', {
            params: {
                user_id: userId,
                ticket_id: ticketId,
            }
        })
        return data
    } catch (err: any) {
        console.error(err)
        return false
    }
}