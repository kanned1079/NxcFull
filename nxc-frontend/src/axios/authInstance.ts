import axios from 'axios';

// axios 实例
const authInstance = axios.create( {
    baseURL: 'http://localhost:8080',
    timeout: 10000 // 设置超时时间
});

export default authInstance;
