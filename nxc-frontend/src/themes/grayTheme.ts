import {type GlobalThemeOverrides} from "naive-ui"

export const grayThemeLight: GlobalThemeOverrides = {
    common: {
        primaryColor: '#9c9c9c',
        primaryColorPressed: '#9c9c9c',
        bodyColor: '#eff2f7',
        modalColor: '#eff2f7',
        tableColor: '#eff2f7',
        primaryColorHover: '#9c9c9c',
        borderRadius: '6px',
    },
    Button: {
        color: '#9c9c9c',
        colorFocusPrimary: '#858585',
        colorHoverPrimary: '#858585',
        colorPressedPrimary: '#858585',
        textColorPrimary: '#fff',
        textColorHoverPrimary: '#fff',
        textColorPressedPrimary: '#fff',
        textColorFocusPrimary: '#fff',
        colorTextPrimary: '#9c9c9c',
        colorTextHoverPrimary: '#858585',
        textColorTextHoverPrimary: '#858585',
        textColorPrimaryDisabled: '#fff',
    },
    Input: {
        borderHover: '1px solid #9c9c9c',
        borderFocus: '1px solid #9c9c9c',
    },
    Card: {
        colorEmbedded: '#fff',  // 比bodyColor略深的颜色
        borderRadius: '6px',
    },
    Switch: {
        railColorActive: '#9c9c9c',
    },
    DataTable: {
        borderRadius: '6px !important',
    },
    Notification: {
        color: 'rgba(255, 255, 255, 1)',
        bordered: '1px solid #9c9c9c',
    },
    Dropdown: {
        color: '#fff',
        borderFocus: '1px solid #9c9c9c',
        borderHover: '1px solid #9c9c9c',
    },
    Select: {
        borderFocus: '1px solid #9c9c9c',
        borderHover: '1px solid #9c9c9c',
    },
    Carousel: {
        borderRadius: '6px',
    },
    Tabs: {
        tabBorderRadius: '6px',
    },
}

export const grayThemeNight: GlobalThemeOverrides = {
    common: {
        primaryColor: '#707070',   // 更深的灰色
        primaryColorPressed: '#707070',
        bodyColor: '#2d2f2f',      // 更深的灰色
        modalColor: '#2d2f2f',     // 更深的灰色
        tableColorStriped: 'rgba(40, 41, 41, 1)',
        tableColor: 'rgba(40, 41, 41, 1)',
        tableHeaderColor: 'rgba(40, 41, 41, 1)',
        tableColorHover: 'rgba(40, 41, 41, 0.5)',
        primaryColorHover: '#707070',
        borderRadius: '6px',
    },
    Button: {
        color: '#707070',
        colorFocusPrimary: '#8d8d8d',
        colorHoverPrimary: '#8d8d8d',
        colorPressedPrimary: '#8d8d8d',
        textColorPrimary: '#fff',
        textColorHoverPrimary: '#fff',
        textColorPressedPrimary: '#fff',
        textColorFocusPrimary: '#fff',
        textColorError: '#ffe3e3',
        textColorHoverError: '#ffe3e3',
        textColorPressedError: '#ffe3e3',
        textColorFocusError: '#ffe3e3',
        textColorTextPrimary: '#fff',
        textColorTextHoverPrimary: '#8d8d8d',
    },
    Input: {
        borderHover: '1px solid #707070',
        borderFocus: '1px solid #707070',
    },
    Card: {
        colorEmbedded: 'rgba(40, 41, 41, 1)',
        borderRadius: '6px',
    },
    Switch: {
        railColorActive: '#707070',
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
        itemTextColorActive: '#9c9c9c',
        itemTextColorActiveHover: '#8a8a8a',
        itemIconColorActive: '#9c9c9c',
        itemIconColorActiveHover: '#8a8a8a',
        itemTextColorChildActive: '#9c9c9c',
        itemTextColorChildActiveHover: '#8a8a8a',
        arrowColorChildActive: '#9c9c9c',
        arrowColorChildActiveHover: '#8a8a8a',
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