import instance from "@/axios/index"
// import {computed} from "vue";


export const commitNewTicket = async (userId: number, newTicketDetails: any) => {
    try {
        let {data} = await instance.post('/api/user/v1/ticket', {
            user_id: userId,
            // ...newTicket.value
            ...newTicketDetails,
        })
        return data
    } catch (error: any) {
        console.log(error)
    }
}

export const getAllMyTickets = async (userId: number, page: number, size: number) => {
    try {
        let {data} = await instance.get('/api/user/v1/ticket', {
            params: {
                user_id: userId,
                page,
                size
            }
        })
        return data
    } catch (error: any) {
        console.log(error)
    }
}

export const closeTicket = async (userId: number, ticketId: number) => {
    try {
        let {data} = await instance.delete('/api/user/v1/ticket', {
            params: {
                user_id: userId,
                ticket_id: ticketId,
            }
        })
        return data
    } catch (error: any) {
        console.log(error)
    }
}