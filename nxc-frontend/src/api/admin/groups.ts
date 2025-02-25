import instance from "@/axios"

export const a1 = async (page: number, size: number) => {
    try {

    }catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleGetAllGroups = async (page: number, size: number) => {
    try {
        let {data} = await instance.get('http://localhost:8081/api/admin/v1/groups', {
            params: {
                page: page,
                size: size,
            }
        })
        return data
    }catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleGetAllGroupsKv = async () => {
    try {
        let {data} = await instance.get('/api/admin/v1/groups/kv')
        return data
    }catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleSubmitNewGroup = async (grpName: string) => {
    try {
        let {data} = await instance.post('/api/admin/v1/groups', {
            group_name: grpName,
        })
        return data
    }catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleUpdateGroup = async (id: number, name: string) => {
    try {
        let {data} = await instance.put('/api/admin/v1/groups', {
            id: id,
            name: name,
        })
        return data
    }catch (err: any) {
        console.error(err)
        return false
    }
}

export const handleDeleteGroup = async (id: number) => {
    try {
        let {data} = await instance.delete('/api/admin/v1/groups', {
            data: {
                id,
            }
        })
        return data
    }catch (err: any) {
        console.error(err)
        return false
    }
}

