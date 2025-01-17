import instance from "@/axios"

export const handleFetchAllAppDownloadList = async (lang?: string) => {
    try {
        let {data} = await instance.get('/api/user/v1/app', {
            params: {
                lang,
            }
        })
        return data
    } catch (err: any) {
        console.log(err)
    }
}