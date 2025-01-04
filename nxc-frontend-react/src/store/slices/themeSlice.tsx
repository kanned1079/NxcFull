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
        theme: {
            token: {
              colorPrimary: ''
            },
            components: {
                Button: {
                },
                Drawer: {

                }
            },
        }
    },
    toggleTheme: () => set((state: ThemeState) => ({
        theme: state.theme = 'dark'
    }))
})

export {
    type ThemeState,
    useThemeSlice
}