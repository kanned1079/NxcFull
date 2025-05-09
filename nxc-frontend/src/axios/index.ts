import axios from 'axios';
import useUserInfoStore from "@/stores/useUserInfoStore";
import {config} from "@/config";

// axios 实例
const instance = axios.create( {
    baseURL: config.apiAddr.axiosAddr,
    // baseURL: 'http://gateway.orb.local:8081',
    timeout: 10000 // 设置超时时间
});


// instance.interceptors.request.use(config => {
//     const token = sessionStorage.getItem('token');
//     // console.log(token)
//     if (token) {
//         config.headers.Authorization = `Bearer ${token}`;
//     }
//     return config;
// }, error => {
//     return Promise.reject(error);
// });
//
// // 响应拦截器
// instance.interceptors.response.use(response => {
//     return response;
// }, error => {
//     if (error.response && error.response.status === 401) {
//         let userInfoStore = useUserInfoStore();
//         // token 过期
//         console.error('Token expired or invalid. Please log in again.');
//         userInfoStore.logout() // 登出操作
//
//     }
//     return Promise.reject(error);
// });

instance.interceptors.request.use(config => {
    const userInfoStore = useUserInfoStore()
    // console.log(token)
    if (userInfoStore.thisUser.token) {
        config.headers.Authorization = `Bearer ${userInfoStore.thisUser.token}`;
    } else {
        userInfoStore.logout()
    }
    return config;
}, error => {
    return Promise.reject(error);
});

instance.interceptors.response.use(response => {
    return response;
}, error => {
    if (error.response && error.response.status === 401) {
        let userInfoStore = useUserInfoStore();
        // token 过期
        console.error('Token expired or invalid. Please log in again.');
        userInfoStore.logout() // 登出操作
    }
    return Promise.reject(error);
});

// 请求拦截器
// instance.interceptors.request.use(config => {
//     const getCookie = (name: string) => {
//         const match = document.cookie.match(new RegExp('(^| )' + name + '=([^;]+)'));
//         return match ? match[2] : null;
//     };
//
//     const token = getCookie('token'); // 从 Cookie 获取 token
//     if (token) {
//         config.headers.Authorization = `Bearer ${token}`;
//     }
//     return config;
// }, error => {
//     return Promise.reject(error);
// });
//
// // 响应拦截器
// instance.interceptors.response.use(response => {
//     return response;
// }, error => {
//     if (error.response && error.response.status === 401) {
//         let userInfoStore = useUserInfoStore();
//         // token 过期
//         console.error('Token expired or invalid. Please log in again.');
//         // 登出操作
//         userInfoStore.logout()
//     }
//     return Promise.reject(error);
// });

export default instance;
