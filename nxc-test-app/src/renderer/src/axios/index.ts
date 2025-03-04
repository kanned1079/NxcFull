// import createAxios from 'axios'

import axios from "axios"

const instance = axios.create( {
  baseURL: 'http://localhost:8081',
  // baseURL: 'http://gateway.orb.local:8081',
  timeout: 4000 // 设置超时时间
});

export default instance
