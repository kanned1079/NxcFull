import {createSlice} from "@reduxjs/toolkit";
import axios from "axios";

const channelStore = createSlice({
    name: 'channel',
    initialState: {
        channelList: [],
    },
    reducers: {
        setChannels(state, action) {    // 同步修改
            state.channelList = action.payload
        }
    }
})

const {setChannels} = channelStore.actions

// 异步请求部分
const fetchChannels = () => {
    return async (dispatch) => {
        const res =  await axios.get('http://geek.itheima.net/v1_0/channels')
        dispatch(setChannels(res.data.data.channels))
    }
}

export {setChannels, fetchChannels};

// 获取reducer
const reducer = channelStore.reducer

// 以默认导出的方式导出reducer
export default reducer
