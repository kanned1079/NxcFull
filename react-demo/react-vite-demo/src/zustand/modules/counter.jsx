const createCounterStore = (set) => {
    return {
        count: 0,
        inc: () => {
            set((state) => ({count: state.count + 1}))  // 基于原数据进行修改
        },
        incd: () => {
            set({count: 100})   // 不需要原数据
        },
    }
}

export default createCounterStore