import {StateCreator} from "zustand";
import {ButtonProps, ConfigProviderProps, type MappingAlgorithm, theme, ThemeConfig} from "antd";


interface ThemeState {
    isMenuFolded: boolean;  // 桌面模式下左侧菜单是否折叠
    setIsMenuFolded: (fold: boolean) => void    // 控制折叠的方法
    isShowMobileMenu: boolean;  // 手机模式下是否显示左侧菜单
    setIsShowMobileMenu: (fold: boolean) => void
    // darkModeEnabled: boolean;
    // theme: 'light' | 'dark';
    // toggleTheme?: () => void;
    glacierBlue?: ConfigProviderProps;
    setTheme: (algorithm: MappingAlgorithm) => void
}

const useThemeSlice: StateCreator<ThemeState, [], [], ThemeState> = (set) => ({
    isMenuFolded: false,
    setIsMenuFolded: (fold: boolean) => set(() => ({isMenuFolded: fold as boolean})),
    isShowMobileMenu: false,
    setIsShowMobileMenu: (fold: boolean) => set(() => ({isShowMobileMenu: fold as boolean})),
    // darkModeEnabled: true,
    // theme: 'light',
    glacierBlue: {  // 主题配置 冰川蓝
        theme: {
            // 根据darkModeEnabled来选择
            // 浅色 theme.defaultAlgorithm
            // 深色 theme.darkAlgorithm
            algorithm: theme.defaultAlgorithm,
            token: {
                colorPrimary: '#6390b9'
            },
            components: {
                Button: {
                    primary: '#00b96b',
                } as ButtonProps,
                Drawer: {}
            },
        } as ThemeConfig,
    },
    setTheme: (newAlgorithm: MappingAlgorithm) => set(state => ({
        glacierBlue: {
            ...state.glacierBlue,
            theme: {
                algorithm: newAlgorithm
            }
        }
    }))

})

export {
    type ThemeState,
    useThemeSlice
}