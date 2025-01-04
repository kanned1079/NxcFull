import {StateCreator} from "zustand";

interface CountState {
    count: number;
    setCount: () => void
}

const useCountSlice: StateCreator<CountState, [], [], CountState> = (set) => ({
    count: 1,
    setCount: () => set(state => ({
        count: state.count + 1,  // 返回一个新的状态对象
    }))
})

export {
    type CountState,
    useCountSlice,
}