import {type GlobalThemeOverrides} from "naive-ui"

export const bambooGreenLight: GlobalThemeOverrides = {
    common: {
        primaryColor: '#5b9c7e',
        primaryColorPressed: '#5b9c7e',
        bodyColor: '#eff2f7',
        modalColor: '#eff2f7',
        tableColor: '#eff2f7',
        primaryColorHover: '#5b9c7e',
        borderRadius: '6px',
    },
    Button: {
        color: '#5b9c7e',
        colorFocusPrimary: '#4d856b',
        colorHoverPrimary: '#4d856b',
        colorPressedPrimary: '#4d856b',
        textColorPrimary: '#fff',
        textColorHoverPrimary: '#fff',
        textColorPressedPrimary: '#fff',
        textColorFocusPrimary: '#fff',
        colorTextPrimary: '#5b9c7e',
        colorTextHoverPrimary: '#4d856b',
        textColorTextHoverPrimary: '#4d856b',
        textColorPrimaryDisabled: '#fff',
    },
    Input: {
        borderHover: '1px solid #5b9c7e',
        borderFocus: '1px solid #5b9c7e',
    },
    Card: {
        colorEmbedded: '#fff',  // 比bodyColor略深的颜色
        borderRadius: '6px',
    },
    Switch: {
        railColorActive: '#5b9c7e',
    },
    DataTable: {
        borderRadius: '6px !important',
    },
    Notification: {
        color: 'rgba(255, 255, 255, 1)',
        bordered: '1px solid #5b9c7e',
    },
    Dropdown: {
        color: '#fff',
        borderFocus: '1px solid #5b9c7e',
        borderHover: '1px solid #5b9c7e',
    },
    Select: {
        borderFocus: '1px solid #5b9c7e',
        borderHover: '1px solid #5b9c7e',
    },
    Carousel: {
        borderRadius: '6px',
    },
    Tabs: {
        tabBorderRadius: '6px',
    },
}

export const bambooGreenNight: GlobalThemeOverrides = {
    common: {
        primaryColor: '#3d6e4a',
        primaryColorPressed: '#3d6e4a',
        bodyColor: '#2d2f2f',
        modalColor: '#2d2f2f',
        tableColorStriped: 'rgba(40, 41, 41, 1)',
        tableColor: 'rgba(40, 41, 41, 1)',
        tableHeaderColor: 'rgba(40, 41, 41, 1)',
        tableColorHover: 'rgba(40, 41, 41, 0.5)',
        primaryColorHover: '#3d6e4a',
        borderRadius: '6px',
    },
    Button: {
        color: '#3d6e4a',
        colorFocusPrimary: '#4a7f62',
        colorHoverPrimary: '#4a7f62',
        colorPressedPrimary: '#4a7f62',
        textColorPrimary: '#fff',
        textColorHoverPrimary: '#fff',
        textColorPressedPrimary: '#fff',
        textColorFocusPrimary: '#fff',
        textColorError: '#ffe3e3',
        textColorHoverError: '#ffe3e3',
        textColorPressedError: '#ffe3e3',
        textColorFocusError: '#ffe3e3',
        textColorTextPrimary: '#fff',
        textColorTextHoverPrimary: '#4a7f62',
    },
    Input: {
        borderHover: '1px solid #3d6e4a',
        borderFocus: '1px solid #3d6e4a',
    },
    Card: {
        colorEmbedded: 'rgba(40, 41, 41, 1)',
        borderRadius: '6px',
    },
    Switch: {
        railColorActive: '#3d6e4a',
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
        itemTextColorActive: '#5b9c7e',
        itemTextColorActiveHover: '#4d855b',
        itemIconColorActive: '#5b9c7e',
        itemIconColorActiveHover: '#4d855b',
        itemTextColorChildActive: '#5b9c7e',
        itemTextColorChildActiveHover: '#4d855b',
        arrowColorChildActive: '#5b9c7e',
        arrowColorChildActiveHover: '#4d855b',
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