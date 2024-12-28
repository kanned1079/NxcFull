import instance from "@/axios/index"

// getAllDocuments 获取所有的文档列表
export const getAllDocuments = async (lang: string, find?: string) => {
    // animated.value = false
    try {
        let {data} = await instance.get('http://localhost:8081/api/user/v1/document', {
            params: {
                lang: lang,
                find: find,
            }
        })
        return data
    } catch (error: any) {
        console.log(error)
    }
}