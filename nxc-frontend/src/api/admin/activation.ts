import instance from "@/axios/index";

export const handleGetAllActivationLog = async (page: number, size: number) => {
    try {
        let {data} = await instance.get('/api/admin/v1/activation', {
            params: {
                page: page,
                size: size,
            }
        });
        return data
    } catch (err: any) {
        console.log(err)
    }
}