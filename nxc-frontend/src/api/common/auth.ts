import instance from "@/axios";
import {encodeToBase64} from "@/utils/encryptor";

export const handleUserLogin = async (
    email: string,
    password: string,
    role: "user" | "admin",
    two_fa_code?: string,
) => {
    try {
        let {data} = await instance.post('/api/user/v1/login', {
            email: email.trim(),
            password: encodeToBase64(password.trim()),
            two_fa_code: two_fa_code || '',
            // password: hashedPwd,
            role: role,
        })
        return data
    } catch (err: any) {
        console.log(err)
        return false
    }
}