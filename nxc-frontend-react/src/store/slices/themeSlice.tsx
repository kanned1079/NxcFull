import {StateCreator} from "zustand";
import type {ConfigProviderProps, ButtonProps, ThemeConfig} from "antd";

interface ThemeState {
    theme: 'light' | 'dark';
    toggleTheme: () => void;
    glacierBlue?: ConfigProviderProps
}

const useThemeSlice: StateCreator<ThemeState, [], [], ThemeState> = (set) => ({
    theme: 'light',
    glacierBlue: {  // 主题配置 冰川蓝
    },
    toggleTheme: () => set((state: ThemeState) => ({
        theme: state.theme = 'dark'
    }))
})

export {
    type ThemeState,
    useThemeSlice
}