// encodeToBase64 将密码进行base64加密
import CryptoJS from "crypto-js";

let encodeToBase64 = (str: string): string => {
    const utf8Encode = new TextEncoder();
    const encoded = utf8Encode.encode(str);
    const base64Encoded = btoa(String.fromCharCode(...encoded));
    return base64Encoded;
}

let hashPassword = (password: string): string => {
    return CryptoJS.SHA256(password).toString();
}

export {
    encodeToBase64,
    hashPassword,
}