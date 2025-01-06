import {computed} from "vue";
import authInstance from "@/axios/authInstance";
import {hashPassword} from "@/utils/encryptor";
import instance from "@/axios";

export const handleSendVerifyCode = async (email: string) => {
    try {
        let {data} = await authInstance.post('/api/user/v1/mail/register/get', {
            email: formValue.value.user.email
        })
        return data
    } catch (error) {
        console.log(error)
    }
}

export const handleVerifyCode = async (email: string, verifyCode: string) => {
    console.log('验证验证码是否正确')
    try {
        let {data} = await authInstance.post('/api/user/v1/mail/register/verify', {
            email: email,
            code: verifyCode
        })
        return data
    } catch (err: any) {
        console.log(error)
        return false
    }
}

export const handleFinalRegister = async (email: string, password: string, inviteUserId: string) => {
    try {
        let hashedPassword = await hashPassword(password.trim() as string)
        let {data} = await instance.post('/api/user/v1/register/register', {
            email: email,
            password: hashedPassword,
            invite_user_id: inviteUserId,
        })
        return data
    } catch (err: any) {
       console.log(err)
    }
}