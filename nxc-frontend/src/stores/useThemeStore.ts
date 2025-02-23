import {computed, type ComputedRef, type Ref, ref} from 'vue'
import {defineStore} from "pinia";
import {darkTheme, type GlobalThemeOverrides} from 'naive-ui';
import useAppInfosStore from "@/stores/useAppInfosStore";
import {glacierBlueLight, glacierBlueNight} from "@/themes/glacierBlue";
import {grayThemeLight, grayThemeNight} from "@/themes/grayTheme";
import {mistyPineLight, mistyPineNight} from "@/themes/mistyPine";
import {bambooGreenLight, bambooGreenNight} from "@/themes/bambooGreen";
import {milkGreenDayLight, milkGreenDayNight} from "@/themes/milkGreenDay";
import {darkBlueDayLight, darkBlueDayNight} from "@/themes/darkBlueDay";

interface CustomThemeConfig {
    topLogoBgColor: string;
    topLogoTextColor: string;
    topHeaderBgColor: string;
    topHeaderTextColor: string;
    globeTheme: {
        asideBgColor: string;
        cardBgColor: string;
    };
    selfOverride: GlobalThemeOverrides;
}

interface Theme {
    enableDarkMode: Ref<boolean>;
    selectedTheme: Ref<string>;
    darkBlueDay: Ref<CustomThemeConfig>;
    milkGreenDay?: Ref<CustomThemeConfig>;
    biliPink?: Ref<CustomThemeConfig>;
    grayTheme: Ref<CustomThemeConfig>;
    bambooGreen?: Ref<CustomThemeConfig>;
    allTheme?: Ref<Record<string, Ref<GlobalThemeOverrides>>>;
    // getTheme: ComputedRef<CustomThemeConfig>;
    getTheme: ComputedRef<CustomThemeConfig>;
    getMainTheme: ComputedRef<any | null>;
    readEnableDarkMode: () => void;
    saveEnableDarkMode: () => void;
    setAdminPageTheme: (theme: string) => void;
    setThemeFromSetting: () => void;
    contentPath: Ref<string>;
    menuSelected: Ref<string>;
    userPath: Ref<string>;
    breadcrumb: Ref<string>;
    menuCollapsed: Ref<boolean>;
    menuIsFlippActive: Ref<boolean>;
}

const useThemeStore = defineStore('themeStore', (): Theme => {
    // 是否启用深色模式
    const enableDarkMode = ref<boolean>(false)
    // 主題名選擇
    // const selectedTheme = ref<string>('glacierBlue')
    const selectedTheme = ref<string>('glacierBlue')

    // 如在个性化中修改了主题 或者app挂载时候需要调用来应用设置
    let setThemeFromSetting = () => selectedTheme.value = useAppInfosStore().appCommonConfig.frontend_theme

    // saveTheme 保存到浏览器localStorage深色配置
    const saveEnableDarkMode = () => localStorage.setItem('themeCode', JSON.stringify({code: enableDarkMode.value as boolean}));

    // 读取浏览器缓存中的配置 处理主题设置
    const readEnableDarkMode = () => !localStorage.getItem('themeCode') ? saveEnableDarkMode() : enableDarkMode.value = JSON.parse(localStorage.getItem('themeCode') as string).code as boolean;

    // 通用部分 计算属性
    const globeTheme = ref({
        asideBgColor: computed(() => enableDarkMode.value ? '#282929' : '#fff'),
        cardBgColor: computed(() => enableDarkMode.value ? '#282929' : '#fff'),
    })

    // 主题配置 深蓝色
    const darkBlueDay = ref({
        topLogoBgColor: computed(() => enableDarkMode.value ? '#2a4564' : '#396a91'),
        topLogoTextColor: '#dbe9f3',
        topHeaderBgColor: computed(() => enableDarkMode.value ? '#2f4f6d' : '#4f7da6'),
        topHeaderTextColor: '#fff',
        globeTheme, // 通用部分
        selfOverride: computed(() => enableDarkMode.value ? {
            // 深色模式
            ...darkBlueDayNight,
        } as GlobalThemeOverrides : {
            // 浅色模式
            ...darkBlueDayLight,
        } as GlobalThemeOverrides)
    })

    // 主题配置 奶绿色
    const milkGreenDay = ref({
        topLogoBgColor: computed(() => enableDarkMode.value ? '#006f61' : '#33a79e'),
        topLogoTextColor: '#dbe9f3',
        topHeaderBgColor: computed(() => enableDarkMode.value ? '#005c4e' : '#1a9e8f'),
        topHeaderTextColor: '#fff',
        globeTheme, // 通用部分
        selfOverride: computed(() => enableDarkMode.value ? {
            // 深色模式
            ...milkGreenDayNight,
        } as GlobalThemeOverrides : {
            // 浅色模式
            ...milkGreenDayLight,
        } as GlobalThemeOverrides)
    })

    // // 主题配置 b站粉色
    // const biliPink = ref({
    //     topLogoBgColor: computed(() => enableDarkMode.value ? '#ec4374' : '#f45987'),
    //     topLogoTextColor: '#bfe1e0',
    //     topHeaderBgColor: computed(() => enableDarkMode.value ? '#f4648d' : '#fb7299'),
    //     topHeaderTextColor: '#fff',
    //     globeTheme, // 通用部分
    //     selfOverride: computed(() => enableDarkMode.value ? {
    //         // 深色模式
    //         common: {
    //             primaryColor: '#f4648d',
    //             bodyColor: '#2d2f2f',
    //             modalColor: '#2d2f2f',
    //             tableColorStriped: 'rgba(40, 41, 41, 1)',
    //             tableColor: 'rgba(40, 41, 41, 1)',
    //             tableHeaderColor: 'rgba(40, 41, 41, 1)',
    //             tableColorHover: 'rgba(40, 41, 41, 0.5)',
    //
    //         },
    //         Button: {
    //             color: '#f4648d',
    //             colorFocusPrimary: '#fda4af',
    //             colorHoverPrimary: '#fda4af',
    //             colorPressedPrimary: '#fda4af',
    //
    //             textColorPrimary: '#fff',
    //             textColorHoverPrimary: '#fff',
    //             textColorPressedPrimary: '#fff',
    //             textColorFocusPrimary: '#fff',
    //             textColorError: '#ffe9e9',
    //             textColorHoverError: '#ffe9e9',
    //             textColorPressedError: '#ffe9e9',
    //             textColorFocusError: '#ffe9e9',
    //         },
    //         Input: {
    //             borderHover: '1px solid #3b4e72',
    //             borderFocus: '1px solid #3b4e72',
    //         },
    //         Card: {
    //             colorEmbedded: 'rgba(40, 41, 41, 1)',
    //         },
    //         Switch: {
    //             railColorActive: '#2a6f6a'
    //         },
    //         DataTable: {
    //             thColor: 'rgba(40, 41, 41, 1)',
    //             tdColor: 'rgba(40, 41, 41, 1)',
    //         },
    //         Notification: {
    //             color: 'rgba(40, 41, 41, 1)',
    //         }
    //     } as GlobalThemeOverrides : {
    //         // 浅色模式
    //         common: {
    //             primaryColor: '#fb7299',
    //             bodyColor: '#eff2f7',
    //             modalColor: '#eff2f7',
    //         },
    //         Button: {
    //             // 主体颜色
    //             color: '#fb7299',
    //             colorFocusPrimary: '#f66790',
    //             colorHoverPrimary: '#f66790',
    //             colorPressedPrimary: '#f66790',
    //             // 颜色主题
    //             textColorPrimary: '#fff',
    //             textColorHoverPrimary: '#fff',
    //             textColorPressedPrimary: '#fff',
    //             textColorFocusPrimary: '#fff',
    //         },
    //         Input: {
    //             borderHover: '1px solid #385894',
    //             borderFocus: '1px solid #385894',
    //         },
    //         Card: {
    //             colorEmbedded: '#fff',
    //         },
    //         Switch: {
    //             railColorActive: '#fb7299'
    //         },
    //         DataTable: {
    //             borderRadius: '3px',
    //         },
    //         Notification: {
    //             color: 'rgba(255, 255, 255, 1)',
    //             bordered: '1px solid #385894',
    //         }
    //     } as GlobalThemeOverrides)
    // })

    // 主题配置 若竹
    const bambooGreen = ref({
        topLogoBgColor: computed(() => enableDarkMode.value ? '#3d6e4a' : '#5b9c7e'),
        topLogoTextColor: '#dbe9f3',
        topHeaderBgColor: computed(() => enableDarkMode.value ? '#355d42' : '#4c8569'),
        topHeaderTextColor: '#fff',
        globeTheme, // 通用部分
        selfOverride: computed(() => enableDarkMode.value ? {
            // 深色模式
            ...bambooGreenNight,
        } as GlobalThemeOverrides : {
            // 浅色模式
            ...bambooGreenLight,
        } as GlobalThemeOverrides)
    })

    const mistyPine = ref({
        topLogoBgColor: computed(() => enableDarkMode.value ? '#4f6d6a' : '#7a9a97'),
        topLogoTextColor: '#dbe9f3',
        topHeaderBgColor: computed(() => enableDarkMode.value ? '#3e5b59' : '#6a8f88'),
        topHeaderTextColor: '#fff',
        globeTheme, // 通用部分
        selfOverride: computed(() => enableDarkMode.value ? {
            // 深色模式
            ...mistyPineNight,
        } as GlobalThemeOverrides : {
            // 浅色模式
           ...mistyPineLight,
        } as GlobalThemeOverrides)
    })

    const glacierBlue = ref({
        topLogoBgColor: computed(() => enableDarkMode.value ? '#497095' : '#5c8db6'),
        topLogoTextColor: '#dbe9f3',
        topHeaderBgColor: computed(() => enableDarkMode.value ? '#4f7698' : '#6390b9'),
        topHeaderTextColor: '#fff',
        globeTheme, // 通用部分
        selfOverride: computed(() => enableDarkMode.value ? {
            //深色模式
            ...glacierBlueNight,
        } as GlobalThemeOverrides : {
            // 浅色模式
            ...glacierBlueLight,
        } as GlobalThemeOverrides)
    })

    const grayTheme = ref({
        topLogoBgColor: computed(() => enableDarkMode.value ? '#636363' : '#8e8e8e'),
        topLogoTextColor: '#dbe9f3',
        topHeaderBgColor: computed(() => enableDarkMode.value ? '#595959' : '#777777'),
        topHeaderTextColor: '#fff',
        globeTheme, // 通用部分
        selfOverride: computed(() => enableDarkMode.value ? {
            // 深色模式
            ...grayThemeNight,
        } as GlobalThemeOverrides : {
            // 浅色模式
           ...grayThemeLight,
        } as GlobalThemeOverrides)
    })

    // 主题配置 默认
    const defaultDay = ref({
        darkBlueDay,
    })

    // 所有的主题集合
    // const allTheme = ref({
    //     defaultDay,     // 默认主题
    //     darkBlueDay,    // 深蓝色
    //     milkGreenDay,   // 奶绿色
    // })

    const setAdminPageTheme = (theme: string) => selectedTheme.value = theme


    // getMainTheme 获取默认的主题
    const getMainTheme = computed((): any => (enableDarkMode.value ? darkTheme : null))

    // getTheme 这个是覆盖的主题
    const getTheme = computed(() => {
        switch (selectedTheme.value) {
            case 'darkBlueDay':
                return darkBlueDay.value;

            case 'milkGreenDay':
                return milkGreenDay.value;

            // case 'biliPink':
            //     return biliPink.value;

            case 'bambooGreen':
                return bambooGreen.value;

            case 'mistyPine':
                return mistyPine.value;

            case 'glacierBlue':
                return glacierBlue.value;

            case 'grayTheme':
                return grayTheme.value;
            default:
                return darkBlueDay.value;

        }
    }) as ComputedRef<CustomThemeConfig>

    // 菜单折叠
    let menuCollapsed = ref<boolean>(false)
    let menuIsFlippActive = ref<boolean>(false)

    // 当前访问的位置
    let userPath = ref<string>('/dashboard/summary')
    let contentPath = ref<string>('/admin/dashboard/summary')
    let menuSelected = ref<string>('dashboard')

    let breadcrumb = ref<string>('')

    return {
        selectedTheme,
        enableDarkMode,
        darkBlueDay,
        milkGreenDay,
        // biliPink,
        bambooGreen,
        // allTheme,
        getTheme,
        getMainTheme,
        readEnableDarkMode,
        saveEnableDarkMode,
        setAdminPageTheme,
        setThemeFromSetting,
        // backgroundUrl,
        contentPath,
        menuSelected,
        userPath,
        menuCollapsed,
        menuIsFlippActive,
        breadcrumb,
        grayTheme
    }

}, {
    persist: true,
})

export default useThemeStore;