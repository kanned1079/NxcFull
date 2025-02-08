import instance from "@/axios/index"

export interface DocItem {
    title: string,
    category: string,
    language: string,
    md_text: string,
    sort: number,
}

export const handleGetAllDoc = async (page: number, size: number) => {
    try {
        let {data} = await instance.get('/api/admin/v1/document', {
            params: {
                page: page,
                size: size,
            }
        })
        return data
    } catch (error: any) {
        console.log(error)
        return false

    }
}

export const handleAddNewDoc = async (doc_data: DocItem) => {
    try {
        let {data} = await instance.post('/api/admin/v1/document', {
            ...doc_data,
        })
        return data
    } catch (err: any) {

    }
}

export const handleEditDocById = async (id: number, doc_data: DocItem) => {
    try {
        let {data} = await instance.put('/api/admin/v1/document', {
            id: id,  //  更新文档的id
            ...doc_data  // 更新的内容
        })
       return data
    } catch (error: any) {
        // message.error(error.toString)
        console.log(error)
        return false
    }
}

export const handleDeleteDocById = async (id: number) => {
    try {
        let {data} = await instance.delete('/api/admin/v1/document', {
            params: {
                id: id,
            }
        })
       return data
    } catch (error: any) {
        console.log(error)
        return false
    }
}

export const handleUpdateDocShow = async (id: number) => {
    try {
        let {data} = await instance.patch('/api/admin/v1/document', {
            id: id,
        })
        return data
    }catch (err: any) {
        console.log(err)
        return false
    }
}
