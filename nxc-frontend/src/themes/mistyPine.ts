import {type GlobalThemeOverrides} from "naive-ui"

export const mistyPineLight: GlobalThemeOverrides = {
    common: {
        primaryColor: '#7a9a97',
        primaryColorPressed: '#7a9a97',
        bodyColor: '#eff2f7',
        modalColor: '#eff2f7',
        tableColor: '#eff2f7',
        primaryColorHover: '#7a9a97',
        borderRadius: '6px',
    },
    Button: {
        color: '#7a9a97',
        colorFocusPrimary: '#6a8f88',
        colorHoverPrimary: '#6a8f88',
        colorPressedPrimary: '#6a8f88',
        textColorPrimary: '#fff',
        textColorHoverPrimary: '#fff',
        textColorPressedPrimary: '#fff',
        textColorFocusPrimary: '#fff',
        colorTextPrimary: '#7a9a97',
        colorTextHoverPrimary: '#6a8f88',
        textColorTextHoverPrimary: '#6a8f88',
        textColorPrimaryDisabled: '#fff',
    },
    Input: {
        borderHover: '1px solid #7a9a97',
        borderFocus: '1px solid #7a9a97',
    },
    Card: {
        colorEmbedded: '#fff',  // 比bodyColor略深的颜色
        borderRadius: '6px',
    },
    Switch: {
        railColorActive: '#7a9a97',
    },
    DataTable: {
        borderRadius: '6px !important',
    },
    Notification: {
        color: 'rgba(255, 255, 255, 1)',
        bordered: '1px solid #7a9a97',
    },
    Dropdown: {
        color: '#fff',
        borderFocus: '1px solid #7a9a97',
        borderHover: '1px solid #7a9a97',
    },
    Select: {
        borderFocus: '1px solid #7a9a97',
        borderHover: '1px solid #7a9a97',
    },
    Carousel: {
        borderRadius: '6px',
    },
    Tabs: {
        tabBorderRadius: '6px',
    },
}

export const mistyPineNight: GlobalThemeOverrides = {
    common: {
        primaryColor: '#4f6d6a',
        primaryColorPressed: '#4f6d6a',
        bodyColor: '#2d2f2f',
        modalColor: '#2d2f2f',
        tableColorStriped: 'rgba(40, 41, 41, 1)',
        tableColor: 'rgba(40, 41, 41, 1)',
        tableHeaderColor: 'rgba(40, 41, 41, 1)',
        tableColorHover: 'rgba(40, 41, 41, 0.5)',
        primaryColorHover: '#4f6d6a',
        borderRadius: '6px',
    },
    Button: {
        color: '#4f6d6a',
        colorFocusPrimary: '#5d8e88',
        colorHoverPrimary: '#5d8e88',
        colorPressedPrimary: '#5d8e88',
        textColorPrimary: '#fff',
        textColorHoverPrimary: '#fff',
        textColorPressedPrimary: '#fff',
        textColorFocusPrimary: '#fff',
        textColorError: '#ffe3e3',
        textColorHoverError: '#ffe3e3',
        textColorPressedError: '#ffe3e3',
        textColorFocusError: '#ffe3e3',
        textColorTextPrimary: '#fff',
        textColorTextHoverPrimary: '#5d8e88',
    },
    Input: {
        borderHover: '1px solid #4f6d6a',
        borderFocus: '1px solid #4f6d6a',
    },
    Card: {
        colorEmbedded: 'rgba(40, 41, 41, 1)',
        borderRadius: '6px',
    },
    Switch: {
        railColorActive: '#4f6d6a',
    },
    DataTable: {
        thColor: 'rgba(40, 41, 41, 1)',
        tdColor: 'rgba(40, 41, 41, 1)',
        borderRadius: '6px !important',
    },
    Notification: {
        color: 'rgba(40, 45, 47, 1)',
    },
    Menu: {
        itemTextColorActive: '#7a9a97',
        itemTextColorActiveHover: '#6b8d8c',
        itemIconColorActive: '#7a9a97',
        itemIconColorActiveHover: '#6b8d8c',
        itemTextColorChildActive: '#7a9a97',
        itemTextColorChildActiveHover: '#6b8d8c',
        arrowColorChildActive: '#7a9a97',
        arrowColorChildActiveHover: '#6b8d8c',
    },
    Dropdown: {
        color: '#252525',
    },
    Carousel: {
        borderRadius: '6px',
    },
    Tabs: {
        tabBorderRadius: '6px',
    },
}