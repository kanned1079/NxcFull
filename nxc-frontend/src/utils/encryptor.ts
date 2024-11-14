// // encodeToBase64 将密码进行base64加密
// import CryptoJS from "crypto-js";
//
// let encodeToBase64 = (str: string): string => {
//     const utf8Encode = new TextEncoder();
//     const encoded = utf8Encode.encode(str);
//     const base64Encoded = btoa(String.fromCharCode(...encoded));
//     return base64Encoded;
// }
//
// let hashPassword = (password: string): string => {
//     return CryptoJS.SHA256(password).toString();
// }
//
// export {
//     encodeToBase64,
//     hashPassword,
// }

import bcrypt from 'bcryptjs';

// import { hash, compare } from 'bcryptjs';

// 加密密码函数，使用 bcryptjs
// let hashPassword = async (password: string): string => {
//     const saltRounds = 10; // bcrypt 的计算强度
//     const hashedPassword = await bcrypt.hash(password, saltRounds); // 直接使用密码和轮数生成哈希值
//     console.log(hashedPassword);
//     return hashedPassword.toString();
// }

let hashPassword = async (password: string): Promise<string> => {
    const saltRounds = 10; // bcrypt 的计算强度
    const hashedPassword = await bcrypt.hash(password, saltRounds); // 直接使用密码和轮数生成哈希值
    console.log(hashedPassword);
    return hashedPassword.toString();
}


// 验证密码函数
let comparePassword = async (password: string, hash: string): Promise<boolean> => {
    return await bcrypt.compare(password, hash); // 比较用户输入的密码与已存储的哈希值
}

// 明文转 Base64
let encodeToBase64 = (plainText: string): string => {
    // 将明文转换为 Base64 编码
    return btoa(plainText);
}

export {
    hashPassword,
    comparePassword,
    encodeToBase64,
};