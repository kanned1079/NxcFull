import instance from "@/axios/index";

export const handleGetAllActivationLog = async (
    page: number,
    size: number,
    email: string,
    valid: boolean,
    sort: string,
    ) => {
    try {
        let {data} = await instance.get('/api/admin/v1/activation', {
            params: {
                page: page,
                size: size,
                email: email,
                valid: valid,
                sort: sort,
            }
        });
        return data
    } catch (err: any) {
        console.log(err)
        return false
    }
}

export const handleShowKeyDetails = async (keyId: number) => {
    try {
        let {data} = await instance.get('/api/user/v1/key/details', {
            params: {
                key_id: keyId,
            }
        })
        return data
    } catch (err: any) {
        console.log(err)
        return false
    }
}