import instance from "@/axios/index"
import {encodeToBase64, hashPassword} from "@/utils/encryptor";
import {computed} from "vue";

// verifyOldPassword 验证旧密码是否正确
export const verifyOldPassword = async (email: string, oldPassword: string): Promise<boolean> => {
    try {
        let {data} = await instance.post('/api/user/v1/auth/passcode/verify', {
            email: email.trim(),
            // old_password: hashPassword(modelRef.value.old_password.trim() as string)
            old_password: encodeToBase64(oldPassword.trim()),
        });
        if (data.code === 200) return data.verified as boolean
        else return false
    } catch
        (error: any) {
        // notify('error', t('userProfile.oldPwdVerifiedFailure'), error.toString())
        console.log(error.toString());
        return false;
    }
}

// saveNewPassword 保存新的密码
export const saveNewPassword = async (email: string, newPassword: string) => {
    try {
        let hashedNewPassword = await hashPassword(newPassword)
        let {data} = await instance.post('/api/user/v1/auth/passcode/update', {
            email: email,
            // new_password: hashPassword(modelRef.value.new_password)
            new_password: hashedNewPassword
        });
        return data
    } catch (err: any) {
        // notify('error', t('userProfile.alertFailure'), error.toString())
        console.log(err)
        return null
    }

}


// handleSetup2FA 创建一个新的2FA新建请求
export const handleSetup2FA = async (userId: number, email: string) => {
    try {
        let {data} = await instance.post('/api/user/v1/auth/2fa/setup', {
            id: userId,
            email: email,
        })
        return data
    } catch (error: any) {
        console.log(error)
    }
}


// handleTest2FA 测试2FA可用性 如果可用则写入进数据库
export const handleTest2FA = async (userId: number, email: string, toFaCode: string) => {
    try {
        let {data} = await instance.post('/api/user/v1/auth/2fa/setup/test', {
            id: userId,
            email: email,
            two_fa_code: toFaCode.trim(),
        })
        return data
    } catch (error: any) {
        console.log(error)
    }
}

// handleCancelSetup2FA 删除redis中的新建临时2FA数据
export const handleCancelSetup2FA = async (email: string) => {
    try {
        let {data} = await instance.delete('/api/user/v1/auth/2fa/setup/cancel', {
            params: {
                // id: userInfoStore.thisUser.id,
                email: email,
            }
        })
        return data
        // showModal.value = false
    } catch (error: any) {
        console.log(error)
    }
}

// handleGet2FAStatus 测试2FA可用性 如果可用则写入进数据库
export const handleGet2FAStatus = async (userId: number) => {
    try {
        let {data} = await instance.get('/api/user/v1/auth/2fa/status', {
            params: {
                id: userId,
            }
        })
        return data
    } catch (error: any) {
        console.log(error)
    }
}

export const handleDisable2FA = async (email: string) => {
    try {
        let {data} = await instance.delete('/api/user/v1/auth/2fa/disable', {
            params: {
                // id: userInfoStore.thisUser.id,
                email: email,
            }
        })
        return data
    } catch (error: any) {
        console.log(error)
    }
}

export const handleDeleteMyAccount = async (userId: number, confirmed: boolean) => {
    try {
        let {data} = await instance.delete('/api/user/v1/user/delete', {
            params: {
                user_id: userId,
                confirmed: confirmed,
            },
        })
        return data
    } catch (err: any) {
        console.log(err)
    }
}