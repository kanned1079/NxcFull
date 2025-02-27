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
    appVersion: 'v0.7.9_patch41',
    codeGithub: `https://github.com/kanned1079/NxcFull`,
    author: 'kanned1079',
    apiAddr: {
        axiosAddr: 'http://localhost:8081',
        wsAddr: 'ws://localhost:8081'
    },
}
