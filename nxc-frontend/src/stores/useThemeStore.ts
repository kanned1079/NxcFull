import {computed, reactive, ref} from 'vue'
import {defineStore} from "pinia";
import {darkTheme, type GlobalThemeOverrides} from 'naive-ui';
import useSettingStore from "@/stores/useSettingStore"

const useThemeStore = defineStore('theme', () => {

    // const darkOverride: GlobalThemeOverrides = {
    //     Notification: {
    //         tex
    //     }
    // }
    // type SelectThemeOverrides = NonNullable<SelectProps['themeOverrides']>


    // 是否启用深色模式
    const enableDarkMode = ref(false)
    // 主題名選擇
    const selectedTheme = ref('')
    // const selectedTheme = useSettingStore().settings.frontend.frontend_theme

    // 如在个性化中修改了主题 或者app挂载时候需要调用来应用设置
    let setThemeFromSetting = () => selectedTheme.value = useSettingStore().settings.frontend.frontend_theme

    // saveTheme 保存到浏览器localStorage深色配置
    const saveEnableDarkMode = () => localStorage.setItem('themeCode', JSON.stringify({code: enableDarkMode.value}));

    // 读取浏览器缓存中的配置 处理主题设置
    const readEnableDarkMode = () => !localStorage.getItem('themeCode') ? saveEnableDarkMode() : enableDarkMode.value = JSON.parse(localStorage.getItem('themeCode') as string).code as boolean;

    // 通用部分 计算属性
    const globeTheme = reactive({
        asideBgColor: computed(() => enableDarkMode.value ? '#282929' : '#fff'),
        //contentBgColor: computed(() => enableDarkMode.value ? '#2d2f2f' : '#eff2f7'),
        //cardBgColor: computed(() => enableDarkMode.value ? 'rgba(40, 41, 41, 1)' : '#fff'),
        //loginCardBgColor: computed(() => enableDarkMode.value ? 'rgba(40, 41, 41, 0.7)' : 'rgba(255, 255, 255, 0.7)'),
    })

    // 等六j
    const backgroundUrl = ref('https://ikanned.com:24444/d/Upload/pexels-martin-p%C3%A9chy-5335217.jpg')

    // 主题配置 深蓝色
    const darkBlueDay = reactive({
        topLogoBgColor: computed(() => enableDarkMode.value ? '#374868' : '#324f85'),
        topLogoTextColor: '#bfc8d9',
        topHeaderBgColor: computed(() => enableDarkMode.value ? '#3b4e72' : '#385894'),
        topHeaderTextColor: '#fff',
        globeTheme, // 通用部分
        selfOverride: computed(() => enableDarkMode.value ? {
            // 深色模式
            common: {
                primaryColor: '#3b4e72',
                bodyColor: '#2d2f2f',
                modalColor: '#2d2f2f',
                tableColorStriped: 'rgba(40, 41, 41, 1)',
                tableColor: 'rgba(40, 41, 41, 1)',
                tableHeaderColor: 'rgba(40, 41, 41, 1)',
                tableColorHover: 'rgba(40, 41, 41, 0.5)',

            },
            Button: {
                color: '#3b4e72',
                colorFocusPrimary: '#3b4e72',
                colorHoverPrimary: '#496097',
                colorPressedPrimary: '#496097',

                textColorPrimary: '#fff',
                textColorHoverPrimary: '#fff',
                textColorPressedPrimary: '#fff',
                textColorFocusPrimary: '#fff',
                textColorError: '#ffe9e9',
                textColorHoverError: '#ffe9e9',
                textColorPressedError: '#ffe9e9',
                textColorFocusError: '#ffe9e9',
            },
            Input: {
                borderHover: '1px solid #3b4e72',
                borderFocus: '1px solid #3b4e72',
            },
            Card: {
                colorEmbedded: 'rgba(40, 41, 41, 1)',
            },
            Switch: {
                railColorActive: '#3b4e72'
            },
            DataTable: {
                thColor: 'rgba(40, 41, 41, 1)',
                tdColor: 'rgba(40, 41, 41, 1)',
                borderRadius: '3px',
            },
            Notification: {
                color: 'rgba(40, 41, 41, 1)',
            }
        } as GlobalThemeOverrides : {
            // 浅色模式
            common: {
                primaryColor: '#385894',
                bodyColor: '#eff2f7',
                modalColor: '#eff2f7',
                tableColor: '#eff2f7',
            },
            Button: {
                // 主体颜色
                color: '#385894',
                colorFocusPrimary: '#385894',
                colorHoverPrimary: '#2f508e',
                colorPressedPrimary: '#2f508e',
                // 颜色主题
                textColorPrimary: '#fff',
                textColorHoverPrimary: '#fff',
                textColorPressedPrimary: '#fff',
                textColorFocusPrimary: '#fff',

                colorTextPrimary: '#385894',
                colorTextHoverPrimary: '#2f508e',
                textColorTextHoverPrimary: '#2f508e'
            },
            Input: {
                borderHover: '1px solid #385894',
                borderFocus: '1px solid #385894',
            },
            Card: {
                colorEmbedded: '#fff',
            },
            Switch: {
                railColorActive: '#385894'
            },
            DataTable: {
                borderRadius: '3px',
            },
            Notification: {
                color: 'rgba(255, 255, 255, 1)',
                bordered: '1px solid #385894',
            },
            Dropdown: {
                color: 'rgba(40, 41, 41, 1)'
            }
        } as GlobalThemeOverrides)
    })

    // 主题配置 奶绿色
    const milkGreenDay = reactive({
        topLogoBgColor: computed(() => enableDarkMode.value ? '#2a6f6a' : '#008784'),
        topLogoTextColor: '#bfe1e0',
        topHeaderBgColor: computed(() => enableDarkMode.value ? '#2d7974' : '#009693'),
        topHeaderTextColor: '#fff',
        globeTheme, // 通用部分
        selfOverride: computed(() => enableDarkMode.value ? {
            // 深色模式
            common: {
                primaryColor: '#2a6f6a',
                bodyColor: '#2d2f2f',
                modalColor: '#2d2f2f',
                tableColorStriped: 'rgba(40, 41, 41, 1)',
                tableColor: 'rgba(40, 41, 41, 1)',
                tableHeaderColor: 'rgba(40, 41, 41, 1)',
                tableColorHover: 'rgba(40, 41, 41, 0.5)',

            },
            Button: {
                color: '#2a6f6a',
                colorFocusPrimary: '#2a6f6a',
                colorHoverPrimary: '#358c85',
                colorPressedPrimary: '#358c85',

                textColorPrimary: '#fff',
                textColorHoverPrimary: '#fff',
                textColorPressedPrimary: '#fff',
                textColorFocusPrimary: '#fff',
                textColorError: '#ffe9e9',
                textColorHoverError: '#ffe9e9',
                textColorPressedError: '#ffe9e9',
                textColorFocusError: '#ffe9e9',
            },
            Input: {
                borderHover: '1px solid #3b4e72',
                borderFocus: '1px solid #3b4e72',
            },
            Card: {
                colorEmbedded: 'rgba(40, 41, 41, 1)',
            },
            Switch: {
                railColorActive: '#2a6f6a'
            },
            DataTable: {
                thColor: 'rgba(40, 41, 41, 1)',
                tdColor: 'rgba(40, 41, 41, 1)',
                borderRadius: '3px',
            },
            Notification: {
                color: 'rgba(40, 41, 41, 1)',
            }
        } as GlobalThemeOverrides : {
            // 浅色模式
            common: {
                primaryColor: '#008784',
                bodyColor: '#eff2f7',
                modalColor: '#eff2f7',
            },
            Button: {
                // 主体颜色
                color: '#008784',
                colorFocusPrimary: '#008784',
                colorHoverPrimary: '#0e807c',
                colorPressedPrimary: '#0e807c',
                // 颜色主题
                textColorPrimary: '#fff',
                textColorHoverPrimary: '#fff',
                textColorPressedPrimary: '#fff',
                textColorFocusPrimary: '#fff',
            },
            Input: {
                borderHover: '1px solid #385894',
                borderFocus: '1px solid #385894',
            },
            Card: {
                colorEmbedded: '#fff',
            },
            Switch: {
                railColorActive: '#008784'
            },
            DataTable: {
                borderRadius: '3px',
            },
            Notification: {
                color: 'rgba(255, 255, 255, 1)',
                bordered: '1px solid #385894',
            }
        } as GlobalThemeOverrides)
    })

    // 主题配置 b站粉色
    const biliPink = reactive({
        topLogoBgColor: computed(() => enableDarkMode.value ? '#ec4374' : '#f45987'),
        topLogoTextColor: '#bfe1e0',
        topHeaderBgColor: computed(() => enableDarkMode.value ? '#f4648d' : '#fb7299'),
        topHeaderTextColor: '#fff',
        globeTheme, // 通用部分
        selfOverride: computed(() => enableDarkMode.value ? {
            // 深色模式
            common: {
                primaryColor: '#f4648d',
                bodyColor: '#2d2f2f',
                modalColor: '#2d2f2f',
                tableColorStriped: 'rgba(40, 41, 41, 1)',
                tableColor: 'rgba(40, 41, 41, 1)',
                tableHeaderColor: 'rgba(40, 41, 41, 1)',
                tableColorHover: 'rgba(40, 41, 41, 0.5)',

            },
            Button: {
                color: '#f4648d',
                colorFocusPrimary: '#fda4af',
                colorHoverPrimary: '#fda4af',
                colorPressedPrimary: '#fda4af',

                textColorPrimary: '#fff',
                textColorHoverPrimary: '#fff',
                textColorPressedPrimary: '#fff',
                textColorFocusPrimary: '#fff',
                textColorError: '#ffe9e9',
                textColorHoverError: '#ffe9e9',
                textColorPressedError: '#ffe9e9',
                textColorFocusError: '#ffe9e9',
            },
            Input: {
                borderHover: '1px solid #3b4e72',
                borderFocus: '1px solid #3b4e72',
            },
            Card: {
                colorEmbedded: 'rgba(40, 41, 41, 1)',
            },
            Switch: {
                railColorActive: '#2a6f6a'
            },
            DataTable: {
                thColor: 'rgba(40, 41, 41, 1)',
                tdColor: 'rgba(40, 41, 41, 1)',
                borderRadius: '3px',
            },
            Notification: {
                color: 'rgba(40, 41, 41, 1)',
            }
        } as GlobalThemeOverrides : {
            // 浅色模式
            common: {
                primaryColor: '#fb7299',
                bodyColor: '#eff2f7',
                modalColor: '#eff2f7',
            },
            Button: {
                // 主体颜色
                color: '#fb7299',
                colorFocusPrimary: '#f66790',
                colorHoverPrimary: '#f66790',
                colorPressedPrimary: '#f66790',
                // 颜色主题
                textColorPrimary: '#fff',
                textColorHoverPrimary: '#fff',
                textColorPressedPrimary: '#fff',
                textColorFocusPrimary: '#fff',
            },
            Input: {
                borderHover: '1px solid #385894',
                borderFocus: '1px solid #385894',
            },
            Card: {
                colorEmbedded: '#fff',
            },
            Switch: {
                railColorActive: '#fb7299'
            },
            DataTable: {
                borderRadius: '3px',
            },
            Notification: {
                color: 'rgba(255, 255, 255, 1)',
                bordered: '1px solid #385894',
            }
        } as GlobalThemeOverrides)
    })

    // 主题配置 若竹
    const bambooGreen = reactive({
        topLogoBgColor: computed(() => enableDarkMode.value ? '#519a6d' : '#5dac81'),
        topLogoTextColor: '#bfe1e0',
        topHeaderBgColor: computed(() => enableDarkMode.value ? '#4d855b' : '#5dac81'),
        topHeaderTextColor: '#fff',
        globeTheme, // 通用部分
        selfOverride: computed(() => enableDarkMode.value ? {
            // 深色模式
            common: {
                primaryColor: '#4d855b',
                bodyColor: '#2d2f2f',
                modalColor: '#2d2f2f',
                tableColorStriped: 'rgba(40, 41, 41, 1)',
                tableColor: 'rgba(40, 41, 41, 1)',
                tableHeaderColor: 'rgba(40, 41, 41, 1)',
                tableColorHover: 'rgba(40, 41, 41, 0.5)',

            },
            Button: {
                color: '#4d855b',
                colorFocusPrimary: '#62a572',
                colorHoverPrimary: '#62a572',
                colorPressedPrimary: '#62a572',

                textColorPrimary: '#fff',
                textColorHoverPrimary: '#fff',
                textColorPressedPrimary: '#fff',
                textColorFocusPrimary: '#fff',
                textColorError: '#ffe9e9',
                textColorHoverError: '#ffe9e9',
                textColorPressedError: '#ffe9e9',
                textColorFocusError: '#ffe9e9',
            },
            Input: {
                borderHover: '1px solid #3b4e72',
                borderFocus: '1px solid #3b4e72',
            },
            Card: {
                colorEmbedded: 'rgba(40, 41, 41, 1)',
            },
            Switch: {
                railColorActive: '#4d855b'
            },
            DataTable: {
                thColor: 'rgba(40, 41, 41, 1)',
                tdColor: 'rgba(40, 41, 41, 1)',
                borderRadius: '3px',
            },
            Notification: {
                color: 'rgba(40, 41, 41, 1)',
            }
        } as GlobalThemeOverrides : {
            // 浅色模式
            common: {
                primaryColor: '#5dac81',
                bodyColor: '#eff2f7',
                modalColor: '#eff2f7',
            },
            Button: {
                // 主体颜色
                color: '#5dac81',
                colorFocusPrimary: '#4ca879',
                colorHoverPrimary: '#4ca879',
                colorPressedPrimary: '#4ca879',
                // 颜色主题
                textColorPrimary: '#fff',
                textColorHoverPrimary: '#fff',
                textColorPressedPrimary: '#fff',
                textColorFocusPrimary: '#fff',
            },
            Input: {
                borderHover: '1px solid #385894',
                borderFocus: '1px solid #385894',
            },
            Card: {
                colorEmbedded: '#fff',
                borderRadius: '3px'
            },
            Switch: {
                railColorActive: '#5dac81s'
            },
            DataTable: {
                borderRadius: '3px',
            },
            Notification: {
                color: 'rgba(255, 255, 255, 1)',
                bordered: '1px solid #385894',
            }
        } as GlobalThemeOverrides)
    })


    // 主题配置 默认
    const defaultDay = reactive({
        darkBlueDay,
    })

    // 所有的主题集合
    const allTheme = reactive({
        defaultDay,     // 默认主题
        darkBlueDay,    // 深蓝色
        milkGreenDay,   // 奶绿色
    })

    const setAdminPageTheme = (theme: string) => selectedTheme.value = theme


    // getMainTheme 获取默认的主题
    const getMainTheme = computed(() => (enableDarkMode.value ? darkTheme : null))

    // getTheme 这个是覆盖的主题
    const getTheme = computed(() => {
        switch (selectedTheme.value) {
            case 'darkBlueDay': {
                return darkBlueDay;
            }
            case 'milkGreenDay': {
                return milkGreenDay;
            }
            case 'biliPink': {
                return biliPink;
            }
            case 'bambooGreen': {
                return bambooGreen;
            }
            default: {
                return darkBlueDay;
            }
        }
    })

    // 菜单折叠
    let menuCollapsed = ref<boolean>(false)

    // 当前访问的位置
    let userPath = ref('/dashboard/summary')
    let contentPath = ref('/admin/dashboard/summary')
    let menuSelected = ref('dashboard')

    return {
        selectedTheme,
        enableDarkMode,
        darkBlueDay,
        milkGreenDay,
        allTheme,
        getTheme,
        getMainTheme,
        readEnableDarkMode,
        saveEnableDarkMode,
        setAdminPageTheme,
        setThemeFromSetting,
        backgroundUrl,
        contentPath,
        menuSelected,
        userPath,
        menuCollapsed
    }

}, {
    // persist: true,
})

export default useThemeStore;