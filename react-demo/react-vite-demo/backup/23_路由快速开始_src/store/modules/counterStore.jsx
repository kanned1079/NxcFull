import { createSlice } from "@reduxjs/toolkit";

const counterStore = createSlice({
    name: "counter",
    initialState: { // 初始数据状态
        count: 0,
    },
    reducers: {     // 修改状态的方法 同步方法 支持直接修改状态
        increment: (state) => {
            state.count++
        },
        decrement: (state) => {
            state.count--
        },
        add2Nums: (state, action) => {
            state.count = action.payload    // 这里的payload就是传递的参数
        }
    }
})

// 解构处actionCreater函数
// 以按需导出的方式导出actionCreater
export const {increment, decrement, add2Nums} = counterStore.actions
// 获取reducer
const reducer = counterStore.reducer

// 以默认导出的方式导出reducer
export default reducer

