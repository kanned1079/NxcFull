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


    // 是否启用深色模式
    const enableDarkMode = ref(false)
    const selectedTheme = ref('darkBlueDay')
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
        // selfOverride1: {
        //     common: {
        //         // 主题色相关
        //         primaryColor: "#385894",
        //         // primaryColor: computed(() => enableDarkMode.value ? '#374868' : '#385894'),
        //
        //         primaryColorHover: "#45639b",
        //         primaryColorPressed: "#0c7a43",
        //         primaryColorSuppl: "#36ad6a",
        //         cardColor: "#ffffff",
        //         bodyColor: '#e3e5e7',
        //         // baseColor: '#7ecf8f',
        //         modalColor: '#ffffff',
        //         tableColor: '#e3e5e7',
        //     },
        //     Button: {
        //         colorPressed: '#fff',
        //         colorPressedPrimary: '#537fd1',
        //         borderPressedPrimary: '#537fd1',
        //         textColorPrimary: '#fff',
        //         textColorPressedPrimary: '#fff',
        //         textColorHoverPrimary: '#fff',
        //         textColorFocusPrimary: '#fff',
        //     },
        //
        // } as GlobalThemeOverrides,

        selfOverride: computed(() => enableDarkMode.value ? {
            // 深色模式
            common: {
                primaryColor: '#3b4e72',
                bodyColor: '#2d2f2f',
                modalColor: '#2d2f2f'
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
            }
        } as GlobalThemeOverrides : {
            // 浅色模式
            common: {
                primaryColor: '#385894',
                bodyColor: '#eff2f7',
                modalColor: '#eff2f7',
            },
            Card: {
                colorEmbedded: '#fff',
            },
        } as GlobalThemeOverrides)
    })

    // 主题配置 奶绿色
    const milkGreenDay = reactive({
        topLogoBgColor: computed(() => enableDarkMode.value ? '#2a6f6a' : '#008784'),
        topLogoTextColor: '#bfe1e0',
        topHeaderBgColor: computed(() => enableDarkMode.value ? '#2d7974' : '#009693'),
        topHeaderTextColor: '#fff',
        globeTheme, // 通用部分
        selfOverride: {
            common: {
                // 主题色相关
                // primaryColor: "#009693",
                primaryColor: computed(() => enableDarkMode.value ? '#2a6f6a' : '#009693'),
                primaryColorHover: "#009693",
                primaryColorPressed: "#0c7a43",
                primaryColorSuppl: "#36ad6a",
                cardColor: computed(() => enableDarkMode.value ? '#2a6f6a' : '#008784'),
            }
        } as GlobalThemeOverrides
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