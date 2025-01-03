import {Component, forwardRef, memo, useImperativeHandle, useRef, useState} from "react"

import {create} from "zustand"

// 1. 创建store
const useStore = create((set) => {
    return {
        // 状态数据
        count: 0,
        // 修改数据的方法
        inc: () => {
            set((state) => ({count: state.count + 1}))  // 基于原数据进行修改
        },
        incd: () => {
            set({count: 100})   // 不需要原数据
        }
    }
})

// 2. 在组件中使用
const App = () => {
    const {count, inc, incd} = useStore()

    return (
        <div>
            <button onClick={() => inc()}>基于原数据进行修改 val: {count}</button>
            <button onClick={() => incd()}>不需要原数据修改 val: {count}</button>
        </div>
    )
}

export default App