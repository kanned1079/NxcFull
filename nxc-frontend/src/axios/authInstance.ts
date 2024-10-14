import axios from 'axios';

// axios 实例
// 这里的 authInstance 没有请求和相应拦截器 仅用于用户注册的邮件鉴权
const authInstance = axios.create( {
    baseURL: 'http://localhost:8081',
    timeout: 10000 // 设置超时时间
});

export default authInstance;
