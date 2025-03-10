type BasicEnv = {
    appVersion: string, // 前端页面版本号
    author: string,
    codeGithub: string,
    apiAddr: {  // 注意
        axiosAddr: string,  // axios请求的后端地址
        wsAddr: string, // websocket请求的后端地址
    }
}

export const config: BasicEnv = {
    appVersion: 'v0.8.1_patch7',
    codeGithub: `https://github.com/kanned1079/NxcFull`,
    author: 'kanned1079',
    apiAddr: {
        // axiosAddr: 'http://localhost:8081',
        // wsAddr: 'ws://localhost:8081',
        // axiosAddr: 'https://192.168.0.243:8081',
        // wsAddr: 'wss://192.168.0.243:8081',
        axiosAddr: 'https://ikanned.com:27081',
        wsAddr: 'wss://ikanned.com:27081'
    },
}
