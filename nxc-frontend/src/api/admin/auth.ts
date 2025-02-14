// import authInstance from "@/axios/authInstance";
import instance from "@/axios";
import {encodeToBase64} from "@/utils/encryptor";

export const handleAdminLoginFunc = async (email: string, password: string, role: 'admin' | 'user') => {
    try {
        let {data} = await instance.post('/api/admin/v1/login', {
            email: email,
            password: encodeToBase64(password),
            role: role, // 限制权限
        })
        return data
    } catch (err: any) {
        console.error(err)
    }
}