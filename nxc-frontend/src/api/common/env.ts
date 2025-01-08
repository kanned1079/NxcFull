import authInstance from "@/axios/authInstance";
import instance from "@/axios/authInstance";
import useAppInfosStore from "@/stores/useAppInfosStore";

export const handleFetchRootRuntimeEnvConfig = async () => {
    const appInfos = useAppInfosStore()
    try {
        let {data} = await instance.get('/api/app/v1/env/root', {
            params: {
                lang: 'en',
            }
        })
        if (data.code === 200) {
            Object.assign(appInfos.appCommonConfig, data.config)
            return true
        }
        return false
    } catch (err: any) {
        console.log(err)
        return false
    }
}

export const handleFetchRegisterConfig = async (lang: string) => {
    const appInfos = useAppInfosStore()
    try {
        let {data} = await authInstance.get('/api/app/v1/env/register', {
            params: {
                lang: lang,
            }
        })
        if (data.code === 200) {
            Object.assign(appInfos.registerPageConfig, data.config)
            return true
        }
        return false
    } catch (err: any) {
        console.log(err)
        return false
    }
}

export const handleFetchWelcomeInfo = async (lang: string) => {
    try {
        let {data} = await authInstance.get('/api/app/v1/env/welcome', {
            params: {
                lang: lang,
            }
        })
        return data
    } catch (err: any) {
        console.log(err + '')
    }
}
