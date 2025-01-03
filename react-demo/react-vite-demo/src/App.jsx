import {create} from "zustand"
import axios from "axios"

import useStore from "./zustand"

// 拆分
// const createCounterStore = (set) => {
//     return {
//         count: 0,
//         inc: () => {
//             set((state) => ({count: state.count + 1}))  // 基于原数据进行修改
//         },
//         incd: () => {
//             set({count: 100})   // 不需要原数据
//         },
//     }
// }

// const createChannelStore = (set) => {
//     return {
//         channelList: [],
//         handleFetchChannels: async () => {
//             let {data} = await axios.get('https://geek.itheima.net/v1_0/channels')
//             console.log(data)
//             if (data.message === 'OK')
//                 set({
//                     channelList: data.data.channels
//                 })
//         },
//         handleClearChannels:  () => {
//             set({
//                 channelList: [],
//             })
//         }
//     }
// }

// 1. 创建store
// const useStore = create((...a) => {
//     return {
//         ...createCounterStore(...a),
//         ...createChannelStore(...a),
//     }
// })

// 2. 在组件中使用
const App = () => {
    // const {count, inc, incd, channelList, handleFetchChannels, handleClearChannels} = useStore()
    const store = useStore()
    return (
        <div>
            <button onClick={() => store.inc()}>基于原数据进行修改 val: {store.count}</button>
            <button onClick={() => store.incd()}>不需要原数据修改 val: {store.count}</button>
            <button onClick={() => store.handleFetchChannels()}>发送异步请求</button>
            <button onClick={() => store.handleClearChannels()}>删除列表</button>
            <ul>
                {store.channelList.map((channel) => (
                    <li key={channel.id}>{channel.name}</li>
                ))}
            </ul>
        </div>
    )
}

export default App